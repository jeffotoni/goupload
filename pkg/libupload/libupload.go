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
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jeffotoni/goupload/pkg/gbolt"
	"github.com/jeffotoni/goupload/pkg/glogs"
)

// using Environment variables and keys
// Are our system settings in memory
var (
	httpConf      *http.Server
	AUTHORIZATION = `tyladfadiwkxceieixweiex747`
	socketfileTmp = `server.red`
	socketfile    = `server.lock`
	Port          = "8080"
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
	fmt.Println("Services successfully tested")
	fmt.Println("Host: " + Host)
	fmt.Println("Scheme:" + Scheme)
	fmt.Println("Port: " + Port)
	fmt.Println("Instance POST ", UrlUpload())
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

// using UploadFile method implemented
// This method will copy the file that is
// coming in by the http.Request handler
// and copy to our physical disk file
func UploadFile(w http.ResponseWriter, r *http.Request) {

	// Setting the maximum size
	// of the mega byte upload
	UploadSize = 500

	// Capturing Authorization of Header
	Autorization := r.Header.Get("Authorization")

	if Autorization == "" {

		fmt.Fprintln(w, "", 500, "Not Authorized")

	} else {

		// Checking Authorization
		// if it is enabled for access
		if Autorization == AUTHORIZATION {

			// Valid user
			acessekey := Autorization

			// Converting byte file size to mega bytes
			sizeMaxUpload := r.ContentLength / 1048576 ///Mb

			// If the file size is larger than allowed,
			// do not allow uploading and send
			// a message to the client
			if sizeMaxUpload > UploadSize {

				fmt.Println("The maximum upload size: ", UploadSize, "Mb is large: ", sizeMaxUpload, "Mb")
				fmt.Fprintln(w, "", 500, "Unsupported file size max: ", UploadSize, "Mb")

			} else {

				// field upload
				file, handler, _ := r.FormFile("nameupload")
				defer file.Close()

				///create dir to key
				pathUpKeyUser := PathLocal + acessekey

				//IF NO EXIST
				if !ExistDir(pathUpKeyUser) {

					os.MkdirAll(pathUpKeyUser, 0777)
				}

				// Mounting the physical path where the
				// file will be uploaded, the folder has
				// nly one level, its access
				// code Autorization + name file
				pathUserAcess := PathLocal + acessekey + "/" + handler.Filename

				// copy file and write
				f, _ := os.OpenFile(pathUserAcess, os.O_WRONLY|os.O_CREATE, 0777)
				defer f.Close()

				bytes, _ := io.Copy(f, file)
				keyfile := acessekey + "/" + handler.Filename

				// Saving the result of the upload
				// in the no-sql database
				SaveDb(keyfile, handler.Filename, bytes, pathUserAcess)

				// Generates a log of everything that
				// happens on the server
				flag.Parse()
				glogs.LogNew(*glogs.PathLog)

				glogs.Log.Printf("......................start upload .........................")
				glogs.Log.Printf("Authorization: %s\n", AUTHORIZATION)
				glogs.Log.Printf("Path dir: %s\n", pathUpKeyUser)
				glogs.Log.Printf("Path file: %s\n", pathUserAcess)
				glogs.Log.Printf("File: %s\n", handler.Filename)
				glogs.Log.Println("Size: ", sizeMaxUpload)
				glogs.Log.Println("Allowed: ", UploadSize, "Mb")
				glogs.Log.Printf("Copied: %v bytes\n", bytes)
				glogs.Log.Printf("Copied: %v Kb\n", bytes/1024)
				glogs.Log.Printf("Copied: %v Mb\n", bytes/1048576)
				glogs.Log.Printf("Database key: %s\n", keyfile)

				glogs.Log.Printf("...........................................................")
				glogs.Log.Printf(" ")

				time.Sleep(1 * time.Second)

				fmt.Fprintln(w, "", 200, "OK")

			}

		} else {

			fmt.Fprintln(w, "", 500, "access denied")
		}
	}
}

// using ExistDir Test if
// directory exists
func ExistDir(name string) bool {

	if _, err := os.Stat(name); err != nil {

		if os.IsNotExist(err) {

			return false
		}
	}

	return true
}

// using UrlUpload Create
// the url of our api Upload
func UrlUpload() string {

	return Scheme + "://" + Host + ":" + Port + "/upload"

}

// usign SaveDb This method is responsible
// for saving all relevant data to upload.
func SaveDb(keyfile string, namefile string, sizefile int64, pathFile string) {

	// Call gbolt method to save data in native bolt format
	err := gbolt.SaveDb(keyfile, namefile, sizefile, pathFile)

	if err == nil {

		fmt.Sprintf("save sucess..")

	} else {

		fmt.Println("Error", err)
	}

}
