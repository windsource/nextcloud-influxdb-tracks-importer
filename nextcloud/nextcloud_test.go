package nextcloud

import (
	"reflect"
	"testing"
	"time"
)

func Test_dateFromTrackFilename(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"Valid filename", args{"2020-06-09.gpx"}, time.Date(2020, 6, 9, 0, 0, 0, 0, time.UTC)},
		{"Missing extension", args{"2020-06-09"}, time.Time{}},
		{"Invalid name 1", args{"Helloworld.gpx"}, time.Time{}},
		{"Invalid name 2", args{"Hello world.txt"}, time.Time{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dateFromTrackFilename(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dateFromTrackFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trackFilenamefromDate(t *testing.T) {
	type args struct {
		year  int
		month time.Month
		day   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Good case 1", args{2020, 11, 12}, "2020-11-12.gpx"},
		{"Good case 2", args{2020, 1, 2}, "2020-01-02.gpx"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trackFilenamefromDate(tt.args.year, tt.args.month, tt.args.day); got != tt.want {
				t.Errorf("trackFilenamefromDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
