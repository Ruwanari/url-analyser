package servers

import (
	"log"
	"net/http"
	"strconv"
	"web-page-analyser/handlers"
)

func Init() {

	go func() {
		log.Print("HTTP server starting on ", 8080)
		err := http.ListenAndServe(":"+strconv.Itoa(8080), handlers.GetRoutes())
		if err != nil {
			log.Fatal("HTTP server running failed.", err)
		}
	}()

	go func() {
		log.Print("HTTP server starting on ", 8081)
		http.Handle("/", http.FileServer(http.Dir("./static")))
		err := http.ListenAndServe(":"+strconv.Itoa(8081), nil)
		if err != nil {
			log.Fatal("HTTP server running failed.", err)
		}
	}()

}
