package main

import (
	"flag"
	"net/http"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	flag.Parse()

	res, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	//io.Copy(os.Stdout, res.Body)

	//body, err := ioutil.ReadAll(res.Body)
	//fmt.Println(body)

	//links, _ := link.
}
