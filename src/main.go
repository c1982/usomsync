package main

import (
	"flag"
	"log"
)

//#sender_host_reject = +include_unknown:lsearch*;/etc/spammers
var USOM_URL string
var SPAMMERIPBLOCKS string
var SPAMMERDOMAINS string

func main() {

	flag.StringVar(&USOM_URL, "url", "https://www.usom.gov.tr/rss/zararli-baglanti.rss", "-url=https://www.usom.gov.tr/rss/zararli-baglanti.rss")
	flag.StringVar(&SPAMMERIPBLOCKS, "ip", "/etc/spammeripblocks", "-ip=/etc/spammeripblocks")
	flag.StringVar(&SPAMMERDOMAINS, "d", "/etc/blockeddomains", "-d=/etc/blockeddomains")
	flag.Parse()

	rsstext, err := GetXmlData(USOM_URL)
	checkError(err)

	rss := DeserializeXml(rsstext)

	err = SaveSpammerIPs(rss.ToIPList(), SPAMMERIPBLOCKS)
	checkError(err)

	err = SaveSpammerHosts(rss.ToDomainList(), SPAMMERDOMAINS)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		flag.Usage()
		log.Fatal(err)
	}
}
