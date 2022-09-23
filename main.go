package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"web-page-analyser/handlers/servers"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	servers.Init()

	message := <-signals
	log.Printf(fmt.Sprintf("Microservice stopped successfully due to signal '%v'", message.String()))

}
