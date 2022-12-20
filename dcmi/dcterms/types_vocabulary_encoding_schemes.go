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

import "github.com/nagare-media/models.go/dc"

const (
	DCMITypeXSIType = "dcterms:DCMIType"
	DDCXSIType      = "dcterms:DDC"
	IMTXSIType      = "dcterms:IMT"
	LCCXSIType      = "dcterms:LCC"
	LCSHXSIType     = "dcterms:LCSH"
	MESHXSIType     = "dcterms:MESH"
	NLMXSIType      = "dcterms:NLM"
	TGNXSIType      = "dcterms:TGN"
	UDCXSIType      = "dcterms:UDC"
)

// The set of classes specified by the DCMI Type Vocabulary, used to categorize the nature or genre of the resource.
//
// See http://purl.org/dc/dcmitype/
// See http://purl.org/dc/terms/DCMIType
type DCMIType dc.SimpleLiteral

// The set of conceptual resources specified by the Dewey Decimal Classification.
//
// See http://www.oclc.org/dewey/
// See http://purl.org/dc/terms/DDC
type DDC dc.SimpleLiteral

// The set of media types specified by the Internet Assigned Numbers Authority.
//
// See http://www.iana.org/assignments/media-types/
// See http://purl.org/dc/terms/IMT
type IMT dc.SimpleLiteral

// The set of conceptual resources specified by the Library of Congress Classification.
//
// See http://lcweb.loc.gov/catdir/cpso/lcco/lcco.html
// See http://purl.org/dc/terms/LCC
type LCC dc.SimpleLiteral

// The set of labeled concepts specified by the Library of Congress Subject Headings.
//
// See http://purl.org/dc/terms/LCSH
type LCSH dc.SimpleLiteral

// The set of labeled concepts specified by the Medical Subject Headings.
//
// See http://www.nlm.nih.gov/mesh/meshhome.html
// See http://purl.org/dc/terms/MESH
type MESH dc.SimpleLiteral

// The set of conceptual resources specified by the National Library of Medicine Classification.
//
// See http://wwwcf.nlm.nih.gov/class/
// See http://purl.org/dc/terms/NLM
//
// TODO: this is not specified in the XML schema
type NLM dc.SimpleLiteral

// The set of places specified by the Getty Thesaurus of Geographic Names.
//
// See http://www.getty.edu/research/tools/vocabulary/tgn/index.html
// See http://purl.org/dc/terms/TGN
type TGN dc.SimpleLiteral

// The set of conceptual resources specified by the Universal Decimal Classification.
//
// See http://www.udcc.org/
// See http://purl.org/dc/terms/UDC
type UDC dc.SimpleLiteral
