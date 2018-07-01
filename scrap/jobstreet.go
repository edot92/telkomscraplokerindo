package scrap

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func checkError(err error) {
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}
func RUnJobStreeet() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("jobstreet.co.id", "www.jobstreet.co.id"),
	)
	// On every a element which has href attribute call callback
	// c.OnHTML(".panel-body", func(e *colly.HTMLElement) {
	// 	link := e.Attr("id")
	// 	if strings.Contains(link, "job_ad_") {
	// 		srcTxt := e.Text
	// 		srcTxt = strings.Replace(srcTxt, " ", "", -1)
	// 		srcTxt = strings.Replace(srcTxt, "\r\n", "", -1)
	// 		srcTxt = strings.Replace(srcTxt, "\n", "", -1)
	// 		srcTxt = strings.Replace(srcTxt, "\r", "", -1)
	// 		log.Println(srcTxt)
	// 		log.Println("_--------------------")
	// 	}
	// })
	count := 0
	c.OnHTML(".position-title-link", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// log.Println(link)
		// if strings.Contains(link, "job_ad_") {
		// 	srcTxt := e.Text
		// 	srcTxt = strings.Replace(srcTxt, " ", "", -1)
		// 	srcTxt = strings.Replace(srcTxt, "\r\n", "", -1)
		// 	srcTxt = strings.Replace(srcTxt, "\n", "", -1)
		// 	srcTxt = strings.Replace(srcTxt, "\r", "", -1)
		// 	log.Println(srcTxt)
		// 	log.Println("_--------------------")
		// }
		// if count == 0 {
		response, err := http.Get(link)
		checkError(err)
		defer response.Body.Close()
		doc, err := goquery.NewDocumentFromReader(io.Reader(response.Body))
		checkError(err)
		// fmt.Println(doc)
		el := doc.Find("#job_description")
		log.Println(el.Text())
		log.Println("--------------")
		// }
		count++
		_ = link
	})
	// Start scraping
	c.Visit("https://www.jobstreet.co.id/id/job-search/job-vacancy.php?key=dokter&area=1&option=1&job-source=1%2C64&classified=0&job-posted=0&sort=2&order=0&pg=3&src=16&srcr=16&ojs=10")
}
