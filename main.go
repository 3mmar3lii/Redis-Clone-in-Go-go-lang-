package main

import (
	//"errors"
	"fmt"
	"time"

	"Redis-Clone/store"
	"Redis-Clone/store/kv"
	"Redis-Clone/store/ttl"
)

func CreateStore() store.Storer {
	plain := kv.NewStore(2)
	logger := NewLoggingMiddleware(plain)
	return logger
}

func main() {
	ttlStore := ttl.NewTTLStore(time.Second * 2)
	logger := CreateStore()
	metrics := NewMatrixMiddleware(logger)

	// Encrypted Store
	encrypted, _ := store.SetKeyWithEncryptionStore(metrics, "key4", "value2")
	encrypted2, _ := store.SetKeyWithEncryptionStore(ttlStore, "key1", "value2")

	if _, err := metrics.Get("fff"); err != nil {
		fmt.Println(err)
	}
	metrics.Report()
	fmt.Println(encrypted, encrypted2)

	//store.Set("key1","val1") // len:0 maxSize:1
	//store.Set("key2","val1") // len:1 maxSize:1

	//if errStoreFull := store.Set("key3","val2"); errStoreFull != nil{ // len:2 maxSize:1
	//	fmt.Println("Key is already exist")
	//	return
	//}

	//if _,err := store.Get("key1"); err != nil {
	//	fmt.Println(err)
	//}
	//// Way 2 of error handling
	//_,err := store.Get("key1");
	//if errors.Is(err,ErrKeyDoesNotExist) {
	//	fmt.Println(err)
	//}
}
