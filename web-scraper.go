package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	"golang.org/x/net/html"

	"github.com/urfave/cli"
)

func searchInHTML(body io.Reader) (links []string) {
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
	app.Version = "0.0.2"
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

	URLs := cli.StringSlice{}

	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "url, u",
			Value: &URLs,
			Usage: "the URL to get data from",
		},
	}

	app.Action = func(c *cli.Context) error {
		if len(URLs) == 0 {
			URLs = append(URLs, "http://google.com") // default value
		}
		var wg sync.WaitGroup
		var mux sync.Mutex
		data := map[string][]string{}
		wg.Add(len(URLs))
		for _, currentUrl := range URLs {

			go func(url string) { //goroutine for each url
				response, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
				}
				defer response.Body.Close()
				defer wg.Done()
				returnedLinks := searchInHTML(response.Body)
				mux.Lock()
				data[url] = returnedLinks
				mux.Unlock()

			}(currentUrl)
		}

		wg.Wait()

		for url, links := range data {
			fmt.Printf("** %s **\n\n", url)
			for _, link := range links {
				fmt.Println(link)
			}
			fmt.Println("")
		}
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
