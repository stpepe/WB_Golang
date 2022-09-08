package main

/*
=== Утилита wget ===
Реализовать утилиту wget с возможностью скачивать сайты целиком
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bytes"
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

func getLinks(data []byte) []string {
	var links []string
	reader := bytes.NewReader(data)
	z := html.NewTokenizer(reader)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						urlStr, err := url.Parse(attr.Val)
						if err != nil {
							continue
						}
						links = append(links, strings.TrimSpace(urlStr.Path))
					}
				}
			}

		}
	}
}

func writeHTML(filename string, data []byte) error {
	err := os.MkdirAll(filename, os.ModePerm)
	if err != nil {
		return err
	}

	return os.WriteFile(filename+"index.html", data, 0644)
}

func wget(url string, r bool) {
	fmt.Println("Downloading ", url)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = writeHTML(path.Join("./output", resp.Request.URL.Path), data)
	if err != nil {
		fmt.Println(err)
		return
	}

	if r {
		links := getLinks(data)

		for _, link := range links {
			re := regexp.MustCompile("^/")
			page := ""
			if re.MatchString(link) {
				page = resp.Request.URL.Scheme + "://" + path.Join(resp.Request.URL.Host, resp.Request.URL.Path, link)
			} else {
				page = resp.Request.URL.Scheme + "://" + path.Join(resp.Request.URL.String(), link)
			}

			fmt.Println(page)

			wget(page, true)
		}
	}
}

func main() {
	flagR := flag.Bool("r", false, "")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		return
	}

	wget(args[0], *flagR)
}