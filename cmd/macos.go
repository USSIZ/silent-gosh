package cmd

import (
	"os/exec"

	"github.com/USSIZ/silent-gosh/utils"
)

// Clears clipboard on MacOS
func ClearBoardMac() {
	// Const to customize the prompted log if success
	const msgbMc string = "Clipboard has been cleared"
	// Array to store the command and arguments to be executed
	var cbMc = [3]string{"/bin/bash", "-c", "pbcopy < /dev/null"}
	// Execute the command and arguments
	_, err := exec.Command(cbMc[0], cbMc[1], cbMc[2]).Output()
	// Check for errors
	utils.CheckError(msgbMc, err)
}

// Clears history of commands on MacOS
func ClearHistoryMac() {
	// Const to customize the prompted log if success
	const msghMc string = "History has been cleared"
	// Array to store the command and arguments to be executed
	//FIXME: THOSE 2 WORKS !! But after executing, we have to close active terminal windows before to check history
	// var chMc = [3]string{"/bin/sh", "-c", "rm -f ~/.zsh_history && rm -f ~/.bash_history"}
	var chMc = [3]string{"/bin/sh", "-c", "cat /dev/null > ~/.bash_history && cat /dev/null > ~/.zsh_history"}
	// Execute the command and arguments
	_, err := exec.Command(chMc[0], chMc[1], chMc[2]).Output()
	// Check for errors
	utils.CheckError(msghMc, err)
}
