package main

import (
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"cbs_gateway/config"
)

func HealthCheck(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("it's work")
}

func main() {
	r := router.New()
	r.GET("/health-check", HealthCheck)

	_, err := config.LoadGlobal()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
