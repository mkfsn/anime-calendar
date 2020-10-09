package main

import (
	"time"
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
