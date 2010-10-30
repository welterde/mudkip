package main

import "crypto/hmac"
import "encoding/base64"
import "time"
import "http"
import "bytes"
import "strings"
import "fmt"
import "strconv"
import "io/ioutil"

func ParseCookies(req *http.Request) (jar map[string]string) {
	jar = make(map[string]string)

	if v, ok := req.Header["Cookie"]; ok {
		cookies := strings.Split(v, ";", -1)
		var seg []string

		for _, cookie := range cookies {
			cookie = strings.TrimSpace(cookie)

			if seg = strings.Split(cookie, "=", 2); len(seg) < 2 {
				continue
			}

			jar[seg[0]] = seg[1]
		}
	}

	return
}

func SetSecureCookie(rw http.ResponseWriter, name, val, path string, timeout int64) {
	var buf bytes.Buffer

	e := base64.NewEncoder(base64.StdEncoding, &buf)
	e.Write([]byte(val))
	e.Close()

	ts := strconv.Itoa64(time.Seconds())
	data := strings.Join([]string{buf.String(), ts, getCookieSig(buf.Bytes(), ts)}, "|")

	var cookie string

	// Timeout of -1 is a special case that omits the entire 'expires' bit.
	// This is used for cookies which expire as soon as a user's session ends.
	if timeout != -1 {
		t := time.UTC()
		t = time.SecondsToUTC(t.Seconds() + timeout)
		ts = t.Format(time.RFC1123)
		ts = ts[0:len(ts)-3] + "GMT"
		cookie = fmt.Sprintf("%s=%s; expires=%s; path=%s", name, data, ts, path)
	} else {
		cookie = fmt.Sprintf("%s=%s; path=%s", name, data, path)
	}

	if context.Config().Secure {
		cookie += "; secure"
	}

	rw.SetHeader("Set-Cookie", cookie)
}

func GetSecureCookie(jar map[string]string, name string) string {
	var cookie string
	var ok bool

	if cookie, ok = jar[name]; !ok {
		return ""
	}

	parts := strings.Split(cookie, "|", 3)

	if getCookieSig([]byte(parts[0]), parts[1]) != parts[2] {
		return ""
	}

	ts, _ := strconv.Atoi64(parts[1])
	if time.Seconds()-2678400 > ts {
		return ""
	}

	buf := bytes.NewBufferString(parts[0])
	e := base64.NewDecoder(base64.StdEncoding, buf)
	d, _ := ioutil.ReadAll(e)
	return string(d)
}

func getCookieSig(val []byte, timestamp string) string {
	hm := hmac.NewSHA1([]byte(context.Config().CookieSalt))
	hm.Write(val)
	hm.Write([]byte(timestamp))
	return fmt.Sprintf("%02x", hm.Sum())
}
