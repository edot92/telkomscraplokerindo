package main

// https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-lang-analyzer.html#indonesian-analyzer
// - cleaning (membersihkan link html)
// - Dibuat lowercase
// - Tokenize (dipotong potong)
// - Stopword (hilangkan kata yg ga penting)
// - Stemming (menghilangkan imbuhan)
// - Pembobotan kata TF-IDF
// - Part of speech ( POS Tagger)
// https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-lang-analyzer.html#indonesian-analyzer
// kawamu@fxprix.com
// jakarta92
import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/edot92/telkomscraplokerindo/scrap"
)

func checkError(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}
func main() {
	// scrap.Demo()
	// return
	// scrap.InitElastic()

	// scrap.PStopWord("saya test")
	// return
	// to change the flags on the default logger
	dat, err := ioutil.ReadFile("stopwords.txt")
	checkError(err)
	tokensizerArray := strings.Split(string(dat), "\r")
	// db.InitDB()
	// scrap.ScrapLockerID2()
	txtbefore := ``

	// log.Println(txtbefore)
	log.Println("--------------------")
	// - cleaning (membersihkan link html)
	// - Dibuat lowercase
	// Tokenize (dipotong potong)
	resp, err := scrap.P1Cleaning(txtbefore)
	if err != nil {
		log.Fatal(err)
	}
	// text1 := (resp.Tokens[0].Token)
	kataProses1 := ""
	for index := 0; index < len(resp.Tokens); index++ {
		text1 := (resp.Tokens[index].Token)
		kataProses1 = kataProses1 + text1 + " "
	}
	log.Printf("banyak array tokenize: %d", len(resp.Tokens))
	scrap.Demo2(kataProses1)
	// LANJUT KE Stopword

	_ = err
	_ = tokensizerArray
}
