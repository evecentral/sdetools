package sdetools

import (
	"log"
	"path/filepath"

	"github.com/boltdb/bolt"
	"gopkg.in/vmihailenco/msgpack.v2"
)

const nameBucket = "invUniqueNames"

type UniqueName struct {
	Id      int         `yaml:"itemID"`
	Name    interface{} `yaml:"itemName"`
	GroupId int         `yaml:"groupID"`
}

type UniqueNames []*UniqueName

// Load unique names into the BoltDB cluster from the JSON file
// Generates the JSON object if required
func (s *SDE) loadNames() error {

	log.Println("loading names...")
	defer log.Println("end loading names")

	path := filepath.Join(s.BaseDir, "bsd/invUniqueNames.yaml")
	var uniqueNames UniqueNames
	err := LoadYamlFile(path, &uniqueNames)

	if err != nil {
		return err
	}

	log.Println("yaml loaded...")

	for _, name := range uniqueNames {
		s.db.Update(func(tx *bolt.Tx) error {
			key := boltKey(name.Id)
			b := tx.Bucket([]byte(nameBucket))
			data, err := msgpack.Marshal(name)
			if err != nil {
				log.Println(err)
				return err
			}

			err = b.Put(key, data)
			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		})
	}

	return nil
}

func (s *SDE) GetSystemNameById(system int) (string, bool) {
	var value string
	var found bool

	s.db.View(func(tx *bolt.Tx) error {
		key := boltKey(system)
		b := tx.Bucket([]byte(nameBucket))
		v := b.Get(key)
		if v == nil {
			return nil
		}
		var uniqueName UniqueName
		msgpack.Unmarshal(v, &uniqueName)
		value = uniqueName.Name.(string)
		found = true
		return nil
	})
	return value, found
}
