package cs

import (
	"context"

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

	resp, err := call.Do()
	if err != nil {
		return nil, err
	}

	var items = make([]Item, 0, len(resp.Items))
	for i := range resp.Items {
		items = append(items, Item{
			Title: resp.Items[i].Title,
			Link:  resp.Items[i].Link,
		})
	}

	return &SearchResult{
		Items: items,
	}, nil
}
