package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/yosuke-furukawa/gojson5/encoding/json5"
	"io/ioutil"
	"os"
)

func main() {
	procArgs()
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: json5 \n")
	fmt.Fprintf(os.Stderr, "Compiles JSON5 file into sibling JSON.\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func procArgs() {
	flag.Usage = usage
	json5Path := flag.String("c", "", "path/to/file.json5")
	outputPath := flag.String("o", "", "path/to/file.json")
	flag.Parse()
	if *json5Path == "" {
		usage()
	}
	file, err := os.Open(*json5Path)
	if err != nil {
		fmt.Println(err)
		usage()
	}

	var data interface{}
	dec := json5.NewDecoder(file)
	err = dec.Decode(&data)
	if err != nil {
		fmt.Println(err)
		usage()
	}
	b, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println(err)
		usage()
	}
	if *outputPath == "" {
		fmt.Println(string(b))
	}
	ioutil.WriteFile(*outputPath, b, 0644)
}
