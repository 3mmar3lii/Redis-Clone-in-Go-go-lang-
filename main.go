package main

import (
	//"errors"
	"fmt"
	"time"

	"Redis-Clone/store/kv"
	"Redis-Clone/store/ttl"
	"Redis-Clone/store"
)




func main(){
	kvStore := kv.NewStore(2)
	ttlStore := ttl.NewTTLStore(time.Second *2)

	// Encrypted Store 
	encrypted,_ := store.SetKeyWithEncryptionStore(kvStore,"key4","value2")
	encrypted2,_ := store.SetKeyWithEncryptionStore(ttlStore,"key1","value2")

	fmt.Println(encrypted,encrypted2)


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