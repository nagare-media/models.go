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

package v1_10_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/nagare-media/models.go/base"
	"github.com/nagare-media/models.go/dcmi/dc"
	ebucore "github.com/nagare-media/models.go/ebu/ebucore/v1.10"
	"github.com/nagare-media/models.go/third_party/encoding/xml"
	"github.com/senseyeio/duration"
)

var (
	lastModified = time.Now()

	want = &ebucore.Main{
		Version:               ebucore.SchemaVersion,
		WritingLibraryName:    "nagare media",
		WritingLibraryVersion: "1.0",
		DateLastModified:      base.Date{Time: lastModified},
		TimeLastModified:      base.Time{Time: lastModified},

		CoreMetadata: ebucore.CoreMetadata{
			Format: []ebucore.Format{{
				Duration: []ebucore.Duration{{NormalPlayTime: &ebucore.DurationValue{Value: base.Duration{Duration: duration.Duration{TM: 12, TS: 4}}}}},

				ContainerFormat: []ebucore.ContainerFormat{{
					ContainerFormatName: "MPEG-4",
					ContainerEncoding:   &ebucore.ContainerEncoding{FormatAttributes: ebucore.FormatAttributes{FormatLabel: "MPEG-4"}},
					Codec:               &ebucore.Codec{CodecIdentifier: &ebucore.Identifier{DCIdentifier: dc.Identifier{Value: "qt  "}}},
					TechnicalAttributes: ebucore.TechnicalAttributes{
						TechnicalAttributeString: []ebucore.String{{
							TypeAttributes: ebucore.TypeAttributes{TypeLabel: "FormatProfile"},
							Value:          "QuickTime",
						}, {
							TypeAttributes: ebucore.TypeAttributes{TypeLabel: "WritingApplication"},
							Value:          "Lavf58.76.100",
						}},
					},
				}},

				VideoFormat: []ebucore.VideoFormat{{
					VideoFormatName: "AVC",
					Width:           []ebucore.Width{{Dimension: ebucore.Dimension{Value: 4096, Unit: "pixel"}}},
					Height:          []ebucore.Height{{Dimension: ebucore.Dimension{Value: 1714, Unit: "pixel"}}},
					FrameRate:       &ebucore.Rational{Value: 24},
					AspectRatio: []ebucore.AspectRatio{{
						TypeAttributes:    ebucore.TypeAttributes{TypeLabel: "display"},
						FactorNumerator:   ebucore.Int{Value: 240},
						FactorDenominator: ebucore.Int{Value: 100},
					}},
					VideoEncoding:  &ebucore.VideoEncoding{TypeAttributes: ebucore.TypeAttributes{TypeLabel: "High 4:4:4 Predictive@L6"}},
					Codec:          &ebucore.Codec{CodecIdentifier: &ebucore.Identifier{DCIdentifier: dc.Identifier{Value: "avc1"}}},
					BitRate:        &ebucore.Dimension{Value: 61229438},
					ScanningFormat: ebucore.ProgressiveScanningFormat,
					VideoTrack:     []ebucore.VideoTrack{{TrackID: "1"}},
					TechnicalAttributes: ebucore.TechnicalAttributes{
						TechnicalAttributeString: []ebucore.String{{
							TypeAttributes: ebucore.TypeAttributes{TypeLabel: "ChromaSubsampling"},
							Value:          "4:4:4",
						}, {
							TypeAttributes: ebucore.TypeAttributes{TypeLabel: "WritingLibrary"},
							Value:          "x264 core 161 r3049 55d517b",
						}},
						TechnicalAttributeInteger: []ebucore.Int32{{
							TypeAttributes: ebucore.TypeAttributes{TypeLabel: "BitDepth"},
							Value:          10,
							Unit:           "bit",
						}},
						TechnicalAttributeLong: []ebucore.Int64{{
							TypeAttributes: ebucore.TypeAttributes{TypeLabel: "StreamSize"},
							Value:          5617800951,
							Unit:           "byte",
						}},
						TechnicalAttributeBoolean: []ebucore.Boolean{{
							TypeAttributes: ebucore.TypeAttributes{TypeLabel: "CABAC"},
							Value:          true,
						}, {
							TypeAttributes: ebucore.TypeAttributes{TypeLabel: "MBAFF"},
							Value:          false,
						}},
					},
				}},

				AudioFormat: []ebucore.AudioFormat{{
					AudioFormatName: "PCM",
					AudioEncoding: &ebucore.AudioEncoding{
						TypeAttributes: ebucore.TypeAttributes{
							TypeLabel: "PCM",
							TypeLink:  "http://www.ebu.ch/metadata/cs/ebu_AudioCompressionCodeCS.xml#11",
						},
					},
					Codec:        &ebucore.Codec{CodecIdentifier: &ebucore.Identifier{DCIdentifier: dc.Identifier{Value: "in24"}}},
					SamplingRate: &ebucore.Int64{Value: 48000},
					SampleSize:   &ebucore.UInt{Value: 24},
					BitRate:      &ebucore.Dimension{Value: 6912000},
					BitRateMode:  ebucore.ConstantBitRateMode,
					AudioTrack: []ebucore.AudioTrack{{
						TrackID:       "2",
						TrackLanguage: "en",
					}},
					Channels: &ebucore.UInt{Value: 6},
					TechnicalAttributes: ebucore.TechnicalAttributes{
						TechnicalAttributeString: []ebucore.String{{
							TypeAttributes: ebucore.TypeAttributes{TypeLabel: "ChannelPositions"},
							Value:          "Front: L C R, Back: L R, LFE",
						}, {
							TypeAttributes: ebucore.TypeAttributes{TypeLabel: "ChannelLayout"},
							Value:          "L R C LFE Ls Rs",
						}, {
							TypeAttributes: ebucore.TypeAttributes{TypeLabel: "Endianness"},
							Value:          "Little",
						}},
						TechnicalAttributeLong: []ebucore.Int64{{
							TypeAttributes: ebucore.TypeAttributes{TypeLabel: "StreamSize"},
							Value:          634176000,
							Unit:           "byte",
						}},
					},
				}, {
					AudioFormatName: "PCM",
					AudioEncoding: &ebucore.AudioEncoding{
						TypeAttributes: ebucore.TypeAttributes{
							TypeLabel: "PCM",
							TypeLink:  "http://www.ebu.ch/metadata/cs/ebu_AudioCompressionCodeCS.xml#11",
						},
					},
					Codec:        &ebucore.Codec{CodecIdentifier: &ebucore.Identifier{DCIdentifier: dc.Identifier{Value: "in24"}}},
					SamplingRate: &ebucore.Int64{Value: 48000},
					SampleSize:   &ebucore.UInt{Value: 24},
					BitRate:      &ebucore.Dimension{Value: 2304000},
					BitRateMode:  ebucore.ConstantBitRateMode,
					AudioTrack: []ebucore.AudioTrack{{
						TrackID:       "3",
						TrackLanguage: "en",
					}},
					Channels: &ebucore.UInt{Value: 2},
					TechnicalAttributes: ebucore.TechnicalAttributes{
						TechnicalAttributeString: []ebucore.String{{
							TypeAttributes: ebucore.TypeAttributes{TypeLabel: "ChannelPositions"},
							Value:          "Front: L R",
						}, {
							TypeAttributes: ebucore.TypeAttributes{TypeLabel: "ChannelLayout"},
							Value:          "L R",
						}, {
							TypeAttributes: ebucore.TypeAttributes{TypeLabel: "Endianness"},
							Value:          "Little",
						}},
						TechnicalAttributeLong: []ebucore.Int64{{
							TypeAttributes: ebucore.TypeAttributes{TypeLabel: "StreamSize"},
							Value:          211392000,
							Unit:           "byte",
						}},
					},
				}},

				DataFormat: []ebucore.DataFormat{{
					DataFormatName: "Timed Text",
					DataTrackID:    "4",
					CaptioningFormat: []ebucore.CaptioningFormat{{
						CaptioningFormatName: "Timed Text",
						TrackID:              "4",
						Language:             "en",
					}},
					Codec: &ebucore.Codec{CodecIdentifier: &ebucore.Identifier{DCIdentifier: dc.Identifier{Value: "text"}}},
				}, {
					DataFormatName: "Timed Text",
					DataTrackID:    "5",
					CaptioningFormat: []ebucore.CaptioningFormat{{
						CaptioningFormatName: "Timed Text",
						TrackID:              "5",
						Language:             "de",
					}},
					Codec: &ebucore.Codec{CodecIdentifier: &ebucore.Identifier{DCIdentifier: dc.Identifier{Value: "text"}}},
				}},

				FileInfo: ebucore.FileInfo{
					FileSize:       &ebucore.Dimension{Value: 6464125480},
					FileName:       &ebucore.String{Value: "tears-of-steel.mov"},
					Locator:        []ebucore.Locator{{URIValue: ebucore.URIValue{Value: "file://path/to/tears-of-steel.mov"}}},
					OverallBitRate: &ebucore.Dimension{Value: 70453684},
				},
			}},
		},
	}
)

// TODO: unmarshal

func TestMarshalXML(t *testing.T) {
	str, err := xml.MarshalIndent(want, "", "   ")
	if err != nil {
		t.Fatal("marshal failed")
	}
	// TODO: check output
	_ = str
}

func TestMarshalJSON(t *testing.T) {
	str, err := json.MarshalIndent(want, "", "   ")
	if err != nil {
		t.Fatal("marshal failed")
	}
	// TODO: check output
	_ = str
}
