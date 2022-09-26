package test

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"testing"
	"web-page-analyser/services"
)

var data = `
	<!DOCTYPE html>
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
	<a href="https://www.digitalocean.com/community/tutorials">Lession.1</a><br />
	<a href="http://www.digitalooocean.com/community/tutorials">Lession.2</a><br />
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
	  <input type="password" name="Pass" id="Pass" placeholder="Password">
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
var d *goquery.Document
var ctx = context.Background()

func Test_findMultipleElementCount(t *testing.T) {
	tag := "h1"
	d, _ = goquery.NewDocumentFromReader(strings.NewReader(data))
	var doc = services.Document{
		Doc: d,
	}
	expected := 2
	result := doc.FindMultipleElementCount(ctx, tag)

	if result != expected {
		t.Errorf("\"FindMultipleElementCount('%v')\" FAILED, expected -> %v, got -> %v", result, expected, result)
	} else {
		t.Logf("\"FindMultipleElementCount('%v')\" SUCCEDED, expected -> %v, got -> %v", result, expected, result)
	}
}

func Test_findMultipleElementCountH2(t *testing.T) {
	tag := "h2"
	d, _ = goquery.NewDocumentFromReader(strings.NewReader(data))
	var doc = services.Document{
		Doc: d,
	}
	expected := 2
	result := doc.FindMultipleElementCount(ctx, tag)

	if result != expected {
		t.Errorf("\"FindMultipleElementCount('%v')\" FAILED, expected -> %v, got -> %v", result, expected, result)
	} else {
		t.Logf("\"FindMultipleElementCount('%v')\" SUCCEDED, expected -> %v, got -> %v", result, expected, result)
	}
}

func Test_FindLinkInfo(t *testing.T) {
	tag := "https:"
	d, _ := goquery.NewDocumentFromReader(strings.NewReader(data))
	var doc = services.Document{
		Doc: d,
	}
	expectedInaccessibleLinkCount := 0
	inaccessibleLinkCount, _, _ := doc.FindLinkInfo(ctx, tag)

	if expectedInaccessibleLinkCount != inaccessibleLinkCount {
		t.Errorf("\"FindMultipleElementCount('%v')\" FAILED, expected -> %v, got -> %v", inaccessibleLinkCount, expectedInaccessibleLinkCount, inaccessibleLinkCount)
	} else {
		t.Logf("\"FindMultipleElementCount('%v')\" SUCCEDED, expected -> %v, got -> %v", inaccessibleLinkCount, expectedInaccessibleLinkCount, inaccessibleLinkCount)
	}
}

func Test_FindLinkInfoLinkCount(t *testing.T) {
	tag := "https:"
	d, _ := goquery.NewDocumentFromReader(strings.NewReader(data))
	var doc = services.Document{
		Doc: d,
	}
	expectedLinkCount := 1
	_, linkCount, _ := doc.FindLinkInfo(ctx, tag)

	if expectedLinkCount != linkCount {
		t.Errorf("\"FindMultipleElementCount('%v')\" FAILED, expected -> %v, got -> %v", linkCount,
			expectedLinkCount, linkCount)
	} else {
		t.Logf("\"FindMultipleElementCount('%v')\" SUCCEDED, expected -> %v, got -> %v", linkCount,
			expectedLinkCount, linkCount)
	}
}

func Test_FindLinkInfoHttp(t *testing.T) {
	tag := "http:"
	d, _ := goquery.NewDocumentFromReader(strings.NewReader(data))
	var doc = services.Document{
		Doc: d,
	}
	expectedInaccessibleLinkCount := 1
	inaccessibleLinkCount, _, _ := doc.FindLinkInfo(ctx, tag)

	if expectedInaccessibleLinkCount != inaccessibleLinkCount {
		t.Errorf("\"FindMultipleElementCount('%v')\" FAILED, expected -> %v, got -> %v", inaccessibleLinkCount, expectedInaccessibleLinkCount, inaccessibleLinkCount)
	} else {
		t.Logf("\"FindMultipleElementCount('%v')\" SUCCEDED, expected -> %v, got -> %v", inaccessibleLinkCount, expectedInaccessibleLinkCount, inaccessibleLinkCount)
	}
}

func Test_FindLinkInfoLinkCountHttp(t *testing.T) {
	tag := "http:"
	d, _ := goquery.NewDocumentFromReader(strings.NewReader(data))
	var doc = services.Document{
		Doc: d,
	}
	expectedLinkCount := 1
	_, linkCount, _ := doc.FindLinkInfo(ctx, tag)

	if expectedLinkCount != linkCount {
		t.Errorf("\"FindMultipleElementCount('%v')\" FAILED, expected -> %v, got -> %v", linkCount,
			expectedLinkCount, linkCount)
	} else {
		t.Logf("\"FindMultipleElementCount('%v')\" SUCCEDED, expected -> %v, got -> %v", linkCount,
			expectedLinkCount, linkCount)
	}
}

func Test_CheckDoctype(t *testing.T) {
	d, _ := goquery.NewDocumentFromReader(strings.NewReader(data))
	var doc = services.Document{
		Doc: d,
	}
	htmlStr, _ := doc.Doc.Html()
	expectedDocType := "HTML 5"
	docType := doc.CheckDoctype(ctx, htmlStr)

	if expectedDocType != docType {
		t.Errorf("\"FindMultipleElementCount('%v')\" FAILED, expected -> %v, got -> %v", docType,
			expectedDocType, docType)
	} else {
		t.Logf("\"FindMultipleElementCount('%v')\" SUCCEDED, expected -> %v, got -> %v", docType,
			expectedDocType, docType)
	}
}

func Test_findLogins(t *testing.T) {
	tag := "body input[type=password]"
	d, _ = goquery.NewDocumentFromReader(strings.NewReader(data))
	var doc = services.Document{
		Doc: d,
	}
	expected := 1
	result := doc.FindMultipleElementCount(ctx, tag)

	if result != expected {
		t.Errorf("\"FindMultipleElementCount('%v')\" FAILED, expected -> %v, got -> %v", result, expected, result)
	} else {
		t.Logf("\"FindMultipleElementCount('%v')\" SUCCEDED, expected -> %v, got -> %v", result, expected, result)
	}
}
