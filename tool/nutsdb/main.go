package main

import (
	"fmt"
	"log"
)
import "github.com/nutsdb/nutsdb"

func main() {
	opt := nutsdb.DefaultOptions
	opt.Dir = "./test.db"
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *nutsdb.Tx) error {
		bucket := "bucketForList"
		key := []byte("myList")
		val := []byte("val2")
		return tx.LPush(bucket, key, val)
	})
	if err != nil {
		log.Fatal(err)
	}
	err = db.View(func(tx *nutsdb.Tx) error {
		bucket := "bucketForList"
		key := []byte("myList")
		v, err := tx.LPeek(bucket, key)
		if err == nil {
			fmt.Println(string(v))
		}
		return err
	})
}
