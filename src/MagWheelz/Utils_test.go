package main

import (
	"testing"
)
//NOTE: UNIT TEST IS INCOMPLETE
func TestGetUserInput(t *testing.T) {

	tables := []struct {
		userQuestion string
		expectedResult string
	} {
		{"Test: Write an entry\n", ""},
	}
	for _, table := range tables {
		var userInput string = GetUserInput(table.userQuestion)

		if userInput != table.expectedResult {
			t.Errorf("Error in getting user input\nExpected:\t%s\nGotten:\t%s", table.expectedResult,
				userInput)
		}
	}

}