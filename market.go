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

type MarketType struct {
	GroupId       int               `yaml:"groupID"`
	MarketGroupId int               `yaml:"margetGroupID"`
	Name          map[string]string `yaml:"name"`
	Published     bool              `yaml:"published"`
	Description   map[string]string `yaml:"description"`
	Volume        float64           `yaml:"volume"`
	Mass          float64           `yaml:"mass"`
}

type MarketTypes map[int64]MarketType

const (
	groupBucket          = "groupIDs"
	marketTypeBucket     = "marketTypes"
	marketTypeNameBucket = "marketTypesName"
)

func (s *SDE) loadMarketTypes() error {
	path := filepath.Join(s.BaseDir, "fsd/typeIDs.yaml")
	var types MarketTypes
	err := LoadYamlFile(path, &types)
	if err != nil {
		log.Println(err)
		return err
	}

	for typeid, mt := range types {
		s.db.Update(func(tx *bolt.Tx) error {
			key := boltKey(int(typeid))
			bucket := tx.Bucket([]byte(marketTypeBucket))
			data, err := msgpack.Marshal(mt)
			if err != nil {
				log.Println(err)
				return err
			}

			err = bucket.Put(key, data)
			if err != nil {
				log.Println(err)
				return err
			}
			key = []byte(mt.Name["en"])
			bucket = tx.Bucket([]byte(marketTypeNameBucket))
			data, err = msgpack.Marshal(mt)
			if err != nil {
				log.Println(err)
				return err
			}

			err = bucket.Put(key, data)
			if err != nil {
				log.Println(err)
			}
			return err

			return err

		})
	}
	return nil
}

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

func (s *SDE) GetGroupById(groupid int) (group *Group, found bool) {
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

func (s* SDE) GetTypeById(typeid int) (mt *MarketType, found bool) {
	s.db.View(func(tx *bolt.Tx) error {
		key := boltKey(int(typeid))
		b := tx.Bucket([]byte(marketTypeBucket))
		v := b.Get(key)
		if v == nil {
			return nil
		}

		err := msgpack.Unmarshal(v, &mt)
		if err != nil {
			log.Println(err)
			return err
		}
		found = true
		return nil
	})
	return
}

func (s *SDE) GetTypeByExactName(name string) (mt *MarketType, found bool) {
	s.db.View(func(tx *bolt.Tx) error {
		key := []byte(name)
		b := tx.Bucket([]byte(marketTypeNameBucket))
		v := b.Get(key)
		if v == nil {
			return nil
		}

		err := msgpack.Unmarshal(v, &mt)
		if err != nil {
			log.Println(err)
			return err
		}
		found = true
		return nil
	})
	return
}
