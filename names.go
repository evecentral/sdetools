package sdetools

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type UniqueName struct {
	Id      int         `json:"itemID"`
	Name    interface{} `json:"itemName"`
	GroupId int         `json:"groupID"`
}

type UniqueNames []*UniqueName

func (s *SDE) loadNames() error {
	path := filepath.Join(s.BaseDir, "bsd/invUniqueNames.yaml.json")
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var uniqueNames UniqueNames
	err = json.NewDecoder(file).Decode(&uniqueNames)
	if err != nil {
		return err
	}

	s.systemNamesById = make(map[int]string)

	for _, name := range uniqueNames {
		if name.GroupId == 5 {
			s.systemNamesById[name.Id] = name.Name.(string)
		}
	}

	s.loadedNames = true

	return nil
}

func (s *SDE) GetSystemNameById(system int) (string, bool) {
	if s.loadedNames == false {
		err := s.loadNames()
		if err != nil {
			log.Fatal(err)
		}
	}
	r, ok := s.systemNamesById[system]
	return r, ok
}
