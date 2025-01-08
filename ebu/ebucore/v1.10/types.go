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

package v1_10

import (
	"github.com/nagare-media/models.go/base"
	"github.com/nagare-media/models.go/dcmi/dc"
	"github.com/nagare-media/models.go/third_party/encoding/xml"
)

const (
	SchemaVersion = "1.10.1"

	XMLNS       = "urn:ebu:metadata-schema:ebucore"
	XMLNSPrefix = "ebucore"
)

// The root element of EBUCore.
type Main struct {
	XMLName xml.Name `xml:"urn:ebu:metadata-schema:ebucore ebucore:ebuCoreMain,omitempty" json:"-"`

	// To allow defining a type of description associated with the metadata document.
	TypeAttributes

	// The body of EBUCore's descriptive metadata. coreMetadata is used to describe a so-called "root" item. The same
	// coreMetadata is also used to describe elements / "parts" of this item (see the "part" element).
	CoreMetadata CoreMetadata `xml:"urn:ebu:metadata-schema:ebucore ebucore:coreMetadata,omitempty" json:"ebucore:coreMetadata,omitempty"`

	// Identifies the metadata provider, e.g. the contributing archive. The organisation Id or name provide the archive ID
	// or name required for e.g. OAI metadata harvesting operation.
	MetadataProvider *Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:metadataProvider,omitempty" json:"ebucore:metadataProvider,omitempty"`

	// The name of the schema for e.g. OAI management.
	Schema string `xml:"schema,attr,omitempty" json:"@schema,omitempty"`

	// The version of the schema for e.g. OAI management.
	// TODO: default="1.8">
	Version string `xml:"version,attr,omitempty" json:"@version,omitempty"`

	// The date of edition of the current metadata instance for e.g. OAI management
	DateLastModified base.Date `xml:"dateLastModified,attr,omitempty" json:"@dateLastModified,omitempty"`

	// The time of edition of the current metadata instance for e.g. OAI management
	TimeLastModified base.Time `xml:"timeLastModified,attr,omitempty" json:"@timeLastModified,omitempty"`

	// The unique Identifier of the metadata instance for e.g. OAI management
	DocumentID string `xml:"documentId,attr,omitempty" json:"@documentId,omitempty"`

	// A location at which the document containing the metadata instance can be found
	DocumentLocation base.URI `xml:"documentLocation,attr,omitempty" json:"@documentLocation,omitempty"`

	// An attribute to specify the dominant language used to express metadata information in the document, which can be
	// superceded each time an language attribute or element is available a different levels of description granularity
	XMLLang string `xml:"xml:lang,attr,omitempty" json:"@xml:lang,omitempty"`

	// To provide information on the name of a tool used to generate data
	WritingLibraryName string `xml:"writingLibraryName,attr,omitempty" json:"@writingLibraryName,omitempty"`

	// To provide information on the version of a tool used to generate data
	WritingLibraryVersion string `xml:"writingLibraryVersion,attr,omitempty" json:"@writingLibraryVersion,omitempty"`
}

// The document containing all the core descriptive information regarding the resource
type CoreMetadata struct {
	// A Title is the 'main' name given to a resource e.g. a media item, a media object, or a sequence as specified by the
	// associated title type. It corresponds for a series to the series title, for a programme to the programme title, for
	// an item to the item title, etc. Titles are recorded as they appear. The Title is the name by which a resource is
	// formally known and that everyone should use to refer to or search for that particular resource. The Title may be
	// provided in several languages. If present, the date attributes indicate when the Title was attributed, used and/or
	// deprecated.
	//
	// dc:title is used to provide the main title by which the resource is known. The title can be provided in different
	// languages. The language in which the title is provided can be provided using dc:elementType's lang attribute.
	Title []Title `xml:"urn:ebu:metadata-schema:ebucore ebucore:title,omitempty" json:"ebucore:title,omitempty"`

	// An Alternative Title is the name other than the 'main' Title given to a resource. The type of title is defined by
	// the typeGroup of attributes. The status of the title is defined by the statusGroup of attributes. Alternative
	// Titles are recorded as they appear. An Alternative Title may be attributed to a resource for several reasons
	// described using the status (e.g. working title) and type (e.g. series title) attributes. The alternativeTitle may
	// be provided in several languages. It is sometimes common practice to put dates into the alternativeTitle. If
	// present, the date (indicating when the alternativeTitle was attributed and used or deprecated) in the date
	// attributes should be consistent.
	// dc:alternativeTitle is used to provide an alternative title by which the resource is known. The alternative title
	// of a particular type can be provided in different languages. The language in which the title is provided can be
	// provided using dc:elementType's lang attribute.
	AlternativeTitle []AlternativeTitle `xml:"urn:ebu:metadata-schema:ebucore ebucore:alternativeTitle,omitempty" json:"ebucore:alternativeTitle,omitempty"`

	// The descriptor creator identifies an 'entity' (a person, group of persons or organisation) primarily responsible
	// for creating the content of the resource - behind the camera. Different roles may be considered as representing a
	// creator, e.g. a producer, an author, etc. Creator is a sub-class of Contributor.
	Creator []Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:creator,omitempty" json:"ebucore:creator,omitempty"`

	// The subject addressed by the intellectual content of the resource. Typically, a subject is expressed by keywords,
	// key phrases. Free text, controlled vocabularies, authorities, or formal classification schemes (codes) may be
	// employed when selecting descriptive subject terms. Persons as subjects are also placed here. IMPORTANT NOTE: The
	// "genre" of the content is defined under element “ebucore:type/ebucore:genre”
	// dc:subject is used to express the subject in the form of free text optionally in different languages.
	Subject []Subject `xml:"urn:ebu:metadata-schema:ebucore ebucore:subject,omitempty" json:"ebucore:subject,omitempty"`

	// The topic covered by the intellectual content of the resource.
	Topic []Topic `xml:"urn:ebu:metadata-schema:ebucore ebucore:topic,omitempty" json:"ebucore:topic,omitempty"`

	// The topic covered by the intellectual content of the resource.
	Theme []Theme `xml:"urn:ebu:metadata-schema:ebucore ebucore:theme,omitempty" json:"ebucore:theme,omitempty"`

	// Free-form text or a narrative to report general notes, abstracts, or summaries about the intellectual content of a
	// resource. The information may be in the form of a paragraph giving an individual program description, anecdotal
	// interpretations, or brief content reviews. The description may also consist of outlines, lists, bullet points, edit
	// decision lists, indexes, or tables of content, a reference to a graphical representation of content or even a
	// pointer (URI, URL) to an external resource. A running order can also be provided as a description. For a Radio or
	// television programme a running order can be used as description. A description can be provided in different
	// languages.
	// dc:description is used to provide the text of a description of the resource, which can be provided in different
	// languages.
	Description []Description `xml:"urn:ebu:metadata-schema:ebucore ebucore:description,omitempty" json:"ebucore:description,omitempty"`

	// A publisher is a person, an organisation, or a service. Typically, the name of a Publisher should be used to
	// indicate the entity primarily responsible for distributing or making a resource available to others e.g. by
	// broadcasting, selling, leasing, renting and other modes of distribution.
	Publisher []Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:publisher,omitempty" json:"ebucore:publisher,omitempty"`

	// The descriptor contributor identifies a person or organisation that has made substantial creative contributions to
	// the content of a resource. Refers particularly (but not only) to participation in front of the camera. If in doubt
	// whether an entity is a creator or contributor use the element contributor.
	Contributor []Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:contributor,omitempty" json:"ebucore:contributor,omitempty"`

	// Dates associated with events occurring during the life of the resource. Typically, Date will be associated with the
	// creation, modification or availability of the resource.
	Date []ResourceDates `xml:"urn:ebu:metadata-schema:ebucore ebucore:date,omitempty" json:"ebucore:date,omitempty"`

	// The nature or genre of the resource. Type includes terms describing general categories, functions, genres, or
	// aggregation levels for content. Recommended best practice is to select a value from a controlled vocabulary. To
	// describe the physical or digital manifestation of the resource, use the FORMAT element.
	Type []Type `xml:"urn:ebu:metadata-schema:ebucore ebucore:type,omitempty" json:"ebucore:type,omitempty"`

	// The physical or digital manifestation of the resource. Use the descriptor Format to identify the format of a
	// particular resource as it exists in its physical or digital form. Physical form = an actual physical form that
	// occupies physical space, e.g. a tape. Digital form = a digital file residing on a server or hard drive. Format may
	// be used to determine the software, hardware or other equipment needed to display or operate the resource.
	Format []Format `xml:"urn:ebu:metadata-schema:ebucore ebucore:format,omitempty" json:"ebucore:format,omitempty"`

	// A unique, unambiguous reference or identifier for a resource within a given context. Best practice is to identify
	// the resource (whether analogue or digital) by means of a string or number corresponding to an established or formal
	// identification system if one exists. Otherwise, use an identification method that is in use within your agency,
	// station, production company, office, or institution. It is also possible to enter more than one, different but
	// still unique, identifier for the same resource.
	Identifier []Identifier `xml:"urn:ebu:metadata-schema:ebucore ebucore:identifier,omitempty" json:"ebucore:identifier,omitempty"`

	// Reference to the resource (s) from which the current resource is derived in whole or in part. If no label or number
	// is available, the title and/or the statement of responsibility etc. of the digitized recording is recorded here.
	// For a digitized radio programme the production number is normally given here. The Recommended best practice is to
	// use a unique identifier to identify the physical source that has been used to create the digital resource. In the
	// case of a digitized analogue recording, it is the recording used for digitization which is the source. For
	// commercial recordings the label and number is normally given here. Example: Eurovision feed
	// 2007-07-16T19:20:30.45+01:00
	DCSource []dc.Source `xml:"http://purl.org/dc/elements/1.1/ dc:source,omitempty" json:"dc:source,omitempty"`

	// Identifies languages and their use in the intellectual content of the resource. Recommended best practice for the
	// values of the Language element is defined by RFC 5646. More contextual information can be provided using the “note”
	// attribute.
	Language []Language `xml:"urn:ebu:metadata-schema:ebucore ebucore:language,omitempty" json:"ebucore:language,omitempty"`

	// Recommended best practice is to reference the resource (to which the current resource under description is related)
	// by means of a string or number conforming to a formal identification system. The Relation element is used to show
	// express a custom relation to another resource. A resource can be identified by its title, or an identifier
	// (possibly a URI). The related item has its own separate metadata record. Relation is used to provide a name, an
	// identification number or ID, or a locator where the related item can be found. The type of relation is defined a
	// term like in the https://www.ebu.ch/metadata/cs/ebu_HowRelatedCS.xml controlled vocabulary.
	Relation []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:relation,omitempty" json:"ebucore:relation,omitempty"`

	// Recommended best practice is to reference the resource (to which the current resource under description is related)
	// by means of a string or number conforming to a formal identification system. isRelatedTo is used to show a generic
	// relation of a resource with another resource. For example, "IsPartOf" is used to show the relation between a part
	// of a radio programme and the whole programme or between a track and a record album. A resource can be identified by
	// its title, or an identifier (possibly a URI). The related item has its own separate metadata record. Relation is
	// used to provide a name, an identification number or ID, or a locator where the related item can be found.
	IsRelatedTo []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:isRelatedTo,omitempty" json:"ebucore:isRelatedTo,omitempty"`

	// A relation use to link a resource to another following resource in an ordered sequence
	IsNextInSequence []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:isNextInSequence,omitempty" json:"ebucore:isNextInSequence,omitempty"`

	// A relation use to link a resource to another preceding resource in an ordered sequence
	FollowsInSequence []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:followsInSequence,omitempty" json:"ebucore:followsInSequence,omitempty"`

	// A reference to the resource that the current resource is a version of
	IsVersionOf []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:isVersionOf,omitempty" json:"ebucore:isVersionOf,omitempty"`

	// A reference to another version of the resource
	HasVersion []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:hasVersion,omitempty" json:"ebucore:hasVersion,omitempty"`

	// A reference to a resource replacing the current resource
	IsReplacedBy []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:isReplacedBy,omitempty" json:"ebucore:isReplacedBy,omitempty"`

	// A reference to a resource that the current resource replaces
	Replaces []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:replaces,omitempty" json:"ebucore:replaces,omitempty"`

	// A reference to a resource requiring the current resource
	IsRequiredBy []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:isRequiredBy,omitempty" json:"ebucore:isRequiredBy,omitempty"`

	// A reference to a resource that the current resource requires
	Requires []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:requires,omitempty" json:"ebucore:requires,omitempty"`

	// A reference to a resource that the current resource is a part of
	IsPartOf []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:isPartOf,omitempty" json:"ebucore:isPartOf,omitempty"`

	// A reference to a resource that forms part of the current resource
	HasPart []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:hasPart,omitempty" json:"ebucore:hasPart,omitempty"`

	// An element to identify a part of a track by a title, a start time and an end time in both the media source and
	// media destination. A timeline track part is a specific type of track part.
	HasTrackPart []HasTrackPart `xml:"urn:ebu:metadata-schema:ebucore ebucore:hasTrackPart,omitempty" json:"ebucore:hasTrackPart,omitempty"`

	// An element to identify a part of a track by a title, a start time and an end time in both the media source and
	// media destination. A timeline track part is a specific type of track part.
	IsTrackPartOf []IsTrackPartOf `xml:"urn:ebu:metadata-schema:ebucore ebucore:isTrackPartOf,omitempty" json:"ebucore:isTrackPartOf,omitempty"`

	// A reference to a resource that the current resource references
	References []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:references,omitempty" json:"ebucore:references,omitempty"`

	// A reference to a resource with which the current resource shares a format
	IsFormatOf []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:isFormatOf,omitempty" json:"ebucore:isFormatOf,omitempty"`

	// A format in which the resource is also available
	HasFormat []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:hasFormat,omitempty" json:"ebucore:hasFormat,omitempty"`

	// A reference to a series the current resource is an episode of
	IsEpisodeOf []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:isEpisodeOf,omitempty" json:"ebucore:isEpisodeOf,omitempty"`

	// A reference to a series the current resource is a season of
	IsSeasonOf []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:isSeasonOf,omitempty" json:"ebucore:isSeasonOf,omitempty"`

	// A reference to a series the current resource is an episode of
	HasEpisode []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:hasEpisode,omitempty" json:"ebucore:hasEpisode,omitempty"`

	// A reference to a series the current resource is a season of
	HasSeason []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:hasSeason,omitempty" json:"ebucore:hasSeason,omitempty"`

	// A reference to a brand the current resource is a series of
	HasSeries []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:hasSeries,omitempty" json:"ebucore:hasSeries,omitempty"`

	// A reference to a brand the current resource is a series of
	IsSeriesOf []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:isSeriesOf,omitempty" json:"ebucore:isSeriesOf,omitempty"`

	// A reference to a group e.g. a brand, the current resource is an member of
	IsMemberOf []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:isMemberOf,omitempty" json:"ebucore:isMemberOf,omitempty"`

	// A reference to members of a group
	HasMember []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:hasMember,omitempty" json:"ebucore:hasMember,omitempty"`

	// To indicate that two resources are identical.
	SameAs []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:sameAs,omitempty" json:"ebucore:sameAs,omitempty"`

	// A reference to a Parent resource
	HasParent []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:hasParent,omitempty" json:"ebucore:hasParent,omitempty"`

	// A reference to a Child resource
	IsParentOf []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:isParentOf,omitempty" json:"ebucore:isParentOf,omitempty"`

	// A reference to a Child resource
	HasChild []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:hasChild,omitempty" json:"ebucore:hasChild,omitempty"`

	// A reference to a Parent resource
	IsChildOf []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:isChildOf,omitempty" json:"ebucore:isChildOf,omitempty"`

	// A reference to a related Master resource
	HasMaster []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:hasMaster,omitempty" json:"ebucore:hasMaster,omitempty"`

	// A reference to a resource derived from a Master
	IsMasterOf []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:isMasterOf,omitempty" json:"ebucore:isMasterOf,omitempty"`

	// A reference to a parent resource from which the current resource was derived. This can be used for provenance purposes.
	IsDerivedFrom []Relation `xml:"urn:ebu:metadata-schema:ebucore ebucore:isDerivedFrom,omitempty" json:"ebucore:isDerivedFrom,omitempty"`

	// A manifestation is the physical embodiment of work e.g. a tape, a file...
	HasManifestation []Manifestation `xml:"urn:ebu:metadata-schema:ebucore ebucore:hasManifestation,omitempty" json:"ebucore:hasManifestation,omitempty"`

	// Coverage is used to show various time and place aspects of the subject of the content. Coverage will typically
	// include spatial location (a place name or geographic coordinates), temporal period (a period label, date, or date
	// range) or jurisdiction (such as a named administrative entity). Recommended best practice is to select a value from
	// a controlled vocabulary (for example, the Thesaurus of Geographic Names) and that, where appropriate, named places
	// or time periods be used in preference to numeric identifiers such as sets of coordinates or date ranges.
	Coverage []Coverage `xml:"urn:ebu:metadata-schema:ebucore ebucore:coverage,omitempty" json:"ebucore:coverage,omitempty"`

	// An all-purpose field to identify information (rights management statement or reference to a service providing such
	// information e.g. via a URL) about copyright, intellectual property rights or other property rights held in and over
	// a resource, stating whether access is open or restricted in some way. If dates, times, territories and availability
	// periods are associated with a right, they should be included. If the Rights element is absent, no assumptions can
	// be made about the status of these and other rights with respect to the resource.
	Rights []Rights `xml:"urn:ebu:metadata-schema:ebucore ebucore:rights,omitempty" json:"ebucore:rights,omitempty"`

	// UK Version, US Version, home video version, etc. Mapping to Dublin Core would be made using a description element.
	// There can be mutiple type fo version references.
	Version []Version `xml:"urn:ebu:metadata-schema:ebucore ebucore:version,omitempty" json:"ebucore:version,omitempty"`

	// To provide information on the publication history.
	PublicationHistory []PublicationHistory `xml:"urn:ebu:metadata-schema:ebucore ebucore:publicationHistory,omitempty" json:"ebucore:publicationHistory,omitempty"`

	// To provide a planning.
	Planning []Planning `xml:"urn:ebu:metadata-schema:ebucore ebucore:planning,omitempty" json:"ebucore:planning,omitempty"`

	// An element to provide rating values attributed by reviewers to evaluate the media resource.
	Rating []Rating `xml:"urn:ebu:metadata-schema:ebucore ebucore:ratingType,omitempty" json:"ebucore:ratingType,omitempty"`

	// An element to define the audience rating for a media resource.
	AudienceRating []Rating `xml:"urn:ebu:metadata-schema:ebucore ebucore:audienceRating,omitempty" json:"ebucore:audienceRating,omitempty"`

	// An element to define an event associated with a media resource.
	Event []Event `xml:"urn:ebu:metadata-schema:ebucore ebucore:event,omitempty" json:"ebucore:event,omitempty"`

	// To identify and describe artefacts present in a media resource or part of a media resource.
	Artefact []Artefact `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefact,omitempty" json:"ebucore:artefact,omitempty"`

	// To identify and describe animals associated with a media resource.
	Animal []Animal `xml:"urn:ebu:metadata-schema:ebucore ebucore:animal,omitempty" json:"ebucore:animal,omitempty"`

	// To identify and describe props present in a media resource or part of a media resource.
	Props []Props `xml:"urn:ebu:metadata-schema:ebucore ebucore:props,omitempty" json:"ebucore:props,omitempty"`

	// To identify and describe costume and wardrobe items present in a media resource or part of a media resource.
	Costume []Costume `xml:"urn:ebu:metadata-schema:ebucore ebucore:costume,omitempty" json:"ebucore:costume,omitempty"`

	// To identify and describe food or derived products present in a media resource or part of a media resource.
	Food []Food `xml:"urn:ebu:metadata-schema:ebucore ebucore:food,omitempty" json:"ebucore:food,omitempty"`

	// To provide text associated with a content, a scene or a part. Additional information can be provided on the source
	// of the text as well as a relation to a contributor.
	TextLine []TextLine `xml:"urn:ebu:metadata-schema:ebucore ebucore:textLine,omitempty" json:"ebucore:textLine,omitempty"`

	// To identify emotions associated with a person or character within a scene or a part.
	Emotion []Emotion `xml:"urn:ebu:metadata-schema:ebucore ebucore:emotion,omitempty" json:"ebucore:emotion,omitempty"`

	// To identify actions associated with a person or character within a scene or a part.
	Action []Action `xml:"urn:ebu:metadata-schema:ebucore ebucore:action,omitempty" json:"ebucore:action,omitempty"`

	// To identify parts/segments/fragments within the resource. See http://www.movielabs.com/md/ratings/v2.2.5/html/Summary.html
	Part []Part `xml:"urn:ebu:metadata-schema:ebucore ebucore:part,omitempty" json:"ebucore:part,omitempty"`
}

type HasTrackPart struct {
	Relation

	// To provide a title for a track part.
	TrackPartTitle AlternativeTitle `xml:"urn:ebu:metadata-schema:ebucore ebucore:trackPartTitle,omitempty" json:"ebucore:trackPartTitle,omitempty"`

	// Tor provide an identifier associated with the target resource.
	DestinationID *URIValue `xml:"urn:ebu:metadata-schema:ebucore ebucore:destinationId,omitempty" json:"ebucore:destinationId,omitempty"`

	// To provide a reference start time in the target resource.
	DestinationStart *Time `xml:"urn:ebu:metadata-schema:ebucore ebucore:destinationStart,omitempty" json:"ebucore:destinationStart,omitempty"`

	// To provide a reference end time in the target resource.
	DestinationEnd *Time `xml:"urn:ebu:metadata-schema:ebucore ebucore:destinationEnd,omitempty" json:"ebucore:destinationEnd,omitempty"`

	// Tor provide an identifier associated with the source.
	SourceID *URIValue `xml:"urn:ebu:metadata-schema:ebucore ebucore:sourceId,omitempty" json:"ebucore:sourceId,omitempty"`

	// To provide a reference start time at the source.
	SourceStart *Time `xml:"urn:ebu:metadata-schema:ebucore ebucore:sourceStart,omitempty" json:"ebucore:sourceStart,omitempty"`

	// To provide a reference end time at the source.
	SourceEnd *Time `xml:"urn:ebu:metadata-schema:ebucore ebucore:sourceEnd,omitempty" json:"ebucore:sourceEnd,omitempty"`
}

type IsTrackPartOf struct {
	Relation

	// To provide a title for a track part.
	TrackPartTitle AlternativeTitle `xml:"urn:ebu:metadata-schema:ebucore ebucore:trackPartTitle,omitempty" json:"ebucore:trackPartTitle,omitempty"`

	// Tor provide an identifier associated with the target resource.
	DestinationID *URIValue `xml:"urn:ebu:metadata-schema:ebucore ebucore:destinationId,omitempty" json:"ebucore:destinationId,omitempty"`

	// To provide a reference start time in the target resource.
	DestinationStart *Time `xml:"urn:ebu:metadata-schema:ebucore ebucore:destinationStart,omitempty" json:"ebucore:destinationStart,omitempty"`

	// To provide a reference end time in the target resource.
	DestinationEnd *Time `xml:"urn:ebu:metadata-schema:ebucore ebucore:destinationEnd,omitempty" json:"ebucore:destinationEnd,omitempty"`

	// Tor provide an identifier associated with the source.
	SourceID *URIValue `xml:"urn:ebu:metadata-schema:ebucore ebucore:sourceId,omitempty" json:"ebucore:sourceId,omitempty"`

	// To provide a reference start time at the source.
	SourceStart *Time `xml:"urn:ebu:metadata-schema:ebucore ebucore:sourceStart,omitempty" json:"ebucore:sourceStart,omitempty"`

	// To provide a reference end time at the source.
	SourceEnd *Time `xml:"urn:ebu:metadata-schema:ebucore ebucore:sourceEnd,omitempty" json:"ebucore:sourceEnd,omitempty"`
}

type Food struct {
	Artefact
}

type Manifestation struct {
	TypeAttributes
	FormatAttributes

	ManifestationTitle              []Title              `xml:"urn:ebu:metadata-schema:ebucore ebucore:manifestationTitle,omitempty" json:"ebucore:manifestationTitle,omitempty"`
	ManifestationIdentifier         []Identifier         `xml:"urn:ebu:metadata-schema:ebucore ebucore:manifestationIdentifier,omitempty" json:"ebucore:manifestationIdentifier,omitempty"`
	ManifestationDate               []ResourceDates      `xml:"urn:ebu:metadata-schema:ebucore ebucore:manifestationDate,omitempty" json:"ebucore:manifestationDate,omitempty"`
	ManifestationPublicationHistory []PublicationHistory `xml:"urn:ebu:metadata-schema:ebucore ebucore:manifestationPublicationHistory,omitempty" json:"ebucore:manifestationPublicationHistory,omitempty"`

	ManifestationId string `xml:"manifestationId,attr,omitempty" json:"@manifestationId,omitempty"`
}

// To describe an emotion and establish a link to a person/character.
type Emotion struct {
	// To identify the type of emotion.
	TypeAttributes

	// To identify the person/character associated with the emotion.
	RelatedAgent Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:relatedAgent,omitempty" json:"ebucore:relatedAgent,omitempty"`

	// A timestamp to precisely define a time point when the emotion was expressed..
	Timestamp Time `xml:"urn:ebu:metadata-schema:ebucore ebucore:timestamp,omitempty" json:"ebucore:timestamp,omitempty"`
}

// To describe an action and establish a link to a person/character.
type Action struct {
	// To identify the type of action.
	TypeAttributes

	// To identify the person/character associated with the ACtion.
	RelatedAgent Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:relatedAgent,omitempty" json:"ebucore:relatedAgent,omitempty"`

	// A timestamp to precisely define a time point when the action happened..
	Timestamp Time `xml:"urn:ebu:metadata-schema:ebucore ebucore:timestamp,omitempty" json:"ebucore:timestamp,omitempty"`
}

type TextLine struct {
	// To define the type / source of text e.g. speech to text
	TypeAttributes

	// A link to an external web resource containing the text.
	TextLink URIValue `xml:"urn:ebu:metadata-schema:ebucore ebucore:textLink,omitempty" json:"ebucore:textLink,omitempty"`

	// The text content.
	Text String `xml:"urn:ebu:metadata-schema:ebucore ebucore:text,omitempty" json:"ebucore:text,omitempty"`

	// A description of an agent referenced or associated with the text.
	RelatedAgent Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:relatedAgent,omitempty" json:"ebucore:relatedAgent,omitempty"`

	// A timestamp to precisely define a time point when the text was used.
	Timestamp Time `xml:"urn:ebu:metadata-schema:ebucore ebucore:timestamp,omitempty" json:"ebucore:timestamp,omitempty"`

	// to describe the position of text on the screen.
	TextLineBoxartefacttyPosition *TextLineBoxartefacttyPosition `xml:"urn:ebu:metadata-schema:ebucore ebucore:textLineBoxartefacttyPosition,omitempty" json:"ebucore:textLineBoxartefacttyPosition,omitempty"`

	// An Id attributed to the text line
	TextLineID string `xml:"textLineId,attr,omitempty" json:"@textLineId,omitempty"`

	// The order in which text appears in the content/part/scene.
	Order string `xml:"order,attr,omitempty" json:"@order,omitempty"`

	// The language used for the text line
	XMLLang string `xml:"xml:lang,attr,omitempty" json:"@xml:lang,omitempty"`
}

type TextLineBoxartefacttyPosition struct {
	// To define the line number of the left top corner position of the text box.
	LeftTopCornerLineNumber *UInt `xml:"urn:ebu:metadata-schema:ebucore ebucore:leftTopCornerLineNumber,omitempty" json:"ebucore:leftTopCornerLineNumber,omitempty"`

	// To define the pixel number of the left top corner position of the text box.
	LeftTopCornerPixelNumber *UInt `xml:"urn:ebu:metadata-schema:ebucore ebucore:leftTopCornerPixelNumber,omitempty" json:"ebucore:leftTopCornerPixelNumber,omitempty"`

	// The line number defining the height of the text box.
	Height *UInt `xml:"urn:ebu:metadata-schema:ebucore ebucore:height,omitempty" json:"ebucore:height,omitempty"`

	// The pixel number defining the width of the text box.
	Width *UInt `xml:"urn:ebu:metadata-schema:ebucore ebucore:width,omitempty" json:"ebucore:width,omitempty"`
}

// A complexType to describe artefacts.
type Artefact struct {
	TypeAttributes

	// To provide a name for the artefact.
	ArtefactName []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefactName,omitempty" json:"ebucore:artefactName,omitempty"`

	// To provide the model of the artefact.
	ArtefactModel *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefactModel,omitempty" json:"ebucore:artefactModel,omitempty"`

	// To provide a description of the artefact.
	ArtefactDescription []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefactDescription,omitempty" json:"ebucore:artefactDescription,omitempty"`

	// To provide the brand of the artefact.
	ArtefactBrand *ArtefactBrand `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefactBrand,omitempty" json:"ebucore:artefactBrand,omitempty"`

	// To provide the type of an artefact.
	ArtefactType *ArtefactType `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefacttype,omitempty" json:"ebucore:artefacttype,omitempty"`

	// To provide the colour of the artefact.
	ArtefactColour *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefactColour,omitempty" json:"ebucore:artefactColour,omitempty"`

	// To provide the weight of the artefact.
	ArtefactWeight *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefactWeight,omitempty" json:"ebucore:artefactWeight,omitempty"`

	// To provide the reference associated with the artefact.
	ArtefactReference *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefactReference,omitempty" json:"ebucore:artefactReference,omitempty"`

	// To provide the price of the artefact.
	ArtefactPrice []ArtefactPrice `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefactPrice,omitempty" json:"ebucore:artefactPrice,omitempty"`

	// To provide a website associated with the artefact.
	ArtefactWebsite *URIValue `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefactWebsite,omitempty" json:"ebucore:artefactWebsite,omitempty"`

	// To provide the date of purchase of the artefact.
	ArtefactDateOfPurchase *Date `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefactDateOfPurchase,omitempty" json:"ebucore:artefactDateOfPurchase,omitempty"`

	// To provide the date of sell of the artefact.
	ArtefactDateOfSell *Date `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefactDateOfSell,omitempty" json:"ebucore:artefactDateOfSell,omitempty"`

	// To provide the date of creation of the artefact.
	ArtefactDateOfCreation *Date `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefactDateOfCreation,omitempty" json:"ebucore:artefactDateOfCreation,omitempty"`

	// To provide the date of acquisition of the artefact.
	ArtefactDateOfAcquisition *Date `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefactDateOfAcquisition,omitempty" json:"ebucore:artefactDateOfAcquisition,omitempty"`

	// To provide the date of destruction of the artefact.
	ArtefactDateOfDestruction *Date `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefactDateOfDestruction,omitempty" json:"ebucore:artefactDateOfDestruction,omitempty"`

	// To indicate the availability of the artefact.
	ArtefactAvailability *Boolean `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefactAvailability,omitempty" json:"ebucore:artefactAvailability,omitempty"`

	// To locate the physical position of the artefact.
	ArtefactLocation *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefactLocation,omitempty" json:"ebucore:artefactLocation,omitempty"`

	// To provide information of the usage history of the artefact.
	ArtefactUsageHistory []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefactUsageHistory,omitempty" json:"ebucore:artefactUsageHistory,omitempty"`

	// To provide information on the style of the artefact.
	ArtefactStyle *Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefactStyle,omitempty" json:"ebucore:artefactStyle,omitempty"`

	// To provide information on the period associated with the artefact.
	ArtefactPeriod *ArtefactPeriod `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefactPeriod,omitempty" json:"ebucore:artefactPeriod,omitempty"`

	// to describe the position of text on the screen.
	ArtefactBoxPosition *ArtefactBoxPosition `xml:"urn:ebu:metadata-schema:ebucore ebucore:artefactBoxPosition,omitempty" json:"ebucore:artefactBoxPosition,omitempty"`

	// To identify agents related to the artefact.
	RelatedAgent []Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:relatedAgent,omitempty" json:"ebucore:relatedAgent,omitempty"`

	// To identify the supplier of the artefact.
	Supplier []Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:supplier,omitempty" json:"ebucore:supplier,omitempty"`

	// To identify the retailer of the artefact.
	Retailer []Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:retailer,omitempty" json:"ebucore:retailer,omitempty"`

	// To identify the buyer of the artefact.
	Buyer []Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:buyer,omitempty" json:"ebucore:buyer,omitempty"`

	// To identify the creator of the artefact.
	Creator []Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:creator,omitempty" json:"ebucore:creator,omitempty"`

	// To identify the maker of the artefact.
	Maker []Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:maker,omitempty" json:"ebucore:maker,omitempty"`

	// To identify the owner of the artefact.
	Owner []Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:owner,omitempty" json:"ebucore:owner,omitempty"`

	// To identify a contact associated with the artefact.
	Contact []Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:contact,omitempty" json:"ebucore:contact,omitempty"`

	// A timestamp to precisely define a time point when the artefact is used..
	Timestamp *Time `xml:"urn:ebu:metadata-schema:ebucore ebucore:timestamp,omitempty" json:"ebucore:timestamp,omitempty"`

	// To provide additional custom information about an artefact.
	AdditionalInformation []AdditionalInformation `xml:"urn:ebu:metadata-schema:ebucore ebucore:additionalInformation,omitempty" json:"ebucore:additionalInformation,omitempty"`

	// The id attributed to an artefact.
	ArtefactID string `xml:"artefactid,attr,omitempty" json:"@artefactid,omitempty"`
}

type ArtefactBrand struct {
	TypeAttributes

	Value string `xml:",chardata" json:"#value"`
}

type ArtefactType struct {
	TypeAttributes

	Value string `xml:",chardata" json:"#value"`
}

type ArtefactPrice struct {
	TypeAttributes

	Value    Float    `xml:"urn:ebu:metadata-schema:ebucore ebucore:value,omitempty" json:"ebucore:value,omitempty"`
	Currency Currency `xml:"urn:ebu:metadata-schema:ebucore ebucore:currency,omitempty" json:"ebucore:currency,omitempty"`
}

type Currency struct {
	TypeAttributes
}

type ArtefactPeriod struct {
	DateAttributes
}

type ArtefactBoxPosition struct {
	// To define the line number of the left top corner position of the box around the artefact.
	LeftTopCornerLineNumber *UInt `xml:"urn:ebu:metadata-schema:ebucore ebucore:leftTopCornerLineNumber,omitempty" json:"ebucore:leftTopCornerLineNumber,omitempty"`

	// To define the pixel number of the left top corner position of of the box around the artefact.
	LeftTopCornerPixelNumber *UInt `xml:"urn:ebu:metadata-schema:ebucore ebucore:leftTopCornerPixelNumber,omitempty" json:"ebucore:leftTopCornerPixelNumber,omitempty"`

	// The line number defining the height of the box around the artefact.
	Height *UInt `xml:"urn:ebu:metadata-schema:ebucore ebucore:height,omitempty" json:"ebucore:height,omitempty"`

	// The pixel number defining the width of the box around the artefact.
	Width *UInt `xml:"urn:ebu:metadata-schema:ebucore ebucore:width,omitempty" json:"ebucore:width,omitempty"`
}

type AdditionalInformation struct {
	Element
	TypeAttributes
}

// A complexType to describe props (artefacts used during the production of a media resource).
type Props struct {
	Artefact
}

// A complexType to describe props (artefacts used during the production of a media resource).
type Costume struct {
	Artefact

	CostumeSizeInformation *CostumeSizeInformation `xml:"urn:ebu:metadata-schema:ebucore ebucore:costumeSizeInformation,omitempty" json:"ebucore:costumeSizeInformation,omitempty"`
	CostumeGender          *CostumeGender          `xml:"urn:ebu:metadata-schema:ebucore ebucore:costumeGender,omitempty" json:"ebucore:costumeGender,omitempty"`
	CostumeTexture         []CostumeTexture        `xml:"urn:ebu:metadata-schema:ebucore ebucore:costumeTexture,omitempty" json:"ebucore:costumeTexture,omitempty"`
	LinkToLogo             []URIValue              `xml:"urn:ebu:metadata-schema:ebucore ebucore:linkToLogo,omitempty" json:"ebucore:linkToLogo,omitempty"`
	LinkToSticker          []URIValue              `xml:"urn:ebu:metadata-schema:ebucore ebucore:linkToSticker,omitempty" json:"ebucore:linkToSticker,omitempty"`
}

type CostumeSizeInformation struct {
	CostumeSizeCategory *CostumeSizeCategory `xml:"urn:ebu:metadata-schema:ebucore ebucore:costumeSizeCategory,omitempty" json:"ebucore:costumeSizeCategory,omitempty"`
	CostumeSize         String               `xml:"urn:ebu:metadata-schema:ebucore ebucore:costumeSize,omitempty" json:"ebucore:costumeSize,omitempty"`
	GeographicalArea    *GeographicalArea    `xml:"urn:ebu:metadata-schema:ebucore ebucore:geographicalArea,omitempty" json:"ebucore:geographicalArea,omitempty"`
}

type CostumeSizeCategory struct {
	TypeAttributes
}

type GeographicalArea struct {
	TypeAttributes
}

type CostumeGender struct {
	TypeAttributes
}

type CostumeTexture struct {
	TypeAttributes
}

// A complexType to describe food and derived products shown in a media resource.
// TODO: Should this be merged with "Food"?
type FoodType struct {
	Artefact

	FoodStyle       *FoodStyle      `xml:"urn:ebu:metadata-schema:ebucore ebucore:foodStyle,omitempty" json:"ebucore:foodStyle,omitempty"`
	FoodCategory    *FoodCategory   `xml:"urn:ebu:metadata-schema:ebucore ebucore:foodCategory,omitempty" json:"ebucore:foodCategory,omitempty"`
	CuisineOrigin   []CuisineOrigin `xml:"urn:ebu:metadata-schema:ebucore ebucore:cuisineOrigin,omitempty" json:"ebucore:cuisineOrigin,omitempty"`
	CuisineStyle    *CuisineStyle   `xml:"urn:ebu:metadata-schema:ebucore ebucore:cuisineStyle,omitempty" json:"ebucore:cuisineStyle,omitempty"`
	DishName        *Element        `xml:"urn:ebu:metadata-schema:ebucore ebucore:dishName,omitempty" json:"ebucore:dishName,omitempty"`
	DishDescription *Element        `xml:"urn:ebu:metadata-schema:ebucore ebucore:dishDescription,omitempty" json:"ebucore:dishDescription,omitempty"`
	FoodIngredient  *Element        `xml:"urn:ebu:metadata-schema:ebucore ebucore:foodIngredient,omitempty" json:"ebucore:foodIngredient,omitempty"`
}

type FoodStyle struct {
	TypeAttributes
}

type FoodCategory struct {
	TypeAttributes
}

type CuisineOrigin struct {
	TypeAttributes
}

type CuisineStyle struct {
	TypeAttributes
}

// To describe e.g. editorial segment / fragments (to be identified or located using e.g. W3C's Media Fragment URIs,
// timelines e.g. for dynamic technical metadata, play lists, elements of a group (itself described at root level) such
// as "tracks" of a record, episodes of a series or season, etc. It can also be sued to split data into user defined
// clusters. It can be used to describe different instantiations of the same asset in different locations possibly in a
// different formats with different rights. Etc.
type Part struct {
	// Each part can be described using the same level of detiled information as provided by the coreMetadata element.
	CoreMetadata

	// A group of attributes used to define the type of part.
	TypeAttributes

	// To indicate the start time of a part if of temporal nature.
	PartStartTime *Time `xml:"urn:ebu:metadata-schema:ebucore ebucore:partStartTime,omitempty" json:"ebucore:partStartTime,omitempty"`

	// To indicate the duration of a part if of temporal nature.
	// TODO: either PartDuration or PartEndTime or neither must be specified
	PartDuration *Duration `xml:"urn:ebu:metadata-schema:ebucore ebucore:partDuration,omitempty" json:"ebucore:partDuration,omitempty"`

	// To indicate the end time of a part if of temporal nature.
	// TODO: either PartDuration or PartEndTime or neither must be specified
	PartEndTime *Time `xml:"urn:ebu:metadata-schema:ebucore ebucore:partEndTime,omitempty" json:"ebucore:partEndTime,omitempty"`

	// To uniquely identify a part.
	PartID string `xml:"partId,attr,omitempty" json:"@partId,omitempty"`

	// Tor provide a human readable name for a part.
	PartName string `xml:"partName,attr,omitempty" json:"@partName,omitempty"`

	// Tor provide a human definition name for a part.
	PartDefinition string `xml:"partDefinition,attr,omitempty" json:"@partDefinition,omitempty"`

	// To provide a number associated with a part.
	PartNumber int `xml:"partNumber,attr,omitempty" json:"@partNumber,omitempty"`

	// To indicate the total number of parts of a particular type.
	PartTotalNumber int `xml:"partTotalNumber,attr,omitempty" json:"@partTotalNumber,omitempty"`
}

// The name given to a resource e.g. a media item, media object, sequence. For a series – use the series title; for a
// programme – a programme title; for an item – an item title. etc. Titles are recorded as they appear.
type Title struct {
	// A group of attributes provided to facilitate mapping from SMPTECore. Typed titles are preferably covered by
	// alternativeTitle. In EBUCore, title should be considered as being the MAIN title by default.
	TypeAttributes

	// Defines the date of attribution and / or deprecation of this Title.
	DateAttributes

	// The EBU core metadata set is built as a refinement of the Dublin Core. Free-text to provide the main title by which
	// the resource is known. The title can be provided in different languages. The language in which the title is
	// provided can be provided using elementType's lang attribute. Example: 'the fifth element'
	// TODO: required
	DCTitle []dc.Title `xml:"http://purl.org/dc/elements/1.1/ dc:title,omitempty" json:"dc:title,omitempty"`

	// To express the maximum length as a number of characters. This information can also be found in the title type e.g.
	// as "Title64"
	Length uint `xml:"length,attr,omitempty" json:"@length,omitempty"`

	// A comma separated list of geographical locations (e.g. expressed using country or region codes) where the title
	// applies.
	GeographicalScope string `xml:"geographicalScope,attr,omitempty" json:"@geographicalScope,omitempty"`

	// A comma separated list of geographical locations (e.g. expressed using country or region codes) where the title
	// shall not be used.
	GeographicalExclusionScope string `xml:"geographicalExclusionScope,attr,omitempty" json:"@geographicalExclusionScope,omitempty"`

	// Optional additional contextual information.
	Note string `xml:"note,attr,omitempty" json:"@note,omitempty"`
}

// The name given to a resource e.g. a media item, media object, sequence. For a series – use the series title; for a
// programme – a programme title; for an item – an item title. etc. Titles are recorded as they appear.
type AlternativeTitle struct {
	// The typeGroup is used to define the type of alternative title. This can be an associated title like a series title.
	TypeAttributes

	// Defines the date of attribution and / or deprecation of this Title.
	DateAttributes

	// The statusGroup is used to define the status of the Title such as short, long, full, abridged, working,
	// transmission, published, international, subtitle, original, secondary, alternative, pledged, etc. The name of the
	// format can be provided in the form of a text label, or a link to a code of a classification scheme, optionally
	// accompanied by a definition. the status 'main' shall not be used for alternativeTitle as this applies to the Title
	// only.
	StatusAttributes

	// The EBU core metadata set is built as a refinement of the Dublin Core. Free-text to provide alternative titles by
	// which the resource is known. The language in which the title is provided can be provided using elementType's lang
	// attribute. Example: 'the fifth element'
	// TODO: required
	DCTitle []dc.Title `xml:"http://purl.org/dc/elements/1.1/ dc:title,omitempty" json:"dc:title,omitempty"`

	// To express the maximum length as a number of characters. This information can also be found in the title type e.g.
	// as "Title64".
	Length uint `xml:"length,attr,omitempty" json:"@length,omitempty"`

	// A comma separated list of geographical locations (e.g. expressed using country or region codes) where the title
	// applies.
	GeographicalScope string `xml:"geographicalScope,attr,omitempty" json:"@geographicalScope,omitempty"`

	// A comma separated list of geographical locations (e.g. expressed using country or region codes) where the title
	// shall not be used.
	GeographicalExclusionScope string `xml:"geographicalExclusionScope,attr,omitempty" json:"@geographicalExclusionScope,omitempty"`

	// Optional additional contextual information.
	Note string `xml:"note,attr,omitempty" json:"@note,omitempty"`
}

// A unique, unambiguous reference or identifier for a resource within a given context. Best practice is to identify the
// resource (whether analogue or digital) by means of a string or number corresponding to an established or formal
// identification system if one exists. Otherwise, use an identification method that is in use within your agency,
// station, production company, office, or institution. It is also possible to enter different but unique identifiers
// for the same resource.
type Identifier struct {
	// The typeGroup is used to define the type of Identifier e.g. Main or Alternative
	TypeAttributes

	// Use to define the format and possibly syntax of the identifier. Used in combination with the resource Identifier.
	// It can denote the agency or institution which specified or assigned it e.g. SMPTE UMID, ISO ISAN, IETF URI, ISRC,
	// custom.
	FormatAttributes

	// The EBU core metadata set is built as a refinement of the Dublin Core. Free text to provide an identifier. Example:
	// 06.0A.2B.34.01.01.01.01
	DCIdentifier dc.Identifier `xml:"http://purl.org/dc/elements/1.1/ dc:identifier,omitempty" json:"dc:identifier,omitempty"`

	// To identify the source of attribution of the identifier.
	Attributor []Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:attributor,omitempty" json:"ebucore:attributor,omitempty"`

	// Optional additional contextual information.
	Note string `xml:"note,attr,omitempty" json:"@note,omitempty"`
}

// The generalised topic that represents the intellectual content of the resource. Typically, a subject is expressed by
// keywords, key phrases, or even specific classification codes. Controlled vocabularies, authorities, or formal
// classification schemes may be employed when selecting descriptive subject terms. It is possible to employ both
// keywords, derived from a formal classification scheme, such as Dewey or UDC, and genres/subgenres such as those
// produced by TV-Anytime or Escort, to cover Subject(s) and Genre(s) and enter as appropriate Subject Type below.
// Persons as subjects are also placed here. Genre of the content is placed under element Type.
type Subject struct {
	// To define the source of reference for subject such as a reference document or classification scheme or the
	// framework within which a tag has been attributed.
	TypeAttributes

	// To express the subject in the form of free text.
	DCSubject []dc.Subject `xml:"http://purl.org/dc/elements/1.1/ dc:subject,omitempty" json:"dc:subject,omitempty"`

	// To alternatively express the subject using predefined terms expressed by classification codes. Reference data: -
	// Library of Congress Subject Heading (LCSH), Library of Congress Classification (LCC), Medical Subject Headings
	// (MeSH), Dewey Decimal Classification (DDC), Dansk decimalklassedeling 5.utgave (DK5), Klassifikasjonssystem för
	// svenska bibliotek (SAB), Universal Decimal Classification (UDC), Norske emneord -
	// http://cv.iptc.org/newscodes/subjectcode/. Example: http://cv.iptc.org/newscodes/subjectcode/#15065000
	SubjectCode *URIValue `xml:"urn:ebu:metadata-schema:ebucore ebucore:subjectCode,omitempty" json:"ebucore:subjectCode,omitempty"`

	// An optional definition. Example: 'the subject is about tennis (sport, game)'
	SubjectDefinition []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:subjectDefinition,omitempty" json:"ebucore:subjectDefinition,omitempty"`

	// To identify the source of attribution of the subject/tag.
	Attributor *Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:attributor,omitempty" json:"ebucore:attributor,omitempty"`

	// Optional additional contextual information.
	Note string `xml:"note,attr,omitempty" json:"@note,omitempty"`
}

// A topic that represents the intellectual content of the resource.
type Topic struct {
	// To define the source of reference for a topic such as a reference document or classification scheme or the
	// framework within which a topic has been defined.
	TypeAttributes

	// To identify the source of attribution of the topic.
	Attributor *Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:attributor,omitempty" json:"ebucore:attributor,omitempty"`

	// Optional additional contextual information.
	Note string `xml:"note,attr,omitempty" json:"@note,omitempty"`
}

// A theme covered in the intellectual content of the resource.
type Theme struct {
	// To define the source of reference for a theme such as a reference document or classification scheme or the
	// framework within which a theme has been defined.
	TypeAttributes

	// To identify the source of attribution of the theme.
	Attributor *Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:attributor,omitempty" json:"ebucore:attributor,omitempty"`

	// Optional additional contextual information.
	Note string `xml:"note,attr,omitempty" json:"@note,omitempty"`
}

// The nature or genre of the content of the resource. Type includes terms
// describing general categories, functions, genres, or aggregation levels for content.
// Recommended best practice is to select a value from a controlled vocabulary. To
// describe the physical or digital manifestation of the resource, use the FORMAT
// element.
type Type struct {
	// To allow defining a type of dc:type other than 'genre', 'objectType', or 'targetAudience'.
	TypeAttributes

	// The EBU core metadata set is built as a refinement of the Dublin Core.
	DCType []dc.Type `xml:"http://purl.org/dc/elements/1.1/ dc:type,omitempty" json:"dc:type,omitempty"`

	// A type element specifically dedicated to the description of a resource genre
	Genre []Genre `xml:"urn:ebu:metadata-schema:ebucore ebucore:genre,omitempty" json:"ebucore:genre,omitempty"`

	// A type element specifically dedicated to the description of type of resource being describe e.g. programme, item,
	// episode, series, record
	ObjectType []ObjectType `xml:"urn:ebu:metadata-schema:ebucore ebucore:objectType,omitempty" json:"ebucore:objectType,omitempty"`

	// A type element specifically dedicated to the description of a target audience. The group of format attributes is
	// used to identify the rating/classification system. The group of type attributes is used to define the
	// classification purpose.
	TargetAudience []TargetAudience `xml:"urn:ebu:metadata-schema:ebucore ebucore:targetAudience,omitempty" json:"ebucore:targetAudience,omitempty"`

	// A type element specifically dedicated to the description of an audience level e.g. a geographical/occupational
	// group or an education level. The group of format attributes is used to identify the rating/classification system.
	// The group of type attributes is used to define the classification purpose.
	AudienceLevel []AudienceLevel `xml:"urn:ebu:metadata-schema:ebucore ebucore:audienceLevel,omitempty" json:"ebucore:audienceLevel,omitempty"`

	// A type element specifically dedicated to the description of a content format e.g. documentary
	ContentFormat []ContentFormat `xml:"urn:ebu:metadata-schema:ebucore ebucore:contentFormat,omitempty" json:"ebucore:contentFormat,omitempty"`

	// Optional additional contextual information.
	Note string `xml:"note,attr,omitempty" json:"@note,omitempty"`
}

type Genre struct {
	// To define the Type reference data.
	TypeAttributes

	// To provide an optional ranking position in case multiple genres are associated with the content.
	Level string `xml:"level,attr,omitempty" json:"@level,omitempty"`
}

type ObjectType struct {
	// To define the Type reference data.
	TypeAttributes
}

type ContentFormat struct {
	// To define the Type reference data.
	TypeAttributes

	// To provide an optional ranking position in case multiple formats are associated with the content.
	Level string `xml:"level,attr,omitempty" json:"@level,omitempty"`
}

type TargetAudience struct {
	// To define the Type of reference data.
	TypeAttributes

	// To define the audience definition system used.
	FormatAttributes

	// To describe the region to which the target audience restrictions apply.
	TargetRegion []Region `xml:"urn:ebu:metadata-schema:ebucore ebucore:targetRegion,omitempty" json:"ebucore:targetRegion,omitempty"`

	// To describe the region to which the target audience restrictions shall not apply.
	TargetExclusionRegion []Region `xml:"urn:ebu:metadata-schema:ebucore ebucore:targetExclusionRegion,omitempty" json:"ebucore:targetExclusionRegion,omitempty"`

	// An attribute to optionnally explain why this target audience has been selected
	Reason string `xml:"reason,attr,omitempty" json:"@reason,omitempty"`

	// A link to a visual representation of the target audience rating, if available
	LinkToLogo base.URI `xml:"linkToLogo,attr,omitempty" json:"@linkToLogo,omitempty"`

	// A flag to signal that content has not been rated (if set to "true")
	NotRated bool `xml:"notRated,attr,omitempty" json:"@notRated,omitempty"`

	// A flag to signal if content is adult content (if set to "true") or not for easy identification
	AdultContent bool `xml:"adultContent,attr,omitempty" json:"@adultContent,omitempty"`
}

type AudienceLevel struct {
	// To define the Type reference data.
	TypeAttributes

	// To define the rating system used.
	FormatAttributes

	// To describe the region to which the target audience restrictions apply.
	TargetRegion []Region `xml:"urn:ebu:metadata-schema:ebucore ebucore:targetRegion,omitempty" json:"ebucore:targetRegion,omitempty"`

	// To describe the region to which the target audience restrictions shall not apply.
	TargetExclusionRegion []Region `xml:"urn:ebu:metadata-schema:ebucore ebucore:targetExclusionRegion,omitempty" json:"ebucore:targetExclusionRegion,omitempty"`

	// An attribute to optionnally explain why this target audience has been selected
	Reason string `xml:"reason,attr,omitempty" json:"@reason,omitempty"`

	// A link to a visual representation of the target audience rating, if available
	LinkToLogo base.URI `xml:"linkToLogo,attr,omitempty" json:"@linkToLogo,omitempty"`

	// A flag to signal that content has not been rated (if set to "true")
	NotRated bool `xml:"notRated,attr,omitempty" json:"@notRated,omitempty"`
}

// Free-form text or a narrative to report general notes, abstracts, or summaries about the intellectual content of a
// resource. The information may be in the form of a paragraph giving an individual program description, anecdotal
// interpretations, or brief content reviews. The description may also consist of outlines, lists, bullet points, edit
// decision lists, indexes, or tables of content, a reference to a graphical representation of content or even a pointer
// (URI, URL) to an external resource. For a Radio or television programme a running order can be used as description. A
// description can be provided in different languages.
type Description struct {
	// To define the form of presentation for the information: Annotation, abstract, summary, review, table of content,
	// synopsis, shot list, edit decision list, promotional information, purpose, script, outline, rundown,
	// selection/excerpt, transcript, bookmarks, theme, highlights, running order, etc.
	TypeAttributes

	// Defines the date of attribution and / or deprecation of this Description.
	DateAttributes

	// The EBU core metadata set is built as a refinement of the Dublin Core. Free text to provide a description of the
	// resource. The description can be repeated in different languages as specified by the entityType's lang attribute.
	// The type of description is defined in the type group of attributes.
	DCDescription []dc.Description `xml:"http://purl.org/dc/elements/1.1/ dc:description,omitempty" json:"dc:description,omitempty"`

	// To identify the source of attribution of the description.
	Attributor *Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:attributor,omitempty" json:"ebucore:attributor,omitempty"`

	// To express the maximum length as a number of characters. This information can also sometimes be found in the
	// description type e.g. as "Description150".
	Length uint `xml:"length,attr,omitempty" json:"@length,omitempty"`

	// A comma separated list of geographical locations (e.g. expressed using country or region codes) where the
	// description applies.
	GeographicalScope string `xml:"geographicalScope,attr,omitempty" json:"@geographicalScope,omitempty"`

	// A comma separated list of geographical locations (e.g. expressed using country or region codes) where the
	// description shall not be used.
	GeographicalExclusionScope string `xml:"geographicalExclusionScope,attr,omitempty" json:"@geographicalExclusionScope,omitempty"`

	// A flag to indicate if cast information is provided in the description.
	CastFlag bool `xml:"castFlag,attr,omitempty" json:"@castFlag,omitempty"`

	// Optional additional contextual information.
	Note string `xml:"note,attr,omitempty" json:"@note,omitempty"`
}

// Coverage will typically include spatial location (a place name or geographic coordinates), temporal period (a period
// label, date, or date range) or jurisdiction (such as a named administrative entity). Recommended best practice is to
// select a value from a controlled vocabulary (for example, the Thesaurus of Geographic Names) and that, where
// appropriate, named places or time periods be used in preference to numeric identifiers such as sets of coordinates or
// date ranges.
type Coverage struct {
	// To specify a type of coverage in association with dc:coverage
	TypeAttributes

	// The EBU core metadata set is built as a refinement of the Dublin Core.
	DCCoverage *dc.Coverage `xml:"http://purl.org/dc/elements/1.1/ dc:coverage,omitempty" json:"dc:coverage,omitempty"`

	// Temporal characteristics of the content of the resource. To indicate e.g. specific date, time or period aspects of
	// the subject of the resource in complement to Description.
	Temporal *Temporal `xml:"urn:ebu:metadata-schema:ebucore ebucore:temporal,omitempty" json:"ebucore:temporal,omitempty"`

	// Spatial characteristics of the content of the resource. To indicate e.g. specific place and location aspects of the
	// subject of the resource in complement to Description.
	Spatial *Spatial `xml:"urn:ebu:metadata-schema:ebucore ebucore:spatial,omitempty" json:"ebucore:spatial,omitempty"`
}

type Temporal struct {
	// To precise the type of temporal information provided.
	TypeAttributes

	// To specify a period of time as a delimited time period or by e.g. by its name.
	PeriodOfTime []PeriodOfTime `xml:"urn:ebu:metadata-schema:ebucore ebucore:PeriodOfTime,omitempty" json:"ebucore:PeriodOfTime,omitempty"`

	// Optional additional contextual information.
	Note string `xml:"note,attr,omitempty" json:"@note,omitempty"`
}

type Spatial struct {
	// To specify a location by its name of geographical coordinates.
	Location []Location `xml:"urn:ebu:metadata-schema:ebucore ebucore:location,omitempty" json:"ebucore:location,omitempty"`
}

// To indicate e.g. specific place and location aspects of the resource in complement to Description.
type Location struct {
	// To precise the type of place and location.
	TypeAttributes

	// Any location name in free text plus type
	Name []Name `xml:"urn:ebu:metadata-schema:ebucore ebucore:name,omitempty" json:"ebucore:name,omitempty"`

	// A description of a location name in free text
	Description []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:description,omitempty" json:"ebucore:description,omitempty"`

	// Optional geospatial coordinates. 'posy' is the latitude. 'posx' is the longitude. Both are expressed in digital
	// degrees
	Coordinates *Coordinates `xml:"urn:ebu:metadata-schema:ebucore ebucore:coordinates,omitempty" json:"ebucore:coordinates,omitempty"`

	// A location identified by a code from a predefined list of locations.
	Code []Code `xml:"urn:ebu:metadata-schema:ebucore ebucore:code,omitempty" json:"ebucore:code,omitempty"`

	// A location specifically defined as a region, i.e. a country or countryRegion
	Region []Region `xml:"urn:ebu:metadata-schema:ebucore ebucore:region,omitempty" json:"ebucore:region,omitempty"`

	// To indicate the altutude at a given location.
	Altitude *Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:altitude,omitempty" json:"ebucore:altitude,omitempty"`

	// To provide additional custom information about a location.
	AdditionalInformation []AdditionalInformation `xml:"urn:ebu:metadata-schema:ebucore ebucore:additionalInformation,omitempty" json:"ebucore:additionalInformation,omitempty"`

	// To provide additional information on the type of location described, e.g. countries, regions, cities
	Note string `xml:"note,attr,omitempty" json:"@note,omitempty"`

	// An identifier to support the management of location in databases and RDF
	LocationID base.URI `xml:"locationId,attr,omitempty" json:"@locationId,omitempty"`
}

type Name struct {
	Element
	TypeAttributes
}

type Coordinates struct {
	// To provide additional information on the format in which the coordiantes are expressed.
	FormatAttributes

	// The ordinate of geographical coordinates.
	PosY Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:posy,omitempty" json:"ebucore:posy,omitempty"`

	// The abscisse of geographical coordinates.
	PosX Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:posx,omitempty" json:"ebucore:posx,omitempty"`
}

type Code struct {
	URIValue
	TypeAttributes
}

// To describe a period of time e.g. historical.
type PeriodOfTime struct {
	// The period of time depicted in the resource expressed by date an time delimitated boundaries.
	DateAttributes

	// To provide a name for a time period e.g. neolithic.
	PeriodName []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:periodName,omitempty" json:"ebucore:periodName,omitempty"`

	// An identifier to support the management of time periods (e.g. historical or repetitive event) in databases and RDF.
	PeriodID base.URI `xml:"periodId,attr,omitempty" json:"@periodId,omitempty"`
}

// An all-purpose field to identify information (rights management statement or reference to a service providing such
// information e.g. via a URL) about copyright, intellectual property rights or other property rights held in and over a
// resource, stating whether open access or restricted in some way. If dates, times, territories and availability
// periods are associated with a right, they should be included. If the Rights element is absent, no assumptions can be
// made about the status of these and other rights with respect to the resource.
type Rights struct {
	// To define the type of rights information provided.
	TypeAttributes

	// The EBU core metadata set is built as a refinement of the Dublin Core. An element to express any form of rights
	// related matters.
	DCRights []dc.Rights `xml:"http://purl.org/dc/elements/1.1/ dc:rights,omitempty" json:"dc:rights,omitempty"`

	// A url pointing to a declaration of rights
	RightsLink *URIValue `xml:"urn:ebu:metadata-schema:ebucore ebucore:rightsLink,omitempty" json:"ebucore:rightsLink,omitempty"`

	// To identify the person or organisation holding or managing the rights related to the resource.
	RightsHolder []Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:rightsHolder,omitempty" json:"ebucore:rightsHolder,omitempty"`

	// Use to state any other restrictions, such as non-rights ones, e.g. legal. State by media, territory, scope
	// (restriction on whole item or extracts) and possibly language. The presence of this information can be used by
	// asset management system implementing traffic lights like mechanism to signal that content may be subject to
	// particular restrictions to be clarified before exploitation.
	ExploitationIssues []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:exploitationIssues,omitempty" json:"ebucore:exploitationIssues,omitempty"`

	// Use to report any copyright statement, specifically
	CopyrightStatement []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:copyrightStatement,omitempty" json:"ebucore:copyrightStatement,omitempty"`

	// Specifies a specific start date, end date or period for the availability of the item or the date from which the
	// rights or exploitation issues apply. It may refer to start dates for the availability of an item that is used
	// within a particular geographical area e.g. broadcast locally, regionally, nationally or internationally, or for
	// web-based distribution. A specific time may also be associated with the date.
	Coverage []Coverage `xml:"urn:ebu:metadata-schema:ebucore ebucore:coverage,omitempty" json:"ebucore:coverage,omitempty"`

	// A flag to signal if content is subject to rights related open issues, or not.
	RightsClearanceFlag *Boolean `xml:"urn:ebu:metadata-schema:ebucore ebucore:rightsClearanceFlag,omitempty" json:"ebucore:rightsClearanceFlag,omitempty"`

	// A flag to signal if content can be processed or transformed and under which conditions. This may include the rights
	// to perform (or not) technical operations like denoising with a potential impact on the artistic perception
	// controlled by the content owner.
	ProcessingRestrictionFlag *ProcessingRestrictionFlag `xml:"urn:ebu:metadata-schema:ebucore ebucore:processingRestrictionFlag,omitempty" json:"ebucore:processingRestrictionFlag,omitempty"`

	// A field for a disclaimer about the content, its content, and its use.
	Disclaimer []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:disclaimer,omitempty" json:"ebucore:disclaimer,omitempty"`

	// A identifier related to rights, e.g. attributed for a particular purpose by a specific agency in the context of use
	// and exploitation.
	RightsAttributedID []Identifier `xml:"urn:ebu:metadata-schema:ebucore ebucore:rightsAttributedId,omitempty" json:"ebucore:rightsAttributedId,omitempty"`

	// Minimum information providing means to further identify and contact the rights manager for the resource in the
	// organisation.
	ContactDetails []ContactDetails `xml:"urn:ebu:metadata-schema:ebucore ebucore:contactDetails,omitempty" json:"ebucore:contactDetails,omitempty"`

	// To define the type/encoding of rights e.g. mpeg-21 ODRL...
	RightsEncoding *RightsEncoding `xml:"urn:ebu:metadata-schema:ebucore ebucore:rightsEncoding,omitempty" json:"ebucore:rightsEncoding,omitempty"`

	// Optional additional contextual information.
	Note string `xml:"note,attr,omitempty" json:"@note,omitempty"`

	// A list of ID references identifying formats in which the content is available and to which the set of rights apply.
	FormatIDRefs base.URI `xml:"formatIDRefs,attr,omitempty" json:"@formatIDRefs,omitempty"`

	// An ID to be used as reference to the associated rights
	RightsID base.URI `xml:"rightsID,attr,omitempty" json:"@rightsID,omitempty"`
}

type ProcessingRestrictionFlag struct {
	TypeAttributes

	Value bool `xml:",chardata" json:"#value"`

	QualityClass string `xml:"qualityClass,attr,omitempty" json:"@qualityClass,omitempty"`
	Restrictions string `xml:"restrictions,attr,omitempty" json:"@restrictions,omitempty"`
}

type RightsEncoding struct {
	TypeAttributes
}

// The description of an event and an optional related location.
type Event struct {
	// To specify a type of coverage in association with dc:coverage
	TypeAttributes

	// The name of an event.
	Name []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:name,omitempty" json:"ebucore:name,omitempty"`

	// The description of an event.
	Description []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:description,omitempty" json:"ebucore:description,omitempty"`

	// A location related to the event.
	Location []Location `xml:"urn:ebu:metadata-schema:ebucore ebucore:location,omitempty" json:"ebucore:location,omitempty"`

	// An identifier to support the management of location in databases and RDF
	EventID base.URI `xml:"eventId,attr,omitempty" json:"@eventId,omitempty"`

	// The date when the event started.
	Start base.Date `xml:"start,attr,omitempty" json:"@start,omitempty"`

	// The date when the event ended.
	End base.Date `xml:"end,attr,omitempty" json:"@end,omitempty"`

	// Optional additional contextual information.
	Note string `xml:"note,attr,omitempty" json:"@note,omitempty"`
}

// Recommended best practice is to reference the resource by means of a string or number conforming to a formal
// identification system. Relation is used to show the relation in content to another resource. For example, "IsPartOf"
// is used to show the relation between a part of a radio programme and the whole programme, or between a track and a
// record album. A resource can be identified by its title, or preferably by an identifier. Relation is used to provide
// a name, locator, accession, identification number or ID where the related item can be obtained or found.
type Relation struct {
	// To show the type of relation to another resource, e.g. identifies
	// ways in which the resource is related by intellectual content to some other
	// resource.
	TypeAttributes

	// The EBU core metadata set is built as a refinement of the Dublin Core. A title would be given using this element.
	// TODO: either DCRelation or RelationIdentifier or RelationLink must be specified
	DCRelation *dc.Relation `xml:"http://purl.org/dc/elements/1.1/ dc:relation,omitempty" json:"dc:relation,omitempty"`

	// An identifier would be given using this element.
	// TODO: either DCRelation or RelationIdentifier or RelationLink must be specified
	RelationIdentifier *Identifier `xml:"urn:ebu:metadata-schema:ebucore ebucore:relationIdentifier,omitempty" json:"ebucore:relationIdentifier,omitempty"`

	// A link to related material.
	// TODO: either DCRelation or RelationIdentifier or RelationLink must be specified
	RelationLink *URIValue `xml:"urn:ebu:metadata-schema:ebucore ebucore:relationLink,omitempty" json:"ebucore:relationLink,omitempty"`

	RelationSource []Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:relationSource,omitempty" json:"ebucore:relationSource,omitempty"`

	// If exists, it provides the ranking/running order within an ordered list.
	RunningOrderNumber int `xml:"runningOrderNumber,attr,omitempty" json:"@runningOrderNumber,omitempty"`

	TotalNumberOfGroupMembers int `xml:"totalNumberOfGroupMembers,attr,omitempty" json:"@totalNumberOfGroupMembers,omitempty"`

	OrderedGroupFlag bool `xml:"orderedGroupFlag,attr,omitempty" json:"@orderedGroupFlag,omitempty"`

	// Optional additional contextual information.
	Note string `xml:"note,attr,omitempty" json:"@note,omitempty"`
}

// Identifies languages and their use in the intellectual content of the resource. Recommended best practice for the
// values of the Language element is defined by RFC 5646.
type Language struct {
	// Indicates the purpose of the language described by the Language element e.g. Main original language, main dubbed
	// language, additional original language, additional dubbed language, descriptive video information, supplemental
	// commentary, Director's commentary, audio description, supplementary audio programme, educational notes, voice over,
	// original commentary, dubbed commentary, original narration, dubbed narration, dubbed dialogue, interviewer
	// language, interviewee language, text description for the hard-of-hearing, titles, subtitles, song lyrics, sign
	// language, dubbed sign language, transcript, caption, open caption, closed caption.
	TypeAttributes

	// The EBU core metadata set is built as a refinement of the Dublin Core.
	DCLanguage []dc.Language `xml:"http://purl.org/dc/elements/1.1/ dc:language,omitempty" json:"dc:language,omitempty"`

	// Optional additional contextual information.
	Note string `xml:"note,attr,omitempty" json:"@note,omitempty"`
}

// Dates associated with events occurring during the life of the resource. Typically, ResourceDates will be associated
// e.g. with the creation or availability of the resource.
type ResourceDates struct {
	// to define a type of date associated with dc:date
	TypeAttributes

	// to define the format of date used in dc:date
	FormatAttributes

	// The EBU core metadata set is built as a refinement of the Dublin Core.
	DCDate []dc.Date `xml:"http://purl.org/dc/elements/1.1/ dc:date,omitempty" json:"dc:date,omitempty"`

	// To specify the creation date for a particular version or rendition of a resource across its life cycle. It is the
	// moment in time that the resource was finalized during its production process and is forwarded to other divisions or
	// agencies to make it ready for publication or distribution. A specific time may also be associated with the date.
	Created *Created `xml:"urn:ebu:metadata-schema:ebucore ebucore:created,omitempty" json:"ebucore:created,omitempty"`

	// Date of formal issuance (e.g. publication) of the resource. Specifies the formal date for a particular version or
	// rendition of a resource has been made ready or officially released for distribution, publication or consumption,
	// e.g. the broadcasting date of a radio programme. A specific time may also be associated with the date.
	Issued *Issued `xml:"urn:ebu:metadata-schema:ebucore ebucore:issued,omitempty" json:"ebucore:issued,omitempty"`

	// The date and optionally time when the resource was last modified
	Modified *Modified `xml:"urn:ebu:metadata-schema:ebucore ebucore:modified,omitempty" json:"ebucore:modified,omitempty"`

	// The date and optionally time when the resource was digitised
	Digitised *Digitised `xml:"urn:ebu:metadata-schema:ebucore ebucore:digitised,omitempty" json:"ebucore:digitised,omitempty"`

	// The date and optionally time when the resource was released
	Released []Released `xml:"urn:ebu:metadata-schema:ebucore ebucore:released,omitempty" json:"ebucore:released,omitempty"`

	// The date and optionally time when the resource was copyrighted
	Copyrighted *Copyrighted `xml:"urn:ebu:metadata-schema:ebucore ebucore:copyrighted,omitempty" json:"ebucore:copyrighted,omitempty"`

	// The date and optionally time when the resource was encoded in a particular format.
	Encoded *Encoded `xml:"urn:ebu:metadata-schema:ebucore ebucore:encoded,omitempty" json:"ebucore:encoded,omitempty"`

	// An alternative particular date and optionally time for which the type can be defined.
	Alternative []AlternativeDate `xml:"urn:ebu:metadata-schema:ebucore ebucore:alternative,omitempty" json:"ebucore:alternative,omitempty"`

	// The date and optionally time when the resource was ingested.
	Ingested *Ingested `xml:"urn:ebu:metadata-schema:ebucore ebucore:ingested,omitempty" json:"ebucore:ingested,omitempty"`

	// The date and optionally time when the resource has been archived.
	Archived *Archived `xml:"urn:ebu:metadata-schema:ebucore ebucore:archived,omitempty" json:"ebucore:archived,omitempty"`

	// The date and optionally time when the resource has been deleted.
	Deleted *Deleted `xml:"urn:ebu:metadata-schema:ebucore ebucore:deleted,omitempty" json:"ebucore:deleted,omitempty"`

	// The date and optionally time when the resource has been produced.
	Produced *Produced `xml:"urn:ebu:metadata-schema:ebucore ebucore:produced,omitempty" json:"ebucore:produced,omitempty"`

	// The date and optionally time when the creation of the resource has been planned.
	Planned *Planned `xml:"urn:ebu:metadata-schema:ebucore ebucore:planned,omitempty" json:"ebucore:planned,omitempty"`

	Note *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:note,omitempty" json:"ebucore:note,omitempty"`

	// To provide information on the precision / exactness of the date being provided.
	Precision string `xml:"precision,attr,omitempty" json:"@precision,omitempty"`
}

type Created struct {
	// A set of attributes to express a date, time or a period.
	DateAttributes
}

type Issued struct {
	// A set of attributes to express a date, time or a period.
	DateAttributes
}

type Modified struct {
	// A set of attributes to express a date, time or a period.
	DateAttributes
}

type Digitised struct {
	// A set of attributes to express a date, time or a period.
	DateAttributes
}

type Released struct {
	// A set of attributes to express a date, time or a period.
	DateAttributes
}

type Copyrighted struct {
	// A set of attributes to express a date, time or a period.
	DateAttributes
}

type Encoded struct {
	// A set of attributes to express a date, time or a period.
	DateAttributes
}

type Ingested struct {
	// A set of attributes to express a date, time or a period.
	DateAttributes
}

type Archived struct {
	// A set of attributes to express a date, time or a period.
	DateAttributes
}

type Deleted struct {
	// A set of attributes to express a date, time or a period.
	DateAttributes
}

type Produced struct {
	// A set of attributes to express a date, time or a period.
	DateAttributes
}

type Planned struct {
	// A set of attributes to express a date, time or a period.
	DateAttributes
}

// An alternative particular date and optionally time for which the type can be defined.
type AlternativeDate struct {
	// A set of attributes to define a custom type of date other than predefined in the schema.
	TypeAttributes

	// A set of attributes to express a date or a date period.
	DateAttributes
}

// To provide information about the publication history.
type PublicationHistory struct {
	// To describe a publication event.
	PublicationEvent []PublicationEvent `xml:"urn:ebu:metadata-schema:ebucore ebucore:publicationEvent,omitempty" json:"ebucore:publicationEvent,omitempty"`

	PublicationHistoryID base.URI `xml:"publicationHistoryId,attr,omitempty" json:"@publicationHistoryId,omitempty"`
}

// To provide information about the publication history.
type Planning struct {
	// To describe a publication event.
	PublicationEvent []PublicationEvent `xml:"urn:ebu:metadata-schema:ebucore ebucore:publicationEvent,omitempty" json:"ebucore:publicationEvent,omitempty"`

	PlanningID base.URI `xml:"planningId,attr,omitempty" json:"@planningId,omitempty"`
}

// To describe when, where, in which formats and under which rights conditions the resource has been distributed.
type PublicationEvent struct {
	// The first publication date
	PublicationDate *Date `xml:"urn:ebu:metadata-schema:ebucore ebucore:publicationDate,omitempty" json:"ebucore:publicationDate,omitempty"`

	// The publication time
	PublicationTime *TimeValue `xml:"urn:ebu:metadata-schema:ebucore ebucore:publicationTime,omitempty" json:"ebucore:publicationTime,omitempty"`

	// The publication duration
	PublicationDuration *DurationValue `xml:"urn:ebu:metadata-schema:ebucore ebucore:publicationDuration,omitempty" json:"ebucore:publicationDuration,omitempty"`

	// To indicate the schedule date, which can differ from the actual publicationDate if it occurs after midnight of the
	// schedule date
	ScheduleDate *Date `xml:"urn:ebu:metadata-schema:ebucore ebucore:scheduleDate,omitempty" json:"ebucore:scheduleDate,omitempty"`

	// The service to which the channel is attached
	PublicationService *PublicationService `xml:"urn:ebu:metadata-schema:ebucore ebucore:publicationService,omitempty" json:"ebucore:publicationService,omitempty"`

	// To provide information on the medium used to distribute the resource.
	PublicationMedium *PublicationMedium `xml:"urn:ebu:metadata-schema:ebucore ebucore:publicationMedium,omitempty" json:"ebucore:publicationMedium,omitempty"`

	// To identify the channel through which the resource has been distributed.
	PublicationChannel *PublicationChannel `xml:"urn:ebu:metadata-schema:ebucore ebucore:publicationChannel,omitempty" json:"ebucore:publicationChannel,omitempty"`

	// To identify the region where the resource has been distributed.
	PublicationRegion []Region `xml:"urn:ebu:metadata-schema:ebucore ebucore:publicationRegion,omitempty" json:"ebucore:publicationRegion,omitempty"`

	// To list publication events related to the current publication event
	RelatedPublicationEvent []PublicationEvent `xml:"urn:ebu:metadata-schema:ebucore ebucore:relatedPublicationEvent,omitempty" json:"ebucore:relatedPublicationEvent,omitempty"`

	// To uniquely identify a pèublication event.
	PublicationEventID base.URI `xml:"publicationEventId,attr,omitempty" json:"@publicationEventId,omitempty"`

	// To provide a human readable name for a publication event.
	PublicationEventName string `xml:"publicationEventName,attr,omitempty" json:"@publicationEventName,omitempty"`

	// A flag indicating if the resource was first shown during this publication event.
	FirstShowing bool `xml:"firstShowing,attr,omitempty" json:"@firstShowing,omitempty"`

	// A flag indicating if the resource was last shown during this publication event.
	LastShowing bool `xml:"lastShowing,attr,omitempty" json:"@lastShowing,omitempty"`

	// A flag indicating if the resource is shown live during this publication event. Production and transmission are
	// live.
	Live bool `xml:"live,attr,omitempty" json:"@live,omitempty"`

	// A flag indicating if the resource was captured live for this publication event. Transmission is not live (see live
	// flag). Content is recorded.
	LiveProduction bool `xml:"liveProduction,attr,omitempty" json:"@liveProduction,omitempty"`

	// A flag indicating if the resource was shown with or wothout access restrictions during this publication event.
	Free bool `xml:"free,attr,omitempty" json:"@free,omitempty"`

	// A note to provide additional contextual information.
	Note string `xml:"note,attr,omitempty" json:"@note,omitempty"`

	// to identify the format in which content has been published
	FormatIdRef base.URI `xml:"formatIdRef,attr,omitempty" json:"@formatIdRef,omitempty"`

	// to identify all the rights associated with the publication event
	RightsIDRefs base.URI `xml:"rightsIDRefs,attr,omitempty" json:"@rightsIDRefs,omitempty"`
}

// To provide information on the service associated with the publication.
type PublicationService struct {
	// A group of attribute to specify the type of service.
	TypeAttributes

	// A human readable name of the service.
	PublicationServiceName *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:publicationServiceName,omitempty" json:"ebucore:publicationServiceName,omitempty"`

	// To provide information on the organisation who provided the resource for distribution.
	PublicationSource *OrganisationDetails `xml:"urn:ebu:metadata-schema:ebucore ebucore:publicationSource,omitempty" json:"ebucore:publicationSource,omitempty"`

	// To uniquely identify the service.
	ServiceID base.URI `xml:"serviceId,attr,omitempty" json:"@serviceId,omitempty"`

	// A logo to visually identifiy the publication service
	LinkToLogo base.URI `xml:"linkToLogo,attr,omitempty" json:"@linkToLogo,omitempty"`
}

// The medium on which the title was published
type PublicationMedium struct {
	// To provide additional information on the type of medium on which the title was published
	TypeAttributes

	Value string `xml:",chardata" json:"#value"`

	// To uniquely identify the medium.
	PublicationMediumID base.URI `xml:"publicationMediumId,attr,omitempty" json:"@publicationMediumId,omitempty"`
}

// The channel on which the title was transmitted
type PublicationChannel struct {
	// To provide additional information on the type of channel on which the title was published by the service e.g.
	// online, broadcast
	TypeAttributes

	Value string `xml:",chardata" json:"#value"`

	// To uniquely identify the publication channel.
	PublicationChannelID base.URI `xml:"publicationChannelId,attr,omitempty" json:"@publicationChannelId,omitempty"`

	// A logo to visually identifiy the publication channel
	LinkToLogo base.URI `xml:"linkToLogo,attr,omitempty" json:"@linkToLogo,omitempty"`
}

// The physical or digital manifestation of the resource. Use the descriptor Format to identify the format of a
// particular resource as it exists in its physical or digital form. Physical form = an actual physical form that
// occupies physical space, e.g. a tape. Digital form = a digital file residing on a server or hard drive. Format may be
// used to determine the software, hardware or other equipment needed to display or operate the resource.
type Format struct {
	// To provide information on the type of format. This can be used to facilitate mapping with other schemas. Example:
	// typeLabel could be given the value "instantiation" to illustrate that format in ebucore is the same as
	// 'instantiation in pbcore
	TypeAttributes

	// A point of extension for customisation using a set of technical attributes of predefined datatypes.
	TechnicalAttributes

	// A group of elements to provide general information on a resource instantiation / file.
	FileInfo

	// The EBU core metadata set is built as a refinement of the Dublin
	// Core. Free text to provide information on the format.
	DCFormat *dc.Format `xml:"http://purl.org/dc/elements/1.1/ dc:format,omitempty" json:"dc:format,omitempty"`

	// The material or physical carrier of the resource. If a file, it should be the carrier format.
	Medium []Medium `xml:"urn:ebu:metadata-schema:ebucore ebucore:medium,omitempty" json:"ebucore:medium,omitempty"`

	// Used to list the characteristics of an image.
	ImageFormat []ImageFormat `xml:"urn:ebu:metadata-schema:ebucore ebucore:imageFormat,omitempty" json:"ebucore:imageFormat,omitempty"`

	// Used to list all the characteristics of the video signal.
	VideoFormat []VideoFormat `xml:"urn:ebu:metadata-schema:ebucore ebucore:videoFormat,omitempty" json:"ebucore:videoFormat,omitempty"`

	// Used to list all the characteristics of the audio signal.
	AudioFormat []AudioFormat `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioFormat,omitempty" json:"ebucore:audioFormat,omitempty"`

	// Used to list all the extended audio model characteristics of the audio signal. The extended audio format model is
	// specified in details in Tech.3364.
	AudioFormatExtended []AudioFormatExtended `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioFormatExtended,omitempty" json:"ebucore:audioFormatExtended,omitempty"`

	// To provide information on the wrapper format in complement to the stream encoding information provided in
	// 'channel', e.g. mp3, wave, Quicktime, ogg.
	ContainerFormat []ContainerFormat `xml:"urn:ebu:metadata-schema:ebucore ebucore:containerFormat,omitempty" json:"ebucore:containerFormat,omitempty"`

	// To provide the characteristics of signing, when present.
	SigningFormat *SigningFormat `xml:"urn:ebu:metadata-schema:ebucore ebucore:signingFormat,omitempty" json:"ebucore:signingFormat,omitempty"`

	// To provide information on the data provided with the resource and its format.
	DataFormat []DataFormat `xml:"urn:ebu:metadata-schema:ebucore ebucore:dataFormat,omitempty" json:"ebucore:dataFormat,omitempty"`

	// To provide information on the timecode provided with the resource and its format.
	TimecodeFormat []TimecodeFormat `xml:"urn:ebu:metadata-schema:ebucore ebucore:timecodeFormat,omitempty" json:"ebucore:timecodeFormat,omitempty"`

	// To provide information on the metadata provided with the resource and its format.
	MetadataFormat []MetadataFormat `xml:"urn:ebu:metadata-schema:ebucore ebucore:metadataFormat,omitempty" json:"ebucore:metadataFormat,omitempty"`

	// A element to allow the serialisation of dynamic acquisition metadata. The user can serialise data in two forms:
	//
	// 1 - ParameterSegmentDataOutput parameter per parameter, each with a timeline of time segments along which the
	//     parameter value varies.
	//
	// 2 - SegmentParameterDataOutput a series of time segments for which all values are listed
	//     for a specifc parameter.
	AcquisitionData *[2]*AcquisitionData `xml:"urn:ebu:metadata-schema:ebucore ebucore:acquisitionData,omitempty" json:"ebucore:acquisitionData,omitempty"`

	// The hdrMetadata will be used in IMF as a separate specific side-car metadata file under the structure
	// ebuCoreMain/coreMetadata/format/hdrmetadata. There should therefore be no conflict with e.g. the videoFormat/height
	// and videoFormat/width as this data shouldn't coexist in the same metadata instance file.
	HDRMetadata *HDRMetadata `xml:"urn:ebu:metadata-schema:ebucore ebucore:hdrMetadata,omitempty" json:"ebucore:hdrMetadata,omitempty"`

	// The beginning point for playback of a time-based media item, such as digital video or audio. Use in combination
	// with end or duration to identify a sequence or segment of a media item that has a fixed start time and end time.
	Start []Time `xml:"urn:ebu:metadata-schema:ebucore ebucore:start,omitempty" json:"ebucore:start,omitempty"`

	// The ending point for playback of a time-based media item, such as digital video or audio.
	End []Time `xml:"urn:ebu:metadata-schema:ebucore ebucore:end,omitempty" json:"ebucore:end,omitempty"`

	// The duration of the resource or part of a resource.
	Duration []Duration `xml:"urn:ebu:metadata-schema:ebucore ebucore:duration,omitempty" json:"ebucore:duration,omitempty"`

	// An element to provide a description of a document
	DocumentFormat *DocumentFormat `xml:"urn:ebu:metadata-schema:ebucore ebucore:documentFormat,omitempty" json:"ebucore:documentFormat,omitempty"`

	// The date when a format has been defined.
	DateCreated *DateCreated `xml:"urn:ebu:metadata-schema:ebucore ebucore:dateCreated,omitempty" json:"ebucore:dateCreated,omitempty"`

	// A date when a format has been modified.
	DateModified *DateModified `xml:"urn:ebu:metadata-schema:ebucore ebucore:dateModified,omitempty" json:"ebucore:dateModified,omitempty"`

	// To uniquely identify a format.
	FormatID base.URI `xml:"formatId,attr,omitempty" json:"@formatId,omitempty"`

	// To provide information on the version of the format.
	FormatVersionID string `xml:"formatVersionId,attr,omitempty" json:"@formatVersionId,omitempty"`

	// To provide a human readable name for the format.
	FormatName string `xml:"formatName,attr,omitempty" json:"@formatName,omitempty"`

	// To provide a human readable name for the format.
	FormatDefinition string `xml:"formatDefinition,attr,omitempty" json:"@formatDefinition,omitempty"`
}

type Medium struct {
	// A group opf attributes nto identify the type of medium.
	TypeAttributes

	// To uniquely identify the medium.
	MediumID base.URI `xml:"mediumId,attr,omitempty" json:"@mediumId,omitempty"`
}

type AcquisitionData struct {
	// The timecode at which extraction starts expressed in timecode / fremes. This also provides additional
	// information on the timecode format.
	ExtractionStartTime *Timecode `xml:"urn:ebu:metadata-schema:ebucore ebucore:extractionStartTime,omitempty" json:"ebucore:extractionStartTime,omitempty"`

	// The duration of the extraction expressed in timecode / fremes. This also provides additional information on
	// the timecode format.
	ExtractionDuration *Timecode `xml:"urn:ebu:metadata-schema:ebucore ebucore:extractionDuration,omitempty" json:"ebucore:extractionDuration,omitempty"`

	// To provide information on the frame rate at which data is acquired.
	AcquisitionFrameRate Rational `xml:"urn:ebu:metadata-schema:ebucore ebucore:acquisitionFrameRate,omitempty" json:"ebucore:acquisitionFrameRate,omitempty"`

	// To export acquisition data with a timline per parameter.
	ParameterSegmentDataOutput *ParameterSegmentDataOutput `xml:"urn:ebu:metadata-schema:ebucore ebucore:parameterSegmentDataOutput,omitempty" json:"ebucore:parameterSegmentDataOutput,omitempty"`

	// To export acquisition data with a parameter per timeline.
	SegmentParameterDataOutput *SegmentParameterDataOutput `xml:"urn:ebu:metadata-schema:ebucore ebucore:segmentParameterDataOutput,omitempty" json:"ebucore:segmentParameterDataOutput,omitempty"`
}

type HDRMetadata struct {
	// To define the width of the image. Please check consistency with the videoFormat/width.
	Width Int `xml:"urn:ebu:metadata-schema:ebucore ebucore:width,omitempty" json:"ebucore:width,omitempty"`

	// To define the height of the image. Please check consistency with the videoFormat/height.
	Height Int `xml:"urn:ebu:metadata-schema:ebucore ebucore:height,omitempty" json:"ebucore:height,omitempty"`

	// To define the active area within the image as the ratio between the width and the height centered knowing
	// that the center of the active active area corrsponds to the centre of the image..
	ActiveArea *ActiveArea `xml:"urn:ebu:metadata-schema:ebucore ebucore:activeArea,omitempty" json:"ebucore:activeArea,omitempty"`

	// To establish the parameters of the master display color volume as defined in SMPTE 2086.
	MasteredColorVolume *MasteredColorVolume `xml:"urn:ebu:metadata-schema:ebucore ebucore:masteredColorVolume,omitempty" json:"ebucore:masteredColorVolume,omitempty"`

	// To set the parameters of the light level as specified in "CEA-861.3 Annex A".
	LightLevel LightLevel `xml:"urn:ebu:metadata-schema:ebucore ebucore:lightLevel,omitempty" json:"ebucore:lightLevel,omitempty"`
}

type ActiveArea struct {
	// The numerator of the expression corresponding to the width of the active area of the image in pixels.
	// TODO: default="1">
	FactorNumerator uint `xml:"factorNumerator,attr,omitempty" json:"@factorNumerator,omitempty"`

	// The denominator of the expression corresponding to the height of the active area of the image in
	// pixels.
	// TODO: default="1">
	FactorDenominator uint `xml:"factorDenominator,attr,omitempty" json:"@factorDenominator,omitempty"`
}

type MasteredColorVolume struct {
	// Red Chromaticity.
	PrimaryRChromaticity PrimaryRChromaticity `xml:"urn:ebu:metadata-schema:ebucore ebucore:primaryRChromaticity,omitempty" json:"ebucore:primaryRChromaticity,omitempty"`

	// Green Chromaticity.
	PrimaryGChromaticity PrimaryGChromaticity `xml:"urn:ebu:metadata-schema:ebucore ebucore:primaryGChromaticity,omitempty" json:"ebucore:primaryGChromaticity,omitempty"`

	// Blue Chromaticity.
	PrimaryBChromaticity PrimaryBChromaticity `xml:"urn:ebu:metadata-schema:ebucore ebucore:primaryBChromaticity,omitempty" json:"ebucore:primaryBChromaticity,omitempty"`

	// White point Chromaticity.
	WhitePointChromaticity WhitePointChromaticity `xml:"urn:ebu:metadata-schema:ebucore ebucore:whitePointChromaticity,omitempty" json:"ebucore:whitePointChromaticity,omitempty"`

	// Minimum luminance in cd/m2 on a monitor on a black point within the active area.
	// TODO: <pattern value="(([0])([.])([4][0-9]))"/>
	LuminanceMin Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:luminanceMin,omitempty" json:"ebucore:luminanceMin,omitempty"`

	// Maximum luminance in cd/m2 on a monitor on a white point within the active area.
	// TODO: <pattern value="(([1])([4][0]))"/>
	LuminanceMax Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:luminanceMax,omitempty" json:"ebucore:luminanceMax,omitempty"`
}

type PrimaryRChromaticity struct {
	// X coordinates
	// TODO: <pattern value="(([0])([.])([4][0-9]))"/>
	ChromaticityCIEx Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:ChromaticityCIEx,omitempty" json:"ebucore:ChromaticityCIEx,omitempty"`

	// Y coordinates
	// TODO: <pattern value="(([0])([.])([4][0-9]))"/>
	ChromaticityCIEy Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:ChromaticityCIEy,omitempty" json:"ebucore:ChromaticityCIEy,omitempty"`
}

type PrimaryGChromaticity struct {
	// X coordinates
	// TODO: <pattern value="(([0])([.])([4][0-9]))"/>
	ChromaticityCIEx Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:ChromaticityCIEx,omitempty" json:"ebucore:ChromaticityCIEx,omitempty"`

	// Y coordinates
	// TODO: <pattern value="(([0])([.])([4][0-9]))"/>
	ChromaticityCIEy Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:ChromaticityCIEy,omitempty" json:"ebucore:ChromaticityCIEy,omitempty"`
}

type PrimaryBChromaticity struct {
	// X coordinates
	// TODO: <pattern value="(([0])([.])([4][0-9]))"/>
	ChromaticityCIEx Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:ChromaticityCIEx,omitempty" json:"ebucore:ChromaticityCIEx,omitempty"`

	// Y coordinates
	// TODO: <pattern value="(([0])([.])([4][0-9]))"/>
	ChromaticityCIEy Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:ChromaticityCIEy,omitempty" json:"ebucore:ChromaticityCIEy,omitempty"`
}

type WhitePointChromaticity struct {
	// X coordinates
	// TODO: <pattern value="(([0])([.])([4][0-9]))"/>
	ChromaticityCIEx Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:ChromaticityCIEx,omitempty" json:"ebucore:ChromaticityCIEx,omitempty"`

	// Y coordinates
	// TODO: <pattern value="(([0])([.])([4][0-9]))"/>
	ChromaticityCIEy Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:ChromaticityCIEy,omitempty" json:"ebucore:ChromaticityCIEy,omitempty"`
}

type LightLevel struct {
	// To define the maximum content light level.
	MaxCLL Int `xml:"urn:ebu:metadata-schema:ebucore ebucore:maxCLL,omitempty" json:"ebucore:maxCLL,omitempty"`

	// To define the maximum frame-average light level.
	MaxFall Int `xml:"urn:ebu:metadata-schema:ebucore ebucore:maxFall,omitempty" json:"ebucore:maxFall,omitempty"`
}

type DateCreated struct {
	DateAttributes
}

type DateModified struct {
	DateAttributes
}

// Used to provide information on the signing format and purpose To provide information to the language used for signing
// and its purpose
type SigningFormat struct {
	// To provide information on the type of signing.
	TypeAttributes

	// To specify the format of signing being used e.g. living person or avatar
	FormatAttributes

	SigningFormatID        base.URI `xml:"signingFormatId,attr,omitempty" json:"@signingFormatId,omitempty"`
	SigningFormatVersionID string   `xml:"signingFormatVersionId,attr,omitempty" json:"@signingFormatVersionId,omitempty"`
	SigningFormatName      string   `xml:"signingFormatName,attr,omitempty" json:"@signingFormatName,omitempty"`

	// The track number
	TrackID string `xml:"trackId,attr,omitempty" json:"@trackId,omitempty"`

	// The track name
	TrackName string `xml:"trackName,attr,omitempty" json:"@trackName,omitempty"`

	// A pointer to the file with the signing in available as a separte resource.
	SigningSourceURI base.URI `xml:"signingSourceUri,attr,omitempty" json:"@signingSourceUri,omitempty"`

	// To providing information on the signing language.
	Language string `xml:"language,attr,omitempty" json:"@language,omitempty"`

	// A flag to signal if signing is present.
	SigningPresenceFlag bool `xml:"signingPresenceFlag,attr,omitempty" json:"@signingPresenceFlag,omitempty"`
}

// To provide information on the wrapper format in complement to the stream encoding information provided in 'channel',
// e.g. mp3, wave, Quicktime, ogg.
type ContainerFormat struct {
	// A point of extension for customisation using a set of technical attributes of predefined datatypes.
	TechnicalAttributes

	// To define the wrapping format of the container.
	ContainerEncoding *ContainerEncoding `xml:"urn:ebu:metadata-schema:ebucore ebucore:containerEncoding,omitempty" json:"ebucore:containerEncoding,omitempty"`

	// To identify the product (hardware / software) used to encode content in the specified encoding format
	Codec *Codec `xml:"urn:ebu:metadata-schema:ebucore ebucore:codec,omitempty" json:"ebucore:codec,omitempty"`

	// To provide additional contextual information.
	Comment []Comment `xml:"urn:ebu:metadata-schema:ebucore ebucore:comment,omitempty" json:"ebucore:comment,omitempty"`

	// To uniquely identify a container.
	ContainerFormatID base.URI `xml:"containerFormatId,attr,omitempty" json:"@containerFormatId,omitempty"`

	// To identify a version of the container format.
	ContainerFormatVersionID base.URI `xml:"containerFormatVersionId,attr,omitempty" json:"@containerFormatVersionId,omitempty"`

	// To attribute a name to the container format.
	ContainerFormatName string `xml:"containerFormatName,attr,omitempty" json:"@containerFormatName,omitempty"`

	// To define a profile of the container format.
	ContainerFormatProfile string `xml:"containerFormatProfile,attr,omitempty" json:"@containerFormatProfile,omitempty"`

	// To define a level of the profile of the container format.
	ContainerFormatProfileLevel string `xml:"containerFormatProfileLevel,attr,omitempty" json:"@containerFormatProfileLevel,omitempty"`

	// To additional information on the container format.
	ContainerFormatDefinition string `xml:"containerFormatDefinition,attr,omitempty" json:"@containerFormatDefinition,omitempty"`
}

type ContainerEncoding struct {
	// A set of attribute to express the audio encoding format.
	FormatAttributes
}

// To provide information on timecode tracks.
type TimecodeFormat struct {
	// A point of extension for customisation using a set of technical attributes of predefined datatypes.
	TechnicalAttributes

	// When a timecode starts for a particular timecode track.
	TimecodeStart []Time `xml:"urn:ebu:metadata-schema:ebucore ebucore:timecodeStart,omitempty" json:"ebucore:timecodeStart,omitempty"`

	// To describe the main features of video tracks such as in mutliview systems
	TimecodeTrack []TimecodeTrack `xml:"urn:ebu:metadata-schema:ebucore ebucore:timecodeTrack,omitempty" json:"ebucore:timecodeTrack,omitempty"`

	// To provide additional contextual information.
	Comment []Comment `xml:"urn:ebu:metadata-schema:ebucore ebucore:comment,omitempty" json:"ebucore:comment,omitempty"`

	// To uniquely identify a timecode format.
	TimecodeFormatID base.URI `xml:"timecodeFormatId,attr,omitempty" json:"@timecodeFormatId,omitempty"`

	// To identify a version of the timecode format.
	TimecodeFormatVersionID base.URI `xml:"timecodeFormatVersionId,attr,omitempty" json:"@timecodeFormatVersionId,omitempty"`

	// To attribute a name to the timecode format.
	TimecodeFormatName base.URI `xml:"timecodeFormatName,attr,omitempty" json:"@timecodeFormatName,omitempty"`

	// To additional information on the timecode format.
	TimecodeFormatDefinition base.URI `xml:"timecodeFormatDefinition,attr,omitempty" json:"@timecodeFormatDefinition,omitempty"`
}

type TimecodeTrack struct {
	// The type of video track e.g. particular view angle.
	TypeAttributes

	// The track ID or track number
	TrackID string `xml:"trackId,attr,omitempty" json:"@trackId,omitempty"`

	// The track name
	TrackName string `xml:"trackName,attr,omitempty" json:"@trackName,omitempty"`
}

// To provide information on metadata formats and tracks.
type MetadataFormat struct {
	// A point of extension for customisation using a set of technical attributes of predefined datatypes.
	TechnicalAttributes

	// To describe the main features of video tracks such as in mutliview systems
	MetadataTrack []MetadataTrack `xml:"urn:ebu:metadata-schema:ebucore ebucore:metadataTrack,omitempty" json:"ebucore:metadataTrack,omitempty"`

	Start *Time `xml:"urn:ebu:metadata-schema:ebucore ebucore:start,omitempty" json:"ebucore:start,omitempty"`

	Duration *Duration `xml:"urn:ebu:metadata-schema:ebucore ebucore:duration,omitempty" json:"ebucore:duration,omitempty"`

	// To provide additional contextual information.
	Comment []Comment `xml:"urn:ebu:metadata-schema:ebucore ebucore:comment,omitempty" json:"ebucore:comment,omitempty"`

	// To uniquely identify a timecode format.
	MetadataFormatID base.URI `xml:"metadataFormatId,attr,omitempty" json:"@metadataFormatId,omitempty"`

	// To identify a version of the timecode format.
	MetadataFormatVersionID base.URI `xml:"metadataFormatVersionId,attr,omitempty" json:"@metadataFormatVersionId,omitempty"`

	// To attribute a name to the timecode format.
	MetadataFormatName base.URI `xml:"metadataFormatName,attr,omitempty" json:"@metadataFormatName,omitempty"`

	// To additional information on the timecode format.
	MetadataFormatDefinition base.URI `xml:"metadataFormatDefinition,attr,omitempty" json:"@metadataFormatDefinition,omitempty"`
}

type MetadataTrack struct {
	// The type of video track e.g. particular view angle.
	TypeAttributes

	// The track ID or track number
	TrackID string `xml:"trackId,attr,omitempty" json:"@trackId,omitempty"`

	// The track name
	TrackName string `xml:"trackName,attr,omitempty" json:"@trackName,omitempty"`
}

type DateAttributes struct {
	// To express the date as a year
	Year int `xml:"year,attr,omitempty" json:"@year,omitempty"`

	// To express the date as a date
	Date base.Date `xml:"date,attr,omitempty" json:"@date,omitempty"`

	// To express the time as a time
	Time base.Time `xml:"time,attr,omitempty" json:"@time,omitempty"`

	// To indicate the starting year of a period of time, a year of allocation / attribution, etc.
	StartYear int `xml:"startYear,attr,omitempty" json:"@startYear,omitempty"`

	// To indicate the starting date of a period of time, a date of allocation / attribution, etc.
	StartDate base.Date `xml:"startDate,attr,omitempty" json:"@startDate,omitempty"`

	// To indicate the starting time of a period of time, a time of allocation / attribution, etc.
	StartTime base.Time `xml:"startTime,attr,omitempty" json:"@startTime,omitempty"`

	// To indicate the closing year of a period of time, a year of deprecation, etc.
	EndYear int `xml:"endYear,attr,omitempty" json:"@endYear,omitempty"`

	// To indicate the closing date of a period of time, a date of deprecation, etc.
	EndDate base.Date `xml:"endDate,attr,omitempty" json:"@endDate,omitempty"`

	// To indicate the closing time of a period of time, a time of deprecation, etc.
	EndTime base.Time `xml:"endTime,attr,omitempty" json:"@endTime,omitempty"`

	// To describe a period using free text. This allows defining a period in complement to the concept of period provided
	// in coverage / temporal (associated with the editorial content or rights)
	Period string `xml:"period,attr,omitempty" json:"@period,omitempty"`
}

// To identify a person, group of persons or organisation
type Entity struct {
	// Minimum information providing means to further identify and contact a person.
	ContactDetails []ContactDetails `xml:"urn:ebu:metadata-schema:ebucore ebucore:contactDetails,omitempty" json:"ebucore:contactDetails,omitempty"`

	// Minimum information providing means to further identify and contact the entity as an organisation.
	OrganisationDetails []OrganisationDetails `xml:"urn:ebu:metadata-schema:ebucore ebucore:organisationDetails,omitempty" json:"ebucore:organisationDetails,omitempty"`

	// Used to identify the function fulfilled by the person, group or organisation described as an entity. This is used
	// to detail the role of a 'contributor'. This also applies to e.g. 'creator' as several functions can be seen as
	// participating to the creative process
	Role []Role `xml:"urn:ebu:metadata-schema:ebucore ebucore:role,omitempty" json:"ebucore:role,omitempty"`

	// To list awards and prizes granted to a contact or an organisation.
	Award []Award `xml:"urn:ebu:metadata-schema:ebucore ebucore:award,omitempty" json:"ebucore:award,omitempty"`

	// To list e.g. key personal / career or other key event related to an entity.
	Event []Event `xml:"urn:ebu:metadata-schema:ebucore ebucore:event,omitempty" json:"ebucore:event,omitempty"`

	// to provide the fees of an agent e.g. per type of prestation.
	AgentFee []AgentFee `xml:"urn:ebu:metadata-schema:ebucore ebucore:agentFee,omitempty" json:"ebucore:agentFee,omitempty"`

	// To provide a unique Id for the entity.
	EntityID base.URI `xml:"entityId,attr,omitempty" json:"@entityId,omitempty"`
}

type Role struct {
	TypeAttributes

	// An attribute to associated an entity to a cost centre in the framework of the defined role/job
	CostCentre string `xml:"costCentre,attr,omitempty" json:"@costCentre,omitempty"`
}

type AgentFee struct {
	TypeAttributes

	Value    Float    `xml:"urn:ebu:metadata-schema:ebucore ebucore:value,omitempty" json:"ebucore:value,omitempty"`
	Currency Currency `xml:"urn:ebu:metadata-schema:ebucore ebucore:currency,omitempty" json:"ebucore:currency,omitempty"`
}

type Animal struct {
	// To identify a type of animal.
	TypeAttributes

	// To name an animal.
	AnimalName []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:animalName,omitempty" json:"ebucore:animalName,omitempty"`

	// To describe an animal.
	AnimalDescription []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:animalDescription,omitempty" json:"ebucore:animalDescription,omitempty"`

	// To associate a code with an animal.
	AnimalCode []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:animalCode,omitempty" json:"ebucore:animalCode,omitempty"`

	// To give the gender of an animal.
	AnimalGender *Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:animalGender,omitempty" json:"ebucore:animalGender,omitempty"`

	// To give the birth year of an animal.
	AnimalBirthYear *Date `xml:"urn:ebu:metadata-schema:ebucore ebucore:animalBirthYear,omitempty" json:"ebucore:animalBirthYear,omitempty"`

	// To give the passport (pet identity, markings, vaccination records...) of an animal.
	AnimalPassport *Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:animalPassport,omitempty" json:"ebucore:animalPassport,omitempty"`

	// To associate a colour code with an animal.
	AnimalColourCode []AnimalColourCode `xml:"urn:ebu:metadata-schema:ebucore ebucore:animalColourCode,omitempty" json:"ebucore:animalColourCode,omitempty"`

	// To associate a breed code with an animal.
	AnimalBreedCode []AnimalBreedCode `xml:"urn:ebu:metadata-schema:ebucore ebucore:animalBreedCode,omitempty" json:"ebucore:animalBreedCode,omitempty"`

	// To identify the owner of an animal.
	AnimalOwner []Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:animalOwner,omitempty" json:"ebucore:animalOwner,omitempty"`

	// To identify the groom / care keeper of an animal.
	AnimalGroom []Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:animalgroom,omitempty" json:"ebucore:animalgroom,omitempty"`

	// To associate a role with an animal.
	AnimalRoleCode []AnimalRoleCode `xml:"urn:ebu:metadata-schema:ebucore ebucore:animalRoleCode,omitempty" json:"ebucore:animalRoleCode,omitempty"`

	// To give a fictitious character name to an animal.
	AnimalCharacterName *Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:animalCharacterName,omitempty" json:"ebucore:animalCharacterName,omitempty"`

	// To associate an id with an animal.
	AnimalID base.URI `xml:"animalId,attr,omitempty" json:"@animalId,omitempty"`
}

type AnimalColourCode struct {
	TypeAttributes
}

type AnimalBreedCode struct {
	TypeAttributes
}

type AnimalRoleCode struct {
	TypeAttributes
}

// To provide all the information necessary to contact and locate a person.
type ContactDetails struct {
	// To define a type fo contact.
	TypeAttributes

	// The contact name is optional as a contact maybe identified by its contactId. The name is a combination of different
	// elements of information for a specific purpose such as a name for display.
	// TODO: either Name or GivenName,FamilyName,OtherGivenName,Suffix,Salutation or none must be specified
	Name []CompoundName `xml:"urn:ebu:metadata-schema:ebucore ebucore:name,omitempty" json:"ebucore:name,omitempty"`

	// Also known as the first name among others.
	// TODO: either Name or GivenName,FamilyName,OtherGivenName,Suffix,Salutation or none must be specified
	GivenName *Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:givenName,omitempty" json:"ebucore:givenName,omitempty"`

	// also known as the last name among others.
	// TODO: either Name or GivenName,FamilyName,OtherGivenName,Suffix,Salutation or none must be specified
	FamilyName *Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:familyName,omitempty" json:"ebucore:familyName,omitempty"`

	// e.g. a middle name.
	// TODO: either Name or GivenName,FamilyName,OtherGivenName,Suffix,Salutation or none must be specified
	OtherGivenName []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:otherGivenName,omitempty" json:"ebucore:otherGivenName,omitempty"`

	// E.g. 'Jr' for junior or 'Sr' for senior.
	// TODO: either Name or GivenName,FamilyName,OtherGivenName,Suffix,Salutation or none must be specified
	Suffix *Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:suffix,omitempty" json:"ebucore:suffix,omitempty"`

	// E.g. M., Ms., Mrs, Pr., Dr.
	// TODO: either Name or GivenName,FamilyName,OtherGivenName,Suffix,Salutation or none must be specified
	Salutation *Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:salutation,omitempty" json:"ebucore:salutation,omitempty"`

	// The date of birth.
	BirthDate []Date `xml:"urn:ebu:metadata-schema:ebucore ebucore:birthDate,omitempty" json:"ebucore:birthDate,omitempty"`

	// The date of death. If absent, means that the contact is still alive.
	DeathDate []Date `xml:"urn:ebu:metadata-schema:ebucore ebucore:deathDate,omitempty" json:"ebucore:deathDate,omitempty"`

	// The place of birth.
	BirthPlace []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:birthPlace,omitempty" json:"ebucore:birthPlace,omitempty"`

	// The place of death.
	DeathPlace []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:deathPlace,omitempty" json:"ebucore:deathPlace,omitempty"`

	// The current nationality.
	Nationality []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:nationality,omitempty" json:"ebucore:nationality,omitempty"`

	// The name used by the contact in a particular context.
	Username []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:username,omitempty" json:"ebucore:username,omitempty"`

	// A name given to the contact.
	Nickname []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:nickname,omitempty" json:"ebucore:nickname,omitempty"`

	// The job function of the contact
	Occupation []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:occupation,omitempty" json:"ebucore:occupation,omitempty"`

	// The detailed contact information of a person/contact.
	Details []Details `xml:"urn:ebu:metadata-schema:ebucore ebucore:details,omitempty" json:"ebucore:details,omitempty"`

	// For example, in the case a contact is a performing actor/actress, the stage name will be the fictitious character's
	// name
	StageName []StageName `xml:"urn:ebu:metadata-schema:ebucore ebucore:stageName,omitempty" json:"ebucore:stageName,omitempty"`

	// An attribute to provide the name of a character/fictional entity typically associated with an Actor/Actress role
	CharacterName []CharacterName `xml:"urn:ebu:metadata-schema:ebucore ebucore:characterName,omitempty" json:"ebucore:characterName,omitempty"`

	// A flag to indicate if the contact contribution/appearance was as a guest
	Guest []Boolean `xml:"urn:ebu:metadata-schema:ebucore ebucore:guest,omitempty" json:"ebucore:guest,omitempty"`

	// An element to facilitate filtering / search on gender, e.g. male or female.
	Gender []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:gender,omitempty" json:"ebucore:gender,omitempty"`

	// To provide links with related information such as Wikipedia or other web pages (e.g. press releases).
	RelatedInformationLink []RelatedInformationLink `xml:"urn:ebu:metadata-schema:ebucore ebucore:relatedInformationLink,omitempty" json:"ebucore:relatedInformationLink,omitempty"`

	// This is used to identify contacts (e.g. an assistant, an agent) related to the contact/person being described.
	RelatedContacts []Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:relatedContacts,omitempty" json:"ebucore:relatedContacts,omitempty"`

	// To provide additional information on a contact/person's particular abilities.
	Skill []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:skill,omitempty" json:"ebucore:skill,omitempty"`

	// To provide information on the affiliation of a contact/person.
	Affiliation []Affiliation `xml:"urn:ebu:metadata-schema:ebucore ebucore:affiliation,omitempty" json:"ebucore:affiliation,omitempty"`

	// To provide additional information about an agent.
	AdditionalInformation []AdditionalInformation `xml:"urn:ebu:metadata-schema:ebucore ebucore:additionalInformation,omitempty" json:"ebucore:additionalInformation,omitempty"`

	// To provide a unique identifier for a contact.
	ContactID base.URI `xml:"contactId,attr,omitempty" json:"@contactId,omitempty"`

	// To provide a date when this information was last known to be valid.
	LastUpdate base.Date `xml:"lastUpdate,attr,omitempty" json:"@lastUpdate,omitempty"`
}

type StageName struct {
	Element
	TypeAttributes
}

type CharacterName struct {
	Element
	TypeAttributes
}

// To provide information on the details needed to contact and locate an organisation.
type OrganisationDetails struct {
	// A group of attributes to specific a type or organisation.
	TypeAttributes

	// The organisation name is optional as a contact maybe identified by its contactId. The name is aa cobination of
	// different elements of information for a specific purpose, e.g. type of name such as the display name or else.
	OrganisationName []CompoundName `xml:"urn:ebu:metadata-schema:ebucore ebucore:organisationName,omitempty" json:"ebucore:organisationName,omitempty"`

	// To provide a code associated with and organisation.
	OrganisationCode []Identifier `xml:"urn:ebu:metadata-schema:ebucore ebucore:organisationCode,omitempty" json:"ebucore:organisationCode,omitempty"`

	// To provide a code associated with and organisation.
	OrganisationDescription []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:organisationDescription,omitempty" json:"ebucore:organisationDescription,omitempty"`

	// To provide the nationality of an organisation.
	OrganisationNationality *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:organisationNationality,omitempty" json:"ebucore:organisationNationality,omitempty"`

	// To identify one or more production area / department / service where the resource was created/originated, in free
	// text
	OrganisationDepartment *OrganisationDepartment `xml:"urn:ebu:metadata-schema:ebucore ebucore:organisationDepartment,omitempty" json:"ebucore:organisationDepartment,omitempty"`

	// To provide detailed contact information for an organisation.
	Details []Details `xml:"urn:ebu:metadata-schema:ebucore ebucore:details,omitempty" json:"ebucore:details,omitempty"`

	// To provide a link to a resource containing related additional information about the organisation, e.g. a Wikipedia
	// page.
	RelatedInformationLink []RelatedInformationLink `xml:"urn:ebu:metadata-schema:ebucore ebucore:relatedInformationLink,omitempty" json:"ebucore:relatedInformationLink,omitempty"`

	// Useful to provide contact information within the organisation particularly is no other person information is
	// otherwise provided.
	Contacts []Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:contacts,omitempty" json:"ebucore:contacts,omitempty"`

	// To provide a unique identifier for an organisation.
	OrganisationID base.URI `xml:"organisationId,attr,omitempty" json:"@organisationId,omitempty"`

	// A logo to visually identifiy an organisation
	LinkToLogo base.URI `xml:"linkToLogo,attr,omitempty" json:"@linkToLogo,omitempty"`

	// To provide a date when this information was known to be valid.
	LastUpdate base.Date `xml:"lastUpdate,attr,omitempty" json:"@lastUpdate,omitempty"`
}

type RelatedInformationLink struct {
	URIValue

	// A group of attributes to define the type of link provided.
	TypeAttributes
}

type OrganisationDepartment struct {
	Element

	// To provide a unique identifier for a department within an organisation.
	DepartmentID base.URI `xml:"departmentId,attr,omitempty" json:"@departmentId,omitempty"`
}

// Detailed contact information for a person or organisation.
type Details struct {
	// A group of attributes to specify the type fo contact details e.g. 'private' or 'professional'.
	TypeAttributes

	// The e-mail address through which the contact can be directly joined.
	EmailAddress []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:emailAddress,omitempty" json:"ebucore:emailAddress,omitempty"`

	// The web address where additional information can be found regarding an organisation or a contact.
	WebAddress *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:webAddress,omitempty" json:"ebucore:webAddress,omitempty"`

	// The address of an organisation or contact.
	Address *Address `xml:"urn:ebu:metadata-schema:ebucore ebucore:address,omitempty" json:"ebucore:address,omitempty"`

	// A telephone number.
	TelephoneNumber *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:telephoneNumber,omitempty" json:"ebucore:telephoneNumber,omitempty"`

	// A mobile telephone number.
	MobileTelephoneNumber *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:mobileTelephoneNumber,omitempty" json:"ebucore:mobileTelephoneNumber,omitempty"`
}

// Provides address details for an organisation or contact.
type Address struct {
	// A line containing details of an address such as the street name and number.
	AddressLine []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:addressLine,omitempty" json:"ebucore:addressLine,omitempty"`

	// The name of the city within which the adress is located.
	AddressTownCity *Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:addressTownCity,omitempty" json:"ebucore:addressTownCity,omitempty"`

	// The name of the county and/or state within which the address is located.
	AddressCountyState *Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:addressCountyState,omitempty" json:"ebucore:addressCountyState,omitempty"`

	// The postal/delivery code corresponding to the address.
	AddressDeliveryCode []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:addressDeliveryCode,omitempty" json:"ebucore:addressDeliveryCode,omitempty"`

	// The country within which the address is located.
	Country *Country `xml:"urn:ebu:metadata-schema:ebucore ebucore:country,omitempty" json:"ebucore:country,omitempty"`
}

type Country struct {
	// A group of attributes to specify a country by it name or code, optionally from a classification scheme.
	TypeAttributes
}

// To allow the definition of a compound name combining different elements of information for a usage defined by the
// typeGroup attributes and in a format defined by the formatGroup attributes. For use for contact and organisation
// names.
type CompoundName struct {
	Element

	// A group of attributes to specify a type of compound name.
	TypeAttributes

	// A group of attributes to define a format/structure of a compound name.
	FormatAttributes
}

// A complexType to describe an affiliation.
type Affiliation struct {
	Organisation OrganisationDetails `xml:"urn:ebu:metadata-schema:ebucore ebucore:organisation,omitempty" json:"ebucore:organisation,omitempty"`
	Period       Period              `xml:"urn:ebu:metadata-schema:ebucore ebucore:period,omitempty" json:"ebucore:period,omitempty"`
}

type Period struct {
	DateAttributes
}

// A copmplexType to define a region as a country and / or one or more regions (e.g. administrative regions or states)
// within a country.
type Region struct {
	// A country defined as a typeLabel, i.e. a free text, or a typeLink pointing to a term within a controlled
	// vocabulary, classification scheme or taxonomy
	Country *Country `xml:"urn:ebu:metadata-schema:ebucore ebucore:country,omitempty" json:"ebucore:country,omitempty"`

	// A region, state or other geographical subdivision defined as a typeLabel, i.e. free text, or a typeLink pointing to
	// a term within a controlled vocabulary, classification scheme or taxonomy.
	CountryRegion []CountryRegion `xml:"urn:ebu:metadata-schema:ebucore ebucore:countryRegion,omitempty" json:"ebucore:countryRegion,omitempty"`
}

type CountryRegion struct {
	// A group of attributes to identify a region.
	TypeAttributes
}

// To describe awards, decorations and prizes.
type Award struct {
	// The name of the award.
	// TODO: required
	Name []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:name,omitempty" json:"ebucore:name,omitempty"`

	// The description of the award.
	Description []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:description,omitempty" json:"ebucore:description,omitempty"`

	// The description of the award.
	Category []Category `xml:"urn:ebu:metadata-schema:ebucore ebucore:category,omitempty" json:"ebucore:category,omitempty"`

	// The name of the ceremony when the award was received.
	Ceremony []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:ceremony,omitempty" json:"ebucore:ceremony,omitempty"`

	// The name of the organisation and/or personality who gave the award.
	Official []Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:official,omitempty" json:"ebucore:official,omitempty"`

	// The date when the award was received.
	Date []Date `xml:"urn:ebu:metadata-schema:ebucore ebucore:date,omitempty" json:"ebucore:date,omitempty"`

	// An identifier associated with the award.
	AwardID base.URI `xml:"awardId,attr,omitempty" json:"@awardId,omitempty"`
}

type Category struct {
	TypeAttributes
}

type Date struct {
	DateAttributes
}

// A complex type to express a number of edit units. An editUnit is the inverse of the edit rate, or corrected edit rate
// as the result of editUnit=1/(editrate*(factorNumerator/factorDenominator)).
type EditUnitNumber struct {
	Value int64 `xml:",chardata" json:"#value"`

	// The base number of frames or samples per seconds. This base number can be corrected by a factor calculated as
	// the result of 'factorNumerator/factorDenominator'.
	EditRate uint `xml:"editRate,attr,omitempty" json:"@editRate,omitempty"`

	// The numerator of the correction factor.
	// TODO: default="1">
	FactorNumerator uint `xml:"factorNumerator,attr,omitempty" json:"@factorNumerator,omitempty"`

	// The denominator of the correction factor.
	// TODO: default="1">
	FactorDenominator uint `xml:"factorDenominator,attr,omitempty" json:"@factorDenominator,omitempty"`
}

// To provide rating using classification scheme or custom rating schemes.
type Rating struct {
	// The typeGroup is used to further describe the purpose of the rating.
	TypeAttributes

	// The formatGroup should be used to provide information on the rating format, if applicable.
	FormatAttributes

	// The value of rating as free text optionnally ina specified language
	RatingValue []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:ratingValue,omitempty" json:"ebucore:ratingValue,omitempty"`

	// A link to a web resource describing a rating value or to a term in a classification scheme.
	RatingLink []URIValue `xml:"urn:ebu:metadata-schema:ebucore ebucore:ratingLink,omitempty" json:"ebucore:ratingLink,omitempty"`

	// The maximum rating value in the system in use (see typeGroup), as free text optionnally ina specified language
	RatingScaleMaxValue []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:ratingScaleMaxValue,omitempty" json:"ebucore:ratingScaleMaxValue,omitempty"`

	// The minimum rating value in the system in use (see typeGroup), as free text optionnally ina specified language
	RatingScaleMinValue []Element `xml:"urn:ebu:metadata-schema:ebucore ebucore:ratingScaleMinValue,omitempty" json:"ebucore:ratingScaleMinValue,omitempty"`

	// To identify the person or organisation who rated the resource.
	RatingProvider *Entity `xml:"urn:ebu:metadata-schema:ebucore ebucore:ratingProvider,omitempty" json:"ebucore:ratingProvider,omitempty"`

	// The region within which the rating applies.
	RatingRegion []Region `xml:"urn:ebu:metadata-schema:ebucore ebucore:ratingRegion,omitempty" json:"ebucore:ratingRegion,omitempty"`

	// The region within which the rating doesn not apply.
	RatingExclusionRegion []Region `xml:"urn:ebu:metadata-schema:ebucore ebucore:ratingExclusionRegion,omitempty" json:"ebucore:ratingExclusionRegion,omitempty"`

	// To provide the name of the system used as a reference for rating.
	RatingSystem string `xml:"ratingSystem,attr,omitempty" json:"@ratingSystem,omitempty"`

	// To provide the name of the environment in which rating applies.
	RatingEnvironment string `xml:"ratingEnvironment,attr,omitempty" json:"@ratingEnvironment,omitempty"`

	// An attribute to optionnally explain why this rating has been selected
	Reason string `xml:"reason,attr,omitempty" json:"@reason,omitempty"`

	// A link to a visual representation of the rating, if available
	LinkToLogo base.URI `xml:"linkToLogo,attr,omitempty" json:"@linkToLogo,omitempty"`

	// A flag to signal that content has not been rated (if set to "true")
	NotRated bool `xml:"notRated,attr,omitempty" json:"@notRated,omitempty"`

	// A flag to signal if content is adult content (if set to "true") or not
	AdultContent bool `xml:"adultContent,attr,omitempty" json:"@adultContent,omitempty"`
}

type TypeAttributes struct {
	// Free text to define the type of the associated element.
	TypeLabel string `xml:"typeLabel,attr,omitempty" json:"@typeLabel,omitempty"`

	// Free text to provide a definition for the type.
	TypeDefinition string `xml:"typeDefinition,attr,omitempty" json:"@typeDefinition,omitempty"`

	// A URI to link e.g. to a type in a classification scheme.
	TypeLink base.URI `xml:"typeLink,attr,omitempty" json:"@typeLink,omitempty"`

	// To identify a source of attribution.
	TypeSource string `xml:"typeSource,attr,omitempty" json:"@typeSource,omitempty"`

	TypeNamespace string `xml:"typeNamespace,attr,omitempty" json:"@typeNamespace,omitempty"`

	// To define the language in which the type information is provided.
	TypeLanguage string `xml:"typeLanguage,attr,omitempty" json:"@typeLanguage,omitempty"`

	// To define the controlled vocabulary where this term is coming from.
	TypeThesaurus string `xml:"typeThesaurus,attr,omitempty" json:"@typeThesaurus,omitempty"`
}

type FormatAttributes struct {
	// Free text to define the format of the associated element.
	FormatLabel string `xml:"formatLabel,attr,omitempty" json:"@formatLabel,omitempty"`

	// Free text to provide a definition for the format.
	FormatDefinition string `xml:"formatDefinition,attr,omitempty" json:"@formatDefinition,omitempty"`

	// A URI to link e.g. to a format in a classification scheme.
	FormatLink base.URI `xml:"formatLink,attr,omitempty" json:"@formatLink,omitempty"`

	// To identify a source of attribution.
	FormatSource string `xml:"formatSource,attr,omitempty" json:"@formatSource,omitempty"`

	FormatNamespace string `xml:"formatNamespace,attr,omitempty" json:"@formatNamespace,omitempty"`

	// To define the language in which the type information is provided.
	FormatLanguage string `xml:"formatLanguage,attr,omitempty" json:"@formatLanguage,omitempty"`

	// To define the controlled vocabulary where this term is coming from.
	FormatThesaurus string `xml:"formatThesaurus,attr,omitempty" json:"@formatThesaurus,omitempty"`
}

type StatusAttributes struct {
	// Free text to define the status of the associated element.
	StatusLabel string `xml:"statusLabel,attr,omitempty" json:"@statusLabel,omitempty"`

	// Free text to provide a definition for the status.
	StatusDefinition string `xml:"statusDefinition,attr,omitempty" json:"@statusDefinition,omitempty"`

	// A URI to link e.g. to a status in a classification scheme.
	StatusLink base.URI `xml:"statusLink,attr,omitempty" json:"@statusLink,omitempty"`

	// To identify a source of attribution.
	StatusSource string `xml:"statusSource,attr,omitempty" json:"@statusSource,omitempty"`

	StatusNamespace string `xml:"statusNamespace,attr,omitempty" json:"@statusNamespace,omitempty"`

	// To define the language in which the type information is provided.
	StatusLanguage string `xml:"statusLanguage,attr,omitempty" json:"@statusLanguage,omitempty"`

	// To define the controlled vocabulary where this term is coming from.
	StatusThesaurus string `xml:"statusThesaurus,attr,omitempty" json:"@statusThesaurus,omitempty"`
}

// To provide a length associated with a unit.
type Length struct {
	Value uint `xml:",chardata" json:"#value"`

	// An attribute to specify the unit in which the width is expressed.
	Unit string `xml:"unit,attr,omitempty" json:"@unit,omitempty"`
}

// To provide a dimension associated with a unit.
type Dimension struct {
	Value uint `xml:",chardata" json:"#value"`

	// An attribute to specify the unit in which the width is expressed.
	Unit string `xml:"unit,attr,omitempty" json:"@unit,omitempty"`
}

// A set of metadata for the description of an image format.
type ImageFormat struct {
	// A point of extension for customisation using technical attributes of predefined datatypes.
	TechnicalAttributes

	// The identification of a region in a document, an image or a video is done by defining the coordinates of the bottom
	// left corner of the region. The region is defined from this point of reference using the width and height
	// properties. regionDelimX uses the same unit as width.
	RegionDelimX *Dimension `xml:"urn:ebu:metadata-schema:ebucore ebucore:regionDelimX,omitempty" json:"ebucore:regionDelimX,omitempty"`

	// The identification of a region in a document, an image or a video is done by defining the coordinates of the bottom
	// left corner of the region. The region is defined from this point of reference using the width and height
	// properties. regionDelimY uses the same unit as height.
	RegionDelimY *Dimension `xml:"urn:ebu:metadata-schema:ebucore ebucore:regionDelimY,omitempty" json:"ebucore:regionDelimY,omitempty"`

	// To express the width of an image
	Width *Dimension `xml:"urn:ebu:metadata-schema:ebucore ebucore:width,omitempty" json:"ebucore:width,omitempty"`

	// To express the height of an image
	Height *Dimension `xml:"urn:ebu:metadata-schema:ebucore ebucore:height,omitempty" json:"ebucore:height,omitempty"`

	// To express the orientation of the image
	Orientation Orientation `xml:"urn:ebu:metadata-schema:ebucore ebucore:orientation,omitempty" json:"ebucore:orientation,omitempty"`

	// A string to define e.g. the ratio of the picture (the width by the height), for instance '4:3' or '16:9'
	// (rational).
	AspectRatio *AspectRatio `xml:"urn:ebu:metadata-schema:ebucore ebucore:aspectRatio,omitempty" json:"ebucore:aspectRatio,omitempty"`

	// Used to express the encoding parameters of the resource e.g. jpg.
	ImageEncoding *ImageEncoding `xml:"urn:ebu:metadata-schema:ebucore ebucore:imageEncoding,omitempty" json:"ebucore:imageEncoding,omitempty"`

	ImageCodec *Codec `xml:"urn:ebu:metadata-schema:ebucore ebucore:imageCodec,omitempty" json:"ebucore:imageCodec,omitempty"`

	// To provide additional contextual information.
	Comment []Comment `xml:"urn:ebu:metadata-schema:ebucore ebucore:comment,omitempty" json:"ebucore:comment,omitempty"`

	// To uniquely identify an image format.
	ImageFormatID base.URI `xml:"imageFormatId,attr,omitempty" json:"@imageFormatId,omitempty"`

	// To provide a version number associated with an image format.
	ImageFormatVersionID string `xml:"imageFormatVersionId,attr,omitempty" json:"@imageFormatVersionId,omitempty"`

	// To provide a human readable name associated with an image format.
	ImageFormatName string `xml:"imageFormatName,attr,omitempty" json:"@imageFormatName,omitempty"`

	// To provide a human readable definition of the image format.
	ImageFormatDefinition string `xml:"imageFormatDefinition,attr,omitempty" json:"@imageFormatDefinition,omitempty"`

	// To provide additional information on a particular format profile
	ImageFormatProfile string `xml:"imageFormatProfile,attr,omitempty" json:"@imageFormatProfile,omitempty"`

	// To provide additional information on the level of a particular format profile
	ImageFormatProfileLevel string `xml:"imageFormatProfileLevel,attr,omitempty" json:"@imageFormatProfileLevel,omitempty"`

	// A flag to presence of an image.
	ImagePresenceFlag bool `xml:"imagePresenceFlag,attr,omitempty" json:"@imagePresenceFlag,omitempty"`
}

type Orientation *String

var (
	LandscapeOrientation Orientation = &String{Value: "landscape"}
	PortraitOrientation  Orientation = &String{Value: "portrait"}
)

type ImageEncoding struct {
	// A group of attributes to specify the encoding type
	TypeAttributes
}

// A set of metadata for the description of a video format.
type VideoFormat struct {
	// A point of extension for customisation using technical attributes of predefined datatype.
	TechnicalAttributes

	// The identification of a region in a document, an image or a video is done by defining the coordinates of the bottom
	// left corner of the region. The region is defined from this point of reference using the width and height
	// properties. regionDelimX uses the same unit as width.
	RegionDelimX *Dimension `xml:"urn:ebu:metadata-schema:ebucore ebucore:regionDelimX,omitempty" json:"ebucore:regionDelimX,omitempty"`

	// The identification of a region in a document, an image or a video is done by defining the coordinates of the bottom
	// left corner of the region. The region is defined from this point of reference using the width and height
	// properties. regionDelimY uses the same unit as height.
	RegionDelimY *Dimension `xml:"urn:ebu:metadata-schema:ebucore ebucore:regionDelimY,omitempty" json:"ebucore:regionDelimY,omitempty"`

	// To define different widths of a video image e.g. display, active or else
	Width []Width `xml:"urn:ebu:metadata-schema:ebucore ebucore:width,omitempty" json:"ebucore:width,omitempty"`

	// To define different heights of a video image e.g. display, active or else
	Height []Height `xml:"urn:ebu:metadata-schema:ebucore ebucore:height,omitempty" json:"ebucore:height,omitempty"`

	// The number of lines or resolution height e.g. 1080, 720, etc.
	Lines *UInt `xml:"urn:ebu:metadata-schema:ebucore ebucore:lines,omitempty" json:"ebucore:lines,omitempty"`

	// the frequencey at which frames are sampled in frame per second
	FrameRate *Rational `xml:"urn:ebu:metadata-schema:ebucore ebucore:frameRate,omitempty" json:"ebucore:frameRate,omitempty"`

	// A string to define e.g. the ratio of the picture (the width by the height), for instance '4:3' or '16:9'
	// (rational).
	AspectRatio []AspectRatio `xml:"urn:ebu:metadata-schema:ebucore ebucore:aspectRatio,omitempty" json:"ebucore:aspectRatio,omitempty"`

	// Used to express the encoding parameters of the resource e.g. H264 for a video channel. videoEncoding can also be
	// used to provide additional information on a particular set of parameters associated with a format or profile of
	// this format.
	VideoEncoding *VideoEncoding `xml:"urn:ebu:metadata-schema:ebucore ebucore:videoEncoding,omitempty" json:"ebucore:videoEncoding,omitempty"`

	// To identify the product (hardware / software) used to encode content in the specified encoding format
	Codec *Codec `xml:"urn:ebu:metadata-schema:ebucore ebucore:codec,omitempty" json:"ebucore:codec,omitempty"`

	// the video bit rate in bits per second (default unit)
	BitRate *Dimension `xml:"urn:ebu:metadata-schema:ebucore ebucore:bitRate,omitempty" json:"ebucore:bitRate,omitempty"`

	// the video maximum bit rate in bits per second (default unit)
	BitRateMax *Dimension `xml:"urn:ebu:metadata-schema:ebucore ebucore:bitRateMax,omitempty" json:"ebucore:bitRateMax,omitempty"`

	// To indicate if if the bit rate is constant or variable
	BitRateMode BitRateMode `xml:"urn:ebu:metadata-schema:ebucore ebucore:bitRateMode,omitempty" json:"ebucore:bitRateMode,omitempty"`

	// To provide information on the picture scanning format e.g. interlaced or progressive.
	ScanningFormat ScanningFormat `xml:"urn:ebu:metadata-schema:ebucore ebucore:scanningFormat,omitempty" json:"ebucore:scanningFormat,omitempty"`

	// The order in which the image is scanned.
	ScanningOrder ScanningOrder `xml:"urn:ebu:metadata-schema:ebucore ebucore:scanningOrder,omitempty" json:"ebucore:scanningOrder,omitempty"`

	// a flag to indicate if a noise filter has been used, and ist nature/type.
	NoiseFilter *NoiseFilter `xml:"urn:ebu:metadata-schema:ebucore ebucore:noiseFilter,omitempty" json:"ebucore:noiseFilter,omitempty"`

	// To describe the main features of video tracks such as in mutliview systems
	VideoTrack []VideoTrack `xml:"urn:ebu:metadata-schema:ebucore ebucore:videoTrack,omitempty" json:"ebucore:videoTrack,omitempty"`

	// A flag to indicate if the visual content is for stereoscopic rendering
	Flag3D *Boolean `xml:"urn:ebu:metadata-schema:ebucore ebucore:flag_3D,omitempty" json:"ebucore:flag_3D,omitempty"`

	// A flag to indicate if the visual content is 360°.
	Flag360 *Boolean `xml:"urn:ebu:metadata-schema:ebucore ebucore:flag_360,omitempty" json:"ebucore:flag_360,omitempty"`

	// A flag to indicate if the visual content is multi-view.
	FlagMultiview *Boolean `xml:"urn:ebu:metadata-schema:ebucore ebucore:flag_multiview,omitempty" json:"ebucore:flag_multiview,omitempty"`

	// To identify a filter e.g. a noise filter.
	Filter []Filter `xml:"urn:ebu:metadata-schema:ebucore ebucore:filter,omitempty" json:"ebucore:filter,omitempty"`

	// HDR (High Dynamic Range) static metadata.
	MasteredColorVolume *DigitalAssetColorVolume `xml:"urn:ebu:metadata-schema:ebucore ebucore:masteredColorVolume,omitempty" json:"ebucore:masteredColorVolume,omitempty"`

	// To indicate the light level. Aligned wit MovieLabs' common metadata set.
	LightLevel *DigitalAssetVideoPictureLightLevel `xml:"urn:ebu:metadata-schema:ebucore ebucore:lightLevel,omitempty" json:"ebucore:lightLevel,omitempty"`

	// The distance between two full images (I-frames), also referred to as the GOP size.
	// TODO: default="1"
	IFrameInterval *Int `xml:"urn:ebu:metadata-schema:ebucore ebucore:iFrameInterval,omitempty" json:"ebucore:iFrameInterval,omitempty"`

	// To provide additional contextual information on the video format.
	Comment []Comment `xml:"urn:ebu:metadata-schema:ebucore ebucore:comment,omitempty" json:"ebucore:comment,omitempty"`

	// To uniquely identify a video format.
	VideoFormatID base.URI `xml:"videoFormatId,attr,omitempty" json:"@videoFormatId,omitempty"`

	// To provide information on the version of a video format.
	VideoFormatVersionID string `xml:"videoFormatVersionId,attr,omitempty" json:"@videoFormatVersionId,omitempty"`

	// To provide a human readable name for the video format.
	VideoFormatName string `xml:"videoFormatName,attr,omitempty" json:"@videoFormatName,omitempty"`

	// To provide a human readable definition for the video format.
	VideoFormatDefinition string `xml:"videoFormatDefinition,attr,omitempty" json:"@videoFormatDefinition,omitempty"`

	// To provide additional information on a particular format profile
	VideoFormatProfile string `xml:"videoFormatProfile,attr,omitempty" json:"@videoFormatProfile,omitempty"`

	// To provide additional information on a particular format level
	VideoFormatProfileLevel string `xml:"videoFormatProfileLevel,attr,omitempty" json:"@videoFormatProfileLevel,omitempty"`

	// A flag to signal the presence of video.
	VideoPresenceFlag bool `xml:"videoPresenceFlag,attr,omitempty" json:"@videoPresenceFlag,omitempty"`
}

type Width struct {
	// A group of attributes to specify the type of dimension used.
	TypeAttributes

	Dimension
}

type Height struct {
	// A group of attributes to specify the type of dimension used.
	TypeAttributes

	Dimension
}

type VideoEncoding struct {
	// A group of attributes to specify the type of encoding.
	TypeAttributes

	// To provide information on the encoding profile.
	VideoEncodingProfile string `xml:"videoEncodingProfile,attr,omitempty" json:"@videoEncodingProfile,omitempty"`

	// To provide information on the encoding level.
	VideoEncodingLevel string `xml:"videoEncodingLevel,attr,omitempty" json:"@videoEncodingLevel,omitempty"`
}

type BitRateMode *String

var (
	NoneBitRateMode     BitRateMode = &String{Value: "none"}
	ConstantBitRateMode BitRateMode = &String{Value: "constant"}
	VariableBitRateMode BitRateMode = &String{Value: "variable"}
)

type ScanningFormat *String

var (
	NoneScanningFormat        ScanningFormat = &String{Value: "none"}
	InterlacedScanningFormat  ScanningFormat = &String{Value: "interlaced"}
	ProgressiveScanningFormat ScanningFormat = &String{Value: "progressive"}
	MixedScanningFormat       ScanningFormat = &String{Value: "mixed"}
)

type ScanningOrder *String

var (
	NoneScanningOrder     ScanningOrder = &String{Value: "none"}
	TopScanningOrder      ScanningOrder = &String{Value: "top"}
	BottomScanningOrder   ScanningOrder = &String{Value: "bottom"}
	PulldownScanningOrder ScanningOrder = &String{Value: "pulldown"}
)

type NoiseFilter struct {
	TypeAttributes

	Value bool `xml:",chardata" json:"#value"`

	// TODO: use="required"/>
	VendorID string `xml:"vendorId,attr,omitempty" json:"@vendorId,omitempty"`
}

type VideoTrack struct {
	// The type of video track e.g. particular view angle.
	TypeAttributes

	// The track ID or track number
	TrackID string `xml:"trackId,attr,omitempty" json:"@trackId,omitempty"`

	// The track name
	TrackName string `xml:"trackName,attr,omitempty" json:"@trackName,omitempty"`
}

// HDR (High Dynamic Range) static metadata. Naming similar to MovieLabs' common metadata set and SMPTE 2086.
type DigitalAssetColorVolume struct {
	PrimaryRChromaticity   DigitalAssetChromaticity `xml:"urn:ebu:metadata-schema:ebucore ebucore:primaryRChromaticity,omitempty" json:"ebucore:primaryRChromaticity,omitempty"`
	PrimaryGChromaticity   DigitalAssetChromaticity `xml:"urn:ebu:metadata-schema:ebucore ebucore:primaryGChromaticity,omitempty" json:"ebucore:primaryGChromaticity,omitempty"`
	PrimaryBChromaticity   DigitalAssetChromaticity `xml:"urn:ebu:metadata-schema:ebucore ebucore:primaryBChromaticity,omitempty" json:"ebucore:primaryBChromaticity,omitempty"`
	WhitePointChromaticity DigitalAssetChromaticity `xml:"urn:ebu:metadata-schema:ebucore ebucore:whitePointChromaticity,omitempty" json:"ebucore:whitePointChromaticity,omitempty"`

	// Minimum luminance is expressed in cd/m2.
	// TODO: restriction: 0.0001 -- 5.0000
	LuminanceMin Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:luminanceMin,omitempty" json:"ebucore:luminanceMin,omitempty"`

	// Maximum luminance is expressed in cd/m2.
	// TODO: restriction: 5.00 -- 10000.00
	LuminanceMax Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:luminanceMax,omitempty" json:"ebucore:luminanceMax,omitempty"`
}

type DigitalAssetChromaticity struct {
	// TODO: restriction: 0.0001 -- 0.7400
	ChromaticityCIEx Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:ChromaticityCIEx,omitempty" json:"ebucore:ChromaticityCIEx,omitempty"`

	// TODO: restriction: 0.0001 -- 0.7400
	ChromaticityCIEy Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:ChromaticityCIEy,omitempty" json:"ebucore:ChromaticityCIEy,omitempty"`
}

type DigitalAssetVideoPictureLightLevel struct {
	ContentMax      []ContentMax      `xml:"urn:ebu:metadata-schema:ebucore ebucore:ContentMax,omitempty" json:"ebucore:ContentMax,omitempty"`
	FrameAverageMax []FrameAverageMax `xml:"urn:ebu:metadata-schema:ebucore ebucore:FrameAverageMax,omitempty" json:"ebucore:FrameAverageMax,omitempty"`
}

type ContentMax struct {
	Value uint `xml:",chardata" json:"#value"`

	Interpretation string `xml:"interpretation,attr,omitempty" json:"@interpretation,omitempty"`
}

type FrameAverageMax struct {
	Value uint `xml:",chardata" json:"#value"`

	Interpretation string `xml:"interpretation,attr,omitempty" json:"@interpretation,omitempty"`
}

// To provide a set of audio technical characteristics.
type AudioFormat struct {
	// A point of extension for customisation using technical attributes of predefined datatypes.
	TechnicalAttributes

	// To define the audio compression format of the resource e.g. AAC for an audio channel.
	AudioEncoding *AudioEncoding `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioEncoding,omitempty" json:"ebucore:audioEncoding,omitempty"`

	// To identify the hardware / software used
	Codec *Codec `xml:"urn:ebu:metadata-schema:ebucore ebucore:codec,omitempty" json:"ebucore:codec,omitempty"`

	// To define the audio track configuration. Used to express the arrangement or audio tracks e.g. 'stereo', '2+1',
	// 'surround', 'surround (7+1)'
	AudioTrackConfiguration *AudioTrackConfiguration `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioTrackConfiguration,omitempty" json:"ebucore:audioTrackConfiguration,omitempty"`

	// the frequency at which audio is being sampled in samples per
	// second
	SamplingRate *Int64 `xml:"urn:ebu:metadata-schema:ebucore ebucore:samplingRate,omitempty" json:"ebucore:samplingRate,omitempty"`

	// The size of each audio sample
	SampleSize *UInt `xml:"urn:ebu:metadata-schema:ebucore ebucore:sampleSize,omitempty" json:"ebucore:sampleSize,omitempty"`

	// The type of audio sample
	SampleType SampleType `xml:"urn:ebu:metadata-schema:ebucore ebucore:sampleType,omitempty" json:"ebucore:sampleType,omitempty"`

	// the audio bit rate in bits per second
	BitRate *Dimension `xml:"urn:ebu:metadata-schema:ebucore ebucore:bitRate,omitempty" json:"ebucore:bitRate,omitempty"`

	// the audio maximum bit rate in bits per second
	BitRateMax *Dimension `xml:"urn:ebu:metadata-schema:ebucore ebucore:bitRateMax,omitempty" json:"ebucore:bitRateMax,omitempty"`

	// To indicate if if the bit rate is constant or variable
	BitRateMode BitRateMode `xml:"urn:ebu:metadata-schema:ebucore ebucore:bitRateMode,omitempty" json:"ebucore:bitRateMode,omitempty"`

	// To describe the track allocation e.g. in conformance with EBU R123.
	AudioTrack []AudioTrack `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioTrack,omitempty" json:"ebucore:audioTrack,omitempty"`

	// The total number of audio channels.
	Channels *UInt `xml:"urn:ebu:metadata-schema:ebucore ebucore:channels,omitempty" json:"ebucore:channels,omitempty"`

	// To identify a filter.
	Filter []Filter `xml:"urn:ebu:metadata-schema:ebucore ebucore:filter,omitempty" json:"ebucore:filter,omitempty"`

	// To provide additional contextual information on the audio format.
	Comment []Comment `xml:"urn:ebu:metadata-schema:ebucore ebucore:comment,omitempty" json:"ebucore:comment,omitempty"`

	// To uniquely identify an audio format.
	AudioFormatID base.URI `xml:"audioFormatId,attr,omitempty" json:"@audioFormatId,omitempty"`

	// To provide information on the version of an audio format.
	AudioFormatVersionID string `xml:"audioFormatVersionId,attr,omitempty" json:"@audioFormatVersionId,omitempty"`

	// To provide a human readable name for an audio format.
	AudioFormatName string `xml:"audioFormatName,attr,omitempty" json:"@audioFormatName,omitempty"`

	// To provide a human readable definition for an audio format.
	AudioFormatDefinition string `xml:"audioFormatDefinition,attr,omitempty" json:"@audioFormatDefinition,omitempty"`

	// To define a profile of the audio format.
	AudioFormatProfile string `xml:"audioFormatProfile,attr,omitempty" json:"@audioFormatProfile,omitempty"`

	// To define a level of the profile of the audio format.
	AudioFormatProfileLevel string `xml:"audioFormatProfileLevel,attr,omitempty" json:"@audioFormatProfileLevel,omitempty"`

	// A flag to signal the presence of audio.
	AudioPresenceFlag bool `xml:"audioPresenceFlag,attr,omitempty" json:"@audioPresenceFlag,omitempty"`

	// A flag to signal the presence of an audio description service for accessibility.
	AudioDescriptionPresenceFlag bool `xml:"audioDescriptionPresenceFlag,attr,omitempty" json:"@audioDescriptionPresenceFlag,omitempty"`
}

type AudioEncoding struct {
	// A set of attribute to express the audio encoding format.
	TypeAttributes

	// To provide information on the encoding profile.
	AudioEncodingProfile string `xml:"audioEncodingProfile,attr,omitempty" json:"@audioEncodingProfile,omitempty"`

	// To provide information on the encoding level.
	AudioEncodingLevel string `xml:"audioEncodingLevel,attr,omitempty" json:"@audioEncodingLevel,omitempty"`
}

type AudioTrackConfiguration struct {
	TypeAttributes
}

type SampleType *String

var (
	FloatSampleType   SampleType = &String{Value: "float"}
	IntegerSampleType SampleType = &String{Value: "integer"}
)

type AudioTrack struct {
	// The track type.
	TypeAttributes

	// The track language.
	TrackLanguage string `xml:"trackLanguage,attr,omitempty" json:"@trackLanguage,omitempty"`

	// The track number.
	TrackID string `xml:"trackId,attr,omitempty" json:"@trackId,omitempty"`

	// The track name.
	TrackName string `xml:"trackName,attr,omitempty" json:"@trackName,omitempty"`
}

// A group of attributes to provide a set of techncial attributes on data provided with the resource.
type DataFormat struct {
	// A point of extension for customisation using technical attributes of predefined datatypes.
	TechnicalAttributes

	// Used to provide information on the captioning format and purpose
	CaptioningFormat []CaptioningFormat `xml:"urn:ebu:metadata-schema:ebucore ebucore:captioningFormat,omitempty" json:"ebucore:captioningFormat,omitempty"`

	// Used to provide information on the subtitling format and purpose, e.g. translation in a different language
	SubtitlingFormat []SubtitlingFormat `xml:"urn:ebu:metadata-schema:ebucore ebucore:subtitlingFormat,omitempty" json:"ebucore:subtitlingFormat,omitempty"`

	// Used to provide information on ancillary data format and purpose. This type provides information on the Ancillary
	// Data packet type. See SMPTE 291M, SMPTE 436M
	AncillaryDataFormat []AncillaryDataFormat `xml:"urn:ebu:metadata-schema:ebucore ebucore:ancillaryDataFormat,omitempty" json:"ebucore:ancillaryDataFormat,omitempty"`

	// To identify the product (hardware / software) used to encode content in the specified encoding format. The
	// dataFormat shall be repeated in order to associate a codec e.g. with captioning and separately for subtitling.
	Codec *Codec `xml:"urn:ebu:metadata-schema:ebucore ebucore:codec,omitempty" json:"ebucore:codec,omitempty"`

	// To provide additional contextual information.
	Comment []Comment `xml:"urn:ebu:metadata-schema:ebucore ebucore:comment,omitempty" json:"ebucore:comment,omitempty"`

	// To uniquely identify a data format.
	DataFormatID base.URI `xml:"dataFormatId,attr,omitempty" json:"@dataFormatId,omitempty"`

	// To provide version information on a data format.
	DataFormatVersionID string `xml:"dataFormatVersionId,attr,omitempty" json:"@dataFormatVersionId,omitempty"`

	// To provide a human readable name for a data format.
	DataFormatName string `xml:"dataFormatName,attr,omitempty" json:"@dataFormatName,omitempty"`

	// To provide a human readable definition for a data format.
	DataFormatDefinition string `xml:"dataFormatDefinition,attr,omitempty" json:"@dataFormatDefinition,omitempty"`

	// The data track number.
	DataTrackID string `xml:"dataTrackId,attr,omitempty" json:"@dataTrackId,omitempty"`

	// The data track name.
	DataTrackName string `xml:"dataTrackName,attr,omitempty" json:"@dataTrackName,omitempty"`

	// A group of attributes to specify the data language.
	DataTrackLanguage string `xml:"dataTrackLanguage,attr,omitempty" json:"@dataTrackLanguage,omitempty"`

	// To provide additional information on a particular format profile
	DataFormatProfile string `xml:"dataFormatProfile,attr,omitempty" json:"@dataFormatProfile,omitempty"`

	// To provide additional information on a particular format level
	DataFormatProfileLevel string `xml:"dataFormatProfileLevel,attr,omitempty" json:"@dataFormatProfileLevel,omitempty"`

	// A flag to signal the presence of data.
	DataPresenceFlag bool `xml:"dataPresenceFlag,attr,omitempty" json:"@dataPresenceFlag,omitempty"`
}

// Used to provide information on ancillary data format and purpose. This type provides information on the Ancillary
// Data packet type. See SMPTE 291M, SMPTE 436M
type AncillaryDataFormat struct {
	// ANC DID Value.
	DID *Int `xml:"urn:ebu:metadata-schema:ebucore ebucore:DID,omitempty" json:"ebucore:DID,omitempty"`

	// ANC SDID Value.
	SDID *Int `xml:"urn:ebu:metadata-schema:ebucore ebucore:SDID,omitempty" json:"ebucore:SDID,omitempty"`

	// Video line number containing the ANC packets of this type.
	LineNumber []Int `xml:"urn:ebu:metadata-schema:ebucore ebucore:lineNumber,omitempty" json:"ebucore:lineNumber,omitempty"`

	// Indicates HANC or VANC, and what field in which packets should be stored. See SMPTE 436M for legal values.
	WrappingType *Int `xml:"urn:ebu:metadata-schema:ebucore ebucore:wrappingType,omitempty" json:"ebucore:wrappingType,omitempty"`

	// To uniquely identify and ancillary data format.
	AncillaryDataFormatID base.URI `xml:"ancillaryDataFormatId,attr,omitempty" json:"@ancillaryDataFormatId,omitempty"`

	// A human readable name for an ancillary data format.
	AncillaryDataFormatName string `xml:"ancillaryDataFormatName,attr,omitempty" json:"@ancillaryDataFormatName,omitempty"`

	// To provide additional information on a particular format profile
	AncillaryDataFormatProfile string `xml:"ancillaryDataFormatProfile,attr,omitempty" json:"@ancillaryDataFormatProfile,omitempty"`
}

// Used to provide information on the subtitling format and purpose, e.g. translation in a different language
type SubtitlingFormat struct {
	// To specify the purpose for subtitling
	TypeAttributes

	// To specify the subtitling format
	FormatAttributes

	// To uniquely identify a subtitling format.
	SubtitlingFormatID base.URI `xml:"subtitlingFormatId,attr,omitempty" json:"@subtitlingFormatId,omitempty"`

	// A human readable name for the subtitling format.
	SubtitlingFormatName string `xml:"subtitlingFormatName,attr,omitempty" json:"@subtitlingFormatName,omitempty"`

	// To provide additional information on a particular format profile
	SubtitlingFormatProfile string `xml:"subtitlingFormatProfile,attr,omitempty" json:"@subtitlingFormatProfile,omitempty"`

	// The track number
	TrackID string `xml:"trackId,attr,omitempty" json:"@trackId,omitempty"`

	// The track name
	TrackName string `xml:"trackName,attr,omitempty" json:"@trackName,omitempty"`

	// A pointer to the file containing the subtitling.
	SubtitlingSourceURI base.URI `xml:"subtitlingSourceUri,attr,omitempty" json:"@subtitlingSourceUri,omitempty"`

	// The language used for subtitling.
	Language string `xml:"language,attr,omitempty" json:"@language,omitempty"`

	// Closed subtitling if true
	Closed bool `xml:"closed,attr,omitempty" json:"@closed,omitempty"`

	// A flag to signal the presence of subtitling.
	SubtitlingPresenceFlag bool `xml:"subtitlingPresenceFlag,attr,omitempty" json:"@subtitlingPresenceFlag,omitempty"`
}

// Used to provide information on the captioning format and purpose
type CaptioningFormat struct {
	// To specify the purpose for captioning
	TypeAttributes

	// To specify the captioning format
	FormatAttributes

	// To uniquely identify a captioning format.
	CaptioningFormatID base.URI `xml:"captioningFormatId,attr,omitempty" json:"@captioningFormatId,omitempty"`

	// To provide a human readable name for the captioning format.
	CaptioningFormatName string `xml:"captioningFormatName,attr,omitempty" json:"@captioningFormatName,omitempty"`

	// To provide additional information on a particular format profile
	CaptioningFormatProfile string `xml:"captioningFormatProfile,attr,omitempty" json:"@captioningFormatProfile,omitempty"`

	// The track number
	TrackID string `xml:"trackId,attr,omitempty" json:"@trackId,omitempty"`

	// The track name
	TrackName string `xml:"trackName,attr,omitempty" json:"@trackName,omitempty"`

	// A pointer to the file containing captioning
	CaptioningSourceURI base.URI `xml:"captioningSourceUri,attr,omitempty" json:"@captioningSourceUri,omitempty"`

	// A group of attributes to specify the captioning language.
	Language string `xml:"language,attr,omitempty" json:"@language,omitempty"`

	// Closed captioning if true
	Closed bool `xml:"closed,attr,omitempty" json:"@closed,omitempty"`

	// A flag to signal the presence of captioning.
	CaptioningPresenceFlag bool `xml:"captioningPresenceFlag,attr,omitempty" json:"@captioningPresenceFlag,omitempty"`
}

// To list the technical characteristics of a document.
type DocumentFormat struct {
	// A group of attributes to define a type of document format.
	TypeAttributes

	// A group of attributes to define the format of a document.
	FormatAttributes

	// A point of extension for customisation using technical attributes of predefined datatypes.
	TechnicalAttributes

	// To provide a word count for the document
	WordCount *Int `xml:"urn:ebu:metadata-schema:ebucore ebucore:wordCount,omitempty" json:"ebucore:wordCount,omitempty"`

	// The identification of a region in a document, an image or a video is done by defining the coordinates of the bottom
	// left corner of the region. The region is defined from this point of reference using the width and height
	// properties. regionDelimX uses the same unit as width.
	RegionDelimX *Dimension `xml:"urn:ebu:metadata-schema:ebucore ebucore:regionDelimX,omitempty" json:"ebucore:regionDelimX,omitempty"`

	// The identification of a region in a document, an image or a video is done by defining the coordinates of the bottom
	// left corner of the region. The region is defined from this point of reference using the width and height
	// properties. regionDelimY uses the same unit as height.
	RegionDelimY *Dimension `xml:"urn:ebu:metadata-schema:ebucore ebucore:regionDelimY,omitempty" json:"ebucore:regionDelimY,omitempty"`

	// The width of an image, video or document.
	Width *Dimension `xml:"urn:ebu:metadata-schema:ebucore ebucore:width,omitempty" json:"ebucore:width,omitempty"`

	// The height of an image, video or document.
	Height *Dimension `xml:"urn:ebu:metadata-schema:ebucore ebucore:height,omitempty" json:"ebucore:height,omitempty"`

	// To provide additional contextual information.
	Comment []Comment `xml:"urn:ebu:metadata-schema:ebucore ebucore:comment,omitempty" json:"ebucore:comment,omitempty"`

	// To uniquely identify a document format.
	DocumentFormatID base.URI `xml:"documentFormatId,attr,omitempty" json:"@documentFormatId,omitempty"`

	// To provide information on the version of a document format.
	DocumentFormatVersionID string `xml:"documentFormatVersionId,attr,omitempty" json:"@documentFormatVersionId,omitempty"`

	// To provide a human readable name for the document format.
	DocumentFormatName string `xml:"documentFormatName,attr,omitempty" json:"@documentFormatName,omitempty"`

	// To provide additional information on a particular format profile
	DocumentFormatProfile string `xml:"documentFormatProfile,attr,omitempty" json:"@documentFormatProfile,omitempty"`

	// To provide a human readable definition for the document format.
	DocumentFormatDefinition string `xml:"documentFormatDefinition,attr,omitempty" json:"@documentFormatDefinition,omitempty"`
}

// To provide information on a hardware of software codec used to encode part or all of the resource.
type Codec struct {
	TypeAttributes

	// the commercial name of the codec used
	CodecIdentifier *Identifier `xml:"urn:ebu:metadata-schema:ebucore ebucore:codecIdentifier,omitempty" json:"ebucore:codecIdentifier,omitempty"`

	// the commercial name of the codec used
	Name *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:name,omitempty" json:"ebucore:name,omitempty"`

	// the name of the vendor
	Vendor *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:vendor,omitempty" json:"ebucore:vendor,omitempty"`

	// the version of the product
	Version *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:version,omitempty" json:"ebucore:version,omitempty"`

	// the family of products to which the codec belongs to
	Family *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:family,omitempty" json:"ebucore:family,omitempty"`

	// A Unique Resource Location at which more information can be found about the codec.
	URL *URIValue `xml:"urn:ebu:metadata-schema:ebucore ebucore:url,omitempty" json:"ebucore:url,omitempty"`
}

// Complex type for describing filters of any kind
type Filter struct {
	TypeAttributes

	// Reference to the audio- or videotrack the filter will be applied to
	// TODO: required
	TrackIdRef []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:trackIdRef,omitempty" json:"ebucore:trackIdRef,omitempty"`

	// The filter profile that should be applied
	FilterProfile FilterProfile `xml:"urn:ebu:metadata-schema:ebucore ebucore:filterProfile,omitempty" json:"ebucore:filterProfile,omitempty"`

	// Additional settings and attributes for a filter
	FilterSetting []FilterSetting `xml:"urn:ebu:metadata-schema:ebucore ebucore:filterSetting,omitempty" json:"ebucore:filterSetting,omitempty"`

	// Order of the appliance of filter
	FilterOrder int `xml:"filterOrder,attr,omitempty" json:"@filterOrder,omitempty"`
}

type FilterProfile struct {
	TypeAttributes
}

type FilterSetting struct {
	TypeAttributes

	TechnicalAttributes

	// Order of the filter attribute
	FilterAttributeOrder int `xml:"filterAttributeOrder,attr,omitempty" json:"@filterAttributeOrder,omitempty"`
}

type Time struct {
	// A group of attributes to specify a type of time.
	TypeAttributes

	// To express time as timecode.
	// TODO: either Timecode or NormalPlayTime or OffsetNormalPlayTime or EditUnitNumber or Time must be specified
	Timecode *Timecode `xml:"urn:ebu:metadata-schema:ebucore ebucore:timecode,omitempty" json:"ebucore:timecode,omitempty"`

	// To express the start time in the form HH:MM:SS.S
	// TODO: either Timecode or NormalPlayTime or OffsetNormalPlayTime or EditUnitNumber or Time must be specified
	NormalPlayTime *TimeValue `xml:"urn:ebu:metadata-schema:ebucore ebucore:normalPlayTime,omitempty" json:"ebucore:normalPlayTime,omitempty"`

	// To express a time offset as a duration in the iso duration format PnYnMnDTnHnMnS e.g. PT1M40.1S
	// TODO: either Timecode or NormalPlayTime or OffsetNormalPlayTime or EditUnitNumber or Time must be specified
	OffsetNormalPlayTime *DurationValue `xml:"urn:ebu:metadata-schema:ebucore ebucore:offsetNormalPlayTime,omitempty" json:"ebucore:offsetNormalPlayTime,omitempty"`

	// The express the start time as a number of edit Units
	// TODO: either Timecode or NormalPlayTime or OffsetNormalPlayTime or EditUnitNumber or Time must be specified
	EditUnitNumber *EditUnitNumber `xml:"urn:ebu:metadata-schema:ebucore ebucore:editUnitNumber,omitempty" json:"ebucore:editUnitNumber,omitempty"`

	// To express the start time in a user defined time format e.g. in a number of seconds
	// TODO: either Timecode or NormalPlayTime or OffsetNormalPlayTime or EditUnitNumber or Time must be specified
	Time *CustomTimeValue `xml:"urn:ebu:metadata-schema:ebucore ebucore:time,omitempty" json:"ebucore:time,omitempty"`
}

type CustomTimeValue struct {
	// To specify the format of a user defined type.
	FormatAttributes

	Value string `xml:",chardata" json:"#value"`
}

// To specify a duration.
type Duration struct {
	// To specify a type of duration.
	TypeAttributes

	// To express the duration using timecode compliant with SMPTE ST 2021-1:2009
	// TODO: either Timecode or NormalPlayTime or EditUnitNumber or Duration must be specified
	Timecode *Timecode `xml:"urn:ebu:metadata-schema:ebucore ebucore:timecode,omitempty" json:"ebucore:timecode,omitempty"`

	// To express the duration in compliance with ISO8601: PnYnMnDTnHnMnS e.g. PT1M40.1S for a duration of 1 minute 40.1
	// seconds
	// TODO: either Timecode or NormalPlayTime or EditUnitNumber or Duration must be specified
	NormalPlayTime *DurationValue `xml:"urn:ebu:metadata-schema:ebucore ebucore:normalPlayTime,omitempty" json:"ebucore:normalPlayTime,omitempty"`

	// The express the duration as a number of edit Units
	// TODO: either Timecode or NormalPlayTime or EditUnitNumber or Duration must be specified
	EditUnitNumber *EditUnitNumber `xml:"urn:ebu:metadata-schema:ebucore ebucore:editUnitNumber,omitempty" json:"ebucore:editUnitNumber,omitempty"`

	// To express the duration in a user defined time format e.g. in a number seconds
	// TODO: either Timecode or NormalPlayTime or EditUnitNumber or Duration must be specified
	Duration *CustomDurationValue `xml:"urn:ebu:metadata-schema:ebucore ebucore:duration,omitempty" json:"ebucore:duration,omitempty"`
}

type CustomDurationValue struct {
	// To specify the format od a user defined duration.
	FormatAttributes

	Value string `xml:",chardata" json:"#value"`
}

// To provide hash details.
type Hash struct {
	// The value calculated for the hash code using the hash function.
	// TODO: required
	// TODO: use []byte?
	Value string `xml:"urn:ebu:metadata-schema:ebucore ebucore:hashValue,omitempty" json:"ebucore:hashValue,omitempty"`

	// The algorythm used to calculate the hash code associated with an essence file.
	// TODO: required
	Function HashFunction `xml:"urn:ebu:metadata-schema:ebucore ebucore:hashFunction,omitempty" json:"ebucore:hashFunction,omitempty"`
}

type HashFunction struct {
	// A group of attributes used to identify a hash function.
	TypeAttributes
}

// To provide generic information on an instantiation/file.
type FileInfo struct {
	// To indicate the storage requirements or file size of a digital resource. The file size is expressed in bytes
	// (default unit).
	FileSize *Dimension `xml:"urn:ebu:metadata-schema:ebucore ebucore:fileSize,omitempty" json:"ebucore:fileSize,omitempty"`

	// The name of the file as it appears in the location path or url.
	FileName *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:fileName,omitempty" json:"ebucore:fileName,omitempty"`

	// Define the main MIME type as defined by IANA: e.g. audio, video, text, application, or a container MIME type
	MimeType []MimeType `xml:"urn:ebu:metadata-schema:ebucore ebucore:mimeType,omitempty" json:"ebucore:mimeType,omitempty"`

	// An "address for a resource". For an organisation or producer acting as caretaker for a media resource, Format
	// Location may contain information about a specific shelf location for an asset, including an organisation's name,
	// departmental name, shelf id. and contact information. The Format Location for a data file or web page may include a
	// complete URI with a domain, path, filename or html URL. Examples: "Archives Building A, Row J, Shelf 2",
	// "d://playout/server/content.mpg", "http://www.ebu.ch/CorporateVideo.avi". The storage structure to be found at the
	// locator address may be complex and form of sub-directories e.g. for video, audio and data.
	Locator []Locator `xml:"urn:ebu:metadata-schema:ebucore ebucore:locator,omitempty" json:"ebucore:locator,omitempty"`

	// A hash code to verify the integrity of an essence file. This can also be used as a fixity checksum.
	Hash *Hash `xml:"urn:ebu:metadata-schema:ebucore ebucore:hash,omitempty" json:"ebucore:hash,omitempty"`

	// To provide the overall bitrate (audio + video + data) in bits per second (default unit)
	OverallBitRate *Dimension `xml:"urn:ebu:metadata-schema:ebucore ebucore:overallBitRate,omitempty" json:"ebucore:overallBitRate,omitempty"`

	// To provide the overall edit rate.
	EditRate *Rational `xml:"urn:ebu:metadata-schema:ebucore ebucore:editRate,omitempty" json:"ebucore:editRate,omitempty"`
}

type MimeType struct {
	// A group of attributes to specify a MIME type.
	TypeAttributes
}

type Locator struct {
	URIValue

	// The typeGroup can be used to define e.g. the storage type from a predefined list. or as a term. The definition can
	// be used to provide any additional information as deemed necessary e.g. on the location and/ or managing entity.
	TypeAttributes
}

// A container with all definitions related to the audioContents and associated components contained in the material.
type AudioFormatExtended struct {
	// One set of content associated with the material.
	AudioProgramme []AudioProgramme `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioProgramme,omitempty" json:"ebucore:audioProgramme,omitempty"`

	// One or more audioContents associated with an audioProgramme. audioContent refers to an audioObject which contains
	// the audio and its format description.
	AudioContent []AudioContent `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioContent,omitempty" json:"ebucore:audioContent,omitempty"`

	// A time limited (if required) set of audio tracks with a format defined by audioPackFormat.
	AudioObject []AudioObject `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioObject,omitempty" json:"ebucore:audioObject,omitempty"`

	// A set of audioChannels that belong together.
	AudioPackFormat []AudioPackFormat `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioPackFormat,omitempty" json:"ebucore:audioPackFormat,omitempty"`

	// A single sequence of audio samples.
	AudioChannelFormat []AudioChannelFormat `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioChannelFormat,omitempty" json:"ebucore:audioChannelFormat,omitempty"`

	// A division of an audioChannel in time.
	AudioBlockFormat []AudioBlockFormat `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioBlockFormat,omitempty" json:"ebucore:audioBlockFormat,omitempty"`

	// A combination of one or more tracks required to represent a channel, an object, or a group.
	AudioStreamFormat []AudioStreamFormat `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioStreamFormat,omitempty" json:"ebucore:audioStreamFormat,omitempty"`

	// A single set of samples of data in the storage medium.
	AudioTrackFormat []AudioTrackFormat `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioTrackFormat,omitempty" json:"ebucore:audioTrackFormat,omitempty"`

	// The UID of a track in the file.
	AudioTrackUID []AudioTrackUID `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioTrackUID,omitempty" json:"ebucore:audioTrackUID,omitempty"`

	// An unique ID for the extended audio format
	AudioFormatExtendedID string `xml:"audioFormatExtendedID,attr,omitempty" json:"@audioFormatExtendedID,omitempty"`

	// A human readable name for the extended audio format
	AudioFormatExtendedName string `xml:"audioFormatExtendedName,attr,omitempty" json:"@audioFormatExtendedName,omitempty"`

	// A human readable definition for the extended audio format
	AudioFormatExtendedDefinition string `xml:"audioFormatExtendedDefinition,attr,omitempty" json:"@audioFormatExtendedDefinition,omitempty"`

	// A flag to indicate the presence and use of extended audio format data
	AudioFormatExtendedPresenceFlag string `xml:"audioFormatExtendedPresenceFlag,attr,omitempty" json:"@audioFormatExtendedPresenceFlag,omitempty"`

	// To provide a version of the extended audio format data
	Version string `xml:"version,attr,omitempty" json:"@version,omitempty"`
}

// A set of one or more audioContent that derive from the same material, i.e. an audioMultiplex, and the definition of
// its multiplexed audioContents (e.g. foreground and commentary, background music).
type AudioProgramme struct {
	// To optionally define a type of audioProgramme.
	TypeAttributes

	// To optionally define a format of audioProgramme.
	FormatAttributes

	// A list of reference to audioContents, each defining one component of an audioProgramme (e.g. background music), its
	// association with an audioPack (e.g. a 2.0 audioPack of audioChannels for stereo reproduction), its association with
	// a audioStream, and its set of loudness parameters. Notice that loudness values of a program are dependent of the
	// associated audioPack mixReproductionFormat.
	// TODO: required
	AudioContentIDRef []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioContentIDRef,omitempty" json:"ebucore:audioContentIDRef,omitempty"`

	// A set of loudness parameters proper to the audio content of the whole programme.
	LoudnessMetadata *LoudnessMetadata `xml:"urn:ebu:metadata-schema:ebucore ebucore:loudnessMetadata,omitempty" json:"ebucore:loudnessMetadata,omitempty"`

	// Specification of a reference/production/monitoring screen size for the audioProgramme. If the reference screen-size
	// is not given, a default screen-size is implicitly defined.
	AudioProgrammeReferenceScreen *AudioProgrammeReferenceScreen `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioProgrammeReferenceScreen,omitempty" json:"ebucore:audioProgrammeReferenceScreen,omitempty"`

	// An unique ID for the programme.
	// TODO: required
	// TODO: <pattern value="APR_[0-9a-fA-F]{4}"/>
	AudioProgrammeID string `xml:"audioProgrammeID,attr,omitempty" json:"@audioProgrammeID,omitempty"`

	// An unique human readable name for the programme.
	AudioProgrammeName string `xml:"audioProgrammeName,attr,omitempty" json:"@audioProgrammeName,omitempty"`

	// The dialogue language used in this audioProgramme.
	AudioProgrammeLanguage string `xml:"audioProgrammeLanguage,attr,omitempty" json:"@audioProgrammeLanguage,omitempty"`

	// Start time for the programme.
	Start TimecodeStringFrame `xml:"start,attr,omitempty" json:"@start,omitempty"`

	// To define the reference edit rate for the start and end timecodes.
	EditRate float32 `xml:"editRate,attr,omitempty" json:"@editRate,omitempty"`

	// End time for the programme.
	End TimecodeStringFrame `xml:"end,attr,omitempty" json:"@end,omitempty"`

	// A flag to indicate whether the start and end timecodes are drop frame when set to "true".
	// TODO: this field was commented out in the XML Schema
	//Dropframe bool `xml:"dropframe,attr,omitempty" json:"@dropframe,omitempty"`

	// The numerator of a correcting factor for the atart and end timecodes.
	// TODO: this field was commented out in the XML Schema
	// TODO: default = 1
	//FactorNumerator int `xml:"factorNumerator,attr,omitempty" json:"@factorNumerator,omitempty"`

	// The denominator of a correcting factor for the start and end timecodes.
	// TODO: this field was commented out in the XML Schema
	// TODO: default = 1
	//FactorDenominator int `xml:"factorDenominator,attr,omitempty" json:"@factorDenominator,omitempty"`

	// The maximum level (in dB) any audio object in the programme can be reduced by when ducking is invoked.
	MaxDuckingDepth float32 `xml:"maxDuckingDepth,attr,omitempty" json:"@maxDuckingDepth,omitempty"`
}

type AudioProgrammeReferenceScreen struct {
	// Aspect ratio of display expressed as a float (e.g. use 1.6 for 16:10).
	AspectRatio *Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:aspectRatio,omitempty" json:"ebucore:aspectRatio,omitempty"`

	// Centre position of the screen in the room/space (default x=0,y=1,z=0, or az=0,el=0,d=1.0).
	ScreenCentrePosition *ScreenCentrePosition `xml:"urn:ebu:metadata-schema:ebucore ebucore:screenCentrePosition,omitempty" json:"ebucore:screenCentrePosition,omitempty"`

	// Width of screen based on normalised room dimensions.
	ScreenWidth *ScreenWidth `xml:"urn:ebu:metadata-schema:ebucore ebucore:screenWidth,omitempty" json:"ebucore:screenWidth,omitempty"`
}

type ScreenCentrePosition struct {
	// Azimuth angle for the centre of the screen.
	Azimuth float32 `xml:"azimuth,attr,omitempty" json:"@azimuth,omitempty"`

	// Elevation angle for the centre of the screen.
	Elevation float32 `xml:"elevation,attr,omitempty" json:"@elevation,omitempty"`

	// Distance to the centre of the screen.
	Distance float32 `xml:"distance,attr,omitempty" json:"@distance,omitempty"`

	// X-coordinate for the centre of the screen.
	X float32 `xml:"X,attr,omitempty" json:"@X,omitempty"`

	// Y-coordinate for the centre of the screen.
	Y float32 `xml:"Y,attr,omitempty" json:"@Y,omitempty"`

	// Z-coordinate for the centre of the screen.
	Z float32 `xml:"Z,attr,omitempty" json:"@Z,omitempty"`
}

type ScreenWidth struct {
	Float

	// Width of the screen in polar coordinates.
	Azimuth string `xml:"azimuth,attr,omitempty" json:"@azimuth,omitempty"`

	// Width of the screen in Cartesian coordinates.
	X string `xml:"X,attr,omitempty" json:"@X,omitempty"`
}

// An audioContent element defines one component of a programme (e.g. background music), its association with an
// audioPackFormat (e.g. a 2.0 audioPackFormat of audioChannelFormats for stereo reproduction) and its set of loudness
// parameters.
type AudioContent struct {
	// A set of references to audio objects.
	// TODO: required
	AudioObjectIDRef []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioObjectIDRef,omitempty" json:"ebucore:audioObjectIDRef,omitempty"`

	// A set of loudness parameters connected with the audio content.
	LoudnessMetadata *LoudnessMetadata `xml:"urn:ebu:metadata-schema:ebucore ebucore:loudnessMetadata,omitempty" json:"ebucore:loudnessMetadata,omitempty"`

	// A set of values to categorise the presence of dialogue and its type.
	Dialogue *DialogueType `xml:"urn:ebu:metadata-schema:ebucore ebucore:dialogue,omitempty" json:"ebucore:dialogue,omitempty"`

	// An unique ID for the audioContent.
	// TODO: <pattern value="ACO_[0-9a-fA-F]{4}"/>
	AudioContentID string `xml:"audioContentID,attr,omitempty" json:"@audioContentID,omitempty"`

	// An unique human readable name for the audioContent.
	AudioContentName string `xml:"audioContentName,attr,omitempty" json:"@audioContentName,omitempty"`

	// The dialogue language used in this audioContent.
	AudioContentLanguage string `xml:"audioContentLanguage,attr,omitempty" json:"@audioContentLanguage,omitempty"`
}

type DialogueType struct {
	Value int `xml:",chardata" json:"#value"`

	// Numerical ID for the kind of non-dialogue content:
	//   0 - undefined
	//   1 - music
	//   2 - effect
	NonDialogueContentKind NonDialogueContentKind `xml:"nonDialogueContentKind,attr,omitempty" json:"@nonDialogueContentKind,omitempty"`

	// Numerical ID for the kind of dialogue content:
	//    0 - undefined
	//    1 - (storyline) dialogue
	//    2 - voiceover
	//    3 - spoken subtitle
	//    4 - audio description/visually impaired
	//    5 - commentary
	//    6 - emergency
	DialogueContentKind DialogueContentKind `xml:"dialogueContentKind,attr,omitempty" json:"@dialogueContentKind,omitempty"`

	// Numerical ID for the kind of mixed content:
	//    0 - undefined
	//    1 - complete main
	//    2 - mixed
	//    3 - hearing impaired
	MixedContentKind MixedContentKind `xml:"mixedContentKind,attr,omitempty" json:"@mixedContentKind,omitempty"`
}

type NonDialogueContentKind int

const (
	UndefinedNonDialogueContentKind NonDialogueContentKind = iota
	MusicNonDialogueContentKind
	EffectNonDialogueContentKind
)

type DialogueContentKind int

const (
	UndefinedDialogueContentKind DialogueContentKind = iota
	StorylineDialogueDialogueContentKind
	VoiceoverDialogueContentKind
	SpokenSubtitleDialogueContentKind
	AudioDescription_VisuallyImpairedDialogueContentKind
	CommentaryDialogueContentKind
	EmergencyDialogueContentKind
)

type MixedContentKind int

const (
	UndefinedMixedContentKind MixedContentKind = iota
	CompleteMainMixedContentKind
	MixedMixedContentKind
	HearingImpairedMixedContentKind
)

// Establish the relationship between the content using track UIDs and the format via audio packs. These can be nested,
// so refer to other audioObjects.
type AudioObject struct {
	// A set of references to audioPacks.
	AudioPackFormatIDRef []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioPackFormatIDRef,omitempty" json:"ebucore:audioPackFormatIDRef,omitempty"`

	// A set of references to audioObjects.
	AudioObjectIDRef []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioObjectIDRef,omitempty" json:"ebucore:audioObjectIDRef,omitempty"`

	// Reference to another audioObject that should be complementary to this one (i.e. only one of them can be used at a
	// time).
	AudioComplementaryObjectIDRef []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioComplementaryObjectIDRef,omitempty" json:"ebucore:audioComplementaryObjectIDRef,omitempty"`

	// Reference to an audioTrackUID that is listed in the 'chna' chunk.
	// TODO: <pattern value="AC_([0-9a-fA-F]{4})([0-9a-fA-F]{4})"/>
	AudioTrackUIDRef []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioTrackUIDRef,omitempty" json:"ebucore:audioTrackUIDRef,omitempty"`

	// Set of parameters for limiting the amount of interaction of an object.
	AudioObjectInteraction []AudioObjectInteraction `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioObjectInteraction,omitempty" json:"ebucore:audioObjectInteraction,omitempty"`

	// A unique ID for this audioObject,
	// TODO: <pattern value="AO_[0-9a-fA-F]{4}"/>
	AudioObjectID string `xml:"audioObjectID,attr,omitempty" json:"@audioObjectID,omitempty"`

	// A unique human readable name for this audioObject,
	AudioObjectName string `xml:"audioObjectName,attr,omitempty" json:"@audioObjectName,omitempty"`

	// The start time associated with this audioObject expressed as a timecode.
	Start TimecodeString `xml:"start,attr,omitempty" json:"@start,omitempty"`

	// The duration associated with this audioObject expressed as a timecode.
	Duration TimecodeString `xml:"duration,attr,omitempty" json:"@duration,omitempty"`

	// To define the reference edit rate for the start and duration timecodes.
	EditRate float32 `xml:"editRate,attr,omitempty" json:"@editRate,omitempty"`

	// End time for the programme.
	End TimecodeString `xml:"end,attr,omitempty" json:"@end,omitempty"`

	// A flag to indicate whether the start and duration timecodes are drop frame when set to "true".
	// TODO: this field was commented out in the XML Schema
	//Dropframe bool `xml:"dropframe,attr,omitempty" json:"@dropframe,omitempty"`

	// The numerator of a correcting factor for the atart and duration timecodes.
	// TODO: this field was commented out in the XML Schema
	// TODO: default = 1
	//FactorNumerator int `xml:"factorNumerator,attr,omitempty" json:"@factorNumerator,omitempty"`

	// The denominator of a correcting factor for the start and duration timecodes.
	// TODO: this field was commented out in the XML Schema
	// TODO: default = 1
	//FactorDenominator int `xml:"factorDenominator,attr,omitempty" json:"@factorDenominator,omitempty"`

	// 1 if object is dialogue, 0 if not, 2 if mixed. Default = 0.
	// TODO: default="0">
	Dialogue Dialogue `xml:"dialogue,attr,omitempty" json:"@dialogue,omitempty"`

	// Importance of an object. Allows a render to discard an object below a certain level of importance. 10 is most, 0
	// least. Default = 10.
	// TODO: 0 -- 10
	Importance int `xml:"importance,attr,omitempty" json:"@importance,omitempty"`

	// Allows interaction by the user. 1 for allowing interaction, 0 if not. Default = 0.
	Interact bool `xml:"interact,attr,omitempty" json:"@interact,omitempty"`

	// Set to true to prevent the object being ducked.
	DisableDucking bool `xml:"disableDucking,attr,omitempty" json:"@disableDucking,omitempty"`
}

type Dialogue int

const (
	DialogueDialogue Dialogue = iota
	NotDialogueDialogue
	MixedDialogue
)

type AudioObjectInteraction struct {
	// Set the range of gain values allowed for interaction.
	GainInteractionRange *[2]*GainInteractionRange `xml:"urn:ebu:metadata-schema:ebucore ebucore:gainInteractionRange,omitempty" json:"ebucore:gainInteractionRange,omitempty"`

	// Set the range of position values allowed for interaction.
	PositionInteractionRange *[6]*PositionInteractionRange `xml:"urn:ebu:metadata-schema:ebucore ebucore:positionInteractionRange,omitempty" json:"ebucore:positionInteractionRange,omitempty"`

	// Set to true to allow interaction, false to disallow.
	OnOffInteract bool `xml:"onOffInteract,attr,omitempty" json:"@onOffInteract,omitempty"`

	// Set to true to allow gain interaction, false to disallow.
	GainInteract bool `xml:"gainInteract,attr,omitempty" json:"@gainInteract,omitempty"`

	// Set to true to allow position interaction, false to disallow.
	PositionInteract bool `xml:"positionInteract,attr,omitempty" json:"@positionInteract,omitempty"`
}

type GainInteractionRange struct {
	Float

	// Set attribute to "min" to provide minimum gain factor of possible user gain interaction (gainMin = gain (or 1.0
	// if not defined) * gainInteractionRangeMin) Set attribute to "max" to provide maximum gain factor of possible
	// user gain interaction (gainMax = gain (or 1.0 if not defined) * gainInteractionRangeMin)
	Bound string `xml:"bound,attr,omitempty" json:"@bound,omitempty"`
}

type PositionInteractionRange struct {
	Float

	// Set attribute to either "X", "Y", "Z", "azimuth", "elevation" or "distance" to set which axis requires limiting.
	Coordinate string `xml:"coordinate,attr,omitempty" json:"@coordinate,omitempty"`

	// Set attribute to either "min" or "max" to set whether the axis needs a minimum or maximum range for positioning
	// limiting.
	Bound string `xml:"bound,attr,omitempty" json:"@bound,omitempty"`
}

// The audioPackFormat brings together one or more audioChannelFormats that belong together. For example 'stereo' and
// '5.1' would be audioPackFormats for channel-based formats. It can also contain references to other packs to allow
// nesting. To help define the type of channels described within the pack, a typeDefintion is used to define them.
//
// The typeDefinition/typeLabel must match those in the referred audioChannelFormats:
//
//	DirectSpeakers: For channel-based audio, where each channel feeds a speaker directly. Type 0001.
//
//	Matrix: For channel-based audio where channels are matrixed together, such as Mid-Side, Lt/Rt. Type 0002.
//
//	Objects: For object-based audio where channels represent audio objects (or parts of objects), so include positional
//	         information. Type 0003.
//
//	HOA: For scene-based audio where Ambisonics and HOA are used. Type 0004.
//
//	Binaural: For binaural audio, where playback is over headphones. Type 0005.
type AudioPackFormat struct {
	// A set of references to audioChannels.
	AudioChannelFormatIDRef []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioChannelFormatIDRef,omitempty" json:"ebucore:audioChannelFormatIDRef,omitempty"`

	// A set of references to audio packs.
	AudioPackFormatIDRef []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioPackFormatIDRef,omitempty" json:"ebucore:audioPackFormatIDRef,omitempty"`

	// TODO: either EncodePackFormatIDRef,DecodePackFormatIDRef,InputPackFormatIDRef,OutputPackFormatIDRef (if audioPackFormat.typeDefinition=="Matrix")
	//       or Normalization,NfcRefDist,ScreenRef (if audioPackFormat.typeDefinition=="HOA") must be specified

	// Reference to an encoding matrix audioPackFormat from a decoding matrix.
	EncodePackFormatIDRef []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:encodePackFormatIDRef,omitempty" json:"ebucore:encodePackFormatIDRef,omitempty"`

	// Reference to a decoding matrix audioPackFormat from an encoding matrix.
	DecodePackFormatIDRef []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:decodePackFormatIDRef,omitempty" json:"ebucore:decodePackFormatIDRef,omitempty"`

	// Reference to a channel-based (DirectSpeakers) input audioPackFormat.
	InputPackFormatIDRef []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:inputPackFormatIDRef,omitempty" json:"ebucore:inputPackFormatIDRef,omitempty"`

	// Reference to a channel-based (DirectSpeakers) matrix decoded audioPackFormat.
	OutputPackFormatIDRef []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:outputPackFormatIDRef,omitempty" json:"ebucore:outputPackFormatIDRef,omitempty"`

	// Indicates the normalization scheme of the HOA content (N3D, SN3D, FuMa).
	Normalization *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:normalization,omitempty" json:"ebucore:normalization,omitempty"`

	// Indicates the reference distance of the loudspeaker setup for near-field compensation (NFC). If no nfcRefDist is
	// defined or the value is 0, NFC is not necessary.
	NfcRefDist *Int `xml:"urn:ebu:metadata-schema:ebucore ebucore:nfcRefDist,omitempty" json:"ebucore:nfcRefDist,omitempty"`

	// Indicates whether the content is screen-related (flag is equal to 1) or not (flag is equal to 0)
	ScreenRef *Boolean `xml:"urn:ebu:metadata-schema:ebucore ebucore:screenRef,omitempty" json:"ebucore:screenRef,omitempty"`

	// A unique ID for an audio pack.
	// TODO: <pattern value="AP_([0-9a-fA-F]{4})([0-9a-fA-F]{4})"/>
	AudioPackFormatID string `xml:"audioPackFormatID,attr,omitempty" json:"@audioPackFormatID,omitempty"`

	// A unique human readable name for an audio pack.
	AudioPackFormatName string `xml:"audioPackFormatName,attr,omitempty" json:"@audioPackFormatName,omitempty"`

	// Free text to define the type of the associated element.
	// TODO: <pattern value="\d{4}"/>
	TypeLabel string `xml:"typeLabel,attr,omitempty" json:"@typeLabel,omitempty"`

	// Free text to provide a definition for the type.
	TypeDefinition AudioPackFormatTypeDefinition `xml:"typeDefinition,attr,omitempty" json:"@typeDefinition,omitempty"`

	// A URI to link e.g. to a type in a classification scheme.
	TypeLink base.URI `xml:"typeLink,attr,omitempty" json:"@typeLink,omitempty"`

	// To identify a source of attribution.
	TypeSource string `xml:"typeSource,attr,omitempty" json:"@typeSource,omitempty"`

	TypeNamespace string `xml:"typeNamespace,attr,omitempty" json:"@typeNamespace,omitempty"`

	// To define the language in which the type information is provided.
	TypeLanguage string `xml:"typeLanguage,attr,omitempty" json:"@typeLanguage,omitempty"`

	// To define the controlled vocabulary where this term is coming from.
	TypeThesaurus string `xml:"typeThesaurus,attr,omitempty" json:"@typeThesaurus,omitempty"`

	// The absolute distance in metres.
	AbsoluteDistance float32 `xml:"absoluteDistance,attr,omitempty" json:"@absoluteDistance,omitempty"`

	// Importance of an audio pack. Allows a render to discard an audio pack below a certain level of importance. 10 is
	// most, 0 least.
	Importance int `xml:"importance,attr,omitempty" json:"@importance,omitempty"`
}

type AudioPackFormatTypeDefinition string

const (
	DirectSpeakersAudioPackFormatTypeDefinition AudioPackFormatTypeDefinition = "DirectSpeakers"
	MatrixAudioPackFormatTypeDefinition         AudioPackFormatTypeDefinition = "Matrix"
	ObjectsAudioPackFormatTypeDefinition        AudioPackFormatTypeDefinition = "Objects"
	HOAAudioPackFormatTypeDefinition            AudioPackFormatTypeDefinition = "HOA"
	BinauralAudioPackFormatTypeDefinition       AudioPackFormatTypeDefinition = "Binaural"
)

// An audio channel represents a single sequence of audio samples.
//
// An audioChannelFormat represents a single sequence of audio samples. It is sub-divided in the time domain into
// audioBlockFormats, which is must contain at least one of. The typeDefintion of the audioChannelFormat specifies the
// type of audio it is describing, and also determines which parameters are used within its audioBlockFormat children.
//
// Currently, there are five different typeDefinitions:
//
//	DirectSpeakers: For channel-based audio, where each audio channel feeds a speaker directly. Type 0001.
//
//	Matrix: For channel-based audio where channels are matrixed together, such as Mid-Side, Lt/Rt. Type 0002.
//
//	Objects: For object-based audio where channels represent audio objects (or parts of objects), so include
//	         positional information. Type 0003.
//
//	HOA: For scene-based audio where Ambisonics and HOA are used. Type 0004.
//
//	Binaural: For binaural audio, where playback is over headphones. Type 0005.
type AudioChannelFormat struct {
	// Sets a high or low cut-off frequency for the audio in Hz.
	Frequency *[2]*Frequency `xml:"urn:ebu:metadata-schema:ebucore ebucore:frequency,omitempty" json:"ebucore:frequency,omitempty"`

	// A division of the channel in time. Multiple blocks in the channel provide the dynamic information.
	// TODO: required
	AudioBlockFormat []AudioBlockFormat `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioBlockFormat,omitempty" json:"ebucore:audioBlockFormat,omitempty"`

	// A unique human readable name for the audio channel
	AudioChannelFormatName string `xml:"audioChannelFormatName,attr,omitempty" json:"@audioChannelFormatName,omitempty"`

	// A unique ID for the audio channel
	// TODO: <pattern value="AC_([0-9a-fA-F]{4})([0-9a-fA-F]{4})"/>
	AudioChannelFormatID string `xml:"audioChannelFormatID,attr,omitempty" json:"@audioChannelFormatID,omitempty"`

	// Free text to define the type of the associated element.
	// TODO: <pattern value="\d{4}"/>
	TypeLabel string `xml:"typeLabel,attr,omitempty" json:"@typeLabel,omitempty"`

	// Free text to provide a definition for the type.
	TypeDefinition AudioPackFormatTypeDefinition `xml:"typeDefinition,attr,omitempty" json:"@typeDefinition,omitempty"`

	// A URI to link e.g. to a type in a classification scheme.
	TypeLink base.URI `xml:"typeLink,attr,omitempty" json:"@typeLink,omitempty"`

	// To identify a source of attribution.
	TypeSource string `xml:"typeSource,attr,omitempty" json:"@typeSource,omitempty"`

	TypeNamespace string `xml:"typeNamespace,attr,omitempty" json:"@typeNamespace,omitempty"`

	// To define the language in which the type information is provided.
	TypeLanguage string `xml:"typeLanguage,attr,omitempty" json:"@typeLanguage,omitempty"`

	// To define the controlled vocabulary where this term is coming from.
	TypeThesaurus string `xml:"typeThesaurus,attr,omitempty" json:"@typeThesaurus,omitempty"`
}

// A frequency expressed in "Hz" defining a value associated to a type e.g. lowPass or highPass as the type attribute.
type Frequency struct {
	Value int `xml:",chardata" json:"#value"`

	TypeLabel FrequencyTypeLabel `xml:"typeLabel,attr,omitempty" json:"@typeLabel,omitempty"`
}

type FrequencyTypeLabel string

const (
	LowPassFrequencyTypeLabel  FrequencyTypeLabel = "lowPass"
	HighPassFrequencyTypeLabel FrequencyTypeLabel = "highPass"
)

// An audioBlockFormat represents a single sequence of audio samples with fixed parameters, including position, within a
// specified time. The sub-elements with audioBlockFormat are dependent upon the typeDefintion or typeLabel of the
// parent audioChannelFormat element.
//
// Currently, there are five different typeDefinitions:
//
//	DirectSpeakers: For channel-based audio, where each channel feeds a speaker directly. Type 0001.
//
//	Matrix: For channel-based audio where channels are matrixed together, such as Mid-Side, Lt/Rt. Type 0002.
//
//	Objects: For object-based audio where channels represent audio objects (or parts of objects), so include positional
//	         information. Type 0003.
//
//	HOA: For scene-based audio where Ambisonics and HOA are used. Type 0004.
//
//	Binaural: For binaural audio, where playback is over headphones. Type 0005.
type AudioBlockFormat struct {
	// TODO: either
	//       or
	//       or none (if audioChannelFormat.typeDefinition == "Binaural")
	//       must be specified

	// TODO: If audioChannelFormat.typeDefinition == "DirectSpeakers"
	//           use speakerLabel and position, Do not use outputChannelIDRef.
	//
	//       If audioChannelFormat.typeDefinition == "Matrix" / encoding or decoding
	//           use outputChannelIdref, matrix and outputChannelIdRef
	//
	//       If audioChannelFormat.typeDefinition == "Matrix" / direct
	//           use speakerLabel, matrix and outputChannelIdRef
	//
	//       If audioChannelFormat.typeDefinition == "Objects"
	//           use cartesian, position, Gain, Diffuse, Width, Height, Depth, ChannelLock, ObjectDivergence,
	//           JumpPosition, ZoneExclusion, Importance
	//
	//       If audioChannelFormat.typeDefinition == "HOA"
	//           use Degree, Order, Equation, NfcRefDist, Normalization

	// A label to indicate which loudspeaker should be used (e.g. M+30).
	SpeakerLabel []SpeakerLabel `xml:"urn:ebu:metadata-schema:ebucore ebucore:speakerLabel,omitempty" json:"ebucore:speakerLabel,omitempty"`

	// A set of user defined parameters to define a location in space.
	Position *[3]*AudioBlockFormatPosition `xml:"urn:ebu:metadata-schema:ebucore ebucore:position,omitempty" json:"ebucore:position,omitempty"`

	// A set of coefficients that refer to other channels to define a matrix.
	Matrix *Matrix `xml:"urn:ebu:metadata-schema:ebucore ebucore:matrix,omitempty" json:"ebucore:matrix,omitempty"`

	// For defining a decoding or direct matrix, this is the DirectSpeakers type output audioChannelFormat that defines
	// the channel being decoded to.
	OutputChannelIDRef *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:outputChannelIDRef,omitempty" json:"ebucore:outputChannelIDRef,omitempty"`

	// Set to true if cartesian coordinates used, false (default) for spherical.
	Cartesian *Boolean `xml:"urn:ebu:metadata-schema:ebucore ebucore:cartesian,omitempty" json:"ebucore:cartesian,omitempty"`

	// Set a multiplication factor for the audio samples in the block.
	// TODO: default="1.0"
	Gain *Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:gain,omitempty" json:"ebucore:gain,omitempty"`

	// Set to 1 if diffuse, 0 if direct.
	// TODO: default=0
	// TODO: 0.0 -- 1.0
	Diffuse *Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:diffuse,omitempty" json:"ebucore:diffuse,omitempty"`

	// Width of object in degrees along azimuth axis if using polar coordinates, and in the X-axis if using Cartesian
	// coordinates.
	// TODO: default=0
	Width *Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:width,omitempty" json:"ebucore:width,omitempty"`

	// Height of object in degrees along elevation axis if using polar coordinates, and in the Z-axis if using Cartesian
	// coordinates.
	// TODO: default=0
	Height *Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:height,omitempty" json:"ebucore:height,omitempty"`

	// Depth of object along distance line as a normalised distance if using polar coordinates, and in the Y-axis if using
	// Cartesian coordinates
	// TODO: default=0
	Depth *Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:depth,omitempty" json:"ebucore:depth,omitempty"`

	// If set to 1 a renderer can lock the object to the nearest channel or speaker, rather than normal rendering. The
	// optional maxDistance attribute defines the radius of a sphere around the object's position. If one or more speakers
	// exist in the defined sphere or on its surface, the object snaps to the nearest speaker. If maxDistance is
	// undefined, a default value of infinity is assumed, meaning that the object should snap to the nearest of all
	// speakers
	ChannelLock *ChannelLock `xml:"urn:ebu:metadata-schema:ebucore ebucore:channelLock,omitempty" json:"ebucore:channelLock,omitempty"`

	// Sets the divergence of an object where 0.0 means all sound in the direction of the object and 1.0 means all sound
	// from the virtual objects placed at an angle set by the azimuthRange attribute value.
	ObjectDivergence *ObjectDivergence `xml:"urn:ebu:metadata-schema:ebucore ebucore:objectDivergence,omitempty" json:"ebucore:objectDivergence,omitempty"`

	// If set to 1 the position will be interpolated over a period set by the attribute interpolationLength. If set to 0
	// then interpolation will take the entire length of the block. An interpolationLength value of zero will mean the
	// object jumps without interpolation.
	JumpPosition *JumpPosition `xml:"urn:ebu:metadata-schema:ebucore ebucore:jumpPosition,omitempty" json:"ebucore:jumpPosition,omitempty"`

	// Allows 3D zones to be exluded from rendering. So any speakers that exist within the excluded zone will not be used.
	ZoneExclusion *ZoneExclusion `xml:"urn:ebu:metadata-schema:ebucore ebucore:zoneExclusion,omitempty" json:"ebucore:zoneExclusion,omitempty"`

	// A value from 0 to 10 (most) to indicate the importance of the object.
	Importance *Int `xml:"urn:ebu:metadata-schema:ebucore ebucore:importance,omitempty" json:"ebucore:importance,omitempty"`

	// Degree for the ambisonic component.
	Degree *Int `xml:"urn:ebu:metadata-schema:ebucore ebucore:degree,omitempty" json:"ebucore:degree,omitempty"`

	// Order for the ambisonic component.
	Order *Int `xml:"urn:ebu:metadata-schema:ebucore ebucore:order,omitempty" json:"ebucore:order,omitempty"`

	// Equation of the ambisonic component.
	Equation *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:equation,omitempty" json:"ebucore:equation,omitempty"`

	// Indicates the reference distance of the loudspeaker setup for near-field compensation (NFC). If no nfcRefDist is
	// defined or the value is 0, NFC is not necessary.
	NfcRefDist *Int `xml:"urn:ebu:metadata-schema:ebucore ebucore:nfcRefDist,omitempty" json:"ebucore:nfcRefDist,omitempty"`

	// Indicates the normalization scheme of the HOA content (N3D, SN3D, FuMa).
	Normalization *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:normalization,omitempty" json:"ebucore:normalization,omitempty"`

	// Set to true if the object is related to the screen, false if not.
	ScreenRef *Boolean `xml:"urn:ebu:metadata-schema:ebucore ebucore:screenRef,omitempty" json:"ebucore:screenRef,omitempty"`

	// A unique ID for the audio block.
	// TODO: <pattern value="AB_([0-9a-fA-F]{4})([0-9a-fA-F]{4})_([0-9a-fA-F]{8})"/>
	AudioBlockFormatID string `xml:"audioBlockFormatID,attr,omitempty" json:"@audioBlockFormatID,omitempty"`

	// Start time of the block expressed as a timecode.
	Rtime TimecodeString `xml:"rtime,attr,omitempty" json:"@rtime,omitempty"`

	// Duration of the block expressed as a timecode.
	Duration TimecodeString `xml:"duration,attr,omitempty" json:"@duration,omitempty"`

	// To define the reference edit rate for the start and duration timecodes.
	EditRate float32 `xml:"editRate,attr,omitempty" json:"@editRate,omitempty"`

	// End time for the programme.
	End TimecodeString `xml:"end,attr,omitempty" json:"@end,omitempty"`

	// A flag to indicate whether the start and duration timecodes are drop frame when set to "true".
	// TODO: this field was commented out in the XML Schema
	//Dropframe bool `xml:"dropframe,attr,omitempty" json:"@dropframe,omitempty"`

	// The numerator of a correcting factor for the atart and duration timecodes.
	// TODO: this field was commented out in the XML Schema
	// TODO: default = 1
	//FactorNumerator int `xml:"factorNumerator,attr,omitempty" json:"@factorNumerator,omitempty"`

	// The denominator of a correcting factor for the start and duration timecodes.
	// TODO: this field was commented out in the XML Schema
	// TODO: default = 1
	//FactorDenominator int `xml:"factorDenominator,attr,omitempty" json:"@factorDenominator,omitempty"`
}

type SpeakerLabel String

type AudioBlockFormatPosition struct {
	Float

	// The axis used, either azimuth, elevation or distance (for polar); or X, Y or Z (for Cartesian)
	Coordinate Coordinate `xml:"coordinate,attr,omitempty" json:"@coordinate,omitempty"`

	// Set to "min" or "max" to set the bound of the position.
	Bound Bound `xml:"bound,attr,omitempty" json:"@bound,omitempty"`

	// Set to "left", "right", "top" or "bottom" to indicate which edge of the screen to lock the position to.
	ScreenEdgeLock ScreenEdgeLock `xml:"screenEdgeLock,attr,omitempty" json:"@screenEdgeLock,omitempty"`
}

type Coordinate string

const (
	AzimuthCoordinate   Coordinate = "azimuth"
	ElevationCoordinate Coordinate = "elevation"
	DistanceCoordinate  Coordinate = "distance"
	XCoordinate         Coordinate = "X"
	YCoordinate         Coordinate = "Y"
	ZCoordinate         Coordinate = "Z"
)

type Bound string

const (
	MinBound Bound = "min"
	MaxBound Bound = "max"
)

type ScreenEdgeLock string

const (
	LeftScreenEdgeLock   ScreenEdgeLock = "left"
	RightScreenEdgeLock  ScreenEdgeLock = "right"
	TopScreenEdgeLock    ScreenEdgeLock = "top"
	BottomScreenEdgeLock ScreenEdgeLock = "bottom"
)

type Matrix struct {
	// Sets a multiplication coefficient (value attrib) with the ID of another channel
	// TODO: required
	Coefficient []Coefficient `xml:"urn:ebu:metadata-schema:ebucore ebucore:coefficient,omitempty" json:"ebucore:coefficient,omitempty"`
}

type Coefficient struct {
	Value string `xml:",chardata" json:"#value"`

	// Multiplication factor of another channel. Constant value.
	Gain float32 `xml:"gain,attr,omitempty" json:"@gain,omitempty"`

	// Multiplication factor of another channel. Variable.
	GainVar string `xml:"gainVar,attr,omitempty" json:"@gainVar,omitempty"`

	// Phase shift of another channel. Constant value.
	Phase float32 `xml:"phase,attr,omitempty" json:"@phase,omitempty"`

	// Phase shift of another channel. Variable.
	PhaseVar string `xml:"phaseVar,attr,omitempty" json:"@phaseVar,omitempty"`

	// Time delay of another channel. Constant value.
	Delay string `xml:"delay,attr,omitempty" json:"@delay,omitempty"`

	// Time delay of another channel. Variable.
	DelayVar string `xml:"delayVar,attr,omitempty" json:"@delayVar,omitempty"`
}

type ObjectDivergence struct {
	Float

	// Adjusts the balance between the object's specified position and two other positions specified by the azimuthRange
	// value (symmetrical on both sides of the object at the object's position +/– azimuthRange). A value of 0 for the
	// objectDivergence means no divergence.
	AzimuthRange float32 `xml:"azimuthRange,attr,omitempty" json:"@azimuthRange,omitempty"`

	// Adjusts the balance between the object's specified position and two other positions specified by the
	// positionRange value (symmetrical on both sides of the object at the object's position +/– positionRange along the
	// X-axis). A value of 0 for the objectDivergence means no divergence.
	PositionRange float32 `xml:"positionRange,attr,omitempty" json:"@positionRange,omitempty"`
}

type ChannelLock struct {
	Value bool `xml:",chardata" json:"#value"`

	// The maximum distance that channel can lock within.
	MaxDistance float32 `xml:"maxDistance,attr,omitempty" json:"@maxDistance,omitempty"`
}

type JumpPosition struct {
	Value bool `xml:",chardata" json:"#value"`

	// The interpolation length in seconds.
	InterpolationLength float32 `xml:"interpolationLength,attr,omitempty" json:"@interpolationLength,omitempty"`
}

type ZoneExclusion struct {
	// Specifies the corner points of a cuboid in the 3D space that will be excluded from rendering. Multiple zone
	// elements can be used to specify more complex exclusion shapes.
	// TODO: required
	Zone []Zone `xml:"urn:ebu:metadata-schema:ebucore ebucore:zone,omitempty" json:"ebucore:zone,omitempty"`
}

type Zone struct {
	Value string `xml:",chardata" json:"#value"`

	MinX         float32 `xml:"minX,attr,omitempty" json:"@minX,omitempty"`
	MaxX         float32 `xml:"maxX,attr,omitempty" json:"@maxX,omitempty"`
	MinY         float32 `xml:"minY,attr,omitempty" json:"@minY,omitempty"`
	MaxY         float32 `xml:"maxY,attr,omitempty" json:"@maxY,omitempty"`
	MinZ         float32 `xml:"minZ,attr,omitempty" json:"@minZ,omitempty"`
	MaxZ         float32 `xml:"maxZ,attr,omitempty" json:"@maxZ,omitempty"`
	MinElevation float32 `xml:"minElevation,attr,omitempty" json:"@minElevation,omitempty"`
	MaxElevation float32 `xml:"maxElevation,attr,omitempty" json:"@maxElevation,omitempty"`
	MinAzimuth   float32 `xml:"minAzimuth,attr,omitempty" json:"@minAzimuth,omitempty"`
	MaxAzimuth   float32 `xml:"maxAzimuth,attr,omitempty" json:"@maxAzimuth,omitempty"`
}

// A combination of tracks (or one track) required to represent a channel or a pack. The audioStreamFormat establishes a
// relation between audioTrackFormats stored in the asset and audioChannelFormats produced by these audioTrackFormats.
// Its main use is to deal with non-PCM encoded tracks, where one or several audioTrackFormats produce several
// audioChannelFormats. The audioStreamFormat allows knowing which audioChannelFormats are available on an asset and how
// to effectively access to them. An audio stream can also be used to link together several PCM audioTrackFormats.
type AudioStreamFormat struct {
	// A format definition of the audio stream.
	FormatAttributes

	// A set of references to audio channels
	AudioChannelFormatIDRef []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioChannelFormatIDRef,omitempty" json:"ebucore:audioChannelFormatIDRef,omitempty"`

	// A set of references to audio packs
	AudioPackFormatIDRef []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioPackFormatIDRef,omitempty" json:"ebucore:audioPackFormatIDRef,omitempty"`

	// A set of references to audio tracks
	// TODO: required
	AudioTrackFormatIDRef []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioTrackFormatIDRef,omitempty" json:"ebucore:audioTrackFormatIDRef,omitempty"`

	// A unique ID for the audio stream.
	// TODO: <pattern value="AS_[0-9a-fA-F]{4}([0-9a-fA-F]{4})"/>
	// TODO: required
	AudioStreamFormatID string `xml:"audioStreamFormatID,attr,omitempty" json:"@audioStreamFormatID,omitempty"`

	// A unique human readable name for the audio stream.
	AudioStreamFormatName string `xml:"audioStreamFormatName,attr,omitempty" json:"@audioStreamFormatName,omitempty"`
}

// An audioTrack object defines a component of an audioStream.
//
// A single set of samples or data in the storage medium.
//
// An audioTrack is the basic audio data container of a medium. Attribute is an unambiguous reference to this container
// in a given medium
//
// Represents a physical container or carrier to hold an audio stream.
type AudioTrackFormat struct {
	FormatAttributes

	// A set of references to an audio stream format.
	AudioStreamFormatIDRef *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioStreamFormatIDRef,omitempty" json:"ebucore:audioStreamFormatIDRef,omitempty"`

	// A unique ID for the audio track
	// TODO: <pattern value="AT_([0-9a-fA-F]{4})([0-9a-fA-F]{4})_([0-9a-fA-F]{2})"/>
	AudioTrackFormatID string `xml:"audioTrackFormatID,attr,omitempty" json:"@audioTrackFormatID,omitempty"`

	// A unique name for the audio track
	AudioTrackFormatName string `xml:"audioTrackFormatName,attr,omitempty" json:"@audioTrackFormatName,omitempty"`
}

type AudioTrackUID struct {
	// Reference to lookup an MXF essence ID.
	AudioMXFLookUp *AudioMXFLookUp `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioMXFLookUp,omitempty" json:"ebucore:audioMXFLookUp,omitempty"`

	// Reference to an audioTrackFormat.
	AudioTrackFormatIDRef *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioTrackFormatIDRef,omitempty" json:"ebucore:audioTrackFormatIDRef,omitempty"`

	// Reference to an audioPackFormat.
	AudioPackFormatIDRef *String `xml:"urn:ebu:metadata-schema:ebucore ebucore:audioPackFormatIDRef,omitempty" json:"ebucore:audioPackFormatIDRef,omitempty"`

	// UID for the track.
	UID string `xml:"UID,attr,omitempty" json:"@UID,omitempty"`

	// The sample rate for the audio track.
	SampleRate int `xml:"sampleRate,attr,omitempty" json:"@sampleRate,omitempty"`

	// The bit depth for sample for the audio track.
	BitDepth int `xml:"bitDepth,attr,omitempty" json:"@bitDepth,omitempty"`
}

// A set of references to lookup MXF essence IDs.
type AudioMXFLookUp struct {
	// Reference to a package.
	// TODO: required
	PackageUIDRef PackageID `xml:"urn:ebu:metadata-schema:ebucore ebucore:packageUIDRef,omitempty" json:"ebucore:packageUIDRef,omitempty"`

	// Reference to a track.
	// TODO: required
	TrackIDRef String `xml:"urn:ebu:metadata-schema:ebucore ebucore:trackIDRef,omitempty" json:"ebucore:trackIDRef,omitempty"`

	// Reference to a channel.
	// TODO: required
	ChannelIDRef String `xml:"urn:ebu:metadata-schema:ebucore ebucore:channelIDRef,omitempty" json:"ebucore:channelIDRef,omitempty"`
}

// The pattern for providing package IDs.
// TODO: <pattern value="urn:smpte:umid:([0-9a-fA-F]{8}\.){7}[0-9a-fA-F]{8}"/>
type PackageID String

// A set of loudness parameters
type LoudnessMetadata struct {
	// Integrated loudness value in LUFS (Loudness Unit Full Scale).
	IntegratedLoudness *Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:integratedLoudness,omitempty" json:"ebucore:integratedLoudness,omitempty"`

	// Loudness range in LU (Loudness Unit).
	LoudnessRange *Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:loudnessRange,omitempty" json:"ebucore:loudnessRange,omitempty"`

	// Maximum true peak in dBTP (decibel True Peak.
	MaxTruePeak *Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:maxTruePeak,omitempty" json:"ebucore:maxTruePeak,omitempty"`

	// Maximum momentary loudness in LU.
	MaxMomentary *Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:maxMomentary,omitempty" json:"ebucore:maxMomentary,omitempty"`

	// Maximum short term loudness in LU.
	MaxShortTerm *Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:maxShortTerm,omitempty" json:"ebucore:maxShortTerm,omitempty"`

	// Dialogue loudness in LU.
	DialogLoudness *Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:dialogLoudness,omitempty" json:"ebucore:dialogLoudness,omitempty"`

	// Method or algorithm used to calculate the loudness.
	LoudnessMethod string `xml:"loudnessMethod,attr,omitempty" json:"@loudnessMethod,omitempty"`

	// Regional recommendend practice for correcting loudness levels.
	LoudnessRecType string `xml:"loudnessRecType,attr,omitempty" json:"@loudnessRecType,omitempty"`

	// Correction type for the audio, e.g. file-based or real-time.
	LoudnessCorrectionType string `xml:"loudnessCorrectionType,attr,omitempty" json:"@loudnessCorrectionType,omitempty"`
}

// The pattern to express the time using timecode compliant with SMPTE ST12-1:2008
// TODO: <pattern value="(([0-1][0-9])|([2][0-3])):[0-5][0-9]:[0-5][0-9](([.,])|([:;]))[0-9]{2,}" />
type TimecodeString string

// The pattern to express the time using timecode compliant with SMPTE ST12-1:2008
// TODO: <pattern value="([0-1][0-9]|[2][0-3]):[0-5][0-9]:[0-5][0-9]([:;][0-5][0-9]|[:;][0-5][0-9][.,][0-9]{5,})" />
type TimecodeStringFrame string

type Timecode struct {
	Value TimecodeString `xml:",chardata" json:"#value"`

	EditRate  int  `xml:"editRate,attr,omitempty" json:"@editRate,omitempty"`
	Dropframe bool `xml:"dropframe,attr,omitempty" json:"@dropframe,omitempty"`

	// The numerator of a correcting factor.
	// TODO: default="1">
	FactorNumerator int `xml:"factorNumerator,attr,omitempty" json:"@factorNumerator,omitempty"`

	// The denominator of a correcting factor.
	// TODO: default="1">
	FactorDenominator int `xml:"factorDenominator,attr,omitempty" json:"@factorDenominator,omitempty"`
}

type TimecodeFrame struct {
	Value TimecodeStringFrame `xml:",chardata" json:"#value"`

	EditRate  int  `xml:"editRate,attr,omitempty" json:"@editRate,omitempty"`
	Dropframe bool `xml:"dropframe,attr,omitempty" json:"@dropframe,omitempty"`

	// The numerator of a correcting factor.
	// TODO: default="1">
	FactorNumerator int `xml:"factorNumerator,attr,omitempty" json:"@factorNumerator,omitempty"`

	// The denominator of a correcting factor.
	// TODO: default="1">
	FactorDenominator int `xml:"factorDenominator,attr,omitempty" json:"@factorDenominator,omitempty"`
}

type TechnicalAttributes struct {
	// A technical attribute ot type string for user customisation/extension.
	TechnicalAttributeString []String `xml:"urn:ebu:metadata-schema:ebucore ebucore:technicalAttributeString,omitempty" json:"ebucore:technicalAttributeString,omitempty"`

	// A technical attribute ot type byte for user customisation/extension.
	TechnicalAttributeByte []Int8 `xml:"urn:ebu:metadata-schema:ebucore ebucore:technicalAttributeByte,omitempty" json:"ebucore:technicalAttributeByte,omitempty"`

	// A technical attribute ot type short for user customisation/extension.
	TechnicalAttributeShort []Int16 `xml:"urn:ebu:metadata-schema:ebucore ebucore:technicalAttributeShort,omitempty" json:"ebucore:technicalAttributeShort,omitempty"`

	// A technical attribute ot type integer for user customisation/extension.
	TechnicalAttributeInteger []Int32 `xml:"urn:ebu:metadata-schema:ebucore ebucore:technicalAttributeInteger,omitempty" json:"ebucore:technicalAttributeInteger,omitempty"`

	// A technical attribute ot type long for user customisation/extension.
	TechnicalAttributeLong []Int64 `xml:"urn:ebu:metadata-schema:ebucore ebucore:technicalAttributeLong,omitempty" json:"ebucore:technicalAttributeLong,omitempty"`

	// A technical attribute ot type unsigned byte for user customisation/extension.
	TechnicalAttributeUnsignedByte []UInt8 `xml:"urn:ebu:metadata-schema:ebucore ebucore:technicalAttributeUnsignedByte,omitempty" json:"ebucore:technicalAttributeUnsignedByte,omitempty"`

	// A technical attribute ot type unsigned short for user customisation/extension.
	TechnicalAttributeUnsignedShort []UInt16 `xml:"urn:ebu:metadata-schema:ebucore ebucore:technicalAttributeUnsignedShort,omitempty" json:"ebucore:technicalAttributeUnsignedShort,omitempty"`

	// A technical attribute ot type unsigned integer for user customisation/extension.
	TechnicalAttributeUnsignedInteger []UInt32 `xml:"urn:ebu:metadata-schema:ebucore ebucore:technicalAttributeUnsignedInteger,omitempty" json:"ebucore:technicalAttributeUnsignedInteger,omitempty"`

	// A technical attribute ot type unsigned long for user customisation/extension.
	TechnicalAttributeUnsignedLong []UInt64 `xml:"urn:ebu:metadata-schema:ebucore ebucore:technicalAttributeUnsignedLong,omitempty" json:"ebucore:technicalAttributeUnsignedLong,omitempty"`

	// A technical attribute ot type boolean for user customisation/extension.
	TechnicalAttributeBoolean []Boolean `xml:"urn:ebu:metadata-schema:ebucore ebucore:technicalAttributeBoolean,omitempty" json:"ebucore:technicalAttributeBoolean,omitempty"`

	// A technical attribute ot type float for user customisation/extension.
	TechnicalAttributeFloat []Float `xml:"urn:ebu:metadata-schema:ebucore ebucore:technicalAttributeFloat,omitempty" json:"ebucore:technicalAttributeFloat,omitempty"`

	// A technical attribute ot type rational for user customisation/extension.
	// TODO: default="1">
	TechnicalAttributeRational []TechnicalAttributeRational `xml:"urn:ebu:metadata-schema:ebucore ebucore:technicalAttributeRational,omitempty" json:"ebucore:technicalAttributeRational,omitempty"`

	// A technical attribute ot type anyURI for user customisation/extension.
	TechnicalAttributeURI []TechnicalAttributeURI `xml:"urn:ebu:metadata-schema:ebucore ebucore:technicalAttributeUri,omitempty" json:"ebucore:technicalAttributeUri,omitempty"`

	// A technical attribute ot type anyURI for user customisation/extension.
	TechnicalAttributeTimecode []TechnicalAttributeTimecode `xml:"urn:ebu:metadata-schema:ebucore ebucore:technicalAttributeTimecode,omitempty" json:"ebucore:technicalAttributeTimecode,omitempty"`
}

// FrameLayout, VideoSamplingRaster, VideoEncoding, or any other information element to be expressed textually.
type String struct {
	// A group of attributes to define the type of attribute.
	TypeAttributes

	// A group of attributes to define the format of the attribute.
	FormatAttributes

	Value string `xml:",chardata" json:"#value"`

	Unit string `xml:"unit,attr,omitempty" json:"@unit,omitempty"`
}

// IEEE 754 floating point number
// see http://www.smpte-ra.org/schemas/434/Types/TypeName
type Boolean struct {
	// A group of attributes to define the type of attribute.
	TypeAttributes

	Value bool `xml:",chardata" json:"#value"`
}

// IEEE 754 floating point number
// see http://www.smpte-ra.org/schemas/434/Types/TypeName
type Float struct {
	// A group of attributes to define the type of attribute.
	TypeAttributes

	Value float64 `xml:",chardata" json:"#value"`

	Unit string `xml:"unit,attr,omitempty" json:"@unit,omitempty"`
}

// Signed integer
type Int struct {
	// A group of attributes to define the type of attribute.
	TypeAttributes

	Value int `xml:",chardata" json:"#value"`

	Unit string `xml:"unit,attr,omitempty" json:"@unit,omitempty"`
}

// Signed 8 bit integer
// see http://www.smpte-ra.org/schemas/434/Types/TypeName
type Int8 struct {
	// A group of attributes to define the type of attribute.
	TypeAttributes

	Value int8 `xml:",chardata" json:"#value"`

	Unit string `xml:"unit,attr,omitempty" json:"@unit,omitempty"`
}

// Signed 16 bit integer
// see http://www.smpte-ra.org/schemas/434/Types/TypeName
type Int16 struct {
	// A group of attributes to define the type of attribute.
	TypeAttributes

	Value int16 `xml:",chardata" json:"#value"`

	Unit string `xml:"unit,attr,omitempty" json:"@unit,omitempty"`
}

// Signed 32 bit integer
// see http://www.smpte-ra.org/schemas/434/Types/TypeName
type Int32 struct {
	// A group of attributes to define the type of attribute.
	TypeAttributes

	Value int32 `xml:",chardata" json:"#value"`

	Unit string `xml:"unit,attr,omitempty" json:"@unit,omitempty"`
}

// Signed 64 bit integer
// see http://www.smpte-ra.org/schemas/434/Types/TypeName
type Int64 struct {
	// A group of attributes to define the type of attribute.
	TypeAttributes

	Value int64 `xml:",chardata" json:"#value"`

	Unit string `xml:"unit,attr,omitempty" json:"@unit,omitempty"`
}

// Unsigned integer
type UInt struct {
	// A group of attributes to define the type of attribute.
	TypeAttributes

	Value uint `xml:",chardata" json:"#value"`

	Unit string `xml:"unit,attr,omitempty" json:"@unit,omitempty"`
}

// Unsigned 8 bit integer
// see http://www.smpte-ra.org/schemas/434/Types/TypeName
type UInt8 struct {
	// A group of attributes to define the type of attribute.
	TypeAttributes

	Value uint8 `xml:",chardata" json:"#value"`

	Unit string `xml:"unit,attr,omitempty" json:"@unit,omitempty"`
}

// Unsigned 16 bit integer
// see http://www.smpte-ra.org/schemas/434/Types/TypeName
type UInt16 struct {
	// A group of attributes to define the type of attribute.
	TypeAttributes

	Value uint16 `xml:",chardata" json:"#value"`

	Unit string `xml:"unit,attr,omitempty" json:"@unit,omitempty"`
}

// Unsigned 32 bit integer
// see http://www.smpte-ra.org/schemas/434/Types/TypeName
type UInt32 struct {
	// A group of attributes to define the type of attribute.
	TypeAttributes

	Value uint64 `xml:",chardata" json:"#value"`

	Unit string `xml:"unit,attr,omitempty" json:"@unit,omitempty"`
}

// Unsigned 64 bit integer
// see http://www.smpte-ra.org/schemas/434/Types/TypeName
type UInt64 struct {
	// A group of attributes to define the type of attribute.
	TypeAttributes

	Value uint64 `xml:",chardata" json:"#value"`

	Unit string `xml:"unit,attr,omitempty" json:"@unit,omitempty"`
}

// A complex Type defining the structure of a technical attribute ot type rational
type Rational struct {
	Value int64 `xml:",chardata" json:"#value"`

	// The numerator of the fraction / rational number.
	// TODO: default="1"
	FactorNumerator int `xml:"factorNumerator,attr,omitempty" json:"@factorNumerator,omitempty"`

	// The denominator of the fraction / rational number.
	// TODO: default="1"
	FactorDenominator int `xml:"factorDenominator,attr,omitempty" json:"@factorDenominator,omitempty"`

	Unit string `xml:"unit,attr,omitempty" json:"@unit,omitempty"`
}

// A string to define e.g. the ratio of the picture (the width by the height), for instance '4:3' or '16:9' (rational).
type AspectRatio struct {
	// The type of aspect ratio.
	TypeAttributes

	// The numerator of the rational.
	// TODO: default="1"
	FactorNumerator Int `xml:"urn:ebu:metadata-schema:ebucore ebucore:factorNumerator,omitempty" json:"ebucore:factorNumerator,omitempty"`

	// The denominator of the rational.
	// TODO: default="1"
	FactorDenominator Int `xml:"urn:ebu:metadata-schema:ebucore ebucore:factorDenominator,omitempty" json:"ebucore:factorDenominator,omitempty"`
}

// A complex Type defining the structure of a technical attribute ot type rational.
type TechnicalAttributeRational struct {
	Rational

	// The type of rational.
	TypeAttributes
}

// A complex Type defining the structure of a technical attribute ot type anyURI.
type TechnicalAttributeURI struct {
	URIValue

	// The type of URI.
	TypeAttributes
}

// A complex Type defining the structure of a technical attribute ot type rational.
type TechnicalAttributeTimecode struct {
	Timecode

	// The type of rational.
	TypeAttributes
}

// A complex type to define a version associated with the type of version being referred to.
type Version struct {
	Element

	// The type of version information.
	TypeAttributes
}

type Element struct {
	dc.SimpleLiteral

	// An EBUCore extension of the dc:elementType (string + xml:lang) to precise the character set used to represent the
	// string. If the character set is different from the set declared for the whole metadata instance, then data
	// contained in the string should be escaped.
	// TODO: this field was commented out in the XML Schema
	//Charset string `xml:"charset,attr,omitempty" json:"@charset,omitempty"`

	// To specify if the string (using a particular character set if specified) has been encoded using e.g. base64Binary
	// encoding.
	// TODO: this field was commented out in the XML Schema
	//Encoding string `xml:"encoding,attr,omitempty" json:"@encoding,omitempty"`

	// To specify the scripting format as defined in in ISO 15924. This allows defining a translation format with
	// pronounciation of compex characters. The other option consists of providing the scripting format as part of the
	// xml:lang tag as defined in RFC 5646.
	// TODO: this field was commented out in the XML Schema
	//Transcription string `xml:"transcription,attr,omitempty" json:"@transcription,omitempty"`
}

// Defines the dynamic parameter/segment metadata extraction format. Each parameter is addressed one by one. Value(s)
// are given in an associated timeline made of aggregated time segments within which values can be listed at a given
// interval.
type ParameterSegmentDataOutput struct {
	// A list of parameters with their associated timelines, which value(s) are listed optionally at a time interval.
	Parameter []ParameterSegment `xml:"urn:ebu:metadata-schema:ebucore ebucore:parameter,omitempty" json:"ebucore:parameter,omitempty"`
}

// Defines the dynamic segment/parameter metadata extraction format. A series of time segments associated to a single
// parameter for which value(s) will be listed optionally at a time interval.
type SegmentParameterDataOutput struct {
	// A list of segments each associated to a single parameter, which value(s) are listed optionally at a time interval.
	Segment []SegmentParameter `xml:"urn:ebu:metadata-schema:ebucore ebucore:segment,omitempty" json:"ebucore:segment,omitempty"`
}

// A complexType to determine a dynamic parameter along an associated timeline. Some parameters are constant over the
// extraction time period. in this case only one segment appears with a simple or structured value. In this case
// interval=”.
type ParameterSegment struct {
	// To define a time segment by its startTime, endTime and optionally a time interval at which the parameter value is
	// extracted. The parameter value(s) is given by the segment element. If interval != '', the segment value is a list
	// of "space" separated results. If interval = '', the segment contains a single value, simple or structured.
	Segment []Segment `xml:"urn:ebu:metadata-schema:ebucore ebucore:segment,omitempty" json:"ebucore:segment,omitempty"`

	Name string `xml:"name,attr,omitempty" json:"@name,omitempty"`
}

type Segment struct {
	Value ValueList `xml:",chardata" json:"#value"`

	// The start time of the segment expressed in timecode / fremes.
	StartTime TimecodeString `xml:"startTime,attr,omitempty" json:"@startTime,omitempty"`

	// The end time of the segment expressed in timecode / fremes.
	EndTime TimecodeString `xml:"endTime,attr,omitempty" json:"@endTime,omitempty"`

	// The interval in frames (default "1" at which the parameter value is extracted. Interval = '' or null
	// means the parameter has a constant value over the time segment.
	Interval int `xml:"interval,attr,omitempty" json:"@interval,omitempty"`

	// An attribute to specify the unit in which the interval is expressed.
	Unit string `xml:"unit,attr,omitempty" json:"@unit,omitempty"`
}

// A complexType to determine a time segment and its associated parameter for which value(s) will be listed optionally
// at a time interval in the parameter element.
type SegmentParameter struct {
	// To define a parameter which value(s) will be returned within the time segment to which it belongs.
	Parameter []Parameter `xml:"urn:ebu:metadata-schema:ebucore ebucore:parameter,omitempty" json:"ebucore:parameter,omitempty"`

	// The startTime of the segment expressed in timecode / fremes.
	StartTime TimecodeString `xml:"startTime,attr,omitempty" json:"@startTime,omitempty"`

	// The endTime of the segment expressed in timecode / fremes.
	EndTime TimecodeString `xml:"endTime,attr,omitempty" json:"@endTime,omitempty"`
}

type Parameter struct {
	Value ValueList `xml:",chardata" json:"#value"`

	// The name of the parameter.
	Name string `xml:"name,attr,omitempty" json:"@name,omitempty"`

	// The interval in frames (default "1" at which the parameter value is extracted. Interval = '' or null
	// means the parameter has a constant value over the time segment.
	Interval int `xml:"interval,attr,omitempty" json:"@interval,omitempty"`

	// An attribute to specify the unit in which the interval is expressed.
	Unit string `xml:"unit,attr,omitempty" json:"@unit,omitempty"`
}

// a common simple type to support an unbounded space separated list of values as strings.
type ValueList string

// To provide additional contextual information.
type Comment struct {
	Element

	// A group of attributes to specify the type of comment.
	TypeAttributes
}

type URIValue struct {
	Value base.URI `xml:",chardata" json:"#value"`
}

type TimeValue struct {
	Value base.Time `xml:",chardata" json:"#value"`
}

type DurationValue struct {
	Value base.Duration `xml:",chardata" json:"#value"`
}
