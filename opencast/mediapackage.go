/*
Copyright 2021 The nagare media authors

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

package opencast

import (
	"time"

	"github.com/nagare-media/models.go/x/encoding/xml"
)

type MediaPackage struct {
	XMLName xml.Name `xml:"http://mediapackage.opencastproject.org mediapackage"`

	ID       string     `xml:"id,attr"`
	Start    *time.Time `xml:"start,attr,omitempty"`
	Duration int64      `xml:"duration,attr,omitempty"`

	Title        string        `xml:"http://mediapackage.opencastproject.org title,omitempty"`
	SeriesTitle  string        `xml:"http://mediapackage.opencastproject.org seriestitle,omitempty"`
	Language     string        `xml:"http://mediapackage.opencastproject.org language,omitempty"`
	Series       string        `xml:"http://mediapackage.opencastproject.org series,omitempty"`
	License      string        `xml:"http://mediapackage.opencastproject.org license,omitempty"`
	Creators     []string      `xml:"http://mediapackage.opencastproject.org creators>creator,omitempty"`
	Contributors []string      `xml:"http://mediapackage.opencastproject.org contributors>contributor,omitempty"`
	Subjects     []string      `xml:"http://mediapackage.opencastproject.org subjects>subject,omitempty"`
	Media        []Track       `xml:"http://mediapackage.opencastproject.org media>track,omitempty"`
	Metadata     []Catalog     `xml:"http://mediapackage.opencastproject.org metadata>catalog,omitempty"`
	Attachments  []Attachment  `xml:"http://mediapackage.opencastproject.org attachments>attachment,omitempty"`
	Publications []Publication `xml:"http://mediapackage.opencastproject.org publications>publication,omitempty"`
}
