package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	storyCyoa "github.com/robsantossilva/gophercises/cyoa/domain"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	story := loadStory()

	var initialChapter string = story.FirstChapter
	var chapterName *string = &initialChapter

storyLoop:
	for {

		chapter, ok := story.Chapters[*chapterName]

		if !ok {
			break storyLoop
		}

		fmt.Printf("\n\n____________________________________________________\n\n%s\n____________________________________________________\n\n", chapter.Title)
		for _, paragraph := range chapter.Paragraphs {
			fmt.Printf("%s\n\n", paragraph)
		}
		fmt.Printf("____________________________________________________\n\n")

		if len(chapter.Options) <= 0 {
			fmt.Println("FIM!")
			break storyLoop
		}

		for n, option := range chapter.Options {
			fmt.Printf("%s - %s\n", strconv.Itoa(n+1), option.Text)
		}

	inputLoop:
		for {
			option, err := strconv.Atoi(getInput(reader))
			if option > len(chapter.Options) || err != nil {
				fmt.Println("Opção inválida")
				continue inputLoop
			}
			*chapterName = chapter.Options[option-1].Chapter
			break inputLoop
		}

		// fmt.Println("Choose your own adventure: ")
		// fmt.Scanf("%s", chapterName)

	}
}

func getInput(r *bufio.Reader) string {
	fmt.Print("Choose your own adventure: ")
	text, _ := r.ReadString('\n')
	// convert CRLF to LF
	return strings.Replace(text, "\n", "", -1)
}

func loadStory() storyCyoa.Story {
	//load story struct by json file
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := storyCyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	return story
}
