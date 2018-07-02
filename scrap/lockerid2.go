package scrap

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/edot92/telkomscraplokerindo/db"
	"github.com/headzoo/surf"
)

func ScrapLockerID2() {
	noPageStr := "1"
	maxPage := 1000
restart:
	time.Sleep(1 * time.Second)
	urlTarget := "https://www.loker.id/cari-lowongan-kerja/page/" + noPageStr + "?q&lokasi=jawa-barat&category=0&pendidikan=0"
	//  + "&numpangsebentaryabuatskripsidoang=https%3A%2F%2Fwww.facebook.com%2Feddot.fu"
	log.Println(urlTarget)
	var dataLocker db.StructDataLocker
	dataLocker.URL = urlTarget
	bow1 := surf.NewBrowser()
	err := bow1.Open(urlTarget)
	if err != nil {
		log.Fatal(err)
	}
	tempCount := 0
	countChildHref := 0
	// log.Println(bow1.Body())
	scrap1 := bow1.Find(".col-action")
	scrap1.Find("*").Each(func(arg1 int, arg2 *goquery.Selection) {
		// ambill tag href
		link, isFound := arg2.Attr("href")
		if isFound && strings.Contains(arg2.Text(), "Selengkapnya") {
			tempCount++
			// split .html agar hilang kalimat setelah html
			tempArray := strings.Split(link, ".html")
			if len(tempArray) > 1 {
				countChildHref++
				link = tempArray[0] + ".html"
				link = strings.Replace(link, "#", "", -1)
				link = strings.Replace(link, " ", "", -1)
				link = strings.Replace(link, "\r\n", "", -1)
				link = strings.Replace(link, "\r\n\r\n", "", -1)
				link = strings.Replace(link, "	", "", -1)
				// log.Println(link)
				bow1 := surf.NewBrowser()
				err := bow1.Open(link)
				if err != nil {
					log.Fatal(err)
				}
				el := bow1.Find(".entry-content")
				e2 := el.ChildrenFiltered("div")
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
		}
	})
	// max href yang di ambil
	log.Println(countChildHref)
	log.Println("selesai")
	// insert ke db
	err = db.InsertDataLocker(dataLocker)
	if err != nil {
		if strings.Contains(err.Error(), "sudah tersedia") {
			log.Println("-------------------")
			countInt, errCount := strconv.Atoi(noPageStr)
			if errCount != nil {
				// error variabel tidak int
				log.Fatal(errCount)
			}
			countInt++
			if countInt < maxPage {
				noPageStr = strconv.Itoa(countInt)
				log.Println("count ke page " + noPageStr)
				time.Sleep(1 * time.Second)
				goto restart
				// goto restart
			} else {
				log.Println("SUCESS ambil max page")
			}
		} else {
			log.Println("eror insert db code : " + err.Error() + " url:" + urlTarget)
			goto restart
		}

	}
	countInt, errCount := strconv.Atoi(noPageStr)
	if errCount != nil {
		log.Fatal(errCount)
	}
	countInt++
	if countInt < maxPage {
		noPageStr = strconv.Itoa(countInt)
		log.Println("count ke page " + noPageStr)
		time.Sleep(1 * time.Second)
		goto restart
		// goto restart
	} else {
		// program selesai
		log.Println("SUCESS ambil max page")
	}
	_ = maxPage
	_ = tempCount
}
