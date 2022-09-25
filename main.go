package main

import (
	"embed"
	"log"
	"os"
	"os/signal"
	"syscall"
	"web-page-analyser/config-reader"
	"web-page-analyser/handlers/servers"
)

var (
	//go:embed static
	res   embed.FS
	pages = map[string]string{
		"/": "static/index.gohtml",
	}
)

func main() {

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	config_reader.ParseEnvConfig()
	servers.Init(res, pages)

	message := <-signals
	log.Printf("Microservice stopped successfully due to signal '%v'", message.String())

}
