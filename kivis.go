package kivis

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/tkivisik/kivis/logwriter"
)

// Notes are there for you.
type Notes struct {
	Path string
}

// Print all current notes that are written to stone.
func (n Notes) Print() {
	f, err := os.OpenFile(n.Path, os.O_RDONLY, 0644)
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
}

// AddNote appends a note to a stone
func (n Notes) AddNote(args []string) {
	f, err := os.OpenFile(n.Path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	lw := logwriter.New(f)
	log.SetFlags(0)
	log.SetOutput(lw)

	for i := 0; i < len(args); i++ {
		note := args[i]
		log.Printf("%s", note)
	}
}

// Reset deletes the notes file
func (n Notes) Reset() {
	err := os.Remove(n.Path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("There is currently nothing to remove")
			os.Exit(0)
		}
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Notes deleted from: " + n.Path)
}
