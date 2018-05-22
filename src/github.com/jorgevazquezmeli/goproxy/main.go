package main

import (
	"flag"
	"github.com/elazarl/goproxy"
	"net/http"
	"fmt"
	"time"
	"bytes"
	"io/ioutil"
	"log"
	"strings"
)

//Is it an HTML response?
var IsHTML goproxy.RespCondition = goproxy.ContentTypeIs("text/html",
	"text/xml",
	"text/json")

func main() {
	addr := flag.String("addr", ":8080", "proxy listen address")
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	//Get the date and time
	currentTime := time.Now()

	//Logs outbound requests
	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			fmt.Print("\n· ", currentTime.Format("2006.01.02 15:04:05"), " - New Request: ")
			fmt.Print(r.URL)
			r.Header.Set("X-GoProxy", "yxorPoG-X")
			return r, nil
		})


	//Logs responses from the web
	//The IsHTML condition makes it so we only log html responses from the server. Remove it to log everything.
	proxy.OnResponse(IsHTML).DoFunc(
		func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			fmt.Print("\n\n================================================================\n")
			fmt.Print("", currentTime.Format("2006.01.02 15:04:05"), " - Response: ", resp.StatusCode)
			fmt.Print(" - ", ctx.Req.URL)
			fmt.Print("\n================================================================\n")


			//Create a copy of the stream to be able to read it and still return it to the client
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			//Trim the body to 100 characters, remove this next line to see the full body (var b)
			newStr := "\t" + string(b)[:100] + "\n...\n\n"
			newStr = strings.Replace(newStr, "\n", "\n\t", -1)
			fmt.Printf("%s", newStr)

			//Close the reader and create a new one, so that we can return it to the client.
			resp.Body.Close()
			resp.Body = ioutil.NopCloser(bytes.NewReader(b))

			return resp
	})

	http.ListenAndServe(*addr, proxy)
	
}
