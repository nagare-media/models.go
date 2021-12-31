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

package opencast_test

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/nagare-media/models.go/opencast"
	"github.com/nagare-media/models.go/x/encoding/xml"
)

var (
	start = time.Date(2021, 12, 31, 15, 20, 2, 422575000, time.FixedZone("CET", 60*60))

	want = &opencast.MediaPackage{
		ID:       "515e94e8-eb31-4400-a2dd-98f96d4cac22",
		Start:    &start,
		Duration: (7 * time.Minute).Milliseconds(),

		Title:        "Example",
		SeriesTitle:  "Example Series",
		Language:     "deu",
		Series:       "0c07eda3-d895-40dc-a92a-44f02ef42cd1",
		License:      "ALLRIGHTS",
		Creators:     []string{"Prof. John Doe"},
		Contributors: []string{"Max Mustermann"},
		Subjects:     []string{"example", "golang", "test"},

		Media: []opencast.Track{{
			MediaPackageElement: opencast.MediaPackageElement{
				ID:       "b2bc6d47-4790-4209-af0e-e035fceed403",
				Flavor:   "presenter/delivery",
				MimeType: "video/mp4",
				URL:      "https://cdn.opencast/assets/assets/515e94e8-eb31-4400-a2dd-98f96d4cac22/b2bc6d47-4790-4209-af0e-e035fceed403/1/presenter.mp4",
				Tags:     []string{"1080p-quality", "archive", "engage-download"},
				Size:     1302193,
				Checksum: &opencast.Checksum{
					Type:  opencast.MD5ChecksumType,
					Value: "18c24857427ed587c787583e85975b04",
				},
				Reference: &opencast.MediaPackageReference{
					Type: opencast.MediaPackageMediaPackageReference,
					ID:   opencast.SelfID,
					Properties: map[string]string{
						"example-for": "tests",
					},
				},
			},
			Duration:  (7 * time.Minute).Milliseconds(),
			Transport: opencast.DownloadStreamingProtocol,
			Audio: []opencast.AudioStream{
				{
					Stream: opencast.Stream{
						ID: "audio-1",
						Encoder: opencast.Encoder{
							Type: "AAC (Advanced Audio Coding)",
						},
						FrameCount: 21000,
					},
					BitDepth:     24,
					Channels:     2,
					SamplingRate: 48000,
					BitRate:      2280.0,
				},
			},
			Video: []opencast.VideoStream{
				{
					Stream: opencast.Stream{
						ID: "video-1",
						Encoder: opencast.Encoder{
							Type: "H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10",
						},
						FrameCount: 10500,
					},
					BitRate:   769473.0,
					FrameRate: 25,
					Resolution: &opencast.Resolution{
						Width:  1920,
						Height: 1080,
					},
					ScanType: &opencast.Scan{
						Type: opencast.ProgressiveScanType,
					},
				},
			},
		}},

		Metadata: []opencast.Catalog{{
			MediaPackageElement: opencast.MediaPackageElement{
				ID:       "305b30fe-c318-4cb9-a1e3-562703fa1df4",
				Flavor:   "dublincore/episode",
				MimeType: "text/xml",
				URL:      "https://cdn.opencast/assets/assets/515e94e8-eb31-4400-a2dd-98f96d4cac22/305b30fe-c318-4cb9-a1e3-562703fa1df4/1/dublincore.xml",
				Tags:     []string{"archive", "engage-download"},
				Checksum: &opencast.Checksum{
					Type:  opencast.MD5ChecksumType,
					Value: "611a78ab33e2270b9f4e6b947edb14f0",
				},
			},
		}},

		Attachments: []opencast.Attachment{{
			MediaPackageElement: opencast.MediaPackageElement{
				ID:       "f22c5937-8139-4c3e-9f89-e33d51830fe8",
				Flavor:   "presenter/timeline+preview",
				MimeType: "image/png",
				URL:      "https://cdn.opencast/assets/assets/515e94e8-eb31-4400-a2dd-98f96d4cac22/f22c5937-8139-4c3e-9f89-e33d51830fe8/1/timelinepreviews.png",
				Tags:     []string{"archive", "engage-download"},
				Size:     3168,
				Checksum: &opencast.Checksum{
					Type:  opencast.MD5ChecksumType,
					Value: "dc83e111ed27effc1d1b4d61aa1c33c1",
				},
				Reference: &opencast.MediaPackageReference{
					Type: opencast.TrackMediaPackageReference,
					ID:   "b2bc6d47-4790-4209-af0e-e035fceed403",
					Properties: map[string]string{
						"example-for": "tests",
					},
				},
			},
			Properties: map[string]string{
				"resolutionY": "-1",
				"imageCount":  "100",
				"resolutionX": "160",
				"imageSizeY":  "10",
				"imageSizeX":  "10",
			},
		}},

		Publications: []opencast.Publication{{
			MediaPackageElement: opencast.MediaPackageElement{
				ID:   "9643d5d3-3e2a-4ab6-8333-9777e916f710",
				Tags: []string{"archive"},
				URL:  "https://portal.opencast/watch/515e94e8-eb31-4400-a2dd-98f96d4cac22",
			},
			Channel: "engage-player",
		}, {
			MediaPackageElement: opencast.MediaPackageElement{
				ID:   "10582745-d600-4ab0-acfb-397e8f37b54d",
				Tags: []string{"archive"},
				URL:  "https://api.opencast/api/event/515e94e8-eb31-4400-a2dd-98f96d4cac22",
			},
			Channel: "api",
		}},
	}
)

func TestUnmarshalMediaPackage(t *testing.T) {
	str, err := ioutil.ReadFile("testdata/fixture_mediapackage_example.xml")
	if err != nil {
		t.Fatal("could not read file")
	}

	got := &opencast.MediaPackage{}
	err = xml.Unmarshal(str, got)
	if err != nil {
		t.Fatalf("could not unmarshal media package: %s", err)
	}

	if diff := cmp.Diff(got, want, cmpopts.IgnoreFields(opencast.MediaPackage{}, "XMLName")); diff != "" {
		t.Fatalf("unexpected value: %s", diff)
	}
}

func TestMarshalMediaPackage(t *testing.T) {
	str, err := xml.MarshalIndent(want, "", "   ")
	if err != nil {
		t.Fatal("marshal failed")
	}
	// TODO: check output
	_ = str
}
