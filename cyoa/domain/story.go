package domain

import (
	"encoding/json"
	"io"
)

type Story struct {
	Chapters     map[string]Chapter
	FirstChapter string
}

type Chapter struct {
	FirstChapter bool     `json:"is-first-chapter"`
	Title        string   `json:"title"`
	Paragraphs   []string `json:"story"`
	Options      []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func (s *Story) setInitialChapter() {

	for i, c := range s.Chapters {
		if c.FirstChapter {
			s.FirstChapter = i
		}
	}

	if s.FirstChapter == "" {
		for i, _ := range s.Chapters {
			s.FirstChapter = i
			break
		}
	}
}

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story.Chapters); err != nil {
		return story, err
	}
	story.setInitialChapter()
	return story, nil
}
