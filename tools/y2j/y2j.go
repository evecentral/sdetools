package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/evecentral/sdetools/convert"
)

var sdepath = flag.String("sde", "sde", "Path to the SDE root")

func visit(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		return nil
	}
	newPath := fmt.Sprintf("%s.%s", path, "json")

	if !strings.HasSuffix(path, ".yaml") && !strings.HasSuffix(path, ".staticdata") {
		return nil
	}

	fmt.Printf("%s ->  %s", path, newPath)
	convert.ConvertDatafile(path, newPath)
	fmt.Printf(" [done]\n")

}

// y2j bulk converts all the YAML to JSON, writing sidecar files
// which are named with .json at the end for .yaml and .staticdata files
//
// Since the go-yaml library is a huge memory hog, this can save
// significant memory when pre-loading the SDE into memory
// for caching and lookup purposes
func main() {
	flag.Parse()

	err := filepath.Walk(*sdepath, visit)
	if err != nil {
		fmt.Println(err)
		return
	}
}
