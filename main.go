package main

import (
	"log"
	"os"
	"time"

	"github.com/windsource/nextcloud-influxdb-tracks-importer/date"
	"github.com/windsource/nextcloud-influxdb-tracks-importer/db"

	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	host := "http://localhost:8080"
	dbName := "owntracks"
	measurement := "owntracks"
	user := "holger"

	theDb, err := db.NewDbReader(host, dbName, measurement, user)
	if err != nil {
		log.Fatal(err)
	}

	first, err := theDb.GetBorderTimestamp(db.FIRST)
	if err != nil {
		log.Fatal(err)
	}

	// last, err := theDb.GetBorderTimestamp(db.LAST)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	first = time.Date(first.Year(), first.Month(), first.Day(), 0, 0, 0, 0, first.Location())

	data, err := theDb.GetDataOfDay(date.FromTime(first))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", data)

}
