package main

import (
	"log"
	"os"

	"github.com/edot92/telkomscraplokerindo/db"
	"github.com/edot92/telkomscraplokerindo/scrap"
)

func checkError(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
func main() {
	// to change the flags on the default logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	db.InitDB()
	scrap.ScrapLockerID()
}
