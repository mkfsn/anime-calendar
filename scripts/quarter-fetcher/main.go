package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/otiai10/opengraph"
)

type Program struct {
	Time string
	Name string
}

type Link struct {
	Title string
	URL string
}

type Song struct {
	Title string
	Details map[string]string
}

type OpenGraphImage struct {
	URL string
}

type OpenGraph struct {
	Images []*OpenGraphImage
	Description string
}

type Anime struct {
	Title string
	Website string
	OpenGraph *OpenGraph
	Staff map[string]string
	Opening *Song
	Ending *Song
	Cast map[string]string
	Programs []*Program
	Links []*Link
}

func main() {
	url := "http://cal.syoboi.jp/quarter/2020q4"
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup
	animes := make([]*Anime, 0)

	doc.Find("ul.titlesDetail > a").Each(func(i int, selection *goquery.Selection) {
		name := selection.AttrOr("name", "")
		selector := fmt.Sprintf("li#TID%s > table tr > td", name)
		link := doc.Find(selector).Eq(0).Find("a").AttrOr("href", "")
		var anime Anime
		if strings.HasPrefix(link, "http") {
			anime.Website = link
		}

		data := doc.Find(selector).Eq(1)

		// Title
		anime.Title = data.Find("a.title").Text()

		// Staff
		if dt := data.Find(".staff dt"); dt.Length() != 0 {
			anime.Staff = make(map[string]string)
			dt.Each(func(i int, selection *goquery.Selection) {
				anime.Staff[selection.Text()] = selection.NextFiltered("dd").Text()
			})
		}

		// Opening
		if op := data.Find(".op"); op.Length() != 0 {
			var song Song
			song.Title = op.Find(".trackname").Text()
			song.Details = make(map[string]string)
			op.Find("dt").Each(func(i int, selection *goquery.Selection) {
				song.Details[selection.Text()] = selection.NextFiltered("dd").Text()
			})
			anime.Opening = &song
		}

		// Ending
		if ed := data.Find(".ed"); ed.Length() != 0 {
			var song Song
			song.Title = ed.Find(".trackname").Text()
			song.Details = make(map[string]string)
			ed.Find("dt").Each(func(i int, selection *goquery.Selection) {
				song.Details[selection.Text()] = selection.NextFiltered("dd").Text()
			})
			anime.Ending = &song
		}

		// Cast
		if dt := data.Find(".cast dt"); dt.Length() != 0 {
			anime.Cast = make(map[string]string)
			dt.Each(func(i int, selection *goquery.Selection) {
				anime.Cast[selection.Text()] = selection.NextFiltered("dd").Text()
			})
		}

		// Programs
		if li := data.Find(".progs > ul > li"); li.Length() != 0 {
			anime.Programs = make([]*Program, 0)
			li.Each(func(i int, selection *goquery.Selection) {
				anime.Programs = append(anime.Programs, &Program{
					Time: selection.Find(".StTime").Text(),
					Name: selection.Find(".ChName").Text(),
				})
			})
		}

		// Links
		if a := data.Find(".links > ul > li > a"); a.Length() != 0 {
			anime.Links = make([]*Link, 0)
			a.Each(func(i int, selection *goquery.Selection) {
				anime.Links = append(anime.Links, &Link{
					Title: selection.Text(),
					URL: selection.AttrOr("href", ""),
				})
			})
		}

		wg.Add(1)
		go fetchOpenGraph(&wg, &anime)

		animes = append(animes, &anime)
	})
	wg.Wait()

	b, err := json.MarshalIndent(animes, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", b)
}

func fetchOpenGraph(wg *sync.WaitGroup, anime *Anime) {
	defer wg.Done()
	if anime.Website == "" {
		return
	}
	og, err := opengraph.Fetch(anime.Website)
	if err != nil {
		log.Printf("error: failed to get og of %s: %s\n", anime.Title, err)
		return
	}

	anime.OpenGraph = &OpenGraph {
		Description: og.Description,
	}

	if og.Image != nil && len(og.Image) != 0 {
		anime.OpenGraph.Images = make([]*OpenGraphImage, 0)
		for _, image := range og.Image {
			anime.OpenGraph.Images = append(anime.OpenGraph.Images, &OpenGraphImage{
				URL: image.URL,
			})
		}
	}
}
