/*
Copyright 2021-2025 The nagare media authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
