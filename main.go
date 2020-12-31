package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/blackcrw/wprecon-api/pkg/config"
	"github.com/blackcrw/wprecon-api/pkg/routers"
)

func main() {
	config.Loading()
	routerNew := routers.Generate()

	fmt.Printf("Listing in port %d\n", config.APIPORT)
	fmt.Printf("URL: http://127.0.0.1:%d\n", config.APIPORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.APIPORT), routerNew))
}
