package main

import (
	//"errors"
	"fmt"
	//"time"

	"Redis-Clone/store"
	"Redis-Clone/store/kv"
	//"Redis-Clone/store/ttl"
)

func CreateStore() store.Storer {
	plain := kv.NewStore(2)
	logger := NewLoggingMiddleware(plain)
	return logger
}

func main() {
	//cmds := []Command{
	//	{Op: "SET", Key: "env", Value: "production"},
	//	{Op: "SET", Key: "version", Value: "0.0.1"},
	//	{Op: "SET", Key: "debug", Value: "true"},
	//	{Op: "GET", Key: "env"},
	//	{Op: "SET", Key: "region", Value: "eu-west-1"},
	//	{Op: "GET", Key: "version"},
	//}

	cmdStore := CreateStore()
	
	cmdStore.Set("pageviews","0")
	const hits = 1000
	cmds := make([]Command,hits)
	
	for i := range cmds {
		cmds[i]=Command{Op: "INCR",Key: "pageviews"}
	}

	RestoreOnBoot(cmdStore,cmds)

	
	got,_ := cmdStore.Get("pageviews")
	fmt.Println("expected: %d\n", hits)
	fmt.Println("got:			 %s\n", got)

	//ttlStore := ttl.NewTTLStore(time.Second * 2)
	//logger := CreateStore()
	//metrics := NewMatrixMiddleware(logger)

	//// Encrypted Store
	//encrypted, _ := store.SetKeyWithEncryptionStore(metrics, "key4", "value2")
	//encrypted2, _ := store.SetKeyWithEncryptionStore(ttlStore, "key1", "value2")

	//if _, err := metrics.Get("fff"); err != nil {
	//	fmt.Println(err)
	//}
	//metrics.Report()
	//fmt.Println(encrypted, encrypted2)

}

func PopulateDefaults(s store.Storer) error {

	defaults := map[string]string{
		"env":     "development",
		"version": "0.0.1",
		"debug":   "true",
	}

	for k, v := range defaults {
		if err := s.Set(k, v); err != nil {
			return fmt.Errorf("PopulateDefaults: %w", err)
		}
	}

	return nil
}
