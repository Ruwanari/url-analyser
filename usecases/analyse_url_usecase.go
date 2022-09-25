package usecases

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"web-page-analyser/entities/response_schemas"
	"web-page-analyser/interfaces"
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
func AnalyseUrlUsecase(url string) (response response_schemas.AnalyseUrlResponse, err error) {
	headingSlice := []string{"h1", "h2", "h3", "h4", "h5", "h6"}
	document := interfaces.Document{}
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("failed to make a request to url" + err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
		return response, errors.New("error response when contacting url")
	}

	document.Doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("response cannot be parsed as html, err: %v", err)
		return
	}

	htmlString, err := document.Doc.Html()
	if err != nil {
		log.Printf("couldnt extract HTML from document +%v", err.Error())
		return response_schemas.AnalyseUrlResponse{}, err
	}
	response.Version = document.CheckDoctype(htmlString)

	title := document.Doc.Find("title").Text()
	response.Title = title

	for _, val := range headingSlice {
		heading := response_schemas.Header{
			HeadingType: val,
			Count:       document.FindMultipleElementCount(val),
		}
		response.Headers = append(response.Headers, heading)

	}

	inaccessibleLinkCountHttps, externalLinkCountHttps := 0, 0
	inaccessibleLinkCountHttps, externalLinkCountHttps, err = document.FindLinkInfo("https:")
	if err != nil {
		log.Printf("Could not retrieve external link information")
		return
	}

	inaccessibleLinkCountHttp, externalLinkCountHttp := 0, 0
	inaccessibleLinkCountHttp, externalLinkCountHttp, err = document.FindLinkInfo("http:")
	if err != nil {
		log.Printf("Could not retrieve external link information")
		return
	}

	inaccessibleLinkCountInternal := 0
	inaccessibleLinkCountInternal, response.InternalLinks, err = document.FindLinkInfo("#")
	if err != nil {
		log.Printf("Could not retrieve internal link information")
		return
	}

	response.InaccessibleLinkCount = inaccessibleLinkCountHttps + inaccessibleLinkCountHttp +
		inaccessibleLinkCountInternal
	response.ExternalLinks = externalLinkCountHttps + externalLinkCountHttp

	document.Doc.Find("body input[type=password]").Each(func(i int, s *goquery.Selection) {
		response.LoginFormPresent = true
	})
	log.Printf("Success response received %#v", response)

	return
}
