package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	// "gioui.org/x/markdown"
	md "github.com/JohannesKaufmann/html-to-markdown"

	// for screenshots from chrome
	"github.com/chromedp/chromedp"
)

func main() {
	fmt.Println("hello")

	// query := "where to get w2"
	query := "i had fun once"
	fmt.Println("query:", query)

	// get a screenshot using
	// https://github.com/chromedp/examples/blob/master/screenshot/main.go

	results := googleSearch(query)
	fmt.Println("results:", results)
}

func urlEncodeQuery(query string) string {

	// fmt.Println("urlEncodeQuery:", query)
	query = url.QueryEscape(query)
	// fmt.Println("urlEncodeQuery:", query)
	// urlEncodeQuery: where+to+get+w2

	return query
}

func googleSearch(query string) string {
	fmt.Println("googleSearch:", query)

	encQuery := urlEncodeQuery(query)

	// https://www.google.com/search?q=where+to+get+w2
	googleUrl := "https://www.google.com/search?q=" + encQuery
	fmt.Println("googleUrl:", googleUrl)
	// https://www.google.com/search?q=where+to+get+w2

	searchResults, _ := getPage(googleUrl)
	fmt.Println("searchResults:", len(searchResults))

	searchMarkdown := webToMarkdown(searchResults)
	fmt.Println("searchMarkdown:", searchMarkdown)

	// writeMarkdown("search.md", searchMarkdown)
	markdownFileName := encQuery + ".md"
	writeMarkdown(markdownFileName, searchMarkdown)
	fmt.Println("saved:", markdownFileName)

	// save image of search page
	// savePageAsImage(url string, imageFilePath string)
	imageFileName := encQuery + ".png"
	savePageAsImage(googleUrl, imageFileName)

	return searchResults
}

func getPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Read body: %v", err)
	}

	return string(data), nil
}

// webToMarkdown give html from webpage, get back markdown
func webToMarkdown(webHtml string) string {
	converter := md.NewConverter("", true, nil)

	//html := `<strong>Important</strong>`

	// markdown, err := converter.ConvertString(html)
	markdown, err := converter.ConvertString(webHtml)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	fmt.Println("md ->", markdown)

	// replace this
	// /url?q=https://
	// with
	// https://

	return markdown
}

func writeMarkdown(filepath string, text string) {
	file, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(text)
}

// from // https://github.com/chromedp/examples/blob/master/screenshot/main.go
//
// save a page as a screenshot
func savePageAsImage(url string, imageFilePath string) {

	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// capture screenshot of an element
	var buf []byte
	// if err := chromedp.Run(ctx, elementScreenshot(`https://pkg.go.dev/`, `img.Homepage-logo`, &buf)); err != nil {
	// 	log.Fatal(err)
	// }
	// if err := os.WriteFile("elementScreenshot.png", buf, 0o644); err != nil {
	// 	log.Fatal(err)
	// }

	// capture entire browser viewport, returning png with quality=90
	// err := chromedp.Run(ctx, fullScreenshot(`https://brank.as/`, 90, &buf))
	err := chromedp.Run(ctx, fullScreenshot(url, 90, &buf))
	if err != nil {
		log.Fatal(err)
	}
	// err = os.WriteFile("fullScreenshot.png", buf, 0o644)
	err = os.WriteFile(imageFilePath, buf, 0o644)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("wrote %s and %s \n", url, imageFilePath)
}

// elementScreenshot takes a screenshot of a specific element.
// func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
// 	return chromedp.Tasks{
// 		chromedp.Navigate(urlstr),
// 		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
// 	}
// }

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Note: chromedp.FullScreenshot overrides the device's emulation settings. Use
// device.Reset to reset the emulation and viewport settings.
func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}
