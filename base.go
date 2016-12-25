package sdetools

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/boltdb/bolt"
)

type SDE struct {
	BaseDir string

	db *bolt.DB

	loadedGroups bool
	groups       Groups
}

func (s *SDE) Init() {
	db, err := bolt.Open(filepath.Join(s.BaseDir, "boltdb.bolt"), 0644, nil)
	db.NoSync = true
	if err != nil {
		log.Fatal(err)
	}
	s.db = db
}

// Pre-build a BoltDB database of all the static data allowing
// fast local node access for all data without requiring it all
// in memory at all times, useful for running on very small
// compute instances
func (s *SDE) BuildBoltDB() {
	s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(nameBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil

	})
	s.loadNames()
	s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(groupBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil

	})
	s.loadGroups()

}
