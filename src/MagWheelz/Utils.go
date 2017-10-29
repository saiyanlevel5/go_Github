package main

import (
	"os"
	"fmt"
)

func GetUserInput(userQuestion string) (string) {

	var userInput string
	fmt.Fprintf(os.Stdout, userQuestion)
	fmt.Fscanf(os.Stdin, "%s", &userInput)
	LOGGER.Debug("Acquired File Name: ", userInput)
	return userInput

}

