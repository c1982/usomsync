package main

import (
	"regexp"
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
		currentLink := r.extractData(item.Link)

		if currentLink != "" && isValidIp(currentLink) {
			list = append(list, currentLink)
		}
	}

	return list
}

func (r *Rss) ToDomainList() []string {
	list := []string{}

	for _, item := range r.RssChannel.Items {
		currentLink := r.extractData(item.Link)

		if currentLink != "" && !isValidIp(currentLink) {
			list = append(list, currentLink)
		}
	}

	return list
}

func (r *Rss) extractData(text string) string {

	rg := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\n]+@)?(?:www\.)?([^:\/\n]+)`)

	matches := rg.FindStringSubmatch(text)

	if len(matches) == 0 {
		return ""
	} else {
		return matches[1]
	}
}
