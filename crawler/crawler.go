package main

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strings"

	"io/ioutil"
	"net/http"
)

// set a baseDomain to start the crawl from
const baseDomain string = "https://example.com"

func getSource(domain string) string {
	// http get the domain, log failure on error
	resp, err := http.Get(domain)
	if err != nil {
		log.Fatal(err)
	}

	// read response body and store in handler
	// log failure on error
	respBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	// return the body html as a string
	return string(respBody)
}

func findLinks(html string) []string {
	var linksArr []string
	// create a new line scanner for the passed in html
	scanner := bufio.NewScanner(strings.NewReader(html))

	// iterate over the string line by line
	// first match <a href> tags
	// then match actual internal URIs
	for scanner.Scan() {
		re := regexp.MustCompile("<a href=\"/.*\"")
		match := re.FindString(scanner.Text())
		if match != "" {
			re := regexp.MustCompile("[/#]{1}[0-9A-Za-z./_#-]+")
			subMatch := re.FindString(match)
			linksArr = append(linksArr, subMatch)
		}
	}

	// return the array of links
	return linksArr
}

func main() {
	// fetch all links on the base domain
	links := findLinks(getSource(baseDomain))
	// initialise empty map to store sitemap
	siteMap := make(map[string][]string)

	for _, link := range links {
		// for each link http get the page and find further links
		// store each new link in the map with the fullURL as the key
		fullURL := baseDomain + link
		siteMap[link] = findLinks(getSource(fullURL))
	}

	// hacky output and format loop
	for key := range siteMap {
		fmt.Printf("%s %s%s%s", "|__", baseDomain, key, "\n")
		for _, value := range siteMap[key] {
			fmt.Printf("%s %s%s", "   |__", value, "\n")
		}
	}
}
