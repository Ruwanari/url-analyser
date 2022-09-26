package services

import (
	"context"
	"github.com/PuerkitoBio/goquery"
)

type Document struct {
	Doc *goquery.Document
}

//ScrapeHtml contains functions to retrieve data in HTML pages.
type ScrapeHtml interface {
	FindLinkInfo(ctx context.Context, prefix ...string) (inaccessibleLinkCount int, count int, err error)
	CheckDoctype(ctx context.Context, html string) string
	FindMultipleElementCount(ctx context.Context, selector string) (count int)
	CheckAccessibility(ctx context.Context, links []string) (inaccessibleLinkCount int, err error)
}
