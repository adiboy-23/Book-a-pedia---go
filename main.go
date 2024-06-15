package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

type book struct {
	Title string `json:"title"`
}
type searchOutput struct {
	Docs []book `json:"docs"`
}

var rootCmd = &cobra.Command{
	Use:   "bookopedia",
	Short: "Bookopedia is a simple CLI to fetch books according to authors",
}

var cmdSearch = &cobra.Command{
	Use:   "find [author]",
	Short: "Search for books by a specific author",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		findBooks(args[0])
	},
}

func findBooks(author string) {
	safeAuthor := url.QueryEscape(author)
	url := fmt.Sprintf("https://openlibrary.org/search.json?author=%s", safeAuthor)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching data: %s", err)
	}
	defer resp.Body.Close()

	var results searchOutput
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		log.Fatalf("Error decoding data: %s", err)
	}

	if len(results.Docs) == 0 {
		fmt.Println("No books found for this author.")
		return
	}

	fmt.Println("Books found:")
	for _, book := range results.Docs {
		fmt.Printf(" - %s\n", book.Title)
	}
}

func main() {
	rootCmd.AddCommand(cmdSearch)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
