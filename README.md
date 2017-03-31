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


## Examples client

Uploading with Authorization

```
curl -H 'Authorization:tyladfadiwkxceieixweiex747' --form nameupload=@nameFile.tar.bz2 http://localhost:8080/upload
```
