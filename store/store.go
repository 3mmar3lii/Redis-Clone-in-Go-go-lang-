package store

import (
	"encoding/base64"
	"errors"
)

var ErrKeyDoesNotExist = errors.New("Key is not exist")
var ErrorKeyIsEmpty = errors.New("Key is mantodory")
var ErrStoreIsFull = errors.New("Store if at max capacity")


type Storer interface {
	Set(string,string) error
	Get(string) (string,error)
	Keys() []string
	Delete(string) error
	Len() int
}

func SetKeyWithEncryptionStore(store Storer, key, value string) (string, error) {

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	if err := store.Set(key, encoded); err != nil {
		return "", err
	}
	return store.Get(key)
}

