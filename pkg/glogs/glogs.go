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

// Package glogs
// The glogs package is responsible for flexibilizing
// and implementing the recording of the logs, the cool
// thing is that we do this with the printf function of fmt.
//
// Very light your implementation, uses pkg golang log,
// we made a * log.Logger instance, and we were able
// to easily integrate our output.
package glogs

import (
	"flag"
	"log"
	"os"
)

// Variables that we instantiate
// to reuse in our packages
var (
	PathLog = flag.String("pathLog", "goupload.log", "Logs goupload")
	Log     *log.Logger
)

// LogNew Instantiating method Log, opens the file, if it
// does not exist it creates, if it already exists it
// does append in the log and returns the instance log.new
func LogNew(pathLog string) {

	// Opening file, if there is no create,
	// if it already exists, append the file,
	// leave full permission
	file, err := os.OpenFile(pathLog, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)

	if err != nil {
		panic(err)
	}

	// Instantiating Log.New to do the recording of logs
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)

}
