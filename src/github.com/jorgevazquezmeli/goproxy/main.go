package main

import (
	"flag"
	"github.com/elazarl/goproxy"
	"net/http"
	"fmt"

)

func main() {
	addr := flag.String("addr", ":8080", "proxy listen address")
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			fmt.Printf("%s\n\n\n", r.Header)
			r.Header.Set("X-GoProxy", "yxorPoG-X")
			return r, nil
		})

	proxy.OnResponse().DoFunc(
		func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			fmt.Printf("%s\n", resp.Header)
			return resp
	})

	http.ListenAndServe(*addr, proxy)
	
}
