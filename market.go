package sdetools

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Group struct {
	Anchorable           bool              `json:"anchorable"`
	Anchored             bool              `json:"anchored"`
	CategoryId           int               `json:"categoryID"`
	FittableNonSingleton bool              `json:"fittableNonSingleton"`
	Name                 map[string]string `json:"name"`
	Published            bool              `json:"published"`
	UseBasePrice         bool              `json:"useBasePrice"`
}

type Groups map[int]Group

func (s *SDE) loadGroups() error {
	path := filepath.Join(s.BaseDir, "fsd/groupIDs.yaml.json")

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	err = json.NewDecoder(file).Decode(&s.groups)
	if err != nil {
		return err
	}

	s.loadedGroups = true
	return nil
}

func (s *SDE) GetGroupById(group int) (Group, bool) {
	if s.loadedGroups != true {
		err := s.loadGroups()
		if err != nil {
			log.Fatal(err)
		}
	}
	g, ok := s.groups[group]
	return g, ok
}
