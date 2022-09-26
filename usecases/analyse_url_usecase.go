package usecases

import (
	"context"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"web-page-analyser/entities/response_schemas"
	"web-page-analyser/services"
)

/*AnalyseUrlUsecase scrapes the HTML content of the web page using Goquery and returns
following information.
- Title of the web page
- HTML version of the web page
- Headings present on the web page and their respective counts
- External link count on the web page
- Internal link count on the web page
- Inaccessible link count on the web page
- If the web page contains a login or not
*/
func AnalyseUrlUsecase(ctx context.Context, url string) (response response_schemas.AnalyseUrlResponse, err error) {
	headingSlice := []string{"h1", "h2", "h3", "h4", "h5", "h6"}
	document := services.Document{}
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("failed to make a request to url ctx : %v, error : %v", ctx, err.Error())
		return
	}
	defer resp.Body.Close()

	if (resp.StatusCode < http.StatusOK) || (resp.StatusCode >= http.StatusMultipleChoices) {
		log.Printf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
		return response, errors.New("error response when contacting url")
	}

	document.Doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("response cannot be parsed as html, ctx : %v, error : %v", ctx, err)
		return
	}

	htmlString, err := document.Doc.Html()
	if err != nil {
		log.Printf("couldnt extract HTML from document ctx : %v, error : %v", ctx, err.Error())
		return response_schemas.AnalyseUrlResponse{}, err
	}
	response.Version = document.CheckDoctype(ctx, htmlString)

	title := document.Doc.Find("title").Text()
	response.Title = title

	for _, val := range headingSlice {
		heading := response_schemas.Header{
			HeadingType: val,
			Count:       document.FindMultipleElementCount(ctx, val),
		}
		response.Headers = append(response.Headers, heading)

	}

	inaccessibleLinkCountHttps, externalLinkCountHttps := 0, 0
	inaccessibleLinkCountHttps, externalLinkCountHttps, err = document.FindLinkInfo(ctx, "https:")
	if err != nil {
		log.Printf("Could not retrieve external link information ctx : %v, error : %v", ctx, err.Error())
		return
	}

	inaccessibleLinkCountHttp, externalLinkCountHttp := 0, 0
	inaccessibleLinkCountHttp, externalLinkCountHttp, err = document.FindLinkInfo(ctx, "http:")
	if err != nil {
		log.Printf("Could not retrieve external link information ctx : %v, error : %v", ctx, err.Error())
		return
	}

	inaccessibleLinkCountInternal := 0
	inaccessibleLinkCountInternal, response.InternalLinks, err = document.FindLinkInfo(ctx, "#")
	if err != nil {
		log.Printf("Could not retrieve internal link information ctx : %v, error : %v", ctx, err.Error())
		return
	}

	response.InaccessibleLinkCount = inaccessibleLinkCountHttps + inaccessibleLinkCountHttp +
		inaccessibleLinkCountInternal
	response.ExternalLinks = externalLinkCountHttps + externalLinkCountHttp

	passwords := document.FindMultipleElementCount(ctx, "body input[type=password]")
	if passwords > 0 {
		response.LoginFormPresent = true
	}
	log.Printf("Success response received %#v, ctx : %v", response, ctx)

	return
}
