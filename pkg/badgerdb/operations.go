package badgerdb

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/sirupsen/logrus"
)

// WriteData encapsulates the logic for writing data to the database,
// including access control checks from access_control.go.
// It now requires the data (key + value) and signature for verification.
func WriteData(db *badger.DB, key []byte, value []byte, h host.Host) error {

	if !isAuthorized(h.ID().String()) {
		logrus.WithFields(logrus.Fields{
			"nodeID": h.ID().String(),
		}).Error("node is not authorized to write to the datastore")
		return fmt.Errorf("unauthorized node is not authorized to access the datastore")
	}

	err := db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to write to the database")
		return err
	}

	return nil
}

// ReadData reads the value for the given key from the database.
// It requires the host for access control verification before reading.
func ReadData(db *badger.DB, key string, h host.Host) []byte {

	// @TODO CAN ANY NODE READ?
	if isAuthorized(h.ID().String()) {
		logrus.WithFields(logrus.Fields{
			"nodeID": h.ID().String(),
		}).Error("authorized node is not authorized to read the datastore")
		return []byte("unauthorized node is not authorized to access the datastore")
	}

	var result []byte
	_ = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		_ = item.Value(func(val []byte) error {
			result = val
			return nil
		})
		return nil
	})
	return result
}