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
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jeffotoni/goupload/pkg/gbolt"
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
	UploadSize    int64
	PathLocal     = "uploads/"
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

	//Host +
	httpConf = &http.Server{

		Handler: router,
		Addr:    ":" + Port,

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

				//To display results on server

				name := strings.Split(handler.Filename, ".")
				fmt.Printf("File name: %s\n", name[0])
				//fmt.Printf("extension: %s\n", name[1])

				fmt.Println("size file: ", sizeMaxUpload)
				fmt.Println("allowed: ", UploadSize, "Mb")

				fmt.Printf("copied: %v bytes\n", bytes)
				fmt.Printf("copied: %v Kb\n", bytes/1024)
				fmt.Printf("copied: %v Mb\n", bytes/1048576)

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

		fmt.Println("save sucess..")

	} else {

		fmt.Println("Error", err)
	}

}
