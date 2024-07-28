package webscraper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	fileparser "hospital-prices/file_parser"

	"golang.org/x/net/html"
)

const TERMINATION_CLAUSE = "standardcharges"

func ScrapeFiles(url string) ([]string, error) {
	csvFiles := make([]string, 0)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	htmlPage, err := html.Parse(resp.Body)

	if err != nil {
		return nil, err
	}

	processHTMLNodes(htmlPage, &csvFiles)

	return csvFiles, nil

}

func processHTMLNodes(node *html.Node, fileNames *[]string) {

	if node.Type == html.TextNode {
		parseTextNode(node, fileNames)
	}
	if node.Type == html.ElementNode {
		parseElementNode(node)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		processHTMLNodes(child, fileNames)
	}
}

func parseTextNode(node *html.Node, csvFileNames *[]string) []string {
	var fileNames []string
	switch node.Data[0] {
	// check if node is json format by verifying the first character is an open curly bracket
	case '{':
		jsonMap := make(map[string]interface{})
		json.Unmarshal([]byte(node.Data), &jsonMap)
		if baseMap, ok := jsonMap["sitecore"]; ok {
			fileparser.LoopInterfaces(baseMap, csvFileNames, TERMINATION_CLAUSE)
		}
	default:
		// do nothing
	}

	return fileNames
}

func parseElementNode(node *html.Node) {
	if node.Data == "a" {
		for _, attribute := range node.Attr {
			if attribute.Key == "href" && strings.Contains(attribute.Val, "csv") {
				fmt.Printf("attribute: %v\n\n", attribute.Val)
			}
		}
	}
}
