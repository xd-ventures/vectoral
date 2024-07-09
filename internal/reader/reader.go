package parser

import (
	"io"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func ReadStdin() ([]byte, error) {
	return io.ReadAll(os.Stdin)
}
