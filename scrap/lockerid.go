package scrap

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/edot92/telkomscraplokerindo/db"
	"github.com/gocolly/colly"
)

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	} else {
		return true
	}
}
func ScrapLockerID() {
	noPageStr := "2"
	// c1 := make(chan string)
	// c2 := make(chan string)
	// go func() {
	// 	// log.Println("Client ACK with data: ", data)
	// 	// bRouter.Data["json"] = data
	// 	// bRouter.ServeJSON()
	// 	// bRouter.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	// 	// bRouter.Ctx.WriteString()
	// 	// return
	// 	c1 <- data
	// }()
	// go func() {
	// 	time.Sleep(10 * time.Second)
	// 	c2 <- `{"error":1,"msg":"client device not response"}`
	// }()

	// select {
	// case responseOK := <-c1:
	// 	fmt.Println("received", responseOK)
	// case responError := <-c2:
	// 	// fmt.Println("received", msg2)
	// 	time.Sleep(5 * time.Second)
	// }
	// goto restart
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.loker.id", "loker.id"),
	)
	noPageStr = "1"
	urlTarget := "https://www.loker.id/cari-lowongan-kerja/page/" + noPageStr + "?q&lokasi=jawa-barat&category=0&pendidikan=0"
	var dataLocker db.StructDataLocker
	dataLocker.URL = urlTarget
	// ambil class
	c.OnHTML(".col-action", func(e *colly.HTMLElement) {
		// link := e.Attr("href") // ambil isi href
		// log.Println(link)
		link := e.ChildAttr("a", "href")
		arrayTemp := strings.Split(link, ".html")
		if len(arrayTemp) > 0 {
			link = arrayTemp[0] + ".html"
			if isValidUrl(link) {
				// request http get
				response, err := http.Get(link)
				checkError(err)
				defer response.Body.Close()
				doc, err := goquery.NewDocumentFromReader(io.Reader(response.Body))
				checkError(err)
				el := doc.Find(".entry-content")
				e2 := el.ChildrenFiltered("div")
				// log.Println(e2.Text())
				textSave := e2.Text()
				textSave = strings.Replace(textSave, "Simpan", "", -1)
				textSave = strings.Replace(textSave, "Lamar Pekerjaan", "", -1)
				textSave = strings.Replace(textSave, `	`, "", -1)    //spasi
				textSave = strings.Replace(textSave, `    `, "", -1) //tab
				dataLocker.Data = append(dataLocker.Data, db.Data{
					Post:  textSave,
					Waktu: "",
				})
			}

			// db.InsertDataLocker()
			// log.Println("--------------")
		}
		// _ = link
	})
	c.OnScraped(func(e *colly.Response) {
		err := db.InsertDataLocker(dataLocker)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("FINISH")
		}
		countInt, _ := strconv.Atoi(noPageStr)
		countInt++
		if countInt < 100 {
			noPageStr = strconv.Itoa(countInt)
			log.Println("count ke page " + noPageStr)
			// goto restart
		} else {
			log.Println("SUCESS 100 page")
		}

	})
	c.OnError(func(e *colly.Response, err error) {
		log.Println("error on error colly")
		log.Println(err)
	})
	// Start scraping
	c.Visit(urlTarget)

}
