package main

import (
	"log"
	"os"
	"time"

	"github.com/windsource/nextcloud-influxdb-tracks-importer/db"
	nc "github.com/windsource/nextcloud-influxdb-tracks-importer/nextcloud"

	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
)

func init() {
	log.SetOutput(os.Stdout)
}

func dateOnly(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

func getOptionalParameter(paramName string, defaultValue string) string {
	if s := os.Getenv(paramName); s != "" {
		return s
	}
	return defaultValue
}

func getMandatoryParameter(paramName string) string {
	if s := os.Getenv(paramName); s != "" {
		return s
	}
	log.Fatalf("Parameter %s not set!", paramName)
	return "" // to satisfy lint
}

func main() {

	influxdbHost := getOptionalParameter("INFLUXDB_HOST", "http://localhost:8080")
	dbName := getOptionalParameter("INFLUXDB_DB_NAME", "owntracks")
	measurement := getOptionalParameter("INFLUXDB_MEASUREMENT_NAME", "owntracks")
	owntracksUser := getOptionalParameter("OWNTRACKS_USER", "holger")
	nextcloudRoot := getMandatoryParameter("NEXTCLOUD_ROOT")
	nextcloudUser := getMandatoryParameter("NEXTCLOUD_USER")
	nextcloudPassword := getMandatoryParameter("NEXTCLOUD_PASSWORD")
	trackDir := getOptionalParameter("TRACKDIR", "/Tracks/owntracks/")

	ncClient := nc.New(nextcloudRoot, nextcloudUser, nextcloudPassword, trackDir)

	latestTrackDateFromNextcloud, err := ncClient.GetLatestTrackDate()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Latest track date from Nextcloud: %v", latestTrackDateFromNextcloud)

	theDb, err := db.NewDbReader(influxdbHost, dbName, measurement, owntracksUser)
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
	if latestTrackDateFromNextcloud.After(startDate) {
		startDate = latestTrackDateFromNextcloud
	}

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

		err := ncClient.StoreTrack(t.Year(), t.Month(), t.Day(), []byte(gpxDoc.ToXml()))
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("Done")
}
