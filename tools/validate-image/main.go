// Copyright © 2021 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func testPNG(path string) bool {
	// Read image
	finput, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	input, err := ioutil.ReadAll(finput)
	if err != nil {
		panic(err)
	}
	in := bytes.NewReader(input)
	img, err := png.Decode(in)
	finput.Close()
	if err != nil {
		return false
	}

	if img != nil {
		return true
	}

	return true
}

func testJPEG(path string) bool {
	// Read image
	finput, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	input, err := ioutil.ReadAll(finput)
	if err != nil {
		panic(err)
	}
	in := bytes.NewReader(input)
	img, err := jpeg.Decode(in)
	finput.Close()
	if err != nil {
		return false
	}

	if img != nil {
		return true
	}

	return true
}

func main() {
	args := os.Args[1:]
	file := args[0]

	if _, err := os.Stat(file); os.IsNotExist(err) {
		log.Fatalf("Decode Image Failed: file %v does not exist", file)
	}

	ext := strings.ToLower(filepath.Ext(file))
	if ext == ".png" || ext == ".jpg" || ext == ".jpeg" {
		success := testPNG(file)
		if !success {
			success = testJPEG(file)
		}

		if !success {
			log.Fatalf("Decode Image Failed: Unknown file format %v", file)
		}
	}
}
