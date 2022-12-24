/*
Copyright 2022 The nagare media authors

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

package v2_test

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/nagare-media/models.go/base"
	nbmp "github.com/nagare-media/models.go/iso/nbmp/v2"
)

func TestUnmarshalConfiguration(t *testing.T) {
	files := []string{
		"testdata/example-parameter-representation.json",
	}

	for _, file := range files {
		str, err := os.ReadFile(file)
		if err != nil {
			t.Fatalf("could not read file %s: %s", file, err)
		}

		cfg := nbmp.Configuration{}
		decoder := json.NewDecoder(bytes.NewReader(str))
		decoder.DisallowUnknownFields()
		err = decoder.Decode(&cfg)
		if err != nil {
			t.Fatalf("could not unmarshal configuration: %s", err)
		}
	}
}

func TestUnmarshalFunction(t *testing.T) {
	files := []string{
		// TODO: These examples are provided by MPEG, but do not follow the normative JSON schema
		//       (see https://github.com/MPEGGroup/NBMP/issues/40)
		// "testdata/function-RTP-360sticher.json",
		// "testdata/function-RTP-cgtranscoder.json",
		// "testdata/function-RTP-cropping.json",
		// "testdata/function-RTP-dash-packager.json",
		// "testdata/function-RTP-fifo.json",
		// "testdata/function-RTP-omaf-packager.json",
		// "testdata/function-RTP-pcdecoder.json",
		// "testdata/function-RTP-pcencoder.json",
		// "testdata/function-RTP-selectorcompositor.json",
		"testdata/function-RTP-merger.json",
		"testdata/function-RTP-splitter.json",
	}

	for _, file := range files {
		str, err := os.ReadFile(file)
		if err != nil {
			t.Fatalf("could not read file %s: %s", file, err)
		}

		fd := nbmp.Function{}
		decoder := json.NewDecoder(bytes.NewReader(str))
		decoder.DisallowUnknownFields()
		err = decoder.Decode(&fd)
		if err != nil {
			t.Fatalf("could not unmarshal function: %s", err)
		}
	}
}

func TestMarshalWorkflow(t *testing.T) {
	var (
		nagareBrand = base.URI("urn:mpeg:mpegi:nbmp:2023:nagare")
		now         = time.Now()
		videoIn1    = "videoIn1"
		videoOut1   = "videoOut1"
		videoOut2   = "videoOut2"
		breakable   = true
		vcpu        = uint64(2)
		ram         = uint64(1 * 1024)
	)

	wf := nbmp.Workflow{
		Scheme: &nbmp.Scheme{nbmp.SchemaURI},
		General: nbmp.General{
			ID:            "0dec3a70-2b4a-4444-8904-387f896bb90b",
			NBMPBrand:     &nagareBrand,
			State:         &nbmp.InstantiatedState,
			Name:          "Test workflow",
			Description:   "This workflow is used as a test case",
			PublishedTime: &now,
		},

		Repository: &nbmp.Repository{
			Mode: &nbmp.StrictRepositoryMode,
			Location: []nbmp.RepositoryLocation{{
				URL:         "https://engine.nagare.media/functions",
				Name:        "Example function repository",
				Description: "This is a function repository example",
			}},
		},

		Input: nbmp.Input{
			MediaParameters: []nbmp.MediaParameter{{
				StreamID:         videoIn1,
				Name:             videoIn1,
				Keywords:         []string{"master"},
				MimeType:         "video/mp4",
				Protocol:         "http",
				Mode:             &nbmp.PullMediaAccessMode,
				CachingServerURL: "https://engine.nagare.media/master.mp4",
				VideoFormat: []nbmp.Parameter{{
					Name:     "stereo3d",
					ID:       1,
					Datatype: nbmp.StringDatatype,
					Values: []nbmp.ParameterValue{nbmp.StringParameterValue{
						Name:         "value",
						ID:           11,
						Restrictions: []string{"side by side"},
					}},
				}},
			}},
		},

		Output: nbmp.Output{
			MediaParameters: []nbmp.MediaParameter{{
				StreamID:         videoOut1,
				Name:             videoOut1,
				Keywords:         []string{"720p"},
				MimeType:         "video/mp4",
				Protocol:         "http",
				Mode:             &nbmp.PushMediaAccessMode,
				CachingServerURL: "http://engine.nagare.media/720p.mp4",
			}, {
				StreamID:         videoOut2,
				Name:             videoOut2,
				Keywords:         []string{"livestream"},
				MimeType:         "video/mp4",
				Protocol:         "rtmp",
				Mode:             &nbmp.PushMediaAccessMode,
				CachingServerURL: "rtmp://engine.nagare.media/live/stream",
			}},
		},

		Processing: nbmp.Processing{
			Keywords: []string{},
			Image:    []nbmp.ProcessingImage{},

			ConnectionMap: []nbmp.ConnectionMapping{{
				ConnectionID: "transcode720p -> loop-stream",
				Breakable:    &breakable,
				From: nbmp.ConnectionMappingPort{
					ID:       "transcode-function",
					Instance: "transcode720p",
					PortName: "video-output",
				},
				To: nbmp.ConnectionMappingPort{
					ID:       "push-rtmp-stream-function",
					Instance: "loop-stream",
					PortName: "video-input",
				},
			}},

			FunctionRestrictions: []nbmp.FunctionRestriction{{
				Instance: "transcode720p",
				General: &nbmp.General{
					ID:          "transcode720p",
					Name:        "transcode720p",
					Description: "Transcode input video to 720p",
					InputPorts: []nbmp.Port{{
						PortName: "video-input",
						Bind: nbmp.PortBinding{
							StreamID: &videoIn1,
							Name:     "video-input",
						},
					}},
					OutputPorts: []nbmp.Port{{
						PortName: "video-output",
						Bind: nbmp.PortBinding{
							StreamID: &videoOut1,
							Name:     "video-output",
						},
					}},
				},
				Requirements: &nbmp.Requirement{
					Hardware: &nbmp.HardwareRequirement{
						VCPU: &vcpu,
						RAM:  &ram,
					},
				},
				Configuration: &nbmp.Configuration{Parameters: []nbmp.Parameter{{
					Name:     "codec",
					ID:       1,
					Datatype: nbmp.StringDatatype,
					Values: []nbmp.ParameterValue{nbmp.StringParameterValue{
						Name:         "codec",
						ID:           2,
						Restrictions: []string{"h264"},
					}},
				}}},
			}, {
				Instance: "loop-stream",
				General: &nbmp.General{
					ID:          "loop-stream",
					Name:        "loop-stream",
					Description: "Stream input video as loop to an RTMP destination",
					InputPorts: []nbmp.Port{{
						PortName: "video-input",
						Bind: nbmp.PortBinding{
							Name: "video-input",
						},
					}},
					OutputPorts: []nbmp.Port{{
						PortName: "video-output",
						Bind: nbmp.PortBinding{
							StreamID: &videoOut2,
							Name:     "video-output",
						},
					}},
				},
			}},
		},

		Requirement: nbmp.Requirement{
			Security: &nbmp.SecurityRequirement{
				TLS: true,
			},
			WorkflowTask: &nbmp.WorkflowTaskRequirement{
				ExecutionMode: &nbmp.StreamingExecutionMode,
			},
		},

		Failover: &nbmp.Failover{
			FailoverMode: nbmp.ContinueWithLastGoodStateFailoverMode,
		},
	}

	_, err := json.Marshal(wf)
	if err != nil {
		t.Fatalf("could not marshal to JSON: %s:", err)
	}
}
