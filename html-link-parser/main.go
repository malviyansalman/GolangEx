package main

import (
	"GolangEx"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

var logger = GolangEx.GetLogger()

func main() {
	url := "https://www.google.com/"
	data, err := http.Get(url)
	if err != nil {
		logger.Info("Unable to connect to URL", zap.Any("Error", err))
	}
	node, err := html.Parse(data.Body)
	if err != nil {
		logger.Info("Unable to Parse HTMlL", zap.Any("Error", err))
	}
	urlData := Parse(node)
	for _, node := range urlData {
		fmt.Printf("%+v\n", node)
	}
}
func Parse(node *html.Node) []URLData {
	var result []URLData
	linkNodes := LinkNode(node)
	for _, nod := range linkNodes {
		result = append(result, buildNode(nod))
	}
	return result
}
func buildNode(node *html.Node) URLData {
	var urlData URLData
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			urlData.Href = attr.Val
			break
		}
	}
	urlData.Text = TextNode(node)
	return urlData
}
func TextNode(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}
	if node.Type != html.ElementNode {
		return ""
	}
	var result string
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		result += TextNode(child) + " "
	}
	return strings.Join(strings.Fields(result), " ")
}
func LinkNode(node *html.Node) []*html.Node {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}
	var result []*html.Node
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		result = append(result, LinkNode(child)...)
	}
	return result
}

type URLData struct {
	Href string
	Text string
}
