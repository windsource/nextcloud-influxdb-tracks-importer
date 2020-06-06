package date

import (
	"reflect"
	"testing"
	"time"
)

func TestDate_ToTime(t *testing.T) {
	type fields struct {
		year  int
		month time.Month
		day   int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"2019", fields{2019, 12, 31}, "2019-12-31T00:00:00.0Z"},
		{"2020", fields{2020, 6, 5}, "2020-06-05T00:00:00.0Z"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Date{
				year:  tt.fields.year,
				month: tt.fields.month,
				day:   tt.fields.day,
			}
			want, _ := time.Parse(time.RFC3339, tt.want)
			if got := d.ToTime(); !reflect.DeepEqual(got, want) {
				t.Errorf("Date.ToTime() = %v, want %v", got, want)
			}
		})
	}
}

func TestFromTime(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want Date
	}{
		{"2019", args{"2019-12-31T17:27:14.82725108Z"}, Date{2019, 12, 31}},
		{"2020", args{"2020-06-05T16:47:54.68934859Z"}, Date{2020, 06, 05}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tIn, _ := time.Parse(time.RFC3339, tt.args.t)
			if got := FromTime(tIn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_ToString(t *testing.T) {
	type fields struct {
		year  int
		month time.Month
		day   int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Arbitrary date single",
			fields{2020, 6, 5},
			"2020-06-05",
		},
		{
			"Arbitrary date double",
			fields{2020, 11, 12},
			"2020-11-12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Date{
				year:  tt.fields.year,
				month: tt.fields.month,
				day:   tt.fields.day,
			}
			if got := d.ToString(); got != tt.want {
				t.Errorf("Date.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
