package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type documentCrawler struct {
	link string
}

func NewDocumentCrawler(link string) Crawler {
	return &documentCrawler{link: link}
}

func (c *documentCrawler) Crawl() (interface{}, error) {
	res, err := http.Get(c.link)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return goquery.NewDocumentFromReader(res.Body)
}
