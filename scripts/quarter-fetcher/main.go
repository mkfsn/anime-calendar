package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := NewDocumentCrawler("http://cal.syoboi.jp/quarter/2020q4").Crawl()
	if err != nil {
		log.Fatalln(err)
	}
	doc := res.(*goquery.Document)

	var animes []*Anime
	if err := NewAnimeParse(doc, &animes).Parse(); err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup
	for _, anime := range animes {
		wg.Add(2)
		go fetchOpenGraph(&wg, anime)
		go fetchTimetable(&wg, anime)
	}
	wg.Wait()

	b, err := json.MarshalIndent(animes, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	if err := ioutil.WriteFile("../../public/data/anime.json", b, 0644); err != nil {
		log.Fatalln(err)
	}
}

func fetchTimetable(wg *sync.WaitGroup, anime *Anime) {
	defer wg.Done()
	res, err := NewDocumentCrawler(fmt.Sprintf("http://cal.syoboi.jp/tid/%s/time", anime.Id)).Crawl()
	if err != nil {
		log.Printf("failed to parse response: %s\n", err)
		return
	}
	doc := res.(*goquery.Document)
	if err := NewTimetableParse(doc.Find("table#ProgList"), anime).Parse(); err != nil {
		log.Printf("failed to parse response: %s\n", err)
		return
	}
}

func fetchOpenGraph(wg *sync.WaitGroup, anime *Anime) {
	defer wg.Done()
	if anime.Website == "" {
		return
	}
	if _, err := NewOpenGraphCrawler(anime.Website, anime).Crawl(); err != nil {
		log.Printf("error: failed to get og of %s: %s\n", anime.Title, err)
	}
}
