package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tkivisik/kivis/notes"
)

var version = "v0.2"

// Version shows the current version of the program
func Version() {
	fmt.Printf("kivis %s\n\n", version)
	flag.Usage()
}

func main() {
	usage := func() {
		fmt.Printf("usage:\n\tkivis { add | show | reset | version }\n")
		os.Exit(2)
	}
	if len(os.Args) < 2 {
		usage()
	}

	notesLocation := MakeDir()
	notes := notes.Notes{Location: notesLocation}

	args := os.Args[1:]
	switch os.Args[1] {
	case "add":
		note := strings.Join(args, " ")
		notes.Append(note)
	case "show":
		notes.Print()
	case "location":
		notes.Locate()
	case "reset":
		notes.Destroy()
	case "version":
		Version()
	}
}

// MakeDir creates an app directory if one does not exist
func MakeDir() (notesLocation string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	notesDir := filepath.Join(homeDir, ".kivis")
	err = os.MkdirAll(notesDir, os.ModePerm)
	if err != nil {
		fmt.Printf("%s: %s", err, "could not set default notes path")
	}

	notesLocation = filepath.Join(notesDir, ".notes")
	return notesLocation
}