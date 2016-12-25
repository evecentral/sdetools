package convert

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

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

func ConvertDatafile(path string, newPath string) error {

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
	return nil
}
