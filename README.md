# goupload
Simple restful server written in Go (golang) to receive uploads, we use boltdb a no-sql native go and using token with authentication.
You can send files encrypted by the client, and ordinary files.

## Used libraries:
- https://github.com/boltdb/bolt - No-sql native go(golang)
- https://github.com/gorilla/mux - Implements a request router and dispatcher for matching incoming requests

---
* [A small summary](#summary)
* [Install](#install)
* [Run](#runprogram)
* [Body function](#body-function)
* [Examples Client](#examples-client)
* [Upload File](#upload-files)

## A small summary 

* [goupload.go]

A simple restful server for receiving common and encrypted files, will store the file on the server.


## Install

```sh
git clone https://github.com/jeffotoni/goupload

```

## Run the program

```go
go run server.go start

Services successfully tested
Host: localhost
Scheme:http
Port: 8080
Instance POST  http://localhost:8080/upload
Loaded service

```

Stopping the server

```go
go run server.go stop

```

Compiling upload server

```go
go build server.go
```

## Body function

Body of main function

```go

/** [Main function] */

func main() {

	// start and stop server

	if len(os.Args) > 1 {

		command := os.Args[1]

		if command != "" {

			if command == "start" {

				// Start server

				libupload.StartUploadServer()

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

```

Body of StartUploadServer

```go

/** [StartUploadServer Will build our route, and make calls to the upload method and validations] */

func StartUploadServer() {

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

			} else if r.Method == "GET" {

				fmt.Fprintln(w, "http ", 500, "Not authorized / Allowed method POST")

			} else {

				fmt.Fprintln(w, "http ", 500, "Not authorized / Allowed method POST")
			}
		})

	httpConf = &http.Server{

		Handler: router,
		Addr:    Host + ":" + Port,

		// Good idea!!! Good live!!!

		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Fatal(httpConf.ListenAndServe())
}

```

Body of main function UploadFile

Uploadfile implemented method will create a folder as the name of the token and stored the file in this created folder, this method checks the maximum size of the upload to allow or not to upload to the server.

```go
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

				os.MkdirAll(pathUpKeyUser, 0777)

				pathUserAcess := PathLocal + acessekey + "/" + handler.Filename

				fmt.Println(pathUserAcess)

				// copy file and write

				f, _ := os.OpenFile(pathUserAcess, os.O_WRONLY|os.O_CREATE, 0666)
				defer f.Close()
				n, _ := io.Copy(f, file)

				//up_size := fmt.Sprintf("%v", r.ContentLength)

				//To display results on server

				name := strings.Split(handler.Filename, ".")
				fmt.Printf("File name: %s\n", name[0])
				fmt.Printf("extension: %s\n", name[1])

				fmt.Println("size file: ", sizeMaxUpload)
				fmt.Println("allowed: ", UploadSize, "Mb")

				fmt.Printf("copied: %v bytes\n", n)
				fmt.Printf("copied: %v Kb\n", n/1024)
				fmt.Printf("copied: %v Mb\n", n/1048576)

				fmt.Fprintln(w, "", 200, "OK")

			}

		} else {

			fmt.Fprintln(w, "", 500, "access denied")
		}
	}
}

```

## Examples client

Uploading with Authorization

```
curl -H 'Authorization:tyladfadiwkxceieixweiex747' --form nameupload=@nameFile.tar.bz2 http://localhost:8080/upload
```
