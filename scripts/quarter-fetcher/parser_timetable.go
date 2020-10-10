package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type timetableParser struct {
	selection *goquery.Selection
	anime     *Anime
}

func NewTimetableParse(selection *goquery.Selection, anime *Anime) Parser {
	return &timetableParser{selection: selection, anime: anime}
}

func (p *timetableParser) Parse() error {
	timetableByChannel := make(map[string]Timetable)
	p.selection.Find("tbody tr").Each(func(i int, tr *goquery.Selection) {
		if i == 0 {
			return
		}

		var episode Episode

		channel := tr.Find("td.ch").Text()
		timetable := timetableByChannel[channel]
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

	p.anime.Channels = make([]string, 0)
	for channel := range timetableByChannel {
		p.anime.Channels = append(p.anime.Channels, channel)
	}

	sort.Strings(p.anime.Channels)
	p.anime.Timetables = make([]Timetable, 0)
	for _, channel := range p.anime.Channels {
		p.anime.Timetables = append(p.anime.Timetables, timetableByChannel[channel])
	}

	return nil
}
