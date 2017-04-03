/**********

 ▄▄▄██▀▀▀▓█████   █████▒ █████▒▒█████  ▄▄▄█████▓ ▒█████   ███▄    █  ██▓
   ▒██   ▓█   ▀ ▓██   ▒▓██   ▒▒██▒  ██▒▓  ██▒ ▓▒▒██▒  ██▒ ██ ▀█   █ ▓██▒
   ░██   ▒███   ▒████ ░▒████ ░▒██░  ██▒▒ ▓██░ ▒░▒██░  ██▒▓██  ▀█ ██▒▒██▒
▓██▄██▓  ▒▓█  ▄ ░▓█▒  ░░▓█▒  ░▒██   ██░░ ▓██▓ ░ ▒██   ██░▓██▒  ▐▌██▒░██░
 ▓███▒   ░▒████▒░▒█░   ░▒█░   ░ ████▓▒░  ▒██▒ ░ ░ ████▓▒░▒██░   ▓██░░██░
 ▒▓▒▒░   ░░ ▒░ ░ ▒ ░    ▒ ░   ░ ▒░▒░▒░   ▒ ░░   ░ ▒░▒░▒░ ░ ▒░   ▒ ▒ ░▓
 ▒ ░▒░    ░ ░  ░ ░      ░       ░ ▒ ▒░     ░      ░ ▒ ▒░ ░ ░░   ░ ▒░ ▒ ░
 ░ ░ ░      ░    ░ ░    ░ ░   ░ ░ ░ ▒    ░      ░ ░ ░ ▒     ░   ░ ░  ▒ ░
 ░   ░      ░  ░                  ░ ░               ░ ░           ░  ░


* project server the Upload
* @package     main
* @author      jeffotoni
* @copyright   Copyright (c) 2017
* @license     --
* @link        --
* @since       Version 0.1
*/

// Package main
// Restful simple server written in Go (golang) to
// receive uploads, we use boltdb a native-non-sql go and
// using token with authentication.
//
// You can send client-encrypted files and common files.
// If you prefer dockerfile available so you can build
// in your image and run the goupload.
package main

import (
	"fmt"
	"os"

	"github.com/jeffotoni/goupload/pkg/libupload"
)

// using func main to server stadalone of Upload
func main() {

	if len(os.Args) > 1 { // Conferring the number of arguments

		command := os.Args[1] // Get Command

		if command != "" {

			if command == "start" {

				// using StartUploadServer
				// Method responsible for raising the restful service
				// Has no arguments
				libupload.StartUploadServer()

			} else if command == "stop" {

				fmt.Println("under development!!!!") // Stop server

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
