package main

import (
	//"errors"
	"fmt"
	"time"
)

func main(){
	store := NewStore(2)
	store2 := NewTTLStore(time.Second *2)

	// Encrypted Store 
	encrypted,_ :=SetKeyWithEncryptionStore(store,"key4","value2")
	encrypted2,_ :=SetKeyWithEncryptionStore(store2,"key1","value2")

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