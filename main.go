package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	returnNotesPath := flag.Bool("p", false, "returns a path where notes are kept")
	resetNotes := flag.Bool("reset", false, "removes all notes")

	flag.Parse()
	args := flag.Args()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	baseDir := filepath.Join(homeDir, ".kivis")
	err = os.MkdirAll(baseDir, os.ModePerm)
	if err != nil {
		fmt.Printf("%s: %s", err, "could not set default notes path")
	}

	notesPath := filepath.Join(baseDir, ".notes")
	if *returnNotesPath {
		fmt.Println(notesPath)
		os.Exit(0)
	}
	if *resetNotes {
		err := os.Remove(notesPath)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("There is currently nothing to remove")
				os.Exit(0)
			}
			fmt.Println(err)
		}
		fmt.Println("Notes deleted from: " + notesPath)
		os.Exit(0)
	}

	if len(args) < 1 {
		// Print all current notes
		f, err := os.OpenFile(notesPath, os.O_RDONLY, 0644)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("There are currently no notes")
				os.Exit(0)
			}
			fmt.Println(err)
			os.Exit(0)
		}
		defer f.Close()

		b, err := ioutil.ReadAll(bufio.NewReader(f))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s", b)
		os.Exit(0)
	}

	f, err := os.OpenFile(filepath.Join(baseDir, ".notes"),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	lw := logWriter{
		out: f,
	}
	log.SetFlags(0)
	log.SetOutput(lw)

	for i := 0; i < len(args); i++ {
		note := args[i]
		log.Printf("%s", note)
	}
}

type logWriter struct {
	out *os.File
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05\t")
	content := string(bytes)

	line := fmt.Sprintf("%s\t%s", timestamp, content)
	return writer.out.WriteString(line)
}
