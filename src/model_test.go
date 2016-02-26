package main

import (
	"testing"
)

var xmlString = `<?xml version='1.0' encoding='UTF-8' ?>
		<rss version="2.0">
		<channel>
		<title>Ulusal Siber Olaylara Müdahale Merkezi (USOM) Zararlı Bağlantılar</title>
		<link>http://usom.gov.tr/</link>
		<description>USOM - Zararlı Bağlantı Listesi</description>
		<language>tr</language>
		<item>
		<title>Zararlı Yazılım Barındıran/Yayan IP - 93.79.65.222</title>
		<link>93.79.65.222</link>
		<pubDate>Thu, 25 Feb 2016 14:57:36 +0000</pubDate>
		<description>SOME</description>
		</item>
		<item>
		<title>Zararlı Yazılım Barındıran/Yayan IP - 93.118.90.254</title>
		<link>93.118.90.254</link>
		<pubDate>Thu, 25 Feb 2016 14:57:25 +0000</pubDate>
		<description>SOME</description>
		</item>
		<item>
		<title>Zararlı Yazılım Barındıran/Yayan IP - zararlidomain.com</title>
		<link>zararlidomain.com</link>
		<pubDate>Thu, 25 Feb 2016 14:57:25 +0000</pubDate>
		<description>SOME</description>
		</item>
		</channel>
	</rss>`

func TestXMLUnmarshal(t *testing.T) {

	rss := DeserializeXml(xmlString)

	if rss.RssChannel.Title == "" {
		t.Error("Title is nil")
	}

	if len(rss.RssChannel.Items) != 3 {
		t.Errorf("Item lenght is %v", len(rss.RssChannel.Items))
	}

	if rss.RssChannel.Items[1].Link != "93.118.90.254" {
		t.Errorf("2nd Items wrong value: %v", rss.RssChannel.Items[1].Link)
	}
}

func TestLinkToStringArray(t *testing.T) {
	rss := DeserializeXml(xmlString)

	domains := rss.ToDomainList()
	t.Log(len(domains))

	for _, it := range domains {
		t.Log(it)
	}

	if len(domains) == 0 {
		t.Error("Domains array cannot be empty")
	}

	iplist := rss.ToIPList()
	t.Log(len(iplist))

	for _, it := range iplist {
		t.Log(it)
	}

	if len(iplist) == 0 {
		t.Error("IP List array cannot be empty")
	}
}
