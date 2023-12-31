package cs

import "time"

type SearchResult struct {
	Items []Item `json:"items"`
}

type Item struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Snippet string `json:"snippet"`
	Post    Post   `json:"pagemap"`
}

type Post struct {
	Question []Question `json:"question"`
	Answers  []Answer   `json:"answer"`
}

type Question struct {
	Upvotecount  string `json:"upvotecount"`
	Commentcount string `json:"commentcount"`
	Answercount  string `json:"answercount"`
	Name         string `json:"name"`
	Datecreated  string `json:"datecreated"`
	Text         string `json:"text"`
	URL          string `json:"url"`
}

type Answer struct {
	Upvotecount string `json:"upvotecount"`
	Text        string `json:"text"`
	Datecreated string `json:"datecreated"`
	URL         string `json:"url"`
}

func (a Answer) CreatedDate() (string, error) {
	tm, err := time.Parse("2006-01-02T15:04:05", a.Datecreated)
	if err != nil {
		return "", err
	}

	return tm.Format("Jan 02, 2006"), nil
}