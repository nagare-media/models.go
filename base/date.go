/*
Copyright 2021-2024 The nagare media authors

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
	"errors"
	"time"
)

const DateFormat = "2006-01-02"

type Date struct {
	time.Time
}

var (
	_ encoding.TextMarshaler   = &Date{}
	_ encoding.TextUnmarshaler = &Date{}
)

func (d Date) MarshalJSON() ([]byte, error) {
	if y := d.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Date.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateFormat)+2)
	b = append(b, '"')
	b = d.AppendFormat(b, DateFormat)
	b = append(b, '"')
	return b, nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	t, err := time.Parse(`"`+DateFormat+`"`, string(data))
	*d = Date{Time: t}
	return err
}

func (d Date) MarshalText() (text []byte, err error) {
	if y := d.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Date.MarshalText: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateFormat))
	return d.AppendFormat(b, DateFormat), nil
}

func (d *Date) UnmarshalText(text []byte) error {
	t, err := time.Parse(DateFormat, string(text))
	*d = Date{Time: t}
	return err
}
