package fakecsv

// copy from https://github.com/gocarina/gocsv/blob/master/safe_csv.go
// Wraps around SafeCSVWriter and makes it thread safe.
import (
	"encoding/csv"
	"sync"
)

// SafeCSVWriter use sync.Mutex protect writing
type SafeCSVWriter struct {
	*csv.Writer
	m sync.Mutex
}

// NewSafeCSVWriter return a new pointer to SafeCSVWriter
func NewSafeCSVWriter(original *csv.Writer) *SafeCSVWriter {
	return &SafeCSVWriter{
		Writer: original,
	}
}

// Write Override write
func (w *SafeCSVWriter) Write(row []string) error {
	w.m.Lock()
	defer w.m.Unlock()
	return w.Writer.Write(row)
}

// Flush Override flush
func (w *SafeCSVWriter) Flush() {
	w.m.Lock()
	w.Writer.Flush()
	w.m.Unlock()
}
