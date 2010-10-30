package main

import "os"
import "time"
import "strings"
import "fmt"
import "http"

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
	log(rw, req)
	cookies := ParseCookies(req)

	var session *Session
	if id := GetSecureCookie(cookies, "mudkip_id"); len(id) > 0 {
		if session = context.GetSession(id); session == nil {
			SetSecureCookie(rw, "mudkip_id", "", "/", -1)
		}
	} else {
		session = context.CreateSession(rw.RemoteAddr())
		SetSecureCookie(rw, "mudkip_id", session.Id, "/", 2592000)
	}
}

func log(rw http.ResponseWriter, req *http.Request) {
	reqtime := time.Nanoseconds() / 1000
	addr := rw.RemoteAddr()

	// Strip port number. Make sure we don't cut off bits of the actual IP if we
	// have an IPv6 address. It will be encased in [ and ]. eg: [::1]:1234
	if idx := strings.LastIndex(addr, ":"); idx > strings.LastIndex(addr, "]") {
		addr = addr[:idx]
	}

	fmt.Fprintf(os.Stdout, "[i] %d %s %s\n", reqtime, addr, req.URL)
}
