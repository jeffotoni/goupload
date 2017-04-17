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
package libupload

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/jeffotoni/goupload/pkg/glogs"
)

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
				file, handler, errf := r.FormFile("nameupload")

				if errf != nil {

					color.Red("Error big file, try again!")
					http.Error(w, "Error parsing uploaded file: "+errf.Error(), http.StatusBadRequest)
					return
				}

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
