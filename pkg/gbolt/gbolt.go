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

package gbolt

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

var (
	Database = []byte("DBGoupload")
	DirDb    = "db"
	PathDb   = "db/gbolt.db"
)

type DB struct {
	*bolt.DB
}

var (
	dbbolt *bolt.DB
	err    error
)

//var djson map[string]interface{}

type JsonDataDb struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	Size    int64  `json:"size"`
	Path    string `json:"path"`
	Created string `json:"key"`
}

var djson JsonDataDb

/** Connect bolt db */

func Connect() *DB {

	// Can not leave singleton the bank has to close every call,
	// save, update, get etc ..

	DataBaseTest(PathDb)

	dbbolt, err := bolt.Open(PathDb, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal("connect error: ", err)
	}

	return &DB{dbbolt}

}

func DataBaseTest(PathDb string) {

	if !ExistDb(DirDb) {

		os.MkdirAll(DirDb, 0755)
	}

	// detect if file exists

	if !ExistDb(PathDb) {

		var file, err = os.Create(PathDb)
		checkError(err)
		defer file.Close()

		w, errx := os.OpenFile(PathDb, os.O_WRONLY|os.O_CREATE, 0644)
		checkError(errx)
		defer w.Close()
	}
}

func ExistDb(name string) bool {

	if _, err := os.Stat(name); err != nil {

		if os.IsNotExist(err) {

			return false
		}
	}

	return true
}

func SaveDb(keyfile string, namefile string, sizefile int64, pathFile string) error {

	times := fmt.Sprintf("%s", time.Now())

	stringJson := JsonDataDb{keyfile, namefile, sizefile, pathFile, times}

	respJson, err := json.Marshal(stringJson)

	respJsonX := string(respJson)

	err = Save(keyfile, respJsonX)

	if err == nil {

		//fmt.Println("save sucess..")
		return nil

	} else {

		//fmt.Println("Error", err)
		return err
	}

}

func JsonGet(key []byte) string {

	db := Connect()

	defer db.Close()

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

	fmt.Println("here: ", djson)

	if errjs != nil {

		log.Fatal(fmt.Println("Error Json: ", errjs))
	}

	return string(valbyte)
}

func Save(keyS string, valueS string) error {

	db := Connect()

	defer db.Close()

	key := []byte(keyS)
	value := []byte(valueS)

	err := db.Update(func(tx *bolt.Tx) error {

		bucket, err := tx.CreateBucketIfNotExists(Database)

		if err != nil {

			return err
		}

		err = bucket.Put(key, value)

		if err != nil {

			return err

		} else {

			//fmt.Println("save sucess")
			return nil
		}
	})

	if err != nil {

		fmt.Println("erro try save ", err)
		os.Exit(1)
	}

	return nil
}

func Get(keyS string) string {

	db := Connect()

	defer db.Close()

	key := []byte(keyS)

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

		log.Fatal("Error open db, ", err)
	}

	return string(valbyte)
}

func checkError(err error) {

	if err != nil {
		fmt.Println("Error Database: ", err.Error())
		os.Exit(0)
	}
}
