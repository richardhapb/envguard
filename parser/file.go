package parser

import (
	"bufio"
	"io"
	"os"
)

// File represents a file with scanning capabilities and proper cleanup
type File struct {
	scanner *bufio.Scanner
	closer	io.Closer
}

// ReadFile opens a file and returns a File instance for scanning
func ReadFile(path string) (*File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &File{
		scanner: bufio.NewScanner(f),
		closer:  f,
	}, nil
}

// Scanner returns the underlying scanner
func (f *File) Scanner() *bufio.Scanner {
	return f.scanner
}

// Close closes the underlying file
func (f *File) Close() error {
	return f.closer.Close()
}

// Err returns any error that occurred during scanning
func (f *File) Err() error {
	return f.scanner.Err()
}
