package usecases

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
	"web-page-analyser/entities/response_schemas"
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

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("response cannot be parsed as html, err: %v", err)
		return
	}

	htmlString, err := doc.Html()
	if err != nil {
		log.Printf("couldnt extract HTML from document +%v", err.Error())
		return response_schemas.AnalyseUrlResponse{}, err
	}
	response.Version = checkDoctype(htmlString)

	title := doc.Find("title").Text()
	response.Title = title

	for _, val := range headingSlice {
		heading := response_schemas.Header{
			HeadingType: val,
			Count:       findMultipleElementCount(doc, val),
		}
		response.Headers = append(response.Headers, heading)

	}

	inaccessibleLinkCountHttps, externalLinkCountHttps := 0, 0
	inaccessibleLinkCountHttps, externalLinkCountHttps, err = findLinkInfo(doc, "https:")
	if err != nil {
		log.Printf("Could not retrieve external link information")
		return
	}

	inaccessibleLinkCountHttp, externalLinkCountHttp := 0, 0
	inaccessibleLinkCountHttp, externalLinkCountHttp, err = findLinkInfo(doc, "http:")
	if err != nil {
		log.Printf("Could not retrieve external link information")
		return
	}

	inaccessibleLinkCountInternal := 0
	inaccessibleLinkCountInternal, response.InternalLinks, err = findLinkInfo(doc, "#")
	if err != nil {
		log.Printf("Could not retrieve internal link information")
		return
	}

	response.InaccessibleLinkCount = inaccessibleLinkCountHttps + inaccessibleLinkCountHttp +
		inaccessibleLinkCountInternal
	response.ExternalLinks = externalLinkCountHttps + externalLinkCountHttp

	doc.Find("body input[type=password]").Each(func(i int, s *goquery.Selection) {
		response.LoginFormPresent = true
	})
	log.Printf("Success response received %#v", response)

	return
}

func findMultipleElementCount(doc *goquery.Document, selector string) (count int) {
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		count++

	})
	return count
}

func findLinkInfo(doc *goquery.Document, prefix ...string) (inaccessibleLinkCount int, count int, err error) {

	count = 0
	inaccessibleLinkCount = 0
	for _, value := range prefix {
		filterFunc := func(i int, s *goquery.Selection) bool {

			link, ok := s.Attr("href")
			if !ok {
				log.Printf("Could not scrape href attribute")
				return false
			}
			return strings.HasPrefix(link, value)
		}

		doc.Find("body a").FilterFunction(filterFunc).Each(func(i int, tag *goquery.Selection) {
			count++
			link, _ := tag.Attr("href")
			linkText := tag.Text()
			fmt.Printf("%s %s\n", linkText, link)

			resp, err := http.Get(link)
			if err != nil {
				inaccessibleLinkCount++
				log.Printf("Cannot access link %v, error %v", link, err.Error())
				return
			}

			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				inaccessibleLinkCount++
				log.Printf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
				return
			}
		})
	}
	return inaccessibleLinkCount, count, err
}

func checkDoctype(html string) string {

	var doctypes = make(map[string]string)

	doctypes["HTML 4.01"] = `"-//W3C//DTD HTML 4.01//EN"`
	doctypes["XHTML 1.0"] = `"-//W3C//DTD XHTML 1.0 Strict//EN"`
	doctypes["XHTML 1.1"] = `"-//W3C//DTD XHTML 1.1//EN"`
	doctypes["HTML 5"] = `<!DOCTYPE html>`

	var version = "UNKNOWN"

	for doctype, matcher := range doctypes {
		match := strings.Contains(html, matcher)

		if match == true {
			version = doctype
		}
	}

	return version
}
