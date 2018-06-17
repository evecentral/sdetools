package sdetools

import (
	"github.com/boltdb/bolt"
	"github.com/evecentral/sdetools/regions"
	"gopkg.in/vmihailenco/msgpack.v2"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Stargate struct {
	Destination int `yaml:"destination"`
	TypeID      int `yaml:"typeID"`
}

type SolarSystem struct {
	Id     int `yaml:"solarSystemID"`
	NameId int `yaml:"solarSystemNameID"`

	Name          string
	Region        *regions.Region
	Constellation string

	Security float64 `yaml:"security"`

	Corridor      bool `yaml:"corridor"`
	Fringe        bool `yaml:"fringe"`
	Hub           bool `yaml:"hub"`
	Border        bool `yaml:"border"`
	International bool `yaml:"international"`

	Stargates map[int]Stargate `yaml:"stargates"`
}

func (s *SDE) loadUniverse() {
	log.Println("loading universe...")
	defer log.Println("done loading universe")
	loadFunc := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		dir, fname := filepath.Split(path)
		if fname != "solarsystem.staticdata" {
			return nil
		}

		dirParts := strings.Split(dir, "/")

		ssName := dirParts[len(dirParts)-2]
		constellation := dirParts[len(dirParts)-3]
		regionName := dirParts[len(dirParts)-4]

		var solarsystem SolarSystem
		err = LoadYamlFile(path, &solarsystem)
		if err != nil {
			log.Printf("regions load failure %v at path %v", err, path)
			return err
		}

		solarsystem.Name = ssName
		solarsystem.Region = regions.Get().ByName(regionName)
		solarsystem.Constellation = constellation

		s.db.Update(func(tx *bolt.Tx) error {
			key := boltKey(solarsystem.Id)
			bucket := tx.Bucket([]byte(solarSystemIdsBucket))
			data, err := msgpack.Marshal(&solarsystem)
			if err != nil {
				log.Println("encoding error: %v", err)
				return err
			}

			err = bucket.Put(key, data)
			if err != nil {
				return err
			}

			sKey := []byte(solarsystem.Name)
			sBucket := tx.Bucket([]byte(solarSystemNamesBucket))
			return sBucket.Put(sKey, data)

		})

		return nil
	}

	err := filepath.Walk(filepath.Join(s.BaseDir, "fsd/universe/eve"), loadFunc)
	if err != nil {
		log.Println("file walking error %v", err)
	}
}
