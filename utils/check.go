package utils

import (
	"log"
)

// Centralizes errors checking to avoid repetition
func CheckError(msg string, err error) {
	// Set prefix depending on error message
	if err != nil {
		log.Fatalln("[ 🔴 ] :::", err)
	} else {
		log.Println("[ 🟢 ] :::", msg)
	}
}

//TODO: Add a function to check if a file exists
//TODO: Add a function to check system architecture
