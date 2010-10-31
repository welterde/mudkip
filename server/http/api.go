package main

import "os"
import "mime"
import "bytes"
import "strings"
import "path"
import "time"
import "io"
import "fmt"

func BindApi(methods *ServiceMethodList) (err os.Error) {
	// TODO: Bind Mudkip API
	methods.Add(NewServiceMethod("home", homeHandler, GET))
	methods.Add(NewServiceMethod("worlds", worldsHandler, GET))
	methods.Add(NewServiceMethod("worlds/play", playWorldHandler, GET))
	methods.Add(NewServiceMethod("worlds/create", createWorldHandler, GET))
	methods.Add(NewServiceMethod("worlds/edit", editWorldHandler, GET))
	methods.Add(NewServiceMethod("account", accountHandler, GET))
	methods.Add(NewServiceMethod("account/register", registerAccountHandler, GET))
	methods.Add(NewServiceMethod("account/login", loginAccountHandler, GET))
	methods.Add(NewServiceMethod("account/edit", editAccountHandler, GET))
	methods.Add(NewServiceMethod("account/logout", logoutAccountHandler, GET))

	// Catch-all handlers for HTTP commands we have not yet covered.
	methods.Add(NewServiceMethod("", homeHandler, GET))
	methods.Add(NewServiceMethod("", homeHandler, HEAD))
	methods.Add(NewServiceMethod("", postHandler, POST))
	methods.Add(NewServiceMethod("", notImplementedHandler, CONNECT))
	methods.Add(NewServiceMethod("", notImplementedHandler, DELETE))
	methods.Add(NewServiceMethod("", notImplementedHandler, OPTIONS))
	methods.Add(NewServiceMethod("", notImplementedHandler, PUT))
	methods.Add(NewServiceMethod("", notImplementedHandler, TRACE))
	return methods.Build()
}

func accountHandler(c *ServiceContext) bool {
	if c.Session.Registered {
		servePage(c, mainMenu, accountMenuB, "account", nil)
	} else {
		servePage(c, mainMenu, accountMenuA, "account", nil)
	}
	return true
}

func registerAccountHandler(c *ServiceContext) bool {
	servePage(c, mainMenu, accountMenuA, "account/register", nil)
	return true
}

func loginAccountHandler(c *ServiceContext) bool {
	servePage(c, mainMenu, accountMenuA, "account/login", nil)
	return true
}

func editAccountHandler(c *ServiceContext) bool {
	servePage(c, mainMenu, accountMenuB, "account/edit", nil)
	return true
}

func logoutAccountHandler(c *ServiceContext) bool {
	servePage(c, mainMenu, accountMenuB, "account/logout", nil)
	return true
}

func worldsHandler(c *ServiceContext) bool {
	servePage(c, mainMenu, worldsMenu, "worlds", nil)
	return true
}

func playWorldHandler(c *ServiceContext) bool {
	servePage(c, mainMenu, worldsMenu, "worlds/play", nil)
	return true
}

func createWorldHandler(c *ServiceContext) bool {
	servePage(c, mainMenu, worldsMenu, "worlds/create", nil)
	return true
}

func editWorldHandler(c *ServiceContext) bool {
	servePage(c, mainMenu, worldsMenu, "worlds/edit", nil)
	return true
}

func homeHandler(c *ServiceContext) bool {
	servePage(c, mainMenu, nil, "home", nil)
	return true
}

func postHandler(c *ServiceContext) bool {
	if err := c.Req.ParseForm(); err != nil {
		c.Status(500)
		return false
	}

	return homeHandler(c)
}

func notImplementedHandler(c *ServiceContext) bool {
	c.Status(501)
	return true
}

func servePage(c *ServiceContext, menu, submenu []*MenuItem, name string, data interface{}) {
	c.SetHeaders(200, 2592000, "text/html", -1)

	var pageinfo struct {
		Title       string
		Style       string
		HaveSubMenu bool
		Menu        []*MenuItem
		SubMenu     []*MenuItem
	}

	pageinfo.Title = name
	pageinfo.Style = c.Session.Style
	pageinfo.Menu = menu
	pageinfo.HaveSubMenu = submenu != nil
	pageinfo.SubMenu = submenu

	target := "/"
	if name != "home" {
		target += name
	}

	for _, mnu := range pageinfo.Menu {
		if mnu.Url == "/" {
			mnu.Selected = mnu.Url == target
		} else {
			mnu.Selected = strings.HasPrefix(target, mnu.Url)
		}
	}

	for _, mnu := range pageinfo.SubMenu {
		mnu.Selected = mnu.Url == target
	}

	var buf bytes.Buffer

	serveTemplate(&buf, "header", &pageinfo)
	serveTemplate(&buf, name, data)
	serveTemplate(&buf, "footer", nil)

	c.Write(buf.Bytes())
}

func serveTemplate(w io.Writer, name string, data interface{}) bool {
	if tmpl := templates.Get(name); tmpl != nil {
		if err := tmpl.Execute(data, w); err != nil {
			fmt.Fprintf(os.Stderr, "[e] %v\n", err)
			return false
		}

		return true
	}

	fmt.Fprintf(os.Stderr, "[e] Template not found: %s\n", name)
	return false
}

func serveFile(c *ServiceContext, file string) bool {
	var err os.Error
	var f *os.File
	var t *time.Time
	var modified int64

	file = path.Clean(file)
	if f, err = os.Open(file, os.O_RDONLY, 0); err != nil {
		c.Status(404)
		return false
	}

	defer f.Close()

	stat, _ := f.Stat()
	modified = stat.Mtime_ns / 1e9

	if v, ok := c.Req.Header["If_Modified_Since"]; ok {
		v = v[0:len(v)-3] + "UTC"

		if t, err = time.Parse(v, time.RFC1123); err != nil {
			fmt.Fprintf(os.Stderr, "Unrecognized time format in If_Modified_Since header: %s", v)
			return false
		}

		if modified > t.Seconds() {
			c.Status(200)
		} else {
			c.Status(304)
		}

		return true
	}

	if ctype := mime.TypeByExtension(path.Ext(file)); ctype != "" {
		c.SetHeaders(200, 2592000, ctype, modified)
	} else {
		var data []byte
		var num int64

		buf := bytes.NewBuffer(data)
		if num, err = io.Copyn(buf, f, 1024); err != nil {
			c.Status(500)
			return false
		}

		data = buf.Bytes()
		if isText(data[0:num]) {
			c.SetHeaders(200, 2592000, "text/plain; charset=utf-8", modified)
		} else {
			c.SetHeaders(200, 2592000, "application/octet-stream", modified)
		}
	}

	_, err = c.SendReadWriter(f)
	return err == nil
}
