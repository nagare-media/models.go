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

package opencast

import (
	"encoding"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Track struct {
	MediaPackageElement

	Duration    int64             `xml:"http://mediapackage.opencastproject.org duration,omitempty"`
	Live        bool              `xml:"http://mediapackage.opencastproject.org live,omitempty"`
	Master      bool              `xml:"http://mediapackage.opencastproject.org master,omitempty"`
	LogicalName string            `xml:"http://mediapackage.opencastproject.org logicalname,omitempty"`
	Transport   StreamingProtocol `xml:"http://mediapackage.opencastproject.org transport,omitempty"`
	Audio       []AudioStream     `xml:"http://mediapackage.opencastproject.org audio,omitempty"`
	Video       []VideoStream     `xml:"http://mediapackage.opencastproject.org video,omitempty"`
	Subtitle    []SubtitleStream  `xml:"http://mediapackage.opencastproject.org subtitle,omitempty"`
}

type Stream struct {
	ID string `xml:"id,attr,omitempty"`

	FrameCount int64   `xml:"http://mediapackage.opencastproject.org framecount,omitempty"`
	Device     Device  `xml:"http://mediapackage.opencastproject.org device,omitempty"`
	Encoder    Encoder `xml:"http://mediapackage.opencastproject.org encoder,omitempty"`
}

type Device struct {
	Type    string `xml:"type,attr,omitempty"`
	Version string `xml:"version,attr,omitempty"`
	Vendor  string `xml:"vendor,attr,omitempty"`
}

type Encoder struct {
	Type    string `xml:"type,attr,omitempty"`
	Version string `xml:"version,attr,omitempty"`
	Vendor  string `xml:"vendor,attr,omitempty"`
}

type AudioStream struct {
	Stream

	BitDepth     int     `xml:"http://mediapackage.opencastproject.org bitdepth,omitempty"`
	Channels     int     `xml:"http://mediapackage.opencastproject.org channels,omitempty"`
	SamplingRate int     `xml:"http://mediapackage.opencastproject.org samplingrate,omitempty"`
	BitRate      float32 `xml:"http://mediapackage.opencastproject.org bitrate,omitempty"`
	PkLevDB      float32 `xml:"http://mediapackage.opencastproject.org peakleveldb,omitempty"`
	RmsLevDB     float32 `xml:"http://mediapackage.opencastproject.org rmsleveldb,omitempty"`
	RmsPKDB      float32 `xml:"http://mediapackage.opencastproject.org rmspeakdb,omitempty"`
}

type VideoStream struct {
	Stream

	BitRate    float32     `xml:"http://mediapackage.opencastproject.org bitrate,omitempty"`
	FrameRate  float32     `xml:"http://mediapackage.opencastproject.org framerate,omitempty"`
	Resolution *Resolution `xml:"http://mediapackage.opencastproject.org resolution,omitempty"`
	ScanType   *Scan       `xml:"http://mediapackage.opencastproject.org scantype,omitempty"`
}

type Resolution struct {
	Width  int `xml:"http://mediapackage.opencastproject.org width,omitempty"`
	Height int `xml:"http://mediapackage.opencastproject.org hight,omitempty"`
}

var (
	_ encoding.TextMarshaler   = &Resolution{}
	_ encoding.TextUnmarshaler = &Resolution{}
)

func (res *Resolution) MarshalText() ([]byte, error) {
	if res == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("%dx%d", res.Width, res.Height)), nil
}

func (ref *Resolution) UnmarshalText(text []byte) (err error) {
	parts := strings.Split(string(text), "x")

	if len(parts) != 2 {
		return errors.New("Resolution.UnmarshalText: invalid resolution format")
	}

	ref.Width, err = strconv.Atoi(parts[0])
	if err != nil {
		return
	}

	ref.Height, err = strconv.Atoi(parts[1])
	return
}

type Scan struct {
	Type  ScanType  `xml:"type,attr,omitempty"`
	Order ScanOrder `xml:"order,attr,omitempty"`
}

type ScanType string

const (
	InterlacedScanType  ScanType = "Interlaced"
	ProgressiveScanType ScanType = "Progressive"
)

type ScanOrder string

const (
	TopFieldFirstScanOrder    ScanOrder = "TopFieldFirst"
	BottomFieldFirstScanOrder ScanOrder = "BottomFieldFirst"
)

type StreamingProtocol int

const (
	DownloadStreamingProtocol StreamingProtocol = iota
	HLSStreamingProtocol
	DASHStreamingProtocol
	HDSStreamingProtocol
	SmoothStreamingProtocol
	MMSStreamingProtocol
	RTPStreamingProtocol
	RTSPStreamingProtocol
	RTMPStreamingProtocol
	RTMPEStreamingProtocol
	PNMStreamingProtocol
	PNAStreamingProtocol
	IcyStreamingProtocol
	BitTorentLiveStreamingProtocol
	FileStreamingProtocol
	UnknownStreamingProtocol
)

type SubtitleStream struct {
	Stream
}
