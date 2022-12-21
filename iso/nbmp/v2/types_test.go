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

	nbmp "github.com/nagare-media/models.go/iso/nbmp/v2"
)

func TestUnmarshalConfiguration(t *testing.T) {
	files := []string{
		"testdata/example-parameter-representation.json",
	}

	for _, file := range files {
		str, err := os.ReadFile(file)
		if err != nil {
			t.Fatalf("could not read file %s", file)
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
		"testdata/function-RTP-360sticher.json",
		"testdata/function-RTP-cgtranscoder.json",
		"testdata/function-RTP-cropping.json",
		"testdata/function-RTP-dash-packager.json",
		"testdata/function-RTP-fifo.json",
		"testdata/function-RTP-merger.json",
		"testdata/function-RTP-omaf-packager.json",
		"testdata/function-RTP-pcdecoder.json",
		"testdata/function-RTP-pcencoder.json",
		"testdata/function-RTP-selectorcompositor.json",
		"testdata/function-RTP-splitter.json",
	}

	for _, file := range files {
		str, err := os.ReadFile(file)
		if err != nil {
			t.Fatalf("could not read file %s", file)
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
