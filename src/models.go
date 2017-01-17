package main

import (
	"regexp"
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

func (r *Rss) captureIpv4(text string) string {

	rg := regexp.MustCompile(`(((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))`)

	matches := rg.FindStringSubmatch(text)

	if len(matches) == 0 {
		return nil
	} else {
		return matches[0]
	}
}
