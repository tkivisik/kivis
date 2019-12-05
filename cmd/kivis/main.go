package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tkivisik/kivis"
)

var version = "v0.1"

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

	notesPath := MakeDir()
	notes := kivis.Notes{Path: notesPath}

	args := os.Args[1:]
	switch os.Args[1] {
	case "add":
		notes.AddNote(args)
	case "show":
		notes.Print()
	case "path":
		fmt.Println(notes.Path)
	case "reset":
		notes.Reset()
	case "version":
		Version()
	}
}

// MakeDir creates an app directory if one does not exist
func MakeDir() (notesPath string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	notesDir := filepath.Join(homeDir, ".kivis")
	err = os.MkdirAll(notesDir, os.ModePerm)
	if err != nil {
		fmt.Printf("%s: %s", err, "could not set default notes path")
	}

	notesPath = filepath.Join(notesDir, ".notes")
	return notesPath
}
