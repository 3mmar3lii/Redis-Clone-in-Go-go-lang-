package main

import (
	"time"
)

type TTLStore struct {
	data map[string]ttlEntry
	ttl time.Duration
}

type ttlEntry struct {
	value     string
	expiresAt time.Time
}


func NewTTLStore(ttl time.Duration) *TTLStore{
	return &TTLStore{
	data:make(map[string]ttlEntry),
	}
}


func (t *TTLStore) Set(key,value string) error {
	// key is missing
	if key == ""{
		return ErrorKeyIsEmpty
	}
	
	// set the key and value
	t.data[key]=ttlEntry{
		value:value,
		expiresAt: time.Now().Add(t.ttl),
	}
	return nil
}



func (t * TTLStore) Get(key string) (string,error){

	val,ok :=t.data[key]
	// ttl expires
	if !ok || time.Now().After(t.data[key].expiresAt) {
		delete(t.data,key)
		return "",ErrKeyDoesNotExist
	}

	return val.value,nil
}