package main

import (
	"flag"
	"fmt"
)

//#sender_host_reject = +include_unknown:lsearch*;/etc/spammers
var USOM_URL string
var SPAMMERIPBLOCKS string
var SPAMMER string

func main() {

	USOM_URL := flag.String("url", "https://www.usom.gov.tr/rss/zararli-baglanti.rss", "-url=https://www.usom.gov.tr/rss/zararli-baglanti.rss")
	SPAMMERIPBLOCKS := flag.String("ipblocks", "/etc/spammeripblocks", "-ipblocks=/etc/spammeripblocks")
	SPAMMER := flag.String("hostblocks", "/etc/spammer", "-hostblocks=/etc/spammer")

	rsstext, err := GetXmlData(*USOM_URL)

	if err != nil {
		panic(err)
	}

	rss := DeserializeXml(rsstext)

	fmt.Println(rss.RssChannel.Description)
	fmt.Println(SPAMMERIPBLOCKS)
	fmt.Println(SPAMMER)

}
