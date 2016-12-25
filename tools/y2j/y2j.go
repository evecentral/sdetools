package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var sdepath = flag.String("sde", "sde", "Path to the SDE root")

func transformData(pIn *interface{}) (err error) {
	switch in := (*pIn).(type) {
	case map[interface{}]interface{}:
		m := make(map[string]interface{}, len(in))
		for k, v := range in {
			if err = transformData(&v); err != nil {
				return err

			}
			var sk string
			switch k.(type) {
			case string:
				sk = k.(string)
			case int:
				sk = strconv.Itoa(k.(int))
			case bool:
				if k.(bool) == true {
					sk = "true"
				} else {
					sk = "false"
				}
			default:
				return fmt.Errorf("type mismatch: expect map key string or int; got: %T", k)

			}
			m[sk] = v

		}
		*pIn = m
	case []interface{}:
		for i := len(in) - 1; i >= 0; i-- {
			if err = transformData(&in[i]); err != nil {
				return err

			}

		}

	}
	return nil

}

func visit(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		return nil
	}

	if !strings.HasSuffix(path, ".yaml") && !strings.HasSuffix(path, ".staticdata") {
		return nil
	}

	newPath := fmt.Sprintf("%s.%s", path, "json")
	fmt.Printf("%s ->  %s", path, newPath)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var yd interface{}
	err = yaml.Unmarshal(data, &yd)
	if err != nil {
		return err
	}

	err = transformData(&yd)
	if err != nil {
		return err
	}

	file, err := os.Create(newPath)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "\t")
	err = enc.Encode(&yd)
	if err != nil {
		return err
	}

	fmt.Printf(" [done]\n")
	return nil
}

func main() {
	flag.Parse()

	err := filepath.Walk(*sdepath, visit)
	if err != nil {
		fmt.Println(err)
		return
	}
}
