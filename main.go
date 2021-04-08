package main

import (
	"encoding/json"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	r := router.New()
	r.GET("/json/hackers", requestHandler)

	log.Fatal(fasthttp.ListenAndServe(":8010", r.Handler))
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "application/json")

	var hackers hackers
	hackers.getHacker()

	json,_ := json.Marshal( hackers )
	str := string(json)

	fmt.Fprintf(ctx, str )
}
