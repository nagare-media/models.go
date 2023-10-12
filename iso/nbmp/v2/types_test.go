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

var (
	nagareBrand = base.URI("urn:nagare-media:engine:schema:nbmp:v1apha1")
	now         = time.Now()
)

func TestUnmarshalConfigurations(t *testing.T) {
	files := []string{
		"testdata/nbmp/example-parameter-representation.json",
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

func TestUnmarshalFunctions(t *testing.T) {
	files := []string{
		// TODO: These examples are provided by MPEG, but do not follow the normative JSON schema
		//       (see https://github.com/MPEGGroup/NBMP/issues/40)
		// "testdata/nbmp/function-RTP-360sticher.json",
		// "testdata/nbmp/function-RTP-cgtranscoder.json",
		// "testdata/nbmp/function-RTP-cropping.json",
		// "testdata/nbmp/function-RTP-dash-packager.json",
		// "testdata/nbmp/function-RTP-fifo.json",
		// "testdata/nbmp/function-RTP-omaf-packager.json",
		// "testdata/nbmp/function-RTP-pcdecoder.json",
		// "testdata/nbmp/function-RTP-pcencoder.json",
		// "testdata/nbmp/function-RTP-selectorcompositor.json",
		"testdata/nbmp/function-RTP-merger.json",
		"testdata/nbmp/function-RTP-splitter.json",
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

func TestUnmarshalWorkflows(t *testing.T) {
	files := []string{
		"testdata/nagare/v2_nbmp_live.wdd",
		"testdata/nagare/v2_nbmp_vod.wdd",
	}

	for _, file := range files {
		str, err := os.ReadFile(file)
		if err != nil {
			t.Fatalf("could not read file %s: %s", file, err)
		}

		wd := nbmp.Workflow{}
		decoder := json.NewDecoder(bytes.NewReader(str))
		decoder.DisallowUnknownFields()
		err = decoder.Decode(&wd)
		if err != nil {
			t.Fatalf("could not unmarshal workflow: %s", err)
		}
	}
}

func TestMarshalWorkflowLive(t *testing.T) {
	var breakable = true

	wf := nbmp.Workflow{
		Scheme: &nbmp.Scheme{nbmp.SchemaURI},
		General: nbmp.General{
			ID:            "7e5b8825-9442-499f-95fe-abf8f03970c7",
			NBMPBrand:     &nagareBrand,
			State:         &nbmp.InstantiatedState,
			Name:          "Live workflow",
			Description:   "A simple live workflow.",
			PublishedTime: &now,
		},

		Input: nbmp.Input{
			MediaParameters: []nbmp.MediaParameter{{
				StreamID:         "input-wf",
				Name:             "input-wf",
				MimeType:         "video/mp4",
				Protocol:         "rtmp",
				Mode:             &nbmp.PullMediaAccessMode,
				CachingServerURL: "rtmp://nagare.media/app/input",
				Keywords:         []string{},
			}},
		},

		Output: nbmp.Output{
			MediaParameters: []nbmp.MediaParameter{{
				StreamID:         "output-wf",
				Name:             "output-wf",
				MimeType:         "video/mp4",
				Protocol:         "dash-cmaf-ingest",
				Mode:             &nbmp.PushMediaAccessMode,
				CachingServerURL: base.URI("http://nagare.media/cmaf/example.str/Switching(video)/Streams(output.cmfv)"),
				Keywords:         []string{},
			}},
		},

		Processing: nbmp.Processing{
			Keywords: []string{},
			Image:    []nbmp.ProcessingImage{},

			ConnectionMap: []nbmp.ConnectionMapping{
				{
					ConnectionID: "watermark -> package-cmaf",
					Breakable:    &breakable,
					From: nbmp.ConnectionMappingPort{
						ID:       "watermark-function",
						Instance: "watermark",
						PortName: "output-watermark",
					},
					To: nbmp.ConnectionMappingPort{
						ID:       "package-function",
						Instance: "package-cmaf",
						PortName: "input-package-cmaf",
					},
				},
			},

			FunctionRestrictions: []nbmp.FunctionRestriction{
				{
					Instance: "watermark",
					General: &nbmp.General{
						ID:          "watermark",
						Name:        "watermark",
						Description: "watermark",
						NBMPBrand:   &nagareBrand,
						InputPorts: []nbmp.Port{{
							PortName: "input-package-cmaf",
							Bind: &nbmp.PortBinding{
								StreamID: "input-wf",
							},
						}},
						OutputPorts: []nbmp.Port{{
							PortName: "output-watermark",
						}},
					},
				},
				{
					Instance: "package-cmaf",
					General: &nbmp.General{
						ID:          "package-cmaf",
						Name:        "package-cmaf",
						Description: "package-cmaf",
						NBMPBrand:   &nagareBrand,
						InputPorts: []nbmp.Port{{
							PortName: "input-package-cmaf",
						}},
						OutputPorts: []nbmp.Port{{
							PortName: "output-package-cmaf",
							Bind: &nbmp.PortBinding{
								StreamID: "output-wf",
							},
						}},
					},
				},
			},
		},

		Requirement: nbmp.Requirement{
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

func TestMarshalWorkflowVoD(t *testing.T) {
	var breakable = true

	wf := nbmp.Workflow{
		Scheme: &nbmp.Scheme{nbmp.SchemaURI},
		General: nbmp.General{
			ID:            "8406e61f-9ff5-4000-9185-3af62b54108c",
			NBMPBrand:     &nagareBrand,
			State:         &nbmp.InstantiatedState,
			Name:          "VoD workflow",
			Description:   "A simple VoD workflow.",
			PublishedTime: &now,
		},

		Input: nbmp.Input{
			MediaParameters: []nbmp.MediaParameter{{
				StreamID:         "input-wf",
				Name:             "input-wf",
				MimeType:         "video/mp4",
				Protocol:         "http",
				Mode:             &nbmp.PullMediaAccessMode,
				CachingServerURL: "https://nagare.media/input.mp4",
				Keywords:         []string{},
			}},
		},

		Output: nbmp.Output{
			MediaParameters: []nbmp.MediaParameter{{
				StreamID:         "output-master-playlist-wf",
				Name:             "output-master-playlist-wf",
				MimeType:         "application/vnd.apple.mpegurl",
				Protocol:         "s3",
				Mode:             &nbmp.PushMediaAccessMode,
				CachingServerURL: base.URI("s3://nagare.media/output/master.m3u8"),
				Keywords:         []string{},
			}, {
				StreamID:         "output-1080p-playlist-wf",
				Name:             "output-1080p-playlist-wf",
				MimeType:         "application/vnd.apple.mpegurl",
				Protocol:         "s3",
				Mode:             &nbmp.PushMediaAccessMode,
				CachingServerURL: base.URI("s3://nagare.media/output/variant-1080p.m3u8"),
				Keywords:         []string{},
			}, {
				StreamID:         "output-1080p-wf",
				Name:             "output-1080p-wf",
				MimeType:         "video/mp4",
				Protocol:         "s3",
				Mode:             &nbmp.PushMediaAccessMode,
				CachingServerURL: base.URI("s3://nagare.media/output/1080p.cmfv"),
				Keywords:         []string{},
			}, {
				StreamID:         "output-720p-playlist-wf",
				Name:             "output-720p-playlist-wf",
				MimeType:         "application/vnd.apple.mpegurl",
				Protocol:         "s3",
				Mode:             &nbmp.PushMediaAccessMode,
				CachingServerURL: base.URI("s3://nagare.media/output/variant-720p.m3u8"),
				Keywords:         []string{},
			}, {
				StreamID:         "output-720p-wf",
				Name:             "output-720p-wf",
				MimeType:         "video/mp4",
				Protocol:         "s3",
				Mode:             &nbmp.PushMediaAccessMode,
				CachingServerURL: base.URI("s3://nagare.media/output/720p.cmfv"),
				Keywords:         []string{},
			}, {
				StreamID:         "output-audio-playlist-wf",
				Name:             "output-audio-playlist-wf",
				MimeType:         "application/vnd.apple.mpegurl",
				Protocol:         "s3",
				Mode:             &nbmp.PushMediaAccessMode,
				CachingServerURL: base.URI("s3://nagare.media/output/variant-audio.m3u8"),
				Keywords:         []string{},
			}, {
				StreamID:         "output-audio-wf",
				Name:             "output-audio-wf",
				MimeType:         "audio/mp4",
				Protocol:         "s3",
				Mode:             &nbmp.PushMediaAccessMode,
				CachingServerURL: base.URI("s3://nagare.media/output/audio.cmfa"),
				Keywords:         []string{},
			}},
		},

		Processing: nbmp.Processing{
			Keywords: []string{},
			Image:    []nbmp.ProcessingImage{},

			ConnectionMap: []nbmp.ConnectionMapping{
				{
					ConnectionID: "transcode-1080p -> package-cmaf-hls",
					Breakable:    &breakable,
					From: nbmp.ConnectionMappingPort{
						ID:       "transcode-function",
						Instance: "transcode-1080p",
						PortName: "output-transcode-1080p",
					},
					To: nbmp.ConnectionMappingPort{
						ID:       "package-function",
						Instance: "package-cmaf-hls",
						PortName: "input-1080p-package-cmaf-hls",
					},
				},
				{
					ConnectionID: "transcode-720p -> package-cmaf-hls",
					Breakable:    &breakable,
					From: nbmp.ConnectionMappingPort{
						ID:       "transcode-function",
						Instance: "transcode-720p",
						PortName: "output-transcode-720p",
					},
					To: nbmp.ConnectionMappingPort{
						ID:       "package-function",
						Instance: "package-cmaf-hls",
						PortName: "input-720p-package-cmaf-hls",
					},
				},
				{
					ConnectionID: "transcode-audio -> package-cmaf-hls",
					Breakable:    &breakable,
					From: nbmp.ConnectionMappingPort{
						ID:       "transcode-function",
						Instance: "transcode-audio",
						PortName: "output-transcode-audio",
					},
					To: nbmp.ConnectionMappingPort{
						ID:       "package-function",
						Instance: "package-cmaf-hls",
						PortName: "input-audio-package-cmaf-hls",
					},
				},
			},

			FunctionRestrictions: []nbmp.FunctionRestriction{
				{
					Instance: "transcode-1080p",
					General: &nbmp.General{
						ID:          "transcode-1080p",
						Name:        "transcode-1080p",
						Description: "transcode-1080p",
						NBMPBrand:   &nagareBrand,
						InputPorts: []nbmp.Port{{
							PortName: "input-transcode-1080p",
							Bind: &nbmp.PortBinding{
								StreamID: "input-wf",
							},
						}},
						OutputPorts: []nbmp.Port{{
							PortName: "output-transcode-1080p",
						}},
					},
				},
				{
					Instance: "transcode-720p",
					General: &nbmp.General{
						ID:          "transcode-720p",
						Name:        "transcode-720p",
						Description: "transcode-720p",
						NBMPBrand:   &nagareBrand,
						InputPorts: []nbmp.Port{{
							PortName: "input-transcode-720p",
							Bind: &nbmp.PortBinding{
								StreamID: "input-wf",
							},
						}},
						OutputPorts: []nbmp.Port{{
							PortName: "output-transcode-720p",
						}},
					},
				},
				{
					Instance: "transcode-audio",
					General: &nbmp.General{
						ID:          "transcode-audio",
						Name:        "transcode-audio",
						Description: "transcode-audio",
						NBMPBrand:   &nagareBrand,
						InputPorts: []nbmp.Port{{
							PortName: "input-transcode-audio",
							Bind: &nbmp.PortBinding{
								StreamID: "input-wf",
							},
						}},
						OutputPorts: []nbmp.Port{{
							PortName: "output-transcode-audio",
						}},
					},
				},
				{
					Instance: "package-cmaf-hls",
					General: &nbmp.General{
						ID:          "package-cmaf-hls",
						Name:        "package-cmaf-hls",
						Description: "package-cmaf-hls",
						NBMPBrand:   &nagareBrand,
						InputPorts: []nbmp.Port{
							{
								PortName: "input-1080p-package-cmaf-hls",
							},
							{
								PortName: "input-720p-package-cmaf-hls",
							},
							{
								PortName: "input-audio-package-cmaf-hls",
							},
						},
						OutputPorts: []nbmp.Port{
							{
								PortName: "output-master-playlist-package-cmaf-hls",
								Bind: &nbmp.PortBinding{
									StreamID: "output-master-playlist-wf",
								},
							},
							{
								PortName: "output-1080p-playlist-package-cmaf-hls",
								Bind: &nbmp.PortBinding{
									StreamID: "output-1080p-playlist-wf",
								},
							},
							{
								PortName: "output-1080p-package-cmaf-hls",
								Bind: &nbmp.PortBinding{
									StreamID: "output-1080p-wf",
								},
							},
							{
								PortName: "output-720p-playlist-package-cmaf-hls",
								Bind: &nbmp.PortBinding{
									StreamID: "output-720p-playlist-wf",
								},
							},
							{
								PortName: "output-720p-package-cmaf-hls",
								Bind: &nbmp.PortBinding{
									StreamID: "output-720p-wf",
								},
							},
							{
								PortName: "output-audio-playlist-package-cmaf-hls",
								Bind: &nbmp.PortBinding{
									StreamID: "output-audio-playlist-wf",
								},
							},
							{
								PortName: "output-audio-package-cmaf-hls",
								Bind: &nbmp.PortBinding{
									StreamID: "output-audio-wf",
								},
							},
						},
					},
				},
			},
		},

		Requirement: nbmp.Requirement{
			Security: &nbmp.SecurityRequirement{
				TLS: true,
			},
			WorkflowTask: &nbmp.WorkflowTaskRequirement{
				ExecutionMode: &nbmp.StepExecutionMode,
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
