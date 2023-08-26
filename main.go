package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use: "stackovimflow <sub command>",
}

func init() {
	root.AddCommand(searchCmd)
	root.AddCommand(fetchCmd)
	root.AddCommand(installCmd)
}

var clientKey = struct{}{}

func main() {
	os.Setenv("HTTP_PROXY", os.Getenv("GCS_HTTP_PROXY"))
	os.Setenv("HTTPS_PROXY", os.Getenv("GCS_HTTPS_PROXY"))

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
