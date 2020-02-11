package main

// TODO: concurrent search for multiple sites

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"

	"github.com/urfave/cli"
)

func searchInHTML(body io.ReadCloser) (links []string) {
	tokenizer := html.NewTokenizer(body)

	curr := tokenizer.Next()
	for ; curr != html.ErrorToken; curr = tokenizer.Next() {

		if curr == html.StartTagToken {

			token := tokenizer.Token()
			if token.Data == "a" {

				for _, a := range token.Attr {
					if a.Key == "href" {
						if strings.Index(a.Val, "http") == 0 {
							links = append(links, a.Val)
							break
						}
					}
				}
			}
		}

	}
	return links
}

func main() {
	app := cli.NewApp()
	app.Name = "web-scraper"
	app.Version = "0.0.1"
	app.Usage = "finds all the links inside HTML"
	app.Authors = []cli.Author{
		{
			Name:  "Ünal Akünal",
			Email: "unal.akunal@gmail.com",
		},
	}

	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "There is no command named %q \n", command)
	}

	var URL string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "url, u",
			Value:       "http://google.com",
			Usage:       "the URL to get data from",
			Destination: &URL,
		},
	}

	app.Action = func(c *cli.Context) error {
		response, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()
		links := searchInHTML(response.Body)

		fmt.Println("")

		for _, link := range links {
			fmt.Println(link)
		}
		fmt.Println("")

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
