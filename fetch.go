package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

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

	txt := render(res.Body)
	fmt.Println(txt)
}

func render(r io.Reader) string {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatalf("error parsing response body: %s", err)
	}

	var converter = md.NewConverter("", true, nil)
	converter.AddRules(
		md.Rule{
			Filter: []string{"button"},
			Replacement: func(content string, selec *goquery.Selection, opts *md.Options) *string {
				return md.String("")
			},
		},
		md.Rule{
			Filter: []string{"a"},
			Replacement: func(content string, selec *goquery.Selection, opts *md.Options) *string {
				if href, exist := selec.Attr("href"); exist && strings.HasPrefix(href, "/") {
					return md.String("")
				}
				return &content
			},
		},
		md.Rule{
			Filter: []string{"div"},
			Replacement: func(content string, selec *goquery.Selection, opts *md.Options) *string {
				if clazz, exist := selec.Attr("class"); exist && strings.Contains(clazz, "user-details") {
					return md.String("")
				}
				return &content
			},
		},
		md.Rule{
			Filter: []string{"pre"},
			Replacement: func(content string, selec *goquery.Selection, opts *md.Options) *string {
				lang := ""
				if clazz, exist := selec.Attr("class"); exist {
					lang = extractLang(clazz)
				}

				codeblock := fmt.Sprintf("```%s\n%s\n```", lang, strings.ReplaceAll(content, "`", ""))
				return &codeblock
			},
		},
	)
	var sb strings.Builder
	doc.Find("div.answercell").Each(func(i int, s *goquery.Selection) {
		mdtxt := converter.Convert(s)
		sb.WriteString("*** I am a separator ***\n")
		sb.WriteString(mdtxt)
		sb.WriteString("\n\n")
	})

	return sb.String()
}

var rexLang = regexp.MustCompile(`lang-(\w+)`)

func extractLang(input string) string {
	result := rexLang.FindStringSubmatch(input)
	if len(result) > 1 {
		return result[1]
	}

	return ""
}
