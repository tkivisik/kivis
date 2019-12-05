package logwriter

import (
	"fmt"
	"os"
	"time"
)

// New is a factory function for *logWriter
func New(f *os.File) *logWriter {
	return &logWriter{
		out: f,
	}
}

type logWriter struct {
	out *os.File
}

// Write writes bytes to out specified in the struct
func (writer logWriter) Write(bytes []byte) (int, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05\t")
	content := string(bytes)

	line := fmt.Sprintf("%s\t%s", timestamp, content)
	return writer.out.WriteString(line)
}
