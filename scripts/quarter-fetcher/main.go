package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/otiai10/opengraph"
)

type Program struct {
	Time string
	Name string
}

type Link struct {
	Title string
	URL   string
}

type Song struct {
	Title   string
	Details map[string]string
}

type OpenGraphImage struct {
	URL string
}

type OpenGraph struct {
	Images      []*OpenGraphImage
	Description string
}

type Episode struct {
	Channel  string
	StartAt  time.Time
	Duration time.Duration
	Number   int
	Subtitle string
}

type Timetable []Episode

type Anime struct {
	Id         string
	Title      string
	Website    string
	OpenGraph  *OpenGraph
	Staff      map[string]string
	Opening    *Song
	Ending     *Song
	Cast       map[string]string
	Programs   []*Program
	Links      []*Link
	Channels   []string
	Timetables []Timetable
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
		var anime Anime

		anime.Id = selection.AttrOr("name", "")
		selector := fmt.Sprintf("li#TID%s > table tr > td", anime.Id)
		link := doc.Find(selector).Eq(0).Find("a").AttrOr("href", "")

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
					URL:   selection.AttrOr("href", ""),
				})
			})
		}

		wg.Add(2)
		go fetchOpenGraph(&wg, &anime)
		go fetchTimetable(&wg, &anime)

		animes = append(animes, &anime)
	})
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

	url := fmt.Sprintf("http://cal.syoboi.jp/tid/%s/time", anime.Id)
	res, err := http.Get(url)
	if err != nil {
		log.Printf("failed to send request: %s\n", err)
		return
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Printf("failed to parse response: %s\n", err)
		return
	}

	table := doc.Find("table#ProgList")
	timetableByChannel := make(map[string]Timetable)
	table.Find("tbody tr").Each(func(i int, tr *goquery.Selection) {
		if i == 0 {
			return
		}
		channel := tr.Find("td.ch").Text()
		timetable := timetableByChannel[channel]

		var episode Episode
		episode.Channel = channel

		tr.Find("td.subtitle").Contents().EachWithBreak(func(i int, s *goquery.Selection) bool {
			if goquery.NodeName(s) == "#text" {
				episode.Subtitle = s.Text()
				return false
			}
			return true
		})

		var year, mon, day, hh, mm, ss, carry int
		var weekDay rune
		fmt.Sscanf(tr.Find("td.start").Text(), "%d-%d-%d(%c) %d:%d:%d ", &year, &mon, &day, &weekDay, &hh, &mm, &ss)
		if hh >= 24 {
			hh = hh - 24
			carry = 1
		}
		parsedTimeString := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d+00:00", year, mon, day, hh, mm, ss)
		if startTime, err := time.Parse(time.RFC3339, parsedTimeString); err == nil {
			episode.StartAt = startTime.Add(time.Duration(carry) * time.Hour * 24)
		}

		if min, err := strconv.Atoi(tr.Find("td.min").Text()); err == nil {
			episode.Duration = time.Duration(min) * time.Minute
		}

		if v, err := strconv.Atoi(tr.Find("td.count").Text()); err == nil {
			episode.Number = v
		}

		timetable = append(timetable, episode)
		timetableByChannel[channel] = timetable
	})

	anime.Channels = make([]string, 0)
	anime.Timetables = make([]Timetable, 0)
	for channel, timetable := range timetableByChannel {
		anime.Channels = append(anime.Channels, channel)
		anime.Timetables = append(anime.Timetables, timetable)
	}
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

	anime.OpenGraph = &OpenGraph{
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
