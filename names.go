package sdetools

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type UniqueName struct {
	Id      int    `yaml:"itemID"`
	Name    string `yaml:"itemName"`
	GroupId int64  `yaml:"groupID"`
}

type UniqueNames []*UniqueName

func (s *SDE) loadNames() error {
	path := filepath.Join(s.BaseDir, "bsd/invUniqueNames.yaml")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var uniqueNames UniqueNames
	err = yaml.Unmarshal(data, &uniqueNames)
	if err != nil {
		return err
	}

	s.systemNamesById = make(map[int]string)

	for _, name := range uniqueNames {
		if name.GroupId == 5 {
			s.systemNamesById[name.Id] = name.Name
		}
	}

	s.loadedNames = true

	return nil
}

func (s *SDE) GetSystemNameById(system int) (string, bool) {
	if s.loadedNames == false {
		s.loadNames()
	}
	r, ok := s.systemNamesById[system]
	return r, ok
}
