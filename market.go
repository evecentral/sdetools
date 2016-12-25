package sdetools

import (
	"log"
	"path/filepath"

	"github.com/boltdb/bolt"
	"gopkg.in/vmihailenco/msgpack.v2"
)

type Group struct {
	Anchorable           bool              `yaml:"anchorable"`
	Anchored             bool              `yaml:"anchored"`
	CategoryId           int               `yaml:"categoryID"`
	FittableNonSingleton bool              `yaml:"fittableNonSingleton"`
	Name                 map[string]string `yaml:"name"`
	Published            bool              `yaml:"published"`
	UseBasePrice         bool              `yaml:"useBasePrice"`
}

type Groups map[int64]Group

const (
	groupBucket = "groupIDs"
)

func (s *SDE) loadGroups() error {
	path := filepath.Join(s.BaseDir, "fsd/groupIDs.yaml")
	var groups Groups
	err := LoadYamlFile(path, &groups)

	if err != nil {
		log.Println(err)
		return err
	}

	for groupId, group := range groups {
		s.db.Update(func(tx *bolt.Tx) error {
			key := boltKey(int(groupId))
			bucket := tx.Bucket([]byte(groupBucket))
			data, err := msgpack.Marshal(group)
			if err != nil {
				log.Println(err)
				return err
			}

			err = bucket.Put(key, data)
			if err != nil {
				log.Println(err)
			}
			return err
		})
	}
	return nil
}

func (s *SDE) GetGroupById(groupid int) (group *Group,found bool) {
	s.db.View(func(tx *bolt.Tx) error {
		key := boltKey(int(groupid))
		b := tx.Bucket([]byte(groupBucket))
		v := b.Get(key)
		if v == nil {
			return nil
		}
		found = true
		err := msgpack.Unmarshal(v, &group)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	})
	return
}
