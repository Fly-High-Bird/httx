package httx

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Context struct {
	Headers map[string]string
	Props   map[string]any
	Cookies []*http.Cookie
	Body    []byte
}

func NewContext() (ctx Context) {
	ctx.Headers = map[string]string{}
	ctx.Props = map[string]any{}
	return ctx
}

func LoadContext() Context {
	ctx := NewContext()
	ctx.ParseStdin()
	json.Unmarshal(ctx.Body, &ctx)
	return ctx
}

func (ctx *Context) ParseStdin() {
	if ctx.Body != nil || len(ctx.Body) != 0 {
		return
	}
	if info, _ := os.Stdin.Stat(); info.Mode()&os.ModeNamedPipe != 0 {
		ctx.Body, _ = ioutil.ReadAll(os.Stdin)
	}
}

func (ctx *Context) Json() (data any, err error) {
	return &data, json.Unmarshal(ctx.Body, &data)
}

func (ctx *Context) FormValue(k string) string {
	body, err := url.ParseQuery(string(ctx.Body))
	if err != nil {
		return ""
	}
	return body[k][0]
}

func (ctx *Context) SetBody(b string) {
	ctx.Body = []byte(b)
}
