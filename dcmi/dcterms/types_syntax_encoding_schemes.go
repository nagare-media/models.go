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

package dcterms

import "github.com/nagare-media/models.go/dcmi/dc"

const (
	BoxXSIType      = "dcterms:Box"
	ISO3166XSIType  = "dcterms:ISO3166"
	ISO639_2XSIType = "dcterms:ISO639-2"
	ISO639_3XSIType = "dcterms:ISO639-3"
	PeriodXSIType   = "dcterms:Period"
	PointXSIType    = "dcterms:Point"
	RFC1766XSIType  = "dcterms:RFC1766"
	RFC3066XSIType  = "dcterms:RFC3066"
	RFC4646XSIType  = "dcterms:RFC4646"
	RFC5646XSIType  = "dcterms:RFC5646"
	URIXSIType      = "dcterms:URI"
	W3CDTFXSIType   = "dcterms:W3CDTF"
)

// The set of regions in space defined by their geographic coordinates according to the DCMI Box Encoding Scheme.
//
// See https://www.dublincore.org/specifications/dublin-core/dcmi-box/
// See http://purl.org/dc/terms/Box
type Box dc.SimpleLiteral

// The set of codes listed in ISO 3166-1 for the representation of names of countries.
//
// See https://www.iso.org/obp/ui/#search
// See http://purl.org/dc/terms/ISO3166
type ISO3166 dc.SimpleLiteral

// The three-letter alphabetic codes listed in ISO639-2 for the representation of names of languages.
//
// See http://lcweb.loc.gov/standards/iso639-2/langhome.html
// See http://purl.org/dc/terms/ISO639_2
type ISO639_2 dc.SimpleLiteral

// The set of three-letter codes listed in ISO 639-3 for the representation of names of languages.
//
// http://www.sil.org/iso639-3/
// See http://purl.org/dc/terms/ISO639_3
type ISO639_3 dc.SimpleLiteral

// The set of time intervals defined by their limits according to the DCMI Period Encoding Scheme.
//
// https://www.dublincore.org/specifications/dublin-core/dcmi-period/
// See http://purl.org/dc/terms/Period
type Period dc.SimpleLiteral

// The set of points in space defined by their geographic coordinates according to the DCMI Point Encoding Scheme.
//
// See https://www.dublincore.org/specifications/dublin-core/dcmi-point/
// See http://purl.org/dc/terms/Point
type Point dc.SimpleLiteral

// The set of tags, constructed according to RFC 1766, for the identification of languages.
//
// See http://www.ietf.org/rfc/rfc1766.txt
// See http://purl.org/dc/terms/RFC1766
type RFC1766 dc.SimpleLiteral

// The set of tags constructed according to RFC 3066 for the identification of languages.
//
// RFC 3066 has been obsoleted by RFC 4646.
//
// See http://www.ietf.org/rfc/rfc3066.txt
// See http://purl.org/dc/terms/RFC3066
type RFC3066 dc.SimpleLiteral

// The set of tags constructed according to RFC 4646 for the identification of languages.
//
// RFC 4646 obsoletes RFC 3066.
//
// See http://www.ietf.org/rfc/rfc4646.txt
// See http://purl.org/dc/terms/RFC4646
type RFC4646 dc.SimpleLiteral

// The set of tags constructed according to RFC 5646 for the identification of languages.
//
// RFC 5646 obsoletes RFC 4646.
//
// See http://www.ietf.org/rfc/rfc4646.txt
// See http://purl.org/dc/terms/RFC5646
//
// TODO: this is not specified in the XML schema
type RFC5646 dc.SimpleLiteral

// The set of identifiers constructed according to the generic syntax for Uniform Resource Identifiers as specified by
// the Internet Engineering Task Force.
//
// See http://www.ietf.org/rfc/rfc3986.txt
// See http://purl.org/dc/terms/URI
type URI dc.SimpleLiteral

// The set of dates and times constructed according to the W3C Date and Time Formats Specification.
//
// See http://www.w3.org/TR/NOTE-datetime
// See http://purl.org/dc/terms/W3CDTF
type W3CDTF dc.SimpleLiteral
