package main

import (
	"errors"
	"fmt"
)

func main(){
	store := NewStore()
	store.Set("key1","val1")
	
	// Way 1 of error handling
	if _,err := store.Get("key1"); err != nil {
		fmt.Println(err)
	}
	// Way 2 of error handling
	_,err := store.Get("key1"); 
	if errors.Is(err,ErrKeyDoesNotExist) {
		fmt.Println(err)
	}
}