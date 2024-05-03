package webscraper

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	httpclient "hospital-prices/http_client"

	"golang.org/x/net/html"
)

func ScrapeFiles(url string) ([]string, error) {
	csvFiles := make([]string, 0)

	reader, err := httpclient.GetURL(url)

	if err != nil {
		return nil, err
	}

	defer reader.Close()

	htmlPage, err := html.Parse(reader)

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
	// check if test is json format by verifying the first character is an open curly bracket
	switch node.Data[0] {
	case '{':
		jsonMap := make(map[string]interface{})
		json.Unmarshal([]byte(node.Data), &jsonMap)
		if baseMap, ok := jsonMap["sitecore"]; ok {
			loopInterfacesAndPrint(baseMap, csvFileNames)
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

func loopInterfacesAndPrint(value interface{}, fileNames *[]string) {
	switch inVal := value.(type) {
	case map[string]interface{}:
		for _, val := range inVal {
			loopInterfacesAndPrint(val, fileNames)
		}
	case []interface{}:
		for _, val := range inVal {
			loopInterfacesAndPrint(val, fileNames)
		}
	case string:
		// https://www.healthonecares.com/-/media/project/hca/hut/cms-files/84-1321373_p-sl-medical-center_standardcharges.csv?rev=04cd82bc2d134e3a82c7ebbd8b942c09
		if strings.Contains(inVal, "csv") && !slices.Contains(*fileNames, inVal) {
			*fileNames = append(*fileNames, inVal)
		}
	default:
		// do nothing
	}
}
