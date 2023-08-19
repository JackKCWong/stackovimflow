package cs

type Item struct {
	Title   string `json:"title"`  
	Link    string `json:"link"`   
	Snippet string `json:"snippet"`
}

type SearchResult struct {
	Items []Item `json:"items"`
}
