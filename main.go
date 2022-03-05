package main

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"cbs_gateway/config"
)

func HealthCheck(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("it's work")
}

func MakeProxyHandler(uri *fasthttp.URI) fasthttp.RequestHandler {
	client := &fasthttp.Client{}
	return func(ctx *fasthttp.RequestCtx) {
		req := &ctx.Request
		req.SetURI(uri)
		err := client.Do(&ctx.Request, &ctx.Response)
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	r := router.New()
	r.GET("/health-check", HealthCheck)

	global, err := config.LoadGlobal()
	if err != nil {
		log.Fatal(err)
	}

	services, err := config.LoadServices()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", services)
	for _, service := range services {
		for _, endpoint := range service.Endpoints {
			uri := &fasthttp.URI{}
			err := uri.Parse(nil, []byte(fmt.Sprintf("%s/%s", service.Host, endpoint.To)))
			if err != nil {
				log.Fatal(err)
			}
			r.Handle(endpoint.Method, endpoint.From, MakeProxyHandler(uri))
		}
	}

	err = fasthttp.ListenAndServe(
		net.JoinHostPort(
			global.Entrypoint.Host,
			strconv.Itoa(global.Entrypoint.Port),
		),
		r.Handler,
	)
	log.Fatal(err)
}
