package cmd

import (
	"os/exec"

	"github.com/USSIZ/silent-gosh/utils"
)

//TODO: For now, these are just copies from the MacOS one
// I have to change them later to work on Windows

// Clears clipboard on Windows
func ClearBoardWindows() {
	// Const to customize the prompted log if success
	const msgbLx string = "Clipboard has been cleared"
	// Array to store the command and arguments to be executed
	var cbLx = [3]string{"/bin/bash", "-c", "pbcopy < /dev/null"}
	// Execute the command and arguments
	_, err := exec.Command(cbLx[0], cbLx[1], cbLx[2]).Output()
	// Check for errors
	utils.CheckError(msgbLx, err)
}

// Clears history of commands on Windows
func ClearHistoryWindows() {
	// Const to customize the prompted log if success
	const msghLx string = "History has been cleared"
	// Array to store the command and arguments to be executed
	var chLx = [3]string{"/bin/bash", "-c", "history -cw"}
	// Execute the command and arguments
	_, err := exec.Command(chLx[0], chLx[1], chLx[2]).Output()
	// Check for errors
	utils.CheckError(msghLx, err)
}
