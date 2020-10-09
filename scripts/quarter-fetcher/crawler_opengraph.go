package main

import (
	"github.com/otiai10/opengraph"
)

type openGraphCrawler struct {
	link  string
	anime *Anime
}

func NewOpenGraphCrawler(link string, anime *Anime) Crawler {
	return &openGraphCrawler{link: link, anime: anime}
}

func (c *openGraphCrawler) Crawl() (interface{}, error) {
	og, err := opengraph.Fetch(c.link)
	if err != nil {
		return nil, err
	}

	c.anime.OpenGraph = &OpenGraph{
		Description: og.Description,
	}

	if og.Image != nil && len(og.Image) != 0 {
		c.anime.OpenGraph.Images = make([]*OpenGraphImage, 0)
		for _, image := range og.Image {
			c.anime.OpenGraph.Images = append(c.anime.OpenGraph.Images, &OpenGraphImage{
				URL: image.URL,
			})
		}
	}
	return nil, nil
}
