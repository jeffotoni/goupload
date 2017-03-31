/***********

 ▄▄▄██▀▀▀▓█████   █████▒ █████▒▒█████  ▄▄▄█████▓ ▒█████   ███▄    █  ██▓
   ▒██   ▓█   ▀ ▓██   ▒▓██   ▒▒██▒  ██▒▓  ██▒ ▓▒▒██▒  ██▒ ██ ▀█   █ ▓██▒
   ░██   ▒███   ▒████ ░▒████ ░▒██░  ██▒▒ ▓██░ ▒░▒██░  ██▒▓██  ▀█ ██▒▒██▒
▓██▄██▓  ▒▓█  ▄ ░▓█▒  ░░▓█▒  ░▒██   ██░░ ▓██▓ ░ ▒██   ██░▓██▒  ▐▌██▒░██░
 ▓███▒   ░▒████▒░▒█░   ░▒█░   ░ ████▓▒░  ▒██▒ ░ ░ ████▓▒░▒██░   ▓██░░██░
 ▒▓▒▒░   ░░ ▒░ ░ ▒ ░    ▒ ░   ░ ▒░▒░▒░   ▒ ░░   ░ ▒░▒░▒░ ░ ▒░   ▒ ▒ ░▓
 ▒ ░▒░    ░ ░  ░ ░      ░       ░ ▒ ▒░     ░      ░ ▒ ▒░ ░ ░░   ░ ▒░ ▒ ░
 ░ ░ ░      ░    ░ ░    ░ ░   ░ ░ ░ ▒    ░      ░ ░ ░ ▒     ░   ░ ░  ▒ ░
 ░   ░      ░  ░                  ░ ░               ░ ░           ░  ░

*
*
* project server the Upload
*
* @package     main
* @author      jeffotoni
* @copyright   Copyright (c) 2017
* @license     --
* @link        --
* @since       Version 0.1
*/

package gobolt

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

var (
	Database = []byte("UploadServerDb")
	PathDb   = "db/bolt.db"
)

type DB struct {
	*bolt.DB
}

var (
	dbbolt *bolt.DB
	err    error
)

var djson map[string]interface{}

func Connect() *DB {

	if dbbolt != nil {

		return &DB{dbbolt}

	} else {

		dbbolt, err = bolt.Open(PathDb, 0644, nil)

		if err != nil {

			log.Fatal(err)
		}

		return &DB{dbbolt}
	}
}

func Save(key []byte, value []byte) {

	db := Connect()

	err := db.Update(func(tx *bolt.Tx) error {

		bucket, err := tx.CreateBucketIfNotExists(Database)

		if err != nil {

			return err
		}

		err = bucket.Put(key, value)

		if err != nil {

			return err

		} else {

			fmt.Println("save sucess")
			return nil
		}
	})

	if err != nil {

		fmt.Println("erro try save ", err)
	}
}

func JsonGet(key []byte) string {

	db := Connect()
	var valbyte []byte

	err = db.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket(Database)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", Database)
		}

		valbyte = bucket.Get(key)

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	byt := []byte(string(valbyte))

	///interface

	errjs := json.Unmarshal(byt, &djson)

	if errjs != nil {

		log.Fatal(fmt.Println("Error Json: ", errjs))
	}

	return string(valbyte)
}

func Get(key []byte) string {

	db := Connect()
	var valbyte []byte

	err = db.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket(Database)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", Database)
		}

		valbyte = bucket.Get(key)

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return string(valbyte)
}
