package main

import "sync"
import "path"
import "template"
import "os"
import "fmt"

const tmplExt = ".ghtml" // Template file extension

type TemplateCache struct {
	data map[string]*template.Template
	lock *sync.RWMutex
}

func NewTemplateCache() *TemplateCache {
	v := new(TemplateCache)
	v.data = make(map[string]*template.Template)
	v.lock = new(sync.RWMutex)
	return v
}

func (this *TemplateCache) Get(name string) *template.Template {
	this.lock.RLock()
	defer this.lock.RUnlock()

	if tmpl, ok := this.data[name]; ok {
		return tmpl
	}

	return nil
}

func (this *TemplateCache) Load(root string) (err os.Error) {
	var fd *os.File
	if fd, err = os.Open(root, os.O_RDONLY, 0600); err != nil {
		return
	}

	var dirs []os.FileInfo
	if dirs, err = fd.Readdir(-1); err != nil {
		fd.Close()
		return
	}

	fd.Close()

	var p string
	var ok bool
	var tmpl *template.Template

	for _, sd := range dirs {
		p = path.Join(root, sd.Name)
		if sd.IsDirectory() {
			if err = this.Load(p); err != nil {
				return
			}
		} else {
			if tmplExt != path.Ext(p) {
				continue
			}

			fmt.Fprintf(os.Stdout, "[i] Parsing template: %s\n", p)
			if tmpl, err = template.ParseFile(p, nil); err != nil {
				return
			}

			// Strip webroot and extension
			p = p[len(context.Config().WebRoot):]
			p = p[0 : len(p)-len(tmplExt)]

			if len(p) == 0 {
				return os.NewError("Invalid template name")
			}

			this.lock.RLock()
			if _, ok = this.data[p]; ok {
				this.lock.RUnlock()
				return os.NewError("Duplicate template definition: " + p)
			}
			this.lock.RUnlock()

			this.lock.Lock()
			this.data[p] = tmpl
			this.lock.Unlock()
		}
	}

	return
}
