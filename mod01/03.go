package mod01

import (
	"net/http"

	"golang.org/x/net/html"
)

// Retrieves all `href`s of existing anchor tags in a specific HTML node.
func crawl(node *html.Node) ([]string, error) {
	anchors := make([]string, 0)
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				anchors = append(anchors, a.Val)
				break
			}
		}
	} else {
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			subAnchors, err := crawl(child)
			if err == nil {
				anchors = append(anchors, subAnchors...)
			}
		}
	}
	return anchors, nil
}

// Retrieves all links from a page's content with its URL.
func CrawlAnchorsFrom(url string) ([]string, error) {
	res, err := http.Get(url) // validations run inside as well
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	node, _ := html.Parse(res.Body)
	return crawl(node)
}
