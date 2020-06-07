package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	influxClient "github.com/influxdata/influxdb1-client/v2"
	"github.com/windsource/nextcloud-influxdb-tracks-importer/date"
	"github.com/windsource/nextcloud-influxdb-tracks-importer/gpx"
)

type DbReader struct {
	client      influxClient.Client
	dbName      string
	measurement string
	user        string
}

type Position int

const (
	FIRST Position = iota
	LAST
)

// NewDbReader create a new database for read operation on host (e.g. "http://localhost:8086")
func NewDbReader(host string, dbName string, measurement string, user string) (*DbReader, error) {
	c, err := influxClient.NewHTTPClient(influxClient.HTTPConfig{
		Addr: host,
	})
	return &DbReader{c, dbName, measurement, user}, err
}

func (d *DbReader) GetBorderTimestamp(pos Position) (time.Time, error) {
	sortOrder := "ASC"
	if pos == LAST {
		sortOrder = "DESC"
	}
	queryString := fmt.Sprintf("SELECT time, _type FROM %s WHERE _type='location' AND \"user\"='%s' ORDER BY time %s LIMIT 1", d.measurement, d.user, sortOrder)
	q := influxClient.NewQuery(queryString, d.dbName, "")
	response, err := d.client.Query(q)
	if err != nil {
		return time.Time{}, err
	}
	if response.Error() != nil {
		return time.Time{}, response.Error()
	}
	if len(response.Results) == 0 || len(response.Results[0].Series) == 0 {
		return time.Time{}, errors.New("No results for query")
	}
	// log.Println(response.Results[0].Series[0].Values[0][0].(string))
	timestamp, err := time.Parse(time.RFC3339, response.Results[0].Series[0].Values[0][0].(string))
	return timestamp, err
}

func (d *DbReader) GetDataOfDay(day date.Date) ([]influxClient.Result, error) {
	timeString := day.ToTime().Format(time.RFC3339)
	queryString := fmt.Sprintf("SELECT time, alt, lat, lon FROM %s "+
		"WHERE _type='location' AND \"user\"='%s' AND time >= '%s' AND time < '%s' + 1d "+
		"ORDER BY time ASC",
		d.measurement, d.user, timeString, timeString)
	q := influxClient.NewQuery(queryString, d.dbName, "")
	response, err := d.client.Query(q)
	if err != nil {
		return nil, err
	}
	if response.Error() != nil {
		return nil, response.Error()
	}
	return response.Results, nil
}

func (d *DbReader) GetGpxOfDay(year int, month time.Month, day int) (*gpx.GpxDoc, error) {
	date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	timeString := date.Format(time.RFC3339)
	queryString := fmt.Sprintf("SELECT time, alt, lat, lon FROM %s "+
		"WHERE _type='location' AND \"user\"='%s' AND time >= '%s' AND time < '%s' + 1d "+
		"ORDER BY time ASC",
		d.measurement, d.user, timeString, timeString)
	q := influxClient.NewQuery(queryString, d.dbName, "")
	response, err := d.client.Query(q)
	if err != nil {
		return nil, err
	}
	if response.Error() != nil {
		return nil, response.Error()
	}
	name := date.Format("2006-01-02")
	gpxDoc := gpx.NewGpxDocument(name)

	if len(response.Results) > 0 && len(response.Results[0].Series) > 0 {
		trkpts := make([]gpx.Trkpt, 0, len(response.Results[0].Series[0].Values))
		for _, v := range response.Results[0].Series[0].Values {
			timestamp := v[0].(string)
			ele, _ := v[1].(json.Number).Float64()
			lat, _ := v[2].(json.Number).Float64()
			lon, _ := v[3].(json.Number).Float64()

			trkpt := gpx.Trkpt{Time: timestamp, Ele: ele, Lat: lat, Lon: lon}
			trkpts = append(trkpts, trkpt)
		}

		gpxDoc.AddTrackpoints(trkpts)
	}
	return gpxDoc, nil
}

func (d *DbReader) Close() {
	d.client.Close()
}
