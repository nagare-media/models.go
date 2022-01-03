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
	"github.com/nagare-media/models.go/third_party/encoding/xml"
)

type Attachment struct {
	MediaPackageElement

	Properties Properties `xml:"http://mediapackage.opencastproject.org additionalProperties,omitempty"`
}

type Properties map[string]string

var (
	_ xml.Marshaler   = &Properties{}
	_ xml.Unmarshaler = &Properties{}
)

func (p *Properties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range *p {
		elem := xml.StartElement{
			Name: xml.Name{Local: "property"},
			Attr: []xml.Attr{{Name: xml.Name{Local: "key"}, Value: k}},
		}
		err = e.EncodeToken(elem)
		if err != nil {
			return err
		}

		err = e.EncodeToken(xml.CharData(v))
		if err != nil {
			return err
		}

		err = e.EncodeToken(xml.EndElement{Name: elem.Name})
		if err != nil {
			return err
		}
	}

	err = e.EncodeToken(xml.EndElement{Name: start.Name})
	if err != nil {
		return err
	}

	return e.Flush()
}

func (p *Properties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var xmlProperties struct {
		Properties []struct {
			Key   string `xml:"key,attr,omitempty"`
			Value string `xml:",chardata"`
		} `xml:"http://mediapackage.opencastproject.org property,omitempty"`
	}

	err = d.DecodeElement(&xmlProperties, &start)
	if err != nil {
		return
	}

	*p = make(map[string]string, len(xmlProperties.Properties))
	for _, entry := range xmlProperties.Properties {
		(*p)[entry.Key] = entry.Value
	}

	return
}
