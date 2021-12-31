package base

import (
	"encoding"
	"time"
)

const TimeFormat = "15:04:05Z07:00"

type Time struct {
	time.Time
}

var (
	_ encoding.TextMarshaler   = &Time{}
	_ encoding.TextUnmarshaler = &Time{}
)

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	t0, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = Time{Time: t0}
	return err
}

func (t Time) MarshalText() (text []byte, err error) {
	b := make([]byte, 0, len(TimeFormat))
	return t.AppendFormat(b, TimeFormat), nil
}

func (t *Time) UnmarshalText(text []byte) error {
	t0, err := time.Parse(TimeFormat, string(text))
	*t = Time{Time: t0}
	return err
}
