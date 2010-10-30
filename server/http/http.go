package main

import "os"
import "io"
import "time"
import "strings"
import "fmt"
import "http"
import "path"

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
	reqtime := time.Nanoseconds() / 1e3
	sc := NewServiceContext(rw, req)

	// If we get a static file request, handle it independantly. 
	// Exclude accessing template files directly.
	if file := path.Join(context.Config().WebRoot, req.URL.Path); path.Ext(file) != tmplExt && fileExists(file) {
		if serveFile(sc, file) {
			log(reqtime, os.Stdout, rw, req)
		} else {
			log(reqtime, os.Stderr, rw, req)
		}
		return
	}

	// No static file. See if we got a service request and handle it.
	if methods.Exec(sc) {
		log(reqtime, os.Stdout, rw, req)
	} else {
		log(reqtime, os.Stderr, rw, req)
	}
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
