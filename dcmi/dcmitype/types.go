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

package dcmitype

const (
	SchemaVersion = "2008-02-11"

	XMLNS       = "http://purl.org/dc/dcmitype/"
	XMLNSPrefix = "dcmitype"
)

type DCMIType string

const (
	// An aggregation of resources.
	//
	// A collection is described as a group; its parts may also be separately described.
	//
	// See http://purl.org/dc/dcmitype/Collection
	Collection DCMIType = "Collection"

	// Data encoded in a defined structure.
	//
	// Examples include lists, tables, and databases. A dataset may be useful for direct machine processing.
	//
	// See http://purl.org/dc/dcmitype/Dataset
	Dataset DCMIType = "Dataset"

	// A non-persistent, time-based occurrence.
	//
	// Metadata for an event provides descriptive information that is the basis for discovery of the purpose, location,
	// duration, and responsible agents associated with an event. Examples include an exhibition, webcast, conference,
	// workshop, open day, performance, battle, trial, wedding, tea party, conflagration.
	//
	// See http://purl.org/dc/dcmitype/Event
	Event DCMIType = "Event"

	// A visual representation other than text.
	//
	// Examples include images and photographs of physical objects, paintings, prints, drawings, other images and
	// graphics, animations and moving pictures, film, diagrams, maps, musical notation. Note that Image may include both
	// electronic and physical representations.
	//
	// See http://purl.org/dc/dcmitype/Image
	Image DCMIType = "Image"

	// A resource requiring interaction from the user to be understood, executed, or experienced.
	//
	// Examples include forms on Web pages, applets, multimedia learning objects, chat services, or virtual reality
	// environments.
	//
	// See http://purl.org/dc/dcmitype/MovingImage
	InteractiveResource DCMIType = "InteractiveResource"

	// A series of visual representations imparting an impression of motion when shown in succession.
	//
	// Examples include animations, movies, television programs, videos, zoetropes, or visual output from a simulation.
	// Instances of the type Moving Image must also be describable as instances of the broader type Image.
	//
	// See http://purl.org/dc/dcmitype/StillImage
	MovingImage DCMIType = "MovingImage"

	// An inanimate, three-dimensional object or substance.
	//
	// Note that digital representations of, or surrogates for, these objects should use Image, Text or one of the other
	// types.
	//
	// See http://purl.org/dc/dcmitype/PhysicalObject
	PhysicalObject DCMIType = "PhysicalObject"

	// A system that provides one or more functions.
	//
	// Examples include a photocopying service, a banking service, an authentication service, interlibrary loans, a Z39.50
	// or Web server.
	//
	// See http://purl.org/dc/dcmitype/Service
	Service DCMIType = "Service"

	// A computer program in source or compiled form.
	//
	// Examples include a C source file, MS-Windows .exe executable, or Perl script.
	//
	// See http://purl.org/dc/dcmitype/Software
	Software DCMIType = "Software"

	// A resource primarily intended to be heard.
	//
	// Examples include a music playback file format, an audio compact disc, and recorded speech or sounds.
	//
	// See http://purl.org/dc/dcmitype/Sound
	Sound DCMIType = "Sound"

	// A static visual representation.
	//
	// Examples include paintings, drawings, graphic designs, plans and maps. Recommended best practice is to assign the
	// type Text to images of textual materials. Instances of the type Still Image must also be describable as instances
	// of the broader type Image.
	//
	// See http://purl.org/dc/dcmitype/InteractiveResource
	StillImage DCMIType = "StillImage"

	// A resource consisting primarily of words for reading.
	//
	// Examples include books, letters, dissertations, poems, newspapers, articles, archives of mailing lists. Note that
	// facsimiles or images of texts are still of the genre Text.
	//
	// See http://purl.org/dc/dcmitype/Text
	Text DCMIType = "Text"
)
