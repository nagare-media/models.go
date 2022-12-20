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

package dc_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/nagare-media/models.go/dcmi/dc"
	"github.com/nagare-media/models.go/third_party/encoding/xml"
)

var (
	want = &dc.Elements{
		XMLName: xml.Name{Space: "http://example.org/myapp/", Local: "metadata"},
		Title: []dc.Title{
			{"TestTitle", "", ""},
			{"TestTitleDe", "de", ""},
			{"TestTitleNl", "nl", ""},
		},
		Creator:     []dc.Creator{{"TestCreator", "", ""}},
		Subject:     []dc.Subject{{"TestSubject", "", ""}},
		Description: []dc.Description{{"TestDescription", "", ""}},
		Publisher:   []dc.Publisher{{"TestPublisher", "", ""}},
		Contributor: []dc.Contributor{{"TestContributor", "", ""}},
		Date:        []dc.Date{{"TestDate", "", ""}},
		Type:        []dc.Type{{"TestType", "", ""}},
		Format:      []dc.Format{{"TestFormat", "", ""}},
		Identifier:  []dc.Identifier{{"TestIdentifier", "", ""}},
		Source:      []dc.Source{{"TestSource", "", ""}},
		Language:    []dc.Language{{"TestLanguage", "", ""}},
		Relation:    []dc.Relation{{"TestRelation", "", ""}},
		Coverage:    []dc.Coverage{{"TestCoverage", "", ""}},
		Rights:      []dc.Rights{{"TestRights", "", ""}},
	}
)

func TestUnmarshalXML(t *testing.T) {
	str, err := ioutil.ReadFile("testdata/fixture_example.xml")
	if err != nil {
		t.Fatal("reading file failed")
	}

	got := &dc.Elements{}
	err = xml.Unmarshal(str, got)
	if err != nil {
		t.Fatalf("unmarshal failed: %s", err)
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Fatalf("unexpected value: %s", diff)
	}
}

func TestMarshalXML(t *testing.T) {
	str, err := xml.MarshalIndent(want, "", "   ")
	if err != nil {
		t.Fatal("marshal failed")
	}
	// TODO: check output
	_ = str
}

func TestUnmarshalJSON(t *testing.T) {
	str, err := ioutil.ReadFile("testdata/fixture_example.json")
	if err != nil {
		t.Fatal("reading file failed")
	}

	got := &dc.Elements{}
	err = json.Unmarshal(str, got)
	if err != nil {
		t.Fatalf("unmarshal failed: %s", err)
	}

	if diff := cmp.Diff(got, want, cmpopts.IgnoreFields(dc.Elements{}, "XMLName")); diff != "" {
		t.Fatalf("unexpected value: %s", diff)
	}
}

func TestMarshalJSON(t *testing.T) {
	str, err := json.MarshalIndent(want, "", "   ")
	if err != nil {
		t.Fatal("marshal failed")
	}
	// TODO: check output
	_ = str
}
