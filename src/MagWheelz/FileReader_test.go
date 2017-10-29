package main

import (
	"testing"
)
//NOTE: UNIT TEST IS INCOMPLETE
func TestGetFileName(t *testing.T) {

	var fileName string = GetFileName()
	LOGGER.Info(fileName)
	return
}

func TestReadMyFile(t *testing.T) {

	tables := []struct {
		expectedResult string
	} {
		{"Test: Write an entry +=123^&*<"},
	}
	for _, table := range tables {
		var actualFileContent string = ReadMyFile("testfile.txt")

		if actualFileContent != table.expectedResult {
			t.Errorf("Error in getting user input\nExpected:\t%s\nGotten:\t%s", table.expectedResult,
				actualFileContent)
		}
	}

}