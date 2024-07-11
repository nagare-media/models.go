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

package opencast

import (
	"bytes"
	"encoding"
	"errors"
	"strings"

	"github.com/nagare-media/models.go/base"
)

type MediaPackageElement struct {
	ID        string                 `xml:"id,attr,omitempty"`
	Flavor    Flavor                 `xml:"type,attr,omitempty"`
	Reference *MediaPackageReference `xml:"ref,attr,omitempty"`

	MimeType MimeType  `xml:"http://mediapackage.opencastproject.org mimetype,omitempty"`
	URL      base.URI  `xml:"http://mediapackage.opencastproject.org url,omitempty"`
	Size     int64     `xml:"http://mediapackage.opencastproject.org size,omitempty"`
	Checksum *Checksum `xml:"http://mediapackage.opencastproject.org checksum,omitempty"`
	Tags     []string  `xml:"http://mediapackage.opencastproject.org tags>tag,omitempty"`
}

type Flavor string

type MimeType string

type ChecksumType string

const (
	MD5ChecksumType ChecksumType = "md5"
)

type Checksum struct {
	Value string       `xml:",chardata"`
	Type  ChecksumType `xml:"type,attr"`
}

type MediaPackageReferenceType string

const (
	MediaPackageMediaPackageReference MediaPackageReferenceType = "mediapackage"
	TrackMediaPackageReference        MediaPackageReferenceType = "track"
	CatalogMediaPackageReference      MediaPackageReferenceType = "catalog"
	AttachmentMediaPackageReference   MediaPackageReferenceType = "attachment"
	SeriesMediaPackageReference       MediaPackageReferenceType = "series"

	SelfID = "self"
	AnyID  = "*"
)

type MediaPackageReference struct {
	Type       MediaPackageReferenceType
	ID         string
	Properties map[string]string
}

var (
	_ encoding.TextMarshaler   = &MediaPackageReference{}
	_ encoding.TextUnmarshaler = &MediaPackageReference{}
)

func (ref *MediaPackageReference) MarshalText() (text []byte, err error) {
	if ref == nil {
		return nil, nil
	}

	buf := bytes.Buffer{}

	if ref.Type == MediaPackageMediaPackageReference && ref.ID == SelfID {
		buf.WriteString(SelfID)
	} else {
		buf.WriteString(string(ref.Type))
		buf.WriteString(":")
		buf.WriteString(ref.ID)
	}

	for k, v := range ref.Properties {
		buf.WriteString(";")
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(v)
	}

	text = buf.Bytes()
	return
}

func (ref *MediaPackageReference) UnmarshalText(text []byte) error {
	parts := strings.Split(string(text), ";")

	if parts[0] == SelfID {
		ref.Type = MediaPackageMediaPackageReference
		ref.ID = SelfID
	} else {
		refParts := strings.Split(parts[0], ":")
		if len(refParts) != 2 {
			return errors.New("MediaPackageReference.UnmarshalText: invalid reference format")
		}
		ref.Type = MediaPackageReferenceType(refParts[0])
		ref.ID = refParts[1]
	}

	ref.Properties = make(map[string]string, len(parts)-1)

	for i := 1; i < len(parts); i++ {
		propertyParts := strings.Split(parts[i], "=")
		if len(propertyParts) != 2 {
			return errors.New("MediaPackageReference.UnmarshalText: invalid reference property format")
		}
		ref.Properties[propertyParts[0]] = propertyParts[1]
	}

	return nil
}
