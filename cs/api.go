package cs

import (
	"context"
	"encoding/json"

	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

type CSClient struct {
	apiKey   string
	engineID string
	css      *customsearch.Service
	cse      *customsearch.CseService
}

type coption func(*CSClient)

// WithEngineID sets the engine ID.
func WithEngineID(engineID string) coption {
	return func(c *CSClient) {
		c.engineID = engineID
	}
}

// WithApiKey
func WithApiKey(apiKey string) coption {
	return func(c *CSClient) {
		c.apiKey = apiKey
	}
}

func NewClient(ctx context.Context, opts ...coption) (*CSClient, error) {
	c := CSClient{}
	for _, o := range opts {
		o(&c)
	}

	css, err := customsearch.NewService(ctx, option.WithAPIKey(c.apiKey))
	if err != nil {
		return nil, err
	}

	c.css = css
	c.cse = customsearch.NewCseService(css)

	return &c, nil
}

func (c CSClient) Search(query string) (*SearchResult, error) {
	call := c.cse.List()
	call.Q(query)
	call.Cx(c.engineID)
	call.Num(5)

	resp, err := call.Do()
	if err != nil {
		return nil, err
	}

	var items = make([]Item, 0, len(resp.Items))
	for i := range resp.Items {
		it := resp.Items[i]
		post := Post{}
		page, err := it.Pagemap.MarshalJSON()
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(page, &post)
		if err != nil {
			return nil, err
		}

		items = append(items, Item{
			Title:   it.Title,
			Link:    it.Link,
			Snippet: it.Snippet,
			Post:    post,
		})
	}

	return &SearchResult{
		Items: items,
	}, nil
}
