package main

import (
	"fmt"
)

func main(){
	store := NewStore()
	store.Set("key1","val1")
	
	val,_ := store.Get("key1")
	fmt.Println(val)
}