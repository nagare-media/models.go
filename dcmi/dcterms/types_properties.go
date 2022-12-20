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

import (
	"github.com/nagare-media/models.go/dc"
	"github.com/nagare-media/models.go/third_party/encoding/xml"
)

const (
	SchemaVersion = "2008-02-11"

	XMLNS       = "http://purl.org/dc/terms"
	XMLNSPrefix = "dcterms"
)

// A name given to the resource.
//
// See http://purl.org/dc/terms/title
type Title dc.Title

// An entity responsible for making the resource.
//
// Recommended practice is to identify the creator with a URI. If this is not possible or feasible, a literal value that
// identifies the creator may be provided.
//
// See http://purl.org/dc/terms/creator
type Creator dc.Creator

// A topic of the resource.
//
// Recommended practice is to refer to the subject with a URI. If this is not possible or feasible, a literal value that
// identifies the subject may be provided. Both should preferably refer to a subject in a controlled vocabulary.
//
// See http://purl.org/dc/terms/subject
type Subject dc.Subject

// An account of the resource.
//
// Description may include but is not limited to: an abstract, a table of contents, a graphical representation, or a
// free-text account of the resource.
//
// See http://purl.org/dc/terms/description
type Description dc.Description

// An entity responsible for making the resource available.
//
// See http://purl.org/dc/terms/publisher
type Publisher dc.Publisher

// An entity responsible for making contributions to the resource.
//
// The guidelines for using names of persons or organizations as creators apply to contributors.
//
// See http://purl.org/dc/terms/contributor
type Contributor dc.Contributor

// A point or period of time associated with an event in the lifecycle of the resource.
//
// Date may be used to express temporal information at any level of granularity. Recommended practice is to express the
// date, date/time, or period of time according to ISO 8601-1 [ISO 8601-1] or a published profile of the ISO standard,
// such as the W3C Note on Date and Time Formats [W3CDTF] or the Extended Date/Time Format Specification [EDTF]. If the
// full date is unknown, month and year (YYYY-MM) or just year (YYYY) may be used. Date ranges may be specified using
// ISO 8601 period of time specification in which start and end dates are separated by a '/' (slash) character. Either
// the start or end date may be missing.
//
// See http://purl.org/dc/terms/date
type Date dc.Date

// The nature or genre of the resource.
//
// Recommended practice is to use a controlled vocabulary such as the DCMI Type Vocabulary [DCMI-TYPE]. To describe the
// file format, physical medium, or dimensions of the resource, use the property Format.
//
// See http://purl.org/dc/terms/type
type Type dc.Type

// The file format, physical medium, or dimensions of the resource.
//
// Recommended practice is to use a controlled vocabulary where available. For example, for file formats one could use
// the list of Internet Media Types [MIME]. Examples of dimensions include size and duration.
//
// See http://purl.org/dc/terms/format
type Format dc.Format

// An unambiguous reference to the resource within a given context.
//
// Recommended practice is to identify the resource by means of a string conforming to an identification system.
// Examples include International Standard Book Number (ISBN), Digital Object Identifier (DOI), and Uniform Resource
// Name (URN). Persistent identifiers should be provided as HTTP URIs.
//
// See http://purl.org/dc/terms/identifier
type Identifier dc.Identifier

// A related resource from which the described resource is derived.
//
// This property is intended to be used with non-literal values. The described resource may be derived from the related
// resource in whole or in part. Best practice is to identify the related resource by means of a URI or a string
// conforming to a formal identification system.
//
// See http://purl.org/dc/terms/source
type Source dc.Source

// A language of the resource.
//
// Recommended practice is to use either a non-literal value representing a language from a controlled vocabulary such
// as ISO 639-2 or ISO 639-3, or a literal value consisting of an IETF Best Current Practice 47 [IETF-BCP47] language
// tag.
//
// See http://purl.org/dc/terms/language
type Language dc.Language

// A related resource.
//
// Recommended practice is to identify the related resource by means of a URI. If this is not possible or feasible, a
// string conforming to a formal identification system may be provided.
//
// See http://purl.org/dc/terms/relation
type Relation dc.Relation

// The spatial or temporal topic of the resource, spatial applicability of the resource, or jurisdiction under which the
// resource is relevant.
//
// Spatial topic and spatial applicability may be a named place or a location specified by its geographic coordinates.
// Temporal topic may be a named period, date, or date range. A jurisdiction may be a named administrative entity or a
// geographic place to which the resource applies. Recommended practice is to use a controlled vocabulary such as the
// Getty Thesaurus of Geographic Names [TGN]. Where appropriate, named places or time periods may be used in preference
// to numeric identifiers such as sets of coordinates or date ranges. Because coverage is so broadly defined, it is
// preferable to use the more specific subproperties Temporal Coverage and Spatial Coverage.
//
// See http://purl.org/dc/terms/coverage
type Coverage dc.Coverage

// Information about rights held in and over the resource.
//
// Typically, rights information includes a statement about various property rights associated with the resource,
// including intellectual property rights. Recommended practice is to refer to a rights statement with a URI. If this is
// not possible or feasible, a literal value (name, label, or short text) may be provided.
//
// See http://purl.org/dc/terms/rights
type Rights dc.Rights

// An alternative name for the resource.
//
// The distinction between titles and alternative titles is application-specific.
//
// See http://purl.org/dc/terms/alternative
type Alternative Title

// A list of subunits of the resource.
//
// See http://purl.org/dc/terms/tableOfContents
type TableOfContents Description

// A summary of the resource.
//
// See http://purl.org/dc/terms/abstract
type Abstract Description

// Date of creation of the resource.
//
// Recommended practice is to describe the date, date/time, or period of time as recommended for the property Date, of
// which this is a subproperty.
//
// See http://purl.org/dc/terms/created
type Created Date

// Date (often a range) of validity of a resource.
//
// Recommended practice is to describe the date, date/time, or period of time as recommended for the property Date, of
// which this is a subproperty.
//
// See http://purl.org/dc/terms/valid
type Valid Date

// Date that the resource became or will become available.
//
// Recommended practice is to describe the date, date/time, or period of time as recommended for the property Date, of
// which this is a subproperty.
//
// See http://purl.org/dc/terms/available
type Available Date

// Date of formal issuance of the resource.
//
// Recommended practice is to describe the date, date/time, or period of time as recommended for the property Date, of
// which this is a subproperty.
//
// See http://purl.org/dc/terms/issued
type Issued Date

// Date on which the resource was changed.
//
// Recommended practice is to describe the date, date/time, or period of time as recommended for the property Date, of
// which this is a subproperty.
//
// See http://purl.org/dc/terms/modified
type Modified Date

// Date of acceptance of the resource.
//
// Recommended practice is to describe the date, date/time, or period of time as recommended for the property Date, of
// which this is a subproperty. Examples of resources to which a date of acceptance may be relevant are a thesis
// (accepted by a university department) or an article (accepted by a journal).
//
// See http://purl.org/dc/terms/dateAccepted
type DateAccepted Date

// Date of copyright of the resource.
//
// Typically a year. Recommended practice is to describe the date, date/time, or period of time as recommended for the
// property Date, of which this is a subproperty.
//
// See http://purl.org/dc/terms/dateCopyrighted
type DateCopyrighted Date

// Date of submission of the resource.
//
// Recommended practice is to describe the date, date/time, or period of time as recommended for the property Date, of
// which this is a subproperty. Examples of resources to which a 'Date Submitted' may be relevant include a thesis
// (submitted to a university department) or an article (submitted to a journal).
//
// See http://purl.org/dc/terms/dateSubmitted
type DateSubmitted Date

// The size or duration of the resource.
//
// Recommended practice is to specify the file size in megabytes and duration in ISO 8601 format.
//
// See http://purl.org/dc/terms/extent
type Extent Format

// The material or physical carrier of the resource.
//
// See http://purl.org/dc/terms/medium
type Medium Format

// A related resource of which the described resource is a version, edition, or adaptation.
//
// Changes in version imply substantive changes in content rather than differences in format. This property is intended
// to be used with non-literal values. This property is an inverse property of Has Version.
//
// See http://purl.org/dc/terms/isVersionOf
type IsVersionOf Relation

// A related resource that is a version, edition, or adaptation of the described resource.
//
// Changes in version imply substantive changes in content rather than differences in format. This property is intended
// to be used with non-literal values. This property is an inverse property of Is Version Of.
//
// See http://purl.org/dc/terms/hasVersion
type HasVersion Relation

// A related resource that supplants, displaces, or supersedes the described resource.
//
// This property is intended to be used with non-literal values. This property is an inverse property of Replaces.
//
// See http://purl.org/dc/terms/isReplacedBy
type IsReplacedBy Relation

// A related resource that is supplanted, displaced, or superseded by the described resource.
//
// This property is intended to be used with non-literal values. This property is an inverse property of Is Replaced By.
//
// See http://purl.org/dc/terms/replaces
type Replaces Relation

// A related resource that requires the described resource to support its function, delivery, or coherence.
//
// This property is intended to be used with non-literal values. This property is an inverse property of Requires.
//
// See http://purl.org/dc/terms/isRequiredBy
type IsRequiredBy Relation

// A related resource that is required by the described resource to support its function, delivery, or coherence.
//
// This property is intended to be used with non-literal values. This property is an inverse property of Is Required By.
//
// See http://purl.org/dc/terms/requires
type Requires Relation

// A related resource in which the described resource is physically or logically included.
//
// This property is intended to be used with non-literal values. This property is an inverse property of Has Part.
//
// See http://purl.org/dc/terms/isPartOf
type IsPartOf Relation

// A related resource that is included either physically or logically in the described resource.
//
// This property is intended to be used with non-literal values. This property is an inverse property of Is Part Of.
//
// See http://purl.org/dc/terms/hasPart
type HasPart Relation

// A related resource that references, cites, or otherwise points to the described resource.
//
// This property is intended to be used with non-literal values. This property is an inverse property of References.
//
// See http://purl.org/dc/terms/isReferencedBy
type IsReferencedBy Relation

// A related resource that is referenced, cited, or otherwise pointed to by the described resource.
//
// This property is intended to be used with non-literal values. This property is an inverse property of Is Referenced
// By.
//
// See http://purl.org/dc/terms/references
type References Relation

// A pre-existing related resource that is substantially the same as the described resource, but in another format.
//
// This property is intended to be used with non-literal values. This property is an inverse property of Has Format.
//
// See http://purl.org/dc/terms/isFormatOf
type IsFormatOf Relation

// A related resource that is substantially the same as the pre-existing described resource, but in another format.
//
// This property is intended to be used with non-literal values. This property is an inverse property of Is Format Of.
//
// See http://purl.org/dc/terms/hasFormat
type HasFormat Relation

// An established standard to which the described resource conforms.
//
// See http://purl.org/dc/terms/conformsTo
type ConformsTo Relation

// Spatial characteristics of the resource.
//
// See http://purl.org/dc/terms/spatial
type Spatial Coverage

// Temporal characteristics of the resource.
//
// See http://purl.org/dc/terms/temporal
type Temporal Coverage

// A class of agents for whom the resource is intended or useful.
//
// Recommended practice is to use this property with non-literal values from a vocabulary of audience types.
//
// See http://purl.org/dc/terms/audience
type Audience dc.Any

// The method by which items are added to a collection.
//
// Recommended practice is to use a value from the Collection Description Accrual Method Vocabulary
// [DCMI-ACCRUALMETHOD].
//
// See http://purl.org/dc/terms/accrualMethod
type AccrualMethod dc.Any

// The frequency with which items are added to a collection.
//
// Recommended practice is to use a value from the Collection Description Frequency Vocabulary [DCMI-COLLFREQ].
//
// See http://purl.org/dc/terms/accrualPeriodicity
type AccrualPeriodicity dc.Any

// The policy governing the addition of items to a collection.
//
// Recommended practice is to use a value from the Collection Description Accrual Policy Vocabulary
// [DCMI-ACCRUALPOLICY].
//
// See http://purl.org/dc/terms/accrualPolicy
type AccrualPolicy dc.Any

// A process, used to engender knowledge, attitudes and skills, that the described resource is designed to support.
//
// Instructional Method typically includes ways of presenting instructional materials or conducting instructional
// activities, patterns of learner-to-learner and learner-to-instructor interactions, and mechanisms by which group and
// individual levels of learning are measured. Instructional methods include all aspects of the instruction and learning
// processes from planning and implementation through evaluation and feedback.
//
// See http://purl.org/dc/terms/instructionalMethod
type InstructionalMethod dc.Any

// A statement of any changes in ownership and custody of the resource since its creation that are significant for its
// authenticity, integrity, and interpretation.
//
// The statement may include a description of any changes successive custodians made to the resource.
//
// See http://purl.org/dc/terms/provenance
type Provenance dc.Any

// A person or organization owning or managing rights over the resource.
//
// Recommended practice is to refer to the rights holder with a URI. If this is not possible or feasible, a literal
// value that identifies the rights holder may be provided.
//
// See http://purl.org/dc/terms/rightsHolder
type RightsHolder dc.Any

// An entity that mediates access to the resource.
//
// In an educational context, a mediator might be a parent, teacher, teaching assistant, or care-giver.
//
// See http://purl.org/dc/terms/mediator
type Mediator Audience

// A class of agents, defined in terms of progression through an educational or training context, for which the
// described resource is intended.
//
// See http://purl.org/dc/terms/educationLevel
type EducationLevel Audience

// Information about who access the resource or an indication of its security status.
//
// Access Rights may include information regarding access or restrictions based on privacy, security, or other policies.
//
// See http://purl.org/dc/terms/accessRights
type AccessRights Rights

// A legal document giving official permission to do something with the resource.
//
// Recommended practice is to identify the license document with a URI. If this is not possible or feasible, a literal
// value that identifies the license may be provided.
//
// See http://purl.org/dc/terms/license
type License Rights

// A bibliographic reference for the resource.
//
// Recommended practice is to include sufficient bibliographic detail to identify the resource as unambiguously as
// possible.
//
// See http://purl.org/dc/terms/bibliographicCitation
type BibliographicCitation Identifier

type Terms struct {
	XMLName xml.Name `json:"-"`

	Title                 []Title                 `xml:"http://purl.org/dc/terms dcterms:title,omitempty"                 json:"dcterms:title,omitempty"`
	Creator               []Creator               `xml:"http://purl.org/dc/terms dcterms:creator,omitempty"               json:"dcterms:creator,omitempty"`
	Subject               []Subject               `xml:"http://purl.org/dc/terms dcterms:subject,omitempty"               json:"dcterms:subject,omitempty"`
	Description           []Description           `xml:"http://purl.org/dc/terms dcterms:description,omitempty"           json:"dcterms:description,omitempty"`
	Publisher             []Publisher             `xml:"http://purl.org/dc/terms dcterms:publisher,omitempty"             json:"dcterms:publisher,omitempty"`
	Contributor           []Contributor           `xml:"http://purl.org/dc/terms dcterms:contributor,omitempty"           json:"dcterms:contributor,omitempty"`
	Date                  []Date                  `xml:"http://purl.org/dc/terms dcterms:date,omitempty"                  json:"dcterms:date,omitempty"`
	Type                  []Type                  `xml:"http://purl.org/dc/terms dcterms:type,omitempty"                  json:"dcterms:type,omitempty"`
	Format                []Format                `xml:"http://purl.org/dc/terms dcterms:format,omitempty"                json:"dcterms:format,omitempty"`
	Identifier            []Identifier            `xml:"http://purl.org/dc/terms dcterms:identifier,omitempty"            json:"dcterms:identifier,omitempty"`
	Source                []Source                `xml:"http://purl.org/dc/terms dcterms:source,omitempty"                json:"dcterms:source,omitempty"`
	Language              []Language              `xml:"http://purl.org/dc/terms dcterms:language,omitempty"              json:"dcterms:language,omitempty"`
	Relation              []Relation              `xml:"http://purl.org/dc/terms dcterms:relation,omitempty"              json:"dcterms:relation,omitempty"`
	Coverage              []Coverage              `xml:"http://purl.org/dc/terms dcterms:coverage,omitempty"              json:"dcterms:coverage,omitempty"`
	Rights                []Rights                `xml:"http://purl.org/dc/terms dcterms:rights,omitempty"                json:"dcterms:rights,omitempty"`
	Alternative           []Alternative           `xml:"http://purl.org/dc/terms dcterms:alternative,omitempty"           json:"dcterms:alternative,omitempty"`
	TableOfContents       []TableOfContents       `xml:"http://purl.org/dc/terms dcterms:tableOfContents,omitempty"       json:"dcterms:tableOfContents,omitempty"`
	Abstract              []Abstract              `xml:"http://purl.org/dc/terms dcterms:abstract,omitempty"              json:"dcterms:abstract,omitempty"`
	Created               []Created               `xml:"http://purl.org/dc/terms dcterms:created,omitempty"               json:"dcterms:created,omitempty"`
	Valid                 []Valid                 `xml:"http://purl.org/dc/terms dcterms:valid,omitempty"                 json:"dcterms:valid,omitempty"`
	Available             []Available             `xml:"http://purl.org/dc/terms dcterms:available,omitempty"             json:"dcterms:available,omitempty"`
	Issued                []Issued                `xml:"http://purl.org/dc/terms dcterms:issued,omitempty"                json:"dcterms:issued,omitempty"`
	Modified              []Modified              `xml:"http://purl.org/dc/terms dcterms:modified,omitempty"              json:"dcterms:modified,omitempty"`
	DateAccepted          []DateAccepted          `xml:"http://purl.org/dc/terms dcterms:dateAccepted,omitempty"          json:"dcterms:dateAccepted,omitempty"`
	DateCopyrighted       []DateCopyrighted       `xml:"http://purl.org/dc/terms dcterms:dateCopyrighted,omitempty"       json:"dcterms:dateCopyrighted,omitempty"`
	DateSubmitted         []DateSubmitted         `xml:"http://purl.org/dc/terms dcterms:dateSubmitted,omitempty"         json:"dcterms:dateSubmitted,omitempty"`
	Extent                []Extent                `xml:"http://purl.org/dc/terms dcterms:extent,omitempty"                json:"dcterms:extent,omitempty"`
	Medium                []Medium                `xml:"http://purl.org/dc/terms dcterms:medium,omitempty"                json:"dcterms:medium,omitempty"`
	IsVersionOf           []IsVersionOf           `xml:"http://purl.org/dc/terms dcterms:isVersionOf,omitempty"           json:"dcterms:isVersionOf,omitempty"`
	HasVersion            []HasVersion            `xml:"http://purl.org/dc/terms dcterms:hasVersion,omitempty"            json:"dcterms:hasVersion,omitempty"`
	IsReplacedBy          []IsReplacedBy          `xml:"http://purl.org/dc/terms dcterms:isReplacedBy,omitempty"          json:"dcterms:isReplacedBy,omitempty"`
	Replaces              []Replaces              `xml:"http://purl.org/dc/terms dcterms:replaces,omitempty"              json:"dcterms:replaces,omitempty"`
	IsRequiredBy          []IsRequiredBy          `xml:"http://purl.org/dc/terms dcterms:isRequiredBy,omitempty"          json:"dcterms:isRequiredBy,omitempty"`
	Requires              []Requires              `xml:"http://purl.org/dc/terms dcterms:requires,omitempty"              json:"dcterms:requires,omitempty"`
	IsPartOf              []IsPartOf              `xml:"http://purl.org/dc/terms dcterms:isPartOf,omitempty"              json:"dcterms:isPartOf,omitempty"`
	HasPart               []HasPart               `xml:"http://purl.org/dc/terms dcterms:hasPart,omitempty"               json:"dcterms:hasPart,omitempty"`
	IsReferencedBy        []IsReferencedBy        `xml:"http://purl.org/dc/terms dcterms:isReferencedBy,omitempty"        json:"dcterms:isReferencedBy,omitempty"`
	References            []References            `xml:"http://purl.org/dc/terms dcterms:references,omitempty"            json:"dcterms:references,omitempty"`
	IsFormatOf            []IsFormatOf            `xml:"http://purl.org/dc/terms dcterms:isFormatOf,omitempty"            json:"dcterms:isFormatOf,omitempty"`
	HasFormat             []HasFormat             `xml:"http://purl.org/dc/terms dcterms:hasFormat,omitempty"             json:"dcterms:hasFormat,omitempty"`
	ConformsTo            []ConformsTo            `xml:"http://purl.org/dc/terms dcterms:conformsTo,omitempty"            json:"dcterms:conformsTo,omitempty"`
	Spatial               []Spatial               `xml:"http://purl.org/dc/terms dcterms:spatial,omitempty"               json:"dcterms:spatial,omitempty"`
	Temporal              []Temporal              `xml:"http://purl.org/dc/terms dcterms:temporal,omitempty"              json:"dcterms:temporal,omitempty"`
	Audience              []Audience              `xml:"http://purl.org/dc/terms dcterms:audience,omitempty"              json:"dcterms:audience,omitempty"`
	AccrualMethod         []AccrualMethod         `xml:"http://purl.org/dc/terms dcterms:accrualMethod,omitempty"         json:"dcterms:accrualMethod,omitempty"`
	AccrualPeriodicity    []AccrualPeriodicity    `xml:"http://purl.org/dc/terms dcterms:accrualPeriodicity,omitempty"    json:"dcterms:accrualPeriodicity,omitempty"`
	AccrualPolicy         []AccrualPolicy         `xml:"http://purl.org/dc/terms dcterms:accrualPolicy,omitempty"         json:"dcterms:accrualPolicy,omitempty"`
	InstructionalMethod   []InstructionalMethod   `xml:"http://purl.org/dc/terms dcterms:instructionalMethod,omitempty"   json:"dcterms:instructionalMethod,omitempty"`
	Provenance            []Provenance            `xml:"http://purl.org/dc/terms dcterms:provenance,omitempty"            json:"dcterms:provenance,omitempty"`
	RightsHolder          []RightsHolder          `xml:"http://purl.org/dc/terms dcterms:rightsHolder,omitempty"          json:"dcterms:rightsHolder,omitempty"`
	Mediator              []Mediator              `xml:"http://purl.org/dc/terms dcterms:mediator,omitempty"              json:"dcterms:mediator,omitempty"`
	EducationLevel        []EducationLevel        `xml:"http://purl.org/dc/terms dcterms:educationLevel,omitempty"        json:"dcterms:educationLevel,omitempty"`
	AccessRights          []AccessRights          `xml:"http://purl.org/dc/terms dcterms:accessRights,omitempty"          json:"dcterms:accessRights,omitempty"`
	License               []License               `xml:"http://purl.org/dc/terms dcterms:license,omitempty"               json:"dcterms:license,omitempty"`
	BibliographicCitation []BibliographicCitation `xml:"http://purl.org/dc/terms dcterms:bibliographicCitation,omitempty" json:"dcterms:bibliographicCitation,omitempty"`
}

type Elements struct {
	XMLName xml.Name `json:"-"`

	DCTitle       []dc.Title       `xml:"http://purl.org/dc/elements/1.1/ dc:title,omitempty"       json:"dc:title,omitempty"`
	DCCreator     []dc.Creator     `xml:"http://purl.org/dc/elements/1.1/ dc:creator,omitempty"     json:"dc:creator,omitempty"`
	DCSubject     []dc.Subject     `xml:"http://purl.org/dc/elements/1.1/ dc:subject,omitempty"     json:"dc:subject,omitempty"`
	DCDescription []dc.Description `xml:"http://purl.org/dc/elements/1.1/ dc:description,omitempty" json:"dc:description,omitempty"`
	DCPublisher   []dc.Publisher   `xml:"http://purl.org/dc/elements/1.1/ dc:publisher,omitempty"   json:"dc:publisher,omitempty"`
	DCContributor []dc.Contributor `xml:"http://purl.org/dc/elements/1.1/ dc:contributor,omitempty" json:"dc:contributor,omitempty"`
	DCDate        []dc.Date        `xml:"http://purl.org/dc/elements/1.1/ dc:date,omitempty"        json:"dc:date,omitempty"`
	DCType        []dc.Type        `xml:"http://purl.org/dc/elements/1.1/ dc:type,omitempty"        json:"dc:type,omitempty"`
	DCFormat      []dc.Format      `xml:"http://purl.org/dc/elements/1.1/ dc:format,omitempty"      json:"dc:format,omitempty"`
	DCIdentifier  []dc.Identifier  `xml:"http://purl.org/dc/elements/1.1/ dc:identifier,omitempty"  json:"dc:identifier,omitempty"`
	DCSource      []dc.Source      `xml:"http://purl.org/dc/elements/1.1/ dc:source,omitempty"      json:"dc:source,omitempty"`
	DCLanguage    []dc.Language    `xml:"http://purl.org/dc/elements/1.1/ dc:language,omitempty"    json:"dc:language,omitempty"`
	DCRelation    []dc.Relation    `xml:"http://purl.org/dc/elements/1.1/ dc:relation,omitempty"    json:"dc:relation,omitempty"`
	DCCoverage    []dc.Coverage    `xml:"http://purl.org/dc/elements/1.1/ dc:coverage,omitempty"    json:"dc:coverage,omitempty"`
	DCRights      []dc.Rights      `xml:"http://purl.org/dc/elements/1.1/ dc:rights,omitempty"      json:"dc:rights,omitempty"`
}

// This is included as a convenience for schema authors who need to define a root or container element for all of the DC
// elements and element refinements.
type ElementsAndTerms struct {
	XMLName xml.Name `json:"-"`

	Elements
	Terms
}
