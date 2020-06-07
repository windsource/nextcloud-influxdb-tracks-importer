package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/windsource/nextcloud-influxdb-tracks-importer/db"

	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
)

func init() {
	log.SetOutput(os.Stdout)
}

func dateOnly(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
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

	last, err := theDb.GetBorderTimestamp(db.LAST)
	if err != nil {
		log.Fatal(err)
	}

	startDate := dateOnly(first)

	for t := startDate; t.Before(last); t = t.AddDate(0, 0, 1) {
		dateString := t.Format("2006-01-02")
		log.Printf("Processing %s...", dateString)
		gpxDoc, _ := theDb.GetGpxOfDay(t.Date())
		if err != nil {
			log.Println(err)
			continue
		}
		if gpxDoc.IsEmpty() {
			log.Printf("No data for %s. Skipping date.\n", dateString)
			continue
		}

		filename := fmt.Sprintf("data/%s.gpx", dateString)
		f, err := os.Create(filename)
		if err != nil {
			log.Println(err)
			continue
		}
		_, _ = f.WriteString(gpxDoc.ToXml())
		f.Close()
	}

}
