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
	"fmt"
	"strconv"
	"testing"
)

var tests = []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

var testString = []string{"3737xxxx", "e93393xx", "37373783x", "ue83838x", "ASDFAX&993", "875ASDFAX", "ASDF735493", "ASDFAX&556", "ASDF4466993", "IEUEUX383"}

func TestSave(t *testing.T) {

	fmt.Println("")
	fmt.Println("Test list all keys boltdb")

	for i, v := range testString {

		fmt.Println("value: ", v)

		//x := "key-" + strconv.FormatInt(i, 10) int64
		//x := "key-" + strconv.ParseInt(i, 10, 16)
		stringx := strconv.FormatInt(int64(i), 10)

		x := "key-" + stringx

		if Save(x, v) == nil {

			fmt.Println("Save sucess: ", x, v)

		} else {

			t.Fatalf("at index %d, expected %d, go val %s", i, v, v)
		}
	}

	ListAllKeys()
}

func TestGet(t *testing.T) {

	///TEST KEY
	//key := []byte("tyladfadiwkxceieixweiex747/file2.pdf")
	//key := []byte("Ping")

	//Save("Ping", "ok")

	fmt.Println("Get: ", Get("key-6"))

	//json := gbolt.JsonGet(key)
	//fmt.Println(jsonx)
}
