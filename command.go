package main

import (
	"Redis-Clone/store"
	"fmt"
	"log"
	"sync"
)

type Command struct {
	Op    string
	Key   string
	Value string
}

func dispatch(s store.Storer, c Command) error {
	switch c.Op {
	case "SET":
		log.Printf("SET key %s", c.Key)
		return s.Set(c.Key, c.Value)
	case "DELETE":
		return s.Delete(c.Key)
	case "GET", "KEYS":
		return nil
	default:
		return fmt.Errorf("unknown operation: %s", c.Op)
	}

}

func RestoreOnBoot(s store.Storer, cmds []Command) error {
	var wg sync.WaitGroup
	for _, cmd := range cmds {
		//wg.Add(1)
		//go dispatch(s, cmd,&wg)
		wg.Go(func (){
			dispatch(s,cmd) // Register + wg.Done() + wg.Add()
		})
	}
	wg.Wait()
	return nil
}
