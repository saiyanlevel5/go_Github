package main

import (
	"os"
	"bytes"
)

func ReadMyFile(fileName string) (string){

	file, errorInOpeningFile := os.Open(fileName)
	if errorInOpeningFile != nil {
		LOGGER.Info("Could not open the file - ", file)
		LOGGER.Fatalf("Stack Trace Of Error:", errorInOpeningFile)
	}
	defer file.Close()

	var bytesBuffer *bytes.Buffer = new(bytes.Buffer)
	bytesBuffer.ReadFrom(file)
	LOGGER.Debug("Printing bytesbuffer:\n", bytesBuffer.String())
	return bytesBuffer.String()

}

func GetFileName() (string) {

	var userQuestion string = "Enter The Name Of The File You Would Like To Parse\n"
	var fileName string = GetUserInput(userQuestion)
	return fileName

}
