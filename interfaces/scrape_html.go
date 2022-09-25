package interfaces

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type Document struct {
	Doc *goquery.Document
}

type ScrapeHtml interface {
	FindLinkInfo(prefix ...string) (inaccessibleLinkCount int, count int, err error)
	CheckDoctype(html string) string
	FindMultipleElementCount(selector string) (count int)
}

func (d Document) FindLinkInfo(prefix ...string) (inaccessibleLinkCount int, count int, err error) {
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

		d.Doc.Find("body a").FilterFunction(filterFunc).Each(func(i int, tag *goquery.Selection) {
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

func (d Document) CheckDoctype(html string) string {

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

func (d Document) FindMultipleElementCount(selector string) (count int) {
	d.Doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		count++

	})
	return count
}
