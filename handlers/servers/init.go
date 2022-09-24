package servers

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"web-page-analyser/config"
	"web-page-analyser/handlers"
)

func Init(res embed.FS, pages map[string]string) {

	go func() {
		log.Print("HTTP server starting on ", config.EnvConf.BackendPort)
		err := http.ListenAndServe(":"+config.EnvConf.BackendPort, handlers.GetRoutes())
		if err != nil {
			log.Fatal("HTTP server running failed.", err)
		}
	}()

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			page, ok := pages[r.URL.Path]
			if !ok {
				//log.Printf()
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
				//log
				return
			}
		})
		http.FileServer(http.FS(res))
		log.Println("server started on port " + config.EnvConf.FrontendPort)
		err := http.ListenAndServe(":"+config.EnvConf.FrontendPort, nil)
		if err != nil {
			log.Fatal("HTTP server running failed.", err)
		}
	}()

}
