package main

import (
	"log"
	"runtime"
	"time"

	"github.com/USSIZ/silent-gosh/cmd"
	"github.com/USSIZ/silent-gosh/utils"
)

//TODO: Push the project on GitHub to save the work done at this point

//TODO: (IDEA) Kill all terminals if success after timer or y/n prompt
//TODO: (IDEA) Add a function to empty the bin if possible/needed
//TODO: (IDEA) Add a function to empty the cache if possible/needed
//TODO: (IDEA) Rework it as a command line tool ???

// Defines actions to execute depending on OS
func main() {
	const location string = "Europe/Kiev"
	// Const to customize the prompted log if success
	const msgTime string = "Printing current time in " + location
	// load time zone
	loc, e := time.LoadLocation(location)
	utils.CheckError(msgTime, e)

	// Get time in that zone
	log.Println("[ 🕓 ] :::", time.Now().UTC().In(loc))

	// Storing the OS name and arch in variables
	const goos string = runtime.GOOS
	const goarch string = runtime.GOARCH
	// Switching between functions applied depending on OS
	switch {
	case goos == "linux":
		log.Printf("[ ⚪️ ] ::: Running on Linux (%s/%s)", goos, goarch)
		cmd.ClearHistoryLinux()
		cmd.ClearBoardLinux()
	case goos == "darwin":
		log.Printf("[ ⚪️ ] ::: Running on MacOS (%s/%s)", goos, goarch)
		cmd.ClearHistoryMac()
		cmd.ClearBoardMac()
	case goos == "windows":
		//TODO: Handle non linux / darwin cases
		log.Fatalf("[ 🔴 ] ::: Not yet supported on Windows (%s/%s)", goos, goarch)
		break
	default:
		log.Fatal("Program not supported on this OS (for now)")
		break
	}
}
