package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	addr string
	text string
}

func (l Link) String() string {
	return fmt.Sprintf("Link{\n\tHref: \"%s\",\n\tText: \"%s\",\n}\n", l.addr, l.text)
}

var links []Link

func Parse(n *html.Node) {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				s := GetText(n.FirstChild)
				if s != "" {
					links = append(links, Link{n.Data, s})
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Parse(c)
	}
}

func GetText(n *html.Node) string {

	var meh string

	var rec func(*html.Node)
	rec = func(n *html.Node) {
		if n.Type == html.TextNode {
			s := n.Data

			meh += s

		}
		if n.FirstChild != nil {
			rec(n.FirstChild)
		}
		if n.NextSibling != nil {
			rec(n.NextSibling)
		}
	}
	rec(n)

	return strings.Join(strings.Fields(meh), " ")
}

func main() {
	filepath := flag.String("file", "ex1.html", "html file to be parsed")
	flag.Parse()

	file, err := os.Open(*filepath)
	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(file)
	if err != nil {
		panic(err)
	}

	Parse(doc)

	fmt.Println(links)
}
