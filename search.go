package main

import (
	"context"
	"log"
	"os"
	"strings"
	"text/template"

	_ "embed"

	"github.com/JackKCWong/stackovimflow/cs"
	"github.com/spf13/cobra"
)

//go:embed soi.template
var tmpl string

var t = template.Must(template.New("tmpl").Parse(tmpl))

var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Search for a question",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		client, err := cs.NewClient(context.Background(),
			cs.WithApiKey(os.Getenv("GCS_API_KEY")),
			cs.WithEngineID(os.Getenv("GCS_CX")),
		)

		if err != nil {
			log.Fatal(err)
		}

		query := strings.Join(args, " ")

		search(client, query)
	},
}

func search(client *cs.CSClient, query string) {
	res, err := client.Search(query)
	if err != nil {
		log.Fatal(err)
	}

	t.ExecuteTemplate(os.Stdout, "tmpl", res)
}
