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
* @package     libupload
* @author      jeffotoni
* @copyright   Copyright (c) 2017
* @license     --
* @link        --
* @since       Version 0.1
*
*/

// Package libupload
// This package is responsible for uploading the restful,
// where will provide a handler / upload, we will manage all
// uploads of files by this handler.
//
// A simple restful server to receive common and encrypted files will
// store the file on the server.
//
// You will use port 8080 for communication
// with your api, will be written a no-sql database boltdb all uploads
// made on the server, will be generated a log on disk of all accesses.
package libupload

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/jeffotoni/goupload/pkg/gbolt"
)

// using Environment variables and keys
// Are our system settings in memory
var (
	httpConf *http.Server

	// Is for the system to be able to release
	// or to authenticate our url in a
	// simple way our upload
	AUTHORIZATION = `tyladfadiwkxceieixweiex747`
	Port          = "8081"
	Scheme        = "http"
	Database      = "ServerUpload"
	Host          = "localhost"

	// Leave blank if you run the docker for example, it will receive
	// connection from other machines
	HostHttp = ""

	UploadSize int64
	PathLocal  = "uploads/"
)

// using startUploadServer restful server upload
// This method will generate the handlerFunc
// for our api upload
func StartUploadServer() {

	// Testing boltdb database
	// Start ping database
	// Creating ping ok
	gbolt.Save("Ping", "ok")

	// Testing whether it was recorded
	// and read on the boltdb, we
	// recorded a Ping and then
	// read it back.
	if gbolt.Get("Ping") != "ok" {

		fmt.Println("Services Error Data Base!")
		os.Exit(1)
	}

	// Showing the status screen
	color.Cyan("Services successfully tested")
	color.Green("Host: " + Host)
	color.Green("Scheme:" + Scheme)
	color.Yellow("Port: " + Port)
	color.White("POST %s\n", UrlUpload())
	fmt.Println("Loaded service")

	// create route
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/", http.FileServer(http.Dir("message")))

	router.
		HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {

			if r.Method == "POST" {

				// Build the method here
				fmt.Fprintln(w, "http ", 200, "ok")
				UploadFile(w, r)

			} else if r.Method == "GET" {

				fmt.Fprintln(w, "http ", 500, "Not authorized / Allowed method POST")

			} else {

				fmt.Fprintln(w, "http ", 500, "Not authorized / Allowed method POST")
			}
		})

	// Configuration of our ListenAndServer, [
	// here we put all our configurations that
	// the server will manage
	httpConf = &http.Server{

		Handler: router,
		Addr:    HostHttp + ":" + Port,

		// Good idea!!! Good live!!!
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Fatal(httpConf.ListenAndServe())
}

// using UrlUpload Create
// the url of our api Upload
func UrlUpload() string {

	return Scheme + "://" + Host + ":" + Port + "/upload"

}
