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

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

/** Environment variables and keys */

var (
	confServer    *http.Server
	AUTHORIZATION = `tyladfadiwkxceieixweiex747`
	socketfileTmp = `server.red`
	socketfile    = `server.lock`
	Port          = "8080"
	Scheme        = "http"
	Database      = "ServerUpload"
	Host          = "localhost"
)

func main() {

	// start and stop server

	if len(os.Args) > 1 {

		command := os.Args[1]

		if command != "" {

			if command == "start" {

				// Start server

				startUploadServer()

			} else if command == "stop" {

				// Stop server

				fmt.Println("under development!!!!")

			} else {

				fmt.Println("Usage: server {start|stop}")
			}

		} else {

			command = ""
			fmt.Println("No command given")
		}
	} else {

		fmt.Println("Usage: server {start|stop}")
	}
}

func startUploadServer() {

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

				fmt.Fprintln(w, "http ", 200, "ok")

			} else if r.Method == "GET" {

				fmt.Fprintln(w, "http ", 500, "Not authorized / Allowed method POST")

			} else {

				fmt.Fprintln(w, "http ", 500, "Not authorized / Allowed method POST")
			}
		})

	confServer = &http.Server{

		Handler: router,
		Addr:    Host + ":" + Port,

		// Good idea!!! Good live!!!

		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Fatal(confServer.ListenAndServe())
}

func UrlUpload() string {

	return Scheme + "://" + Host + ":" + Port + "/upload"

}
