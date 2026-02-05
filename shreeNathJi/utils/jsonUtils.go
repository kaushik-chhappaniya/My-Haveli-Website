package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	logger "github.com/kaushik-chhappaniya/myHaweli/middleware/logger"
)

type Store struct {
	FilePath string
	mu       sync.Mutex
}

// Value can be anything JSON-serializable
type Record map[string]any
type Data map[string]any

// Ensure the JSON file exists; if not, create it with an empty JSON object
func (s *Store) ensureFile() error {
	dir := filepath.Dir(s.FilePath)
	logger.Debugf("Filepath: %v, and after dir: %v", s.FilePath, dir)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	if _, err := os.Stat(s.FilePath); errors.Is(err, os.ErrNotExist) {
		return os.WriteFile(s.FilePath, []byte(`{}`), 0644)
	}
	return nil
}

// ReadAll reads all records from the JSON file
func (s *Store) ReadAll() (Data, error) {
	logger.Debugf("Read All function invoked with path: %s", s.FilePath)
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.ensureFile(); err != nil {
		return nil, err
	}
	logger.Debugf("File existence ensured for path: %s", s.FilePath)
	bytes, err := os.ReadFile(s.FilePath)
	if err != nil {
		logger.Errorf("Error reading file: %v", err)
		return nil, err
	}

	var data Data
	if err := json.Unmarshal(bytes, &data); err != nil {
		logger.Errorf("Error unmarshaling JSON: %v", err)
		return nil, err
	}

	logger.Debugf("Data read from file: %v", data)
	return data, nil
}

// WriteAll writes all records to the JSON file (exported version)
func (s *Store) WriteAll(data Data) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.ensureFile(); err != nil {
		return err
	}

	return s.writeAll(data)
}

// WriteAll writes all records to the JSON file
func (s *Store) writeAll(data Data) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.FilePath, bytes, 0644)
}

// Find key in a case-insensitive manner
func findKeyCI(data Data, key string) (string, bool) {
	key = strings.ToLower(key)
	for k := range data {
		if strings.ToLower(k) == key {
			return k, true
		}
	}
	return "", false
}

// Upsert adds or updates a record by key in a case-insensitive manner
func (s *Store) Upsert(key string, value Record) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.ensureFile(); err != nil {
		return err
	}

	bytes, err := os.ReadFile(s.FilePath)
	if err != nil {
		return err
	}

	var data Data
	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}

	if existingKey, found := findKeyCI(data, key); found {
		data[existingKey] = value
	} else {
		data[key] = value
	}

	return s.writeAll(data)
}

// Get retrieves a record by key in a case-insensitive manner
func (s *Store) Get(key string) (Record, bool, error) {
	data, err := s.ReadAll()
	if err != nil {
		return nil, false, err
	}

	if existingKey, found := findKeyCI(data, key); found {
		if record, ok := data[existingKey].(map[string]any); ok {
			return record, true, nil
		}
		return nil, false, fmt.Errorf("value at key %s is not a valid Record", key)
	}

	return nil, false, nil
}

// Delete removes a record by key in a case-insensitive manner
func (s *Store) Delete(key string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	bytes, err := os.ReadFile(s.FilePath)
	if err != nil {
		return false, err
	}

	var data Data
	if err := json.Unmarshal(bytes, &data); err != nil {
		return false, err
	}

	if existingKey, found := findKeyCI(data, key); found {
		delete(data, existingKey)
		return true, s.writeAll(data)
	}

	return false, nil
}
