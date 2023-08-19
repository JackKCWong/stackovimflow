package main

import (
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
		ctx := cmd.Context()
		client := ctx.Value(clientKey)
		query := strings.Join(args, " ")

		search(client.(*cs.CSClient), query)
	},
}

func search(client *cs.CSClient, query string) {
	res, err := client.Search(query)
	if err != nil {
		log.Fatal(err)
	}

	t.ExecuteTemplate(os.Stdout, "tmpl", res)
}
