package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type animeParser struct {
	doc       *goquery.Document
	animeList *[]*Anime
}

func NewAnimeParse(doc *goquery.Document, animeList *[]*Anime) Parser {
	return &animeParser{doc: doc, animeList: animeList}
}

func (p *animeParser) Parse() error {
	*p.animeList = make([]*Anime, 0)
	p.doc.Find("ul.titlesDetail > a").Each(func(i int, selection *goquery.Selection) {
		var anime Anime
		p.parseAnime(&anime, selection)
		*p.animeList = append(*p.animeList, &anime)
	})
	return nil
}

func (p *animeParser) parseAnime(anime *Anime, selection *goquery.Selection) {
	p.parseAnimeId(anime, selection)
	selector := fmt.Sprintf("li#TID%s > table tr > td", anime.Id)
	col1 := p.doc.Find(selector).Eq(0)
	col2 := p.doc.Find(selector).Eq(1)
	p.parseAnimeWebsite(anime, col1)
	p.parseAnimeTitle(anime, col2)
	p.parseAnimeStaff(anime, col2)
	p.parseAnimeOpening(anime, col2)
	p.parseAnimeEnding(anime, col2)
	p.parseAnimeCast(anime, col2)
	p.parseAnimePrograms(anime, col2)
	p.parseAnimeLinks(anime, col2)
}

func (p *animeParser) parseAnimeId(anime *Anime, selection *goquery.Selection) {
	anime.Id = selection.AttrOr("name", "")
}

func (p *animeParser) parseAnimeWebsite(anime *Anime, selection *goquery.Selection) {
	link := selection.Find("a").AttrOr("href", "")
	if !strings.HasPrefix(link, "http") {
		// TODO: logger
		return
	}
	anime.Website = link
}

func (p *animeParser) parseAnimeTitle(anime *Anime, selection *goquery.Selection) {
	anime.Title = selection.Find("a.title").Text()
}

func (p *animeParser) parseAnimeStaff(anime *Anime, selection *goquery.Selection) {
	dt := selection.Find(".staff dt")
	if dt.Length() == 0 {
		return
	}
	anime.Staff = make(map[string]string)
	dt.Each(func(i int, selection *goquery.Selection) {
		anime.Staff[selection.Text()] = selection.NextFiltered("dd").Text()
	})
}

func (p *animeParser) parseAnimeOpening(anime *Anime, selection *goquery.Selection) {
	op := selection.Find(".op")
	if op.Length() == 0 {
		return
	}
	song := p.parseAnimeSong(op)
	anime.Opening = &song
}

func (p *animeParser) parseAnimeEnding(anime *Anime, selection *goquery.Selection) {
	ed := selection.Find(".ed")
	if ed.Length() == 0 {
		return
	}
	song := p.parseAnimeSong(ed)
	anime.Ending = &song
}

func (p *animeParser) parseAnimeCast(anime *Anime, selection *goquery.Selection) {
	dt := selection.Find(".cast dt")
	if dt.Length() == 0 {
		return
	}
	anime.Cast = make(map[string]string)
	dt.Each(func(i int, selection *goquery.Selection) {
		anime.Cast[selection.Text()] = selection.NextFiltered("dd").Text()
	})
}

func (p *animeParser) parseAnimePrograms(anime *Anime, selection *goquery.Selection) {
	li := selection.Find(".progs > ul > li")
	if li.Length() == 0 {
		return
	}
	anime.Programs = make([]*Program, 0)
	li.Each(func(i int, selection *goquery.Selection) {
		anime.Programs = append(anime.Programs, &Program{
			Time: selection.Find(".StTime").Text(),
			Name: selection.Find(".ChName").Text(),
		})
	})
}

func (p *animeParser) parseAnimeLinks(anime *Anime, selection *goquery.Selection) {
	a := selection.Find(".links > ul > li > a")
	if a.Length() == 0 {
		return
	}
	anime.Links = make([]*Link, 0)
	a.Each(func(i int, selection *goquery.Selection) {
		anime.Links = append(anime.Links, &Link{
			Title: selection.Text(),
			URL:   selection.AttrOr("href", ""),
		})
	})
}

func (p *animeParser) parseAnimeSong(op *goquery.Selection) Song {
	var song Song
	song.Title = op.Find(".trackname").Text()
	song.Details = make(map[string]string)
	op.Find("dt").Each(func(i int, selection *goquery.Selection) {
		song.Details[selection.Text()] = selection.NextFiltered("dd").Text()
	})
	return song
}
