package main

import (
	"errors"
	"fmt"
)

func main(){
	store := NewStore(2)
	store.Set("key1","val1") // len:0 maxSize:1
	store.Set("key2","val1") // len:1 maxSize:1

	if errStoreFull := store.Set("key3","val2"); errStoreFull != nil{ // len:2 maxSize:1
		fmt.Println("Key is already exist")
		return 
	}

	if _,err := store.Get("key1"); err != nil {
		fmt.Println(err)
	}
	// Way 2 of error handling
	_,err := store.Get("key1");
	if errors.Is(err,ErrKeyDoesNotExist) {
		fmt.Println(err)
	}
}