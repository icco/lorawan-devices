package main

import (
  "bytes"
  "image/png"
  "image/jpeg"
  "io/ioutil"
  "os"
  "path/filepath"
  "strings"
	"log"
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
