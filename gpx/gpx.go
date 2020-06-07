package gpx

import (
	"encoding/xml"
	"log"
)

type (
	GpxDoc struct {
		XMLName           xml.Name `xml:"http://www.topografix.com/GPX/1/1 gpx"`
		Version           string   `xml:"version,attr"`
		Creator           string   `xml:"creator,attr"`
		XmlnsXsi          string   `xml:"xmlns:xsi,attr"`
		XsiSchemaLocation string   `xml:"xsi:schemaLocation,attr"`

		Metadata Metadata
		Trk      Trk
	}

	Metadata struct {
		XMLName xml.Name `xml:"metadata"`
		Name    string   `xml:"name"`
	}

	Trk struct {
		XMLName xml.Name `xml:"trk"`
		Name    string   `xml:"name"`
		TrkSeg  TrkSeg
	}

	TrkSeg struct {
		XMLName xml.Name `xml:"trkseg"`
		Trkpts  []Trkpt
	}

	Trkpt struct {
		XMLName xml.Name `xml:"trkpt"`
		Lat     float64  `xml:"lat,attr"`
		Lon     float64  `xml:"lon,attr"`
		Ele     float64  `xml:"ele"`
		Time    string   `xml:"time"`
	}
)

func NewGpxDocument(name string) *GpxDoc {
	return &GpxDoc{
		Version:           "1.1",
		Creator:           "windsource",
		XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
		XsiSchemaLocation: "http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd",
		Metadata:          Metadata{Name: name},
		Trk:               Trk{Name: name},
	}
}

func (d *GpxDoc) AddTrackpoints(trkpts []Trkpt) {
	d.Trk.TrkSeg.Trkpts = trkpts
}

func (d *GpxDoc) IsEmpty() bool {
	return len(d.Trk.TrkSeg.Trkpts) <= 1
}

func (d *GpxDoc) ToXml() string {
	output, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	return xml.Header + string(output)
}
