package notes

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// ErrNoNotes is used when there is no notes location, or no notes stored.
var ErrNoNotes = fmt.Errorf("there are no notes")

type Noter interface {
	Append(text string) error
	IterNote(fn func(string)) error
	Locate() error
	Destroy() error
}

// Notes are there for you.
type Notes struct {
	Location string
}

// Print all current notes that are written to stone.
// TODO: rewrite into InterateNote(fn func) error
func (n Notes) Print() error {
	f, err := os.OpenFile(n.Location, os.O_RDONLY, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrNoNotes
		}
		return err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(bufio.NewReader(f))
	if err != nil {
		return err
	}
	fmt.Printf("%s", b)
	return nil
}

// Append a note
func (n Notes) Append(text string) error {
	f, err := os.OpenFile(n.Location,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05\t")
	line := fmt.Sprintf("%s\t%q", timestamp, text)
	fmt.Println(line)
	return nil
}

// Destroy all notes, remove them forever
func (n Notes) Destroy() error {
	err := os.Remove(n.Location)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrNoNotes
		}
		return err
	}
	fmt.Println("Notes deleted from: " + n.Location)
	return nil
}

// Locate where the notes are kept
func (n Notes) Locate() error {
	fmt.Println(n.Location)
	return nil
}
