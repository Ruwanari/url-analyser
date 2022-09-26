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
func EncodeSuccessPayloadResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	analyseUrlResponse := response.(response_schemas.AnalyseUrlResponse)

	fp := path.Join("static", "success_response_template.gohtml")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		if err != nil {
			log.Printf("Error parsing template ctx: %v, err: %v", ctx, err.Error())
			return err
		}
	}

	err = tmpl.Execute(w, analyseUrlResponse)
	if err != nil {
		log.Printf("Error executing template ctx: %v, err: %v ", ctx, err.Error())
		return err
	}
	return err
}
