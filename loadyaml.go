package sdetools

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func LoadYamlFile(path string, to interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, to)
	return err
}
