package main

import (
	"flag"
	"os"

	"github.com/almostmoore/kadastr_tg/server"
	"github.com/almostmoore/kadastr/api_server"
)

var tgToken string
var apiAddr string

func main() {
	flag.StringVar(&tgToken, "tgtoken", os.Getenv("TG_TOKEN"), "")
	flag.StringVar(&apiAddr, "addr", os.Getenv("ADDR"), "Listen address")
	flag.Parse()

	tg := &server.Server{
		APIToken: tgToken,
		ApiClient: api_server.NewClient("http://" + apiAddr),
	}

	tg.Run()
}