package usecases

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
	"web-page-analyser/response_schemas"
)

func AnalyseUrlUsecase(url string) (response response_schemas.AnalyseUrlResponse, err error) {
	headingSlice := []string{"h1", "h2", "h3", "h4", "h5", "h6"}

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("failed to make a request to url" + err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
		return
	}
	response.Protocol = resp.Request.Proto

	hasInAccessibleLink := false

	data := `
	<html lang="en">
	<body>
	<p>List of words</p>
	<h1>dark</h1>
	<h2>smart</h2>
	<h3>war</h3>
		<h3>war</h3>
		<h1>dark</h1>
	<h2>smart</h2>
	<h3>war</h3>
		<h3>war</h3>
		<h1>dark</h1>
	<h2>smart</h2>
	<h3>war</h3>
		<h3>war</h3>
	<a href="#lession1">Lession.1</a><br />
	<a href="#lession2">Lession.2</a><br />
	<a href="#lession3">Lession.3</a><br />
	<a href="#lession4">Lession.4</a><br />
	<form id="login" method="get" action="login.php">
	   <label><b>User Name
	   </b>
	   </label>
	   <input type="text" name="Uname" id="Uname" placeholder="Username">
	   <br><br>
	   <label><b>Password
	   </b>
	   </label>
	   <input type="Password" name="Pass" id="Pass" placeholder="Password">
	   <br><br>
	   <input type="button" name="log" id="log" value="Log In Here">
	   <br><br>
	   <input type="checkbox" id="check">
	   <span>Remember me</span>
	   <br><br>
	   Forgot <a href="#">Password</a>
	</form>
	<footer>footer for words</footer>
	</body>
	</html>
	`
	//strings.NewReader(data)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		log.Printf("response cannot be parsed as html, err: %v", err)
		return
	}

	title := doc.Find("title").Text()
	response.Title = title

	for _, val := range headingSlice {
		switch val {
		case "h1":
			response.Headers.H1 = findMultipleElementCount(doc, val)

		case "h2":
			response.Headers.H2 = findMultipleElementCount(doc, val)

		case "h3":
			response.Headers.H3 = findMultipleElementCount(doc, val)

		case "h4":
			response.Headers.H4 = findMultipleElementCount(doc, val)

		case "h5":
			response.Headers.H5 = findMultipleElementCount(doc, val)

		case "h6":
			response.Headers.H6 = findMultipleElementCount(doc, val)
		}

	}

	//doc.Find("h1").Each(func(i int, s *goquery.Selection) {
	//	// For each item found, get the name.
	//	if i == 0 {
	//		response.Headers.H1 = 1
	//	} else {
	//		response.Headers.H1 = response.Headers.H1 + 1
	//	}
	//
	//})
	//
	//doc.Find("h2").Each(func(i int, s *goquery.Selection) {
	//	// For each item found, get the name.
	//	if i == 0 {
	//		response.Headers.H2 = 1
	//	} else {
	//		response.Headers.H2 = response.Headers.H2 + 1
	//	}
	//
	//})
	//
	//doc.Find("h3").Each(func(i int, s *goquery.Selection) {
	//	// For each item found, get the name.
	//	if i == 0 {
	//		response.Headers.H3 = 1
	//	} else {
	//		response.Headers.H3 = response.Headers.H3 + 1
	//	}
	//
	//})
	//
	//doc.Find("h4").Each(func(i int, s *goquery.Selection) {
	//	if i == 0 {
	//		response.Headers.H4 = 1
	//	} else {
	//		response.Headers.H4 = response.Headers.H4 + 1
	//	}
	//
	//})
	//
	//doc.Find("h5").Each(func(i int, s *goquery.Selection) {
	//	if i == 0 {
	//		response.Headers.H5 = response.Headers.H5 + i + 1
	//	} else {
	//		response.Headers.H5 = response.Headers.H5 + i
	//	}
	//
	//})
	//
	//doc.Find("h6").Each(func(i int, s *goquery.Selection) {
	//	if i == 0 {
	//		response.Headers.H6 = 1
	//	} else {
	//		response.Headers.H6 = response.Headers.H6 + 1
	//	}
	//
	//})

	f := func(i int, s *goquery.Selection) bool {

		link, _ := s.Attr("href")
		return strings.HasPrefix(link, "https")
	}

	doc.Find("body a").FilterFunction(f).Each(func(i int, tag *goquery.Selection) {

		link, _ := tag.Attr("href")
		linkText := tag.Text()
		fmt.Printf("%s %s\n", linkText, link)

		resp2, err := http.Get(link)
		if err != nil {
			hasInAccessibleLink = true
			log.Printf(err.Error())
			return
		}

		defer resp2.Body.Close()

		if resp2.StatusCode != 200 {
			hasInAccessibleLink = true
			log.Printf("failed to fetch data: %d %s", resp2.StatusCode, resp2.Status)
			return
		}
		if i == 0 {
			response.ExternalLinks = 1
		} else {
			response.ExternalLinks = response.ExternalLinks + 1
		}
	})

	f = func(i int, s *goquery.Selection) bool {

		link, _ := s.Attr("href")
		return strings.HasPrefix(link, "#")
	}

	doc.Find("body a").FilterFunction(f).Each(func(i int, tag *goquery.Selection) {

		link, _ := tag.Attr("href")
		linkText := tag.Text()
		fmt.Printf("%s %s\n", linkText, link)

		if i == 0 {
			response.InternalLinks = 1
		} else {
			response.InternalLinks = response.InternalLinks + 1
		}
		resp3, err := http.Get(url + "/" + link)
		if err != nil {
			hasInAccessibleLink = true
			log.Printf(err.Error())
			return
		}

		defer resp3.Body.Close()

		if resp3.StatusCode != 200 {
			hasInAccessibleLink = true
			log.Printf("failed to fetch data: %d %s", resp3.StatusCode, resp3.Status)
			return
		}
	})

	doc.Find("div input[type=password]").Each(func(i int, s *goquery.Selection) {
		response.LoginFormPresent = true

	})

	inp := doc.Find("input")
	v := inp.AttrOr("type", "")
	if v != "" {
		response.LoginFormPresent = true
	}

	response.InaccessibleLinks = hasInAccessibleLink

	return
}

func findMultipleElementCount(doc *goquery.Document, selector string) (count int) {
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			count = 1
		} else {
			count = count + 1
		}

	})
	return count
}
