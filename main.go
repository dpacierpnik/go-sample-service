package main

import (
	"flag"
	"fmt"
	"github.com/dpacierpnik/go-sample-service/html"
	"github.com/dpacierpnik/go-sample-service/rest/echo"
	"github.com/dpacierpnik/go-sample-service/rest/github"
	"log"
	"net/http"
)

func main() {

	port := flag.Int("port", 8080, "Port for listening")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/github/zen", github.Zen)
	mux.HandleFunc("/echo/body", echo.Body)
	mux.HandleFunc("/echo/headers", echo.Headers)
	mux.HandleFunc("/", html.Index)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
