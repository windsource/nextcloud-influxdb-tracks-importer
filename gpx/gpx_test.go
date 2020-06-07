package gpx

import (
	"testing"
)

func TestGpxDocument_ToXml(t *testing.T) {
	want := `<?xml version="1.0" encoding="UTF-8"?>
<gpx xmlns="http://www.topografix.com/GPX/1/1" version="1.1" creator="windsource" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd">
  <metadata>
    <name>The name</name>
  </metadata>
  <trk>
    <name>The name</name>
    <trkseg>
      <trkpt lat="49.445237" lon="11.092556">
        <ele>310.067861</ele>
        <time>2020-05-31T10:28:48Z</time>
      </trkpt>
      <trkpt lat="49.445162" lon="11.092667">
        <ele>310.067861</ele>
        <time>2020-05-31T10:28:55Z</time>
      </trkpt>
    </trkseg>
  </trk>
</gpx>`

	d := NewGpxDocument("The name")
	var trkpts []Trkpt
	trkpts = append(trkpts, *NewTrackpoint(49.445237, 11.092556, 310.067861, "2020-05-31T10:28:48.756Z"))
	trkpts = append(trkpts, *NewTrackpoint(49.445162, 11.092667, 310.067861, "2020-05-31T10:28:55.756Z"))
	d.SetTrackpoints(trkpts)

	got := d.ToXml()

	t.Run("GPX file creation", func(t *testing.T) {
		if got != want {
			t.Errorf("GpxDocument.ToXml() = %v, want %v", got, want)
		}
	})
}

func TestGpxDoc_IsEmpty(t *testing.T) {
	d := NewGpxDocument("some name")

	t.Run("Empty document", func(t *testing.T) {
		if !d.IsEmpty() {
			t.Error()
		}
	})

	var trkpts []Trkpt
	trkpts = append(trkpts, *NewTrackpoint(49.445237, 11.092556, 310.067861, "2020-05-31T10:28:48.756Z"))
	d.SetTrackpoints(trkpts)
	t.Run("One trackpoint", func(t *testing.T) {
		if !d.IsEmpty() {
			t.Error()
		}
	})

	trkpts = make([]Trkpt, 0)
	trkpts = append(trkpts, *NewTrackpoint(49.445237, 11.092556, 310.067861, "2020-05-31T10:28:48.756Z"))
	trkpts = append(trkpts, *NewTrackpoint(49.445162, 11.092667, 310.067861, "2020-05-31T10:28:55.756Z"))
	d.SetTrackpoints(trkpts)
	t.Run("One trackpoint", func(t *testing.T) {
		if d.IsEmpty() {
			t.Error()
		}
	})
}
