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
	Database = []byte("UploadServerDb")
	DirDb    = "db"
	PathDb   = "db/bolt.db"
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

func Connect() *DB {

	if dbbolt != nil {

		return &DB{dbbolt}

	} else {

		// If the directory and file db does not exist create

		if !ExistDb(PathDb) {

			///CREATED

			os.MkdirAll(DirDb, 0755)
			CreateFileDb(PathDb)
		}

		dbbolt, err = bolt.Open(PathDb, 0644, nil)

		if err != nil {

			log.Fatal(err)
		}

		return &DB{dbbolt}

		// } else {

		// 	return &DB{}
		// }
	}
}

func CreateFileDb(path string) {

	// detect if file exists

	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {

		var file, err = os.Create(path)

		checkError(err)

		defer file.Close()
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

	//fmt.Println("json erro :> ", err)

	keyx := []byte(keyfile)
	valuex := []byte(respJson)

	// store some data

	err = Save(keyx, valuex)

	if err == nil {

		//fmt.Println("save sucess..")
		return nil

	} else {

		//fmt.Println("Error", err)
		return err
	}

}

func Save(key []byte, value []byte) error {

	db := Connect()

	defer db.Close()

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

func Get(key []byte) string {

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

	return string(valbyte)
}

func checkError(err error) {

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
