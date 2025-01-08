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

package dc

import "github.com/nagare-media/models.go/third_party/encoding/xml"

const (
	SchemaVersion = "2008-02-11"

	XMLNS       = "http://purl.org/dc/elements/1.1/"
	XMLNSPrefix = "dc"
)

// This is the default type for all of the DC elements. It permits text content only with optional xml:lang attribute.
// Text is allowed because mixed="true", but sub-elements are disallowed because minOccurs="0" and maxOccurs="0" are on
// the xs:any tag.
//
// This complexType allows for restriction or extension permitting child elements.
type SimpleLiteral struct {
	Value   string `xml:",chardata" json:"#value,omitempty"`
	XMLLang string `xml:"http://www.w3.org/XML/1998/namespace xml:lang,attr,omitempty" json:"@xml:lang,omitempty"`
	XSIType string `xml:"http://www.w3.org/2001/XMLSchema-instance xsi:type,attr,omitempty" json:"@xsi:type,omitempty"`
}

type Any SimpleLiteral

// A name given to the resource.
//
// See http://purl.org/dc/elements/1.1/title
type Title Any

// An entity primarily responsible for making the resource.
//
// Examples of a Creator include a person, an organization, or a service. Typically, the name of a Creator should be
// used to indicate the entity.
//
// See http://purl.org/dc/elements/1.1/creator
type Creator Any

// The topic of the resource.
//
// Typically, the subject will be represented using keywords, key phrases, or classification codes. Recommended best
// practice is to use a controlled vocabulary.
//
// See http://purl.org/dc/elements/1.1/subject
type Subject Any

// An account of the resource.
//
// Description may include but is not limited to: an abstract, a table of contents, a graphical representation, or a
// free-text account of the resource.
//
// See http://purl.org/dc/elements/1.1/description
type Description Any

// An entity responsible for making the resource available.
//
// Examples of a Publisher include a person, an organization, or a service. Typically, the name of a Publisher should be
// used to indicate the entity.
//
// See http://purl.org/dc/elements/1.1/publisher
type Publisher Any

// An entity responsible for making contributions to the resource.
//
// The guidelines for using names of persons or organizations as creators also apply to contributors. Typically, the
// name of a Contributor should be used to indicate the entity.
//
// See http://purl.org/dc/elements/1.1/contributor
type Contributor Any

// A point or period of time associated with an event in the lifecycle of the resource.
//
// Date may be used to express temporal information at any level of granularity. Recommended practice is to express the
// date, date/time, or period of time according to ISO 8601-1 [ISO 8601-1] or a published profile of the ISO standard,
// such as the W3C Note on Date and Time Formats [W3CDTF] or the Extended Date/Time Format Specification [EDTF]. If the
// full date is unknown, month and year (YYYY-MM) or just year (YYYY) may be used. Date ranges may be specified using
// ISO 8601 period of time specification in which start and end dates are separated by a '/' (slash) character. Either
// the start or end date may be missing.
//
// See http://purl.org/dc/elements/1.1/date
type Date Any

// The nature or genre of the resource.
//
// Recommended practice is to use a controlled vocabulary such as the DCMI Type Vocabulary [DCMI-TYPE]. To describe the
// file format, physical medium, or dimensions of the resource, use the Format element.
//
// See http://purl.org/dc/elements/1.1/type
type Type Any

// The file format, physical medium, or dimensions of the resource.
//
// Recommended practice is to use a controlled vocabulary where available. For example, for file formats one could use
// the list of Internet Media Types [MIME].
//
// See http://purl.org/dc/elements/1.1/format
type Format Any

// An unambiguous reference to the resource within a given context.
//
// Recommended practice is to identify the resource by means of a string conforming to an identification system.
//
// See http://purl.org/dc/elements/1.1/identifier
type Identifier Any

// A related resource from which the described resource is derived.
//
// The described resource may be derived from the related resource in whole or in part. Recommended best practice is to
// identify the related resource by means of a string conforming to a formal identification system.
//
// See http://purl.org/dc/elements/1.1/source
type Source Any

// A language of the resource.
//
// Recommended practice is to use either a non-literal value representing a language from a controlled vocabulary such
// as ISO 639-2 or ISO 639-3, or a literal value consisting of an IETF Best Current Practice 47 [IETF-BCP47] language
// tag.
//
// See http://purl.org/dc/elements/1.1/language
type Language Any

// A related resource.
//
// Recommended practice is to identify the related resource by means of a URI. If this is not possible or feasible, a
// string conforming to a formal identification system may be provided.
//
// See http://purl.org/dc/elements/1.1/relation
type Relation Any

// The spatial or temporal topic of the resource, spatial applicability of the resource, or jurisdiction under which the
// resource is relevant.
//
// Spatial topic and spatial applicability may be a named place or a location specified by its geographic coordinates.
// Temporal topic may be a named period, date, or date range. A jurisdiction may be a named administrative entity or a
// geographic place to which the resource applies. Recommended practice is to use a controlled vocabulary such as the
// Getty Thesaurus of Geographic Names [TGN]. Where appropriate, named places or time periods may be used in preference
// to numeric identifiers such as sets of coordinates or date ranges.
//
// See http://purl.org/dc/elements/1.1/coverage
type Coverage Any

// Information about rights held in and over the resource.
//
// Typically, rights information includes a statement about various property rights associated with the resource,
// including intellectual property rights.
//
// See http://purl.org/dc/elements/1.1/rights
type Rights Any

// This complexType is included as a convenience for schema authors who need to define a root or container element for
// all of the DC elements.
type Elements struct {
	XMLName xml.Name `json:"-"`

	Title       []Title       `xml:"http://purl.org/dc/elements/1.1/ dc:title,omitempty"       json:"dc:title,omitempty"`
	Creator     []Creator     `xml:"http://purl.org/dc/elements/1.1/ dc:creator,omitempty"     json:"dc:creator,omitempty"`
	Subject     []Subject     `xml:"http://purl.org/dc/elements/1.1/ dc:subject,omitempty"     json:"dc:subject,omitempty"`
	Description []Description `xml:"http://purl.org/dc/elements/1.1/ dc:description,omitempty" json:"dc:description,omitempty"`
	Publisher   []Publisher   `xml:"http://purl.org/dc/elements/1.1/ dc:publisher,omitempty"   json:"dc:publisher,omitempty"`
	Contributor []Contributor `xml:"http://purl.org/dc/elements/1.1/ dc:contributor,omitempty" json:"dc:contributor,omitempty"`
	Date        []Date        `xml:"http://purl.org/dc/elements/1.1/ dc:date,omitempty"        json:"dc:date,omitempty"`
	Type        []Type        `xml:"http://purl.org/dc/elements/1.1/ dc:type,omitempty"        json:"dc:type,omitempty"`
	Format      []Format      `xml:"http://purl.org/dc/elements/1.1/ dc:format,omitempty"      json:"dc:format,omitempty"`
	Identifier  []Identifier  `xml:"http://purl.org/dc/elements/1.1/ dc:identifier,omitempty"  json:"dc:identifier,omitempty"`
	Source      []Source      `xml:"http://purl.org/dc/elements/1.1/ dc:source,omitempty"      json:"dc:source,omitempty"`
	Language    []Language    `xml:"http://purl.org/dc/elements/1.1/ dc:language,omitempty"    json:"dc:language,omitempty"`
	Relation    []Relation    `xml:"http://purl.org/dc/elements/1.1/ dc:relation,omitempty"    json:"dc:relation,omitempty"`
	Coverage    []Coverage    `xml:"http://purl.org/dc/elements/1.1/ dc:coverage,omitempty"    json:"dc:coverage,omitempty"`
	Rights      []Rights      `xml:"http://purl.org/dc/elements/1.1/ dc:rights,omitempty"      json:"dc:rights,omitempty"`
}
