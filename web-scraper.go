package main

/*
TODO:
- add some commands using github page of codegangsta
- understand the concurrent search from tutorial page
- get all the 2016 albums w/ ratings

*/

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/html"

	"github.com/codegangsta/cli"
)

func searchInHTML(url string) []string {
	// HTTP request
	var links []string

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return links
	}

	tokenizer := html.NewTokenizer(response.Body)

	for {
		curr := tokenizer.Next()

		if curr == html.ErrorToken {
			response.Body.Close()
			return links
		}

		if curr == html.StartTagToken {
			token := tokenizer.Token()

			if token.Data != "a" {
				continue
			}

			href := ""

			for _, a := range token.Attr {
				if a.Key == "href" {
					href = a.Val
					//fmt.Println("Found link: ", href)
					links = append(links, href)
					break
				}
			}

			if href == "" {
				continue
			}
		}

	}
}

func main() {
	app := cli.NewApp()
	app.Name = "readHTML"
	app.Version = "0.0.1"
	app.Compiled = time.Now()
	app.Usage = "give me a URL and I'll give you a HTML"
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

	httpFlag := false // w/ http if true, false otherwise

	fmt.Println("app.Flags: \n", app.Flags)

	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			URL = c.Args().Get(0)
		}

		fmt.Println("url is \n", URL, "\n")

		links := searchInHTML(URL)

		fmt.Println("links are:")

		for _, link := range links {
			if strings.Index(link, "http") == 0 { //fetch only those that begin w/ http
				fmt.Println(link)
			}
		}

		fmt.Println("")

		return nil
	}

	app.Run(os.Args)
}
