package main

import "time"
import "http"
import "fmt"
import "compress/gzip"
import "strings"
import "os"
import "io"

type ServiceContext struct {
	Conn       http.ResponseWriter
	Req        *http.Request
	Params     ParamList
	Cookies    map[string]string
	Session    *Session
	Compressed bool
}

func NewServiceContext(rw http.ResponseWriter, req *http.Request) *ServiceContext {
	sc := new(ServiceContext)
	sc.Conn = rw
	sc.Req = req
	sc.Compressed = false
	sc.Cookies = ParseCookies(req)

	if id := GetSecureCookie(sc.Cookies, "mudkip_id"); len(id) > 0 {
		if sc.Session = context.GetSession(id); sc.Session == nil {
			SetSecureCookie(rw, "mudkip_id", "", "/", -1)
		}
	} else {
		sc.Session = context.CreateSession(rw.RemoteAddr())
		SetSecureCookie(rw, "mudkip_id", sc.Session.Id, "/", 2592000)
	}

	if v, ok := req.Header["Accept-Encoding"]; ok {
		sc.Compressed = strings.Index(v, "gzip") != -1
	}

	return sc
}

func (this *ServiceContext) Status(code int) {
	this.Conn.WriteHeader(code)
}

func (this *ServiceContext) SetHeaders(code, timeout int, ctype string, modified int64) {
	t := time.UTC()
	t = time.SecondsToUTC(t.Seconds() + int64(timeout))
	ts := t.Format(time.RFC1123)

	this.Conn.SetHeader("Cache-Control", fmt.Sprintf("max-age=%d; must-revalidate", timeout))
	this.Conn.SetHeader("Expires", fmt.Sprintf("%s GMT", ts[0:len(ts)-4]))

	if modified > 0 {
		t = time.SecondsToUTC(modified)
	} else {
		t = time.UTC()
	}

	ts = t.Format(time.RFC1123)

	this.Conn.SetHeader("Last-Modified", fmt.Sprintf("%s GMT", ts[0:len(ts)-4]))
	this.Conn.SetHeader("Content-Type", ctype)
	this.Conn.SetHeader("Server", context.Config().ServerName)

	if this.Compressed {
		this.Conn.SetHeader("Content-Encoding", "gzip")
	}

	this.Conn.WriteHeader(code)
}

func (this *ServiceContext) Send(data []byte) bool {
	if this.Compressed {
		var cmp *gzip.Compressor
		var err os.Error

		if cmp, err = gzip.NewWriterLevel(this.Conn, gzip.DefaultCompression); err != nil {
			fmt.Fprintf(os.Stderr, "[e] %v\n", err)
			return false
		}

		cmp.Write(data)
		cmp.Close()
	} else {
		this.Conn.Write(data)
	}

	return true
}

func (this *ServiceContext) SendReadWriter(w io.ReadWriter) bool {
	if this.Compressed {
		var cmp *gzip.Compressor
		var err os.Error

		if cmp, err = gzip.NewWriterLevel(this.Conn, gzip.DefaultCompression); err != nil {
			return false
		}

		io.Copy(cmp, w)
		cmp.Close()
	} else {
		io.Copy(this.Conn, w)
	}

	return true
}
