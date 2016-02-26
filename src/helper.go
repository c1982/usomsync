package main

import (
	"bufio"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//GetXmlData Receives content from a URL
func GetXmlData(url string) (content string, err error) {

	client := createHttpClient()
	r, err := client.Get(url)

	if err != nil {
		return "", err
	}

	defer r.Body.Close()

	bt, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return "", err
	}

	return string(bt), err
}

func createHttpClient() http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return http.Client{Transport: tr}
}

//DeserializeXml Deserialize XML string to Rss type
func DeserializeXml(xmlData string) (rss Rss) {
	xml.Unmarshal([]byte(xmlData), &rss)
	return rss
}

//SaveSpammerIPs save ipaddr list to SPAMMERIPBLOCKS file
func SaveSpammerIPs(ipaddr []string) (err error) {

	lines, err := GetAllLinesFromFile(SPAMMERIPBLOCKS)

	if err != nil {
		return err
	}

	newIps := ClearExistsIpAddr(lines, ipaddr)
	err = AppendNewLinesToFile(SPAMMERIPBLOCKS, newIps)

	return err
}

//ClearExistsIpAddr check exists IP or Hostname in spammer file
func ClearExistsIpAddr(exists []string, newips []string) []string {

	result := []string{}
	isExists := false

	for i := 0; i <= len(newips)-1; i++ {

		isExists = false
		for e := 0; e <= len(exists)-1; e++ {
			if exists[e] == newips[i] {
				isExists = true
				break
			}
		}

		if !isExists {
			result = append(result, newips[i])
		}
	}

	return result
}

//GetAllLinesFromFile get all lines in text file
func GetAllLinesFromFile(filePath string) (lines []string, err error) {

	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines = scannerToLines(scanner)

	return lines, scanner.Err()
}

func scannerToLines(s *bufio.Scanner) (lines []string) {

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines

}

//AppendNewLinesToFile add new lines to text file
func AppendNewLinesToFile(filePath string, lines []string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND, 0666)

	if err != nil {
		return err
	}

	defer file.Close()

	for _, line := range lines {
		_, err = file.WriteString(fmt.Sprintln(line))

		if err != nil {
			break
		}
	}

	return err
}
