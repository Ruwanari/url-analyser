package services

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

/*FindLinkInfo finds the links available in HTML page and returns inaccessible link count, external/internal link count.
Prefix (https/http/#) should be sent as a parameter.*/
func (d Document) FindLinkInfo(ctx context.Context, prefix ...string) (inaccessibleLinkCount int, count int, err error) {
	count = 0
	inaccessibleLinkCount = 0
	for _, value := range prefix {
		filterFunc := func(i int, s *goquery.Selection) bool {

			link, ok := s.Attr("href")
			if !ok {
				log.Printf("Could not scrape href attribute ctx: %v", ctx)
				return false
			}
			return strings.HasPrefix(link, value)
		}

		d.Doc.Find("body a").FilterFunction(filterFunc).Each(func(i int, tag *goquery.Selection) {
			count++
			link, ok := tag.Attr("href")
			if !ok {
				log.Printf("Cannot access tag %v, error %v ctx : %v", link, ctx)
				return
			}
			linkText := tag.Text()
			fmt.Printf("%s %s\n", linkText, link)

			resp, err := http.Get(link)
			if err != nil {
				inaccessibleLinkCount++
				log.Printf("Cannot access link %v, error %v ctx : %v", link, err.Error(), ctx)
				return
			}

			defer resp.Body.Close()

			if (resp.StatusCode < http.StatusOK) || (resp.StatusCode >= http.StatusMultipleChoices) {
				inaccessibleLinkCount++
				log.Printf("failed to fetch data %d %s ctx : %v", resp.StatusCode, resp.Status, ctx)
				return
			}
		})
	}
	return inaccessibleLinkCount, count, err
}

//CheckDoctype returns the HTML version of a web page
func (d Document) CheckDoctype(ctx context.Context, html string) string {

	var docTypes = make(map[string]string)

	docTypes["HTML 4.01"] = `"-//W3C//DTD HTML 4.01//EN"`
	docTypes["XHTML 1.0"] = `"-//W3C//DTD XHTML 1.0 Strict//EN"`
	docTypes["XHTML 1.1"] = `"-//W3C//DTD XHTML 1.1//EN"`
	docTypes["HTML 5"] = `<!DOCTYPE html>`

	var version = "UNKNOWN"

	for docType, matcher := range docTypes {
		match := strings.Contains(html, matcher)

		if match {
			version = docType
			break
		}
	}

	return version
}

//FindMultipleElementCount finds the number of elements in an HTML page when the selector is sent as a parameter.
func (d Document) FindMultipleElementCount(ctx context.Context, selector string) (count int) {
	d.Doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		count++

	})
	return count
}
