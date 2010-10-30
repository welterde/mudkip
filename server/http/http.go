package main

import "os"
import "io"
import "time"
import "strings"
import "fmt"
import "http"
import "path"
import "mime"
import "bytes"

func Run(cfg *Config) {
	http.HandleFunc("/", httpHandler)

	if cfg.Secure {
		fmt.Fprintf(os.Stdout, "[i] Listening on %s (secure)\n", cfg.ListenAddr)

		if err := http.ListenAndServeTLS(cfg.ListenAddr, cfg.ServerCert, cfg.ServerKey, nil); err != nil {
			fmt.Fprintf(os.Stderr, "[e] %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Fprintf(os.Stdout, "[i] Listening on %s (non-secure)\n", cfg.ListenAddr)

		if err := http.ListenAndServe(cfg.ListenAddr, nil); err != nil {
			fmt.Fprintf(os.Stderr, "[e] %v\n", err)
			os.Exit(1)
		}
	}
}

func httpHandler(rw http.ResponseWriter, req *http.Request) {
	reqtime := time.Nanoseconds() / 1000
	sc := NewServiceContext(rw, req)

	if file := path.Join(context.Config().WebRoot, req.URL.Path); fileExists(file) {
		if serveFile(sc, file) {
			log(reqtime, os.Stdout, rw, req)
		} else {
			log(reqtime, os.Stderr, rw, req)
		}
		return
	}

	if methods.Exec(sc) {
		log(reqtime, os.Stdout, rw, req)
	} else {
		log(reqtime, os.Stderr, rw, req)
		rw.WriteHeader(404)
	}
}

func postHandler(c *ServiceContext) bool {
	if err := c.Req.ParseForm(); err != nil {
		c.Status(500)
		return false
	}

	return getHandler(c)
}

func getHandler(c *ServiceContext) bool {
	if file := path.Join(context.Config().WebRoot, "index.html"); fileExists(file) {
		return serveFile(c, file)
	}

	c.Status(404)
	return true
}

func notImplementedHandler(c *ServiceContext) bool {
	c.Status(501)
	return true
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

	return c.SendReadWriter(f)
}

func log(stamp int64, w io.Writer, rw http.ResponseWriter, req *http.Request) {
	addr := rw.RemoteAddr()

	// Strip port number. Make sure we don't cut off bits of the actual IP if we
	// have an IPv6 address. It will be encased in [ and ]. eg: [::1]:1234
	if idx := strings.LastIndex(addr, ":"); idx > strings.LastIndex(addr, "]") {
		addr = addr[:idx]
	}

	fmt.Fprintf(w, "%d %s %s\n", stamp, addr, req.URL)
}
