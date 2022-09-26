package servers

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"web-page-analyser/config-reader"
	"web-page-analyser/handlers"
)

func Init(res embed.FS, pages map[string]string) {

	go func() {
		log.Print("HTTP server starting on ", config_reader.EnvConf.BackendPort)
		err := http.ListenAndServe(":"+config_reader.EnvConf.BackendPort, handlers.GetRoutes())
		if err != nil {
			log.Fatal("HTTP server running failed.", err)
		}
	}()

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			page, ok := pages[r.URL.Path]
			if !ok {
				log.Printf("page is not found")
				w.WriteHeader(http.StatusNotFound)
				return
			}
			tpl, err := template.ParseFS(res, page)
			if err != nil {
				log.Printf("page %s not found in pages cache...", r.RequestURI)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			data := map[string]interface{}{
				"userAgent": r.UserAgent(),
			}
			if err := tpl.Execute(w, data); err != nil {
				log.Printf("Error executing template. %v ", err)
				return
			}
		})
		http.FileServer(http.FS(res))
		log.Println("server started on port " + config_reader.EnvConf.FrontendPort)
		err := http.ListenAndServe(":"+config_reader.EnvConf.FrontendPort, nil)
		if err != nil {
			log.Fatal("HTTP server running failed.", err)
		}
	}()

}
