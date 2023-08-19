package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:     "fetch",
	Short:   "fetch a Stackoverflow url to terminal",
	Aliases: []string{"f"},
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fetch(args[0])
	},
}

func fetch(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		log.Fatalf("status code error: %s %s", res.Status, body)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalf("error parsing response body: %s", err)
	}

	var converter = md.NewConverter("", true, nil)
	converter.AddRules(
		md.Rule{
			Filter: []string{"a"},
			Replacement: func(content string, selec *goquery.Selection, opts *md.Options) *string {
				return md.String("")
			},
		},
	)
	doc.Find("div.answercell").Each(func(i int, s *goquery.Selection) {
		mdtxt := converter.Convert(s)
		fmt.Println("######################################################")
		fmt.Println()
		fmt.Println(mdtxt)
		fmt.Println()
	})
}
