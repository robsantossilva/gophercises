package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/robsantossilva/gophercises/urlshort"
)

func main() {

	yamlFile := flag.String("yaml", "paths_urls.yaml", "A yaml file with a list of paths ans urls")
	jsonFile := flag.String("json", "paths_urls.json", "A Json file with a list of paths ans urls")

	yamlFileByte, err := readYamlFile(*yamlFile)
	if err != nil {
		log.Fatal(err)
	}

	jsonFileByte, err := readJsonFile(*jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	// 	yaml := `
	// - path: /urlshort
	//   url: https://github.com/gophercises/urlshort
	// - path: /urlshort-final
	//   url: https://github.com/gophercises/urlshort/tree/solution
	// `
	YAMLHandler, err := urlshort.YAMLHandler(yamlFileByte, mapHandler)
	if err != nil {
		panic(err)
	}

	JSONHandler, err := urlshort.JSONHandler(jsonFileByte, YAMLHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	if err := http.ListenAndServe(":8080", JSONHandler); err != nil {
		log.Fatal(err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func readYamlFile(yamlFileName string) ([]byte, error) {
	yamlFileByte, err := ioutil.ReadFile(yamlFileName)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return nil, err
	}
	return yamlFileByte, nil
}

func readJsonFile(jsonFileName string) ([]byte, error) {
	jsonFileByte, err := ioutil.ReadFile(jsonFileName)
	if err != nil {
		fmt.Printf("Error reading JSON file: %s\n", err)
		return nil, err
	}
	return jsonFileByte, nil
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
