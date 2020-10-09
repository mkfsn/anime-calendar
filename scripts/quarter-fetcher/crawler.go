package main

type Crawler interface {
	Crawl() (interface{}, error)
}
