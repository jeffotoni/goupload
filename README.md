# goupload
Restful simple server written in Go (golang) to receive uploads, we use boltdb a native-non-sql go and using token with authentication.
You can send client-encrypted files and common files.
If you prefer dockerfile available so you can build in your image and run the goupload.

## Used libraries:
- https://github.com/boltdb/bolt - No-sql native go(golang)
- https://github.com/gorilla/mux - Implements a request router and dispatcher for matching incoming requests
- https://github.com/fatih/color - Help define the colors in the terminal

---
* [A small summary](#summary)
* [Docker](#docker)
* [Clone](#clone)
* [Get dependencies](#dependencies)
* [Run](#runprogram)
* [Logs Server](#rlogsserver)
* [Build](#Build)
* [Install](#install)
* [Body function](#body-function)
* [Examples Client](#examples-client)
* [Upload File](#upload-files)

## A small summary 

[goupload.go]

A simple restful server to receive common and encrypted files will store the file on the server.
You will use port 8080 for communication with your api, will be written a no-sql database boltdb all 
uploads made on the server, will be generated a log on disk of all accesses.


## Docker 

docker [`Installing`] (Docker https://docs.docker.com/engine/installation)


## Clone this repo into your GOPATH

```
git clone https://github.com/jeffotoni/goupload
```

dockerfile-golang
```
FROM jeffotoni/golang
MAINTAINER Jefferson Otoni Lima <jeff.otoni@gmail.com>
ADD . /go/src/github.com/jeffotoni/goupload
RUN go install github.com/jeffotoni/goupload
RUN mkdir -p /go/goupload
WORKDIR /go/goupload
ENTRYPOINT /go/bin/goupload start
EXPOSE 8080
```

Building image
```
# docker build -f dockerfile-golang -t jeffotoni/goupload:v1 .

```

Start image and creating container
```
# docker run -it --rm --name goupload-run -p 8080:8080 -v /yourpath/docker:/go/goupload  jeffotoni/goupload:v1

```

There is another possibility, which is of the pull in a complete image

Pull in the image
```
# docker pull jeffotoni/golang-goupload

```

Start image and creating container
```
# docker run -it --rm --name goupload-run -p 8080:8080 -v /yourpath/docker:/go/goupload  jeffotoni/golang-goupload:v1

```

Now is to test and see if everything is ok

Sending a file to the server
```
# curl -X POST -H 'Authorization:tyladfadiwkxceieixweiex747' --form nameupload=@Yourfile http://localhost:8080/upload

```

Visualizing all logs generated in real time
```
# docker exec id-container tail -f /go/goupload/goupload.log
# or
# tail -f /yourpath/docker/goupload.log

```

Listing container ports
```
# docker exec id-container nmap localhost
```

Listing the processes of your container
```
# docker exec id-container ps aux

```

Here is where the running program is located, here will be generated the no-sql boltdb database, 
the access log and where it will store the uploads made by the client

```
# docker exec id-container ls -lh /go/goupload
# or
# ls -lh /yourpath/docker/
```

Here are all sources of goupload
```
# docker exec id-container ls -lh /go/src/github.com/jeffotoni

```

If you want to get into the container to have a look or change something that you think
Enter the container
```
# docker exec -ti id-container /bin/sh
```

## Install local on your machine 

Clone this repo into your GOPATH
```
git clone https://github.com/jeffotoni/goupload

```

## Get dependencies

```sh
got get -u github.com/boltdb/bolt
got get -u github.com/gorilla/mux
got get -u github.com/jeffotoni/goupload

```

## Run the program local on your machine

```go
# go run goupload.go start
Services successfully tested
Host: localhost
Scheme:http
Port: 8080
Instance POST  http://localhost:8080/upload
Loaded service

```

## Logs generated from the server 

```
#tail -f goupload.log
2017/04/01 23:12:11 libupload.go:186: ......................start upload .........................
2017/04/01 23:12:11 libupload.go:187: Authorization: tyladfadiwkxceieixweiex747
2017/04/01 23:12:11 libupload.go:188: Path dir: uploads/tyladfadiwkxceieixweiex747
2017/04/01 23:12:11 libupload.go:189: Path file: uploads/tyladfadiwkxceieixweiex747/file2.pdf
2017/04/01 23:12:11 libupload.go:190: File: file2.pdf
2017/04/01 23:12:11 libupload.go:191: Size:  0
2017/04/01 23:12:11 libupload.go:192: Allowed:  500 Mb
2017/04/01 23:12:11 libupload.go:193: Copied: 246281 bytes
2017/04/01 23:12:11 libupload.go:194: Copied: 240 Kb
2017/04/01 23:12:11 libupload.go:195: Copied: 0 Mb
2017/04/01 23:12:11 libupload.go:196: Database key: tyladfadiwkxceieixweiex747/file2.pdf
2017/04/01 23:12:11 libupload.go:198: ...........................................................

```

Stopping the server

```go
# go run goupload.go stop

```

## You can also build goupload.go 

```go
# go build goupload.go
# ./goupload.go
Services successfully tested
Host: localhost
Scheme:http
Port: 8080
Instance POST  http://localhost:8080/upload
Loaded service

```

## Install this repo into your GOPATH
Check out your GOPATH, when install goupload the executable will go to your bin

```sh
# cd goupload && go install
# goupload
Services successfully tested
Host: localhost
Scheme:http
Port: 8080
Instance POST  http://localhost:8080/upload
Loaded service

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


```

Body of main function UploadFile

Uploadfile implemented method will create a folder as the name of the token and stored the file in this created folder, this method checks the maximum size of the upload to allow or not to upload to the server.

```go

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

```

Body of main function Connect
```go
// using Connect
// This method is what does
// and returns our instance for access
// to our no-sql database.
func Connect() *DB {

	if db2 == nil {

		// Singleton the bank has to close every call,
		// save, update, get etc ..
		//
		// Testing and verifying if there is a directory
		// and file of our bucket, if it does not exist
		// it creates the directory and the file so that
		// we can manipulate all our bucket.
		//
		// Remember that boltdb with the open function also creates.
		if err := DataBaseTest(PathDb); err != nil {

			log.Fatal("Error Test database", err)
		}

		// Here is the object responsible for
		// allowing calls to the methods, such as Get, Save, etc.
		dbbolt, err = bolt.Open(PathDb, 0600, &bolt.Options{Timeout: 1 * time.Second})

		if err != nil {

			log.Fatal("connect error: ", err)
		}

		// We create a new reference
		// just to facilitate
		// understanding and syntax
		db2 = &DB{dbbolt}
	}

	return db2
}

```

Body of main function SaveDb
```go

func SaveDb(keyfile string, namefile string, sizefile int64, pathFile string) error {

	times := fmt.Sprintf("%s", time.Now())

	stringJson := JsonDataDb{keyfile, namefile, sizefile, pathFile, times}

	respJson, err := json.Marshal(stringJson)

	respJsonX := string(respJson)

	err = Save(keyfile, respJsonX)

	if err == nil {

		//fmt.Println("save sucess..")
		return nil

	} else {

		//fmt.Println("Error", err)
		return err
	}

}

```

## Examples client

Uploading with Authorization

```
curl -H 'Authorization:tyladfadiwkxceieixweiex747' --form nameupload=@nameFile.tar.bz2 http://localhost:8080/upload

```
