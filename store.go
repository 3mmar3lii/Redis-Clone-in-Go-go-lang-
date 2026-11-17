package main

import (
	"errors"
	"fmt"
	"slices"
)

var ErrKeyDoesNotExist = errors.New("Key is not exist")
var ErrorKeyIsEmpty = errors.New("Key is mantodory")
var ErrStoreIsFull = errors.New("Store if at max capacity")

type Store struct {
	data map[string]string
	maxSize int
}

func (s *Store) Keys() []string{
	var keys []string 
	for key := range s.data {
		keys = append(keys,key)
	}
	//sort.Strings(keys)
	slices.Sort(keys)
	return keys
}

func (s *Store) Get(key string) (string,error){
	if key == ""{
		return "",ErrorKeyIsEmpty
	}
	val,ok := s.data[key]
	if !ok {
		return "",ErrKeyDoesNotExist
	}
	return val,nil
}


func (s *Store) Set(key,value string) error{

	if key == ""{
		return ErrorKeyIsEmpty
	}
	_,exist := s.data[key]
	
	if s.maxSize > 0 && len(s.data) >= s.maxSize && !exist{
		// throw error based on what type of error happen 
		return fmt.Errorf("Set(%q): %w",key,ErrStoreIsFull)
	}

	s.data[key]=value
	return nil

}


func (s * Store) Delete(key string) {
	if _,ok := s.data[key]; !ok {
		return
	}
	delete(s.data,key)
}

func (s *Store)Len()int{
	return len(s.data)
}

func NewStore(maxSizeKey int) *Store{
	return &Store{
		data :make(map[string]string),
		maxSize: int(maxSizeKey),
	}
}

