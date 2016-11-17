package sdetools

import (
	"gopkg.in/yaml.v2"
)

type GroupId struct {
	Anchorable bool `yaml:"anchorable"`
	Anchored bool `yaml:"anchored"`
	CategoryId int `yaml:"categoryID"`
	FittableNonSingleton bool `yaml:"fittableNonSingleton"`
	Name map[string]string `yaml:"name"`
	Published bool `yaml:"published"`
	UseBasePrice bool `yaml:"useBasePrice"`
}

type Groups map[int]GroupId

func LoadGroups(data []byte) (*Groups, error) {
	var group Groups
	err := yaml.Unmarshal(data, &group)
	if err != nil {
		return nil, err
	}
	return &group, nil
}
