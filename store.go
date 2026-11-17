package main

import (
	"errors"
	"slices"
)

var ErrKeyDoesNotExist = errors.New("Key is not exist")

type Store struct {
	data map[string]string
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
	val,ok := s.data[key]
	if !ok {
		return "",ErrKeyDoesNotExist
	}
	return val,nil
}


func (s *Store) Set(key,value string){
	s.data[key]=value
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

func NewStore() *Store{
	return &Store{
		data :make(map[string]string),
	}
}

