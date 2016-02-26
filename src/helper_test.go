package main

import (
	"testing"
)

func TestReadUsomXml(t *testing.T) {

	xmlstr, err := GetXmlData("https://www.usom.gov.tr/rss/zararli-baglanti.rss")

	if err != nil {
		t.Error(err.Error())
	}

	if xmlstr == "" {
		t.Error("Content is nil")
	}
}

func TestClearExistsIpAddr(t *testing.T) {

	spammeripblocks_file := []string{"10.7.7.8", "192.168.5.2", "55.4.3.1", "60.1.1.12", "10.5.5.80"}
	snom_malware_ips := []string{"10.7.7.8", "10.5.5.80", "4.2.2.1", "4.2.2.7"}

	new_ips_list := ClearExistsItems(spammeripblocks_file, snom_malware_ips)

	if len(new_ips_list) != 2 {
		t.Errorf("new array lenght is %v", len(new_ips_list))
	}

	for _, item := range new_ips_list {
		if item == "10.5.5.80" || item == "10.7.7.8" {
			t.Error("10.7.7.8 or 10.5.5.80 cannot be exists in array")
		}
	}
}

func TestAppendNewLinesToFile(t *testing.T) {
	new_ips := []string{"10.7.7.8", "10.5.5.80", "4.2.2.1", "4.2.2.7"}
	err := AppendNewLinesToFile(".\\test_append_spammeripblocks", new_ips)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetAllLinesFromFile(t *testing.T) {
	lines, err := GetAllLinesFromFile(".\\test_lines_spammeripblocks")

	if err != nil {
		t.Error(err)
	}

	if len(lines) != 5 {
		t.Error("Lines count must be 5. Current:", len(lines))
	}
}
