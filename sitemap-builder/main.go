package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	link "github.com/gophercises/link/students/ccallergard"
)

func main() {

	/**
	1. Get the webpage
	2. parse all the links on the page
	3. build proper urls with our links
	4. filter out any links w/ a diff domain
	5. Find all pages (BFS)
	*/

	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	flag.Parse()

	res, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	reqUrl := res.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	links, _ := link.Parse(res.Body)
	var hrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}

	for _, href := range hrefs {
		fmt.Println(href)
	}
}
