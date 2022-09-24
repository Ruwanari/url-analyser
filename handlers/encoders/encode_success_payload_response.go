package encoders

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"path"
	"web-page-analyser/entities/response_schemas"
)

//EncodeSuccessPayloadResponse encodes and renders success responses to success html template.
func EncodeSuccessPayloadResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	analyseUrlResponse := response.(response_schemas.AnalyseUrlResponse)

	fp := path.Join("static", "responseTemplate.gohtml")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		if err != nil {
			log.Printf("Error parsing template " + err.Error())
			return err
		}
	}

	err = tmpl.Execute(w, analyseUrlResponse)
	if err != nil {
		log.Printf("Error executing template " + err.Error())
		return err
	}
	return err
}
