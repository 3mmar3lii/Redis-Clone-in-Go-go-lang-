package main

import "encoding/base64"

type Storer interface {
	Set  (string,string) error
	Get  (string) (string,error)
}

func SetKeyWithEncryptionStore(store Storer, key, value string) (string, error) {

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	if err := store.Set(key, encoded); err != nil {
		return "", err
	}
	return store.Get(key)
}
