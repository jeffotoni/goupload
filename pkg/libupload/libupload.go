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

/** Environment variables and keys */

var (
	httpConf      *http.Server
	AUTHORIZATION = `tyladfadiwkxceieixweiex747`
	socketfileTmp = `server.red`
	socketfile    = `server.lock`
	Port          = "8080"
	Scheme        = "http"
	Database      = "ServerUpload"
	Host          = "localhost"

	//Leave blank if you run the docker for example, it will receive connection from other machines
	HostHttp = ""

	UploadSize int64
	PathLocal  = "uploads/"
)

/** [startUploadServer restful server upload] */

func StartUploadServer() {

	// Start ping database

	gbolt.Save("Ping", "ok")

	if gbolt.Get("Ping") != "ok" {

		fmt.Println("Services Error Data Base!")
		os.Exit(1)
	}

	fmt.Println("Services successfully tested")

	fmt.Println("Host: " + Host)
	fmt.Println("Scheme:" + Scheme)
	fmt.Println("Port: " + Port)

	fmt.Println("Instance POST ", UrlUpload())
	fmt.Println("Loaded service")

	///create route

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

	httpConf = &http.Server{

		Handler: router,
		Addr:    HostHttp + ":" + Port,

		// Good idea!!! Good live!!!

		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Fatal(httpConf.ListenAndServe())
}

/** [UploadFile method implemented] */

func UploadFile(w http.ResponseWriter, r *http.Request) {

	// SET SIZE UPLOAD

	UploadSize = 500 //MB

	Autorization := r.Header.Get("Authorization")

	if Autorization == "" {

		fmt.Fprintln(w, "", 500, "Not Authorized")

	} else {

		////check database get id user

		if Autorization == AUTHORIZATION {

			///Valid user
			acessekey := Autorization

			sizeMaxUpload := r.ContentLength / 1048576 ///Mb

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

				pathUserAcess := PathLocal + acessekey + "/" + handler.Filename

				// copy file and write

				f, _ := os.OpenFile(pathUserAcess, os.O_WRONLY|os.O_CREATE, 0777)
				defer f.Close()

				bytes, _ := io.Copy(f, file)
				keyfile := acessekey + "/" + handler.Filename

				SaveDb(keyfile, handler.Filename, bytes, pathUserAcess)

				// Generates a log of everything that happens on the server

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
				glogs.Log.Printf("Database key: \n", keyfile)

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

func ExistDir(name string) bool {

	if _, err := os.Stat(name); err != nil {

		if os.IsNotExist(err) {

			return false
		}
	}

	return true
}

func UrlUpload() string {

	return Scheme + "://" + Host + ":" + Port + "/upload"

}

func SaveDb(keyfile string, namefile string, sizefile int64, pathFile string) {

	err := gbolt.SaveDb(keyfile, namefile, sizefile, pathFile)

	if err == nil {

		fmt.Sprintf("save sucess..")

	} else {

		fmt.Println("Error", err)
	}

}
