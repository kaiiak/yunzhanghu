package yunzhanghu

import (
	"errors"
	"time"
)

type (
	Time struct{ time.Time }
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

var (
	ShangHaiTimeLocation, _ = time.LoadLocation("Asia/Shanghai")
)

func (t *Time) UnmarshalJson(data []byte) (err error) {
	if string(data) == "null" {
		return nil
	}
	t.Time, err = time.Parse(timeFormat, string(data))
	if err != nil {
		return
	}
	t.Time = time.Date(t.Time.Year(), t.Time.Month(), t.Time.Day(), t.Time.Hour(), t.Time.Minute(), t.Time.Second(), t.Time.Nanosecond(), ShangHaiTimeLocation)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}
