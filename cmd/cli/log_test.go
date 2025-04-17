package cli

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func safe() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("Something bad happened")
}


func TestFileLog(t *testing.T) {
	// Set up log file
	logFile := "log/log.txt"
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	// Log a message
	logMessage := "Test log message"
	log.Println(logMessage)

	// Read the log file
	content, err := ioutil.ReadFile(logFile)
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	// Check if the log message is in the file
	if !strings.Contains(string(content), logMessage) {
		t.Errorf("log message not found in file")
	}

	// Clean up log file (optional)
	err = os.Truncate(logFile, 0)
	if err != nil {
		t.Fatalf("error truncating file: %v", err)
	}
}

func TestSafe(t *testing.T) {
	safe()
}
