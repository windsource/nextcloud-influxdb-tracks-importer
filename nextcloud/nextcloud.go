package nextcloud

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"time"

	"log"

	"github.com/studio-b12/gowebdav"
)

type Nextcloud struct {
	webdavClient *gowebdav.Client
	httpClient   *http.Client
	uri          string
	trackDir     string
	user         string
	password     string
}

func New(uri string, user string, password string, trackDir string) *Nextcloud {
	webdavClient := gowebdav.NewClient(uri, user, password)
	httpClient := http.Client{}
	return &Nextcloud{webdavClient, &httpClient, uri, trackDir, user, password}
}

func trackFilenamefromDate(year int, month time.Month, day int) string {
	return fmt.Sprintf("%04d-%02d-%02d.gpx", year, month, day)
}

func dateFromTrackFilename(filename string) time.Time {
	re := regexp.MustCompile(`(\d{4}-\d{2}-\d{2})\.gpx`)
	s := re.FindStringSubmatch(filename)
	if len(s) == 2 {
		t, err := time.Parse("2006-01-02", s[1])
		if err == nil {
			return t
		}
	}
	log.Printf("Could not parse filname %s\n", filename)
	return time.Time{}
}

func (nc *Nextcloud) GetLatestTrackDate() (time.Time, error) {
	var result time.Time
	files, err := nc.webdavClient.ReadDir(nc.trackDir)
	if err != nil {
		return result, err
	}
	for _, file := range files {
		t := dateFromTrackFilename(file.Name())
		if t.After(result) {
			result = t
		}
	}
	return result, nil
}

func (nc *Nextcloud) StoreTrack(year int, month time.Month, day int, contents []byte) error {
	filePath := path.Join(nc.trackDir, trackFilenamefromDate(year, month, day))
	log.Printf("Storing file %s\n", filePath)

	// The webdav client does not set the content-length and thus Storageshare creates an
	// empt file. So we have to use our own upload method
	// return nc.client.Write(filePath, contents, 0644)

	url, err := url.Parse(nc.uri)
	if err != nil {
		return err
	}
	url.Path = path.Join(url.Path, filePath)
	return uploadFile(nc.httpClient, nc.user, nc.password, url.String(), contents)
}

func uploadFile(httpClient *http.Client, user, password, url string, contents []byte) error {
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(contents))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/gpx+xml; charset=utf-8")
	req.SetBasicAuth(user, password)
	_, err = httpClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}
