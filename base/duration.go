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
	"bytes"
	"encoding"
	"encoding/json"
	"text/template"

	"github.com/senseyeio/duration"
)

// TODO: re-implement as github.com/senseyeio/duration does not support subseconds, start time, etc.

// See duration.tmpl
var TemplateISO8601 = template.Must(template.New("duration").Parse(`P{{if .Y}}{{.Y}}Y{{end}}{{if .M}}{{.M}}M{{end}}{{if .W}}{{.W}}W{{end}}{{if .D}}{{.D}}D{{end}}{{if .HasTimePart}}T{{end }}{{if .TH}}{{.TH}}H{{end}}{{if .TM}}{{.TM}}M{{end}}{{if .TS}}{{.TS}}S{{end}}`))

type Duration struct {
	duration.Duration
}

var (
	_ encoding.TextMarshaler   = &Duration{}
	_ encoding.TextUnmarshaler = &Duration{}
	_ json.Marshaler           = &Duration{}
	_ json.Unmarshaler         = &Duration{}
)

func (d Duration) MarshalText() ([]byte, error) {
	if d.IsZero() {
		return []byte("P0D"), nil
	}

	var s bytes.Buffer
	err := TemplateISO8601.Execute(&s, d)
	if err != nil {
		return nil, err
	}

	return s.Bytes(), nil
}

func (d *Duration) UnmarshalText(text []byte) error {
	d0, err := duration.ParseISO8601(string(text))
	if err != nil {
		return err
	}
	*d = Duration{d0}
	return nil
}

func (d Duration) MarshalJSON() ([]byte, error) {
	s, err := d.MarshalText()
	if err != nil {
		return nil, err
	}
	return json.Marshal(s)
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var s []byte
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return d.UnmarshalText(s)
}
