package db

import (
	"errors"
	"log"

	"github.com/globalsign/mgo/bson"
)

type Data struct {
	Waktu string `bson:"waktu"`
	Post  string `bson:"post"`
	Raw   string `bson:"raw"`
}

// StructDataLocker ...
type StructDataLocker struct {
	URL  string `bson:"url"`
	Data []struct {
		Waktu string `bson:"waktu"`
		Post  string `bson:"post"`
		Raw   string `bson:"raw"`
	} `bson:"data"`
}

var (
	ErrUrlSUdahTersedia error = errors.New("alamat url sudah tersedia")
)

// InsertDataLocker ...
func InsertDataLocker(data StructDataLocker) error {
	// SessionMgo.
	sesMgo := DBCopyMGO()
	defer sesMgo.Close()
	col := sesMgo.DB(DBNAME).C(ColLockerId)
	query := bson.M{
		"url": data.URL,
	}
	count, _ := col.Find(query).Count()
	if count > 0 {
		return ErrUrlSUdahTersedia
	}
	err := col.Insert(data)
	if err != nil {
		return err
	}
	log.Println("succes insert")
	return nil
}
