package main

import (
	"fmt"

	"github.com/jeffotoni/goupload/pkg/gbolt"
)

/** [main list database] */
func main() {

	///TEST KEY

	//key := []byte("tyladfadiwkxceieixweiex747/file2.pdf")
	//key := []byte("Ping")

	gbolt.Save("Ping", "ok")

	fmt.Println("Get: ", gbolt.Get("Ping"))
	//fmt.Println("ok + >", gbolt.Get("Ping"))

	//json := gbolt.JsonGet(key)
	//fmt.Println(jsonx)
}
