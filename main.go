package main

import (
	"context"
	"log"
	"os"

	"github.com/JackKCWong/stackovimflow/cs"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use: "stackovimflow <sub command>",
}

func init() {
	root.AddCommand(searchCmd)
	root.AddCommand(fetchCmd)
}

var clientKey = struct{}{}

func main() {
	os.Setenv("HTTP_PROXY", os.Getenv("GCS_HTTP_PROXY"))
	os.Setenv("HTTPS_PROXY", os.Getenv("GCS_HTTPS_PROXY"))

	client, err := cs.NewClient(context.Background(),
		cs.WithApiKey(os.Getenv("GCS_API_KEY")),
		cs.WithEngineID(os.Getenv("GCS_CX")),
	)

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.WithValue(context.Background(), clientKey, client)

	if err := root.ExecuteContext(ctx); err != nil {
		log.Fatal(err)
	}
}
