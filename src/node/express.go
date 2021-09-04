package express

import (
	"log"

	"github.com/valyala/fasthttp"
)

var proxyClient = &fasthttp.HostClient{
	Addr: "upstream.host:port",
}

func ReverseProxyHandler(ctx *fasthttp.RequestCtx) {
	req := &ctx.Request
	resp := &ctx.Response
	prepareRequest(req)
	if err := proxyClient.Do(req, resp); err != nil {
		ctx.Logger().Printf("error when proxying the request: %s", err)
	}
	postprocessResponse(resp)
}

func prepareRequest(req *fasthttp.Request) {
	req.Header.Del("Connection")
}

func postprocessResponse(resp *fasthttp.Response) {
	resp.Header.Del("Connection")
	resp.Header.Set("X-Powered-By", "ParakeetCloud")
}

func main() {
	if err := fasthttp.ListenAndServe(":8080", ReverseProxyHandler); err != nil {
		log.Fatalf("error in fasthttp server: %s", err)
	}
}
