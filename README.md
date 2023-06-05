
# Search the web

fast with no bloat, ads, or trackers

gui using go and fltk

- https://golangexample.com/a-simple-go-wrapper-for-fltk2/
- https://github.com/zot/go-fltk
- https://github.com/pwiecz/go-fltk

do we want:
1. a screenshot of the page
2. or the markdown?

take a screenshot of a website
- https://github.com/chromedp/examples/blob/master/screenshot/main.go

this project drives chrome: https://github.com/chromedp

this does pdf: https://github.com/chromedp/examples/blob/master/pdf/main.go

how to parse/scrape websites?
https://github.com/gocolly/colly

    func main() {
        c := colly.NewCollector()

        // Find and visit all links
        c.OnHTML("a[href]", func(e *colly.HTMLElement) {
            e.Request.Visit(e.Attr("href"))
        })

        c.OnRequest(func(r *colly.Request) {
            fmt.Println("Visiting", r.URL)
        })

        c.Visit("http://go-colly.org/")
    }

can we turn webpages into markdown?
https://github.com/JohannesKaufmann/html-to-markdown

    import (
        "fmt"
        "log"

        md "github.com/JohannesKaufmann/html-to-markdown"
    )

    converter := md.NewConverter("", true, nil)

    html := `<strong>Important</strong>`

    markdown, err := converter.ConvertString(html)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println("md ->", markdown)

.

a pdf will hold the text, but maybe bloated