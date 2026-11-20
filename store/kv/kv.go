package kv

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"sync"
	"time"
)

var ErrKeyDoesNotExist = errors.New("Key is not exist")
var ErrorKeyIsEmpty = errors.New("Key is mantodory")
var ErrStoreIsFull = errors.New("Store if at max capacity")

type Store struct {
	data    map[string]string
	maxSize int
	mu      sync.RWMutex
}

func (s *Store) Keys() []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	var keys []string
	for key := range s.data {
		keys = append(keys, key)
	}
	//sort.Strings(keys)
	slices.Sort(keys)
	return keys
}

func (s *Store) Get(key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if key == "" {
		return "", ErrorKeyIsEmpty
	}
	val, ok := s.data[key]
	if !ok {
		return "", ErrKeyDoesNotExist
	}
	return val, nil
}

func (s *Store) Set(key, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if key == "" {
		return ErrorKeyIsEmpty
	}
	_, exist := s.data[key]

	time.Sleep(time.Second)

	if s.maxSize > 0 && len(s.data) >= s.maxSize && !exist {
		// throw error based on what type of error happen
		return fmt.Errorf("Set(%q): %w", key, ErrStoreIsFull)
	}

	s.data[key] = value
	return nil

}

func (s *Store) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[key]; !ok {
		return ErrKeyDoesNotExist
	}
	delete(s.data, key)
	return nil
}

func (s *Store) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.data)
}

func NewStore(maxSizeKey int) *Store {
	return &Store{
		data:    make(map[string]string),
		maxSize: int(maxSizeKey),
	}
}

func (s *Store) Clone() *Store {
	newDate := make(map[string]string, len(s.data))
	for key, val := range s.data {
		newDate[key] = val
	}
	return &Store{
		data:    newDate,
		maxSize: s.maxSize,
	}
}

func (s *Store) INCR(key string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// READ
	current := 0
	if raw, ok := s.data[key]; ok {
		val, err := strconv.Atoi(raw) // int
		if err != nil {
			return 0, fmt.Errorf("INCR %q: %w", key, err)
		}
		current = val
	}

	// MODIFY
	current++

	// WRITE
	value := strconv.Itoa(current) // string
	s.data[key] = value

	return current, nil
}
