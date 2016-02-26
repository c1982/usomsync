package main

import (
	"strings"
)

type Rss struct {
	RssChannel Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Language    string `xml:"language"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
}

func (r *Rss) ToIPList() []string {

	list := []string{}

	for _, item := range r.RssChannel.Items {
		if isValidIp(item.Link) {
			list = append(list, item.Link)
		}
	}

	return list
}

func (r *Rss) ToDomainList() []string {
	list := []string{}

	for _, item := range r.RssChannel.Items {
		if !isValidIp(item.Link) {
			if !strings.HasSuffix(item.Link, "http://") {
				list = append(list, item.Link)
			}
		}
	}

	return list
}
