package sdetools

import (
	"encoding/gob"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/boltdb/bolt"
	"github.com/evecentral/sdetools/convert"
)

const nameBucket = "invUniqueNames"

type UniqueName struct {
	Id      int         `json:"itemID"`
	Name    interface{} `json:"itemName"`
	GroupId int         `json:"groupID"`
}

func init() {
	var u UniqueName
	gob.Register(&u)
}

type UniqueNames []*UniqueName

// Load unique names into the BoltDB cluster from the JSON file
// Generates the JSON object if required
func (s *SDE) loadNames() error {

	path := filepath.Join(s.BaseDir, "bsd/invUniqueNames.yaml.json")

	if _, err := os.Stat(path); err != nil {
		oldPath := filepath.Join(s.BaseDir, "bsd/invUniqueNames.yaml")
		convert.ConvertDatafile(oldPath, path)
	}

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

	for _, name := range uniqueNames {
		s.db.Update(func(tx *bolt.Tx) error {
			key := boltKey(name.Id)
			b := tx.Bucket([]byte(nameBucket))
			err := b.Put(key, gobToBytes(name))
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
	return "", false
}
