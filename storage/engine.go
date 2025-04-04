package storage

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

// Database structure
type Database struct {
	Name  string
	Data  map[string]string // Key-Value storage (for simplicity)
	Mutex sync.RWMutex
}

// Create a new database
func CreateTable(name string) (*Database, error) {
	db := &Database{
		Name: name,
		Data: make(map[string]string),
	}
	return db, nil
}

// Insert data into the database
func (db *Database) Insert(key, value string) error {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	// Prevent duplicate entries
	if _, exists := db.Data[key]; exists {
		return errors.New("record already exists")
	}
	db.Data[key] = value
	return nil
}

// Retrieve data
func (db *Database) Retrieve(key string) (string, bool) {
	db.Mutex.RLock()
	defer db.Mutex.RUnlock()

	value, found := db.Data[key]
	return value, found
}

// Save data to a file (JSON format)
func (db *Database) SaveToFile() error {
	db.Mutex.RLock()
	defer db.Mutex.RUnlock()

	file, err := os.Create(db.Name + ".json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(db.Data)
}

// Load data from a file
func (db *Database) LoadFromFile() error {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	file, err := os.Open(db.Name + ".json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&db.Data)
}
