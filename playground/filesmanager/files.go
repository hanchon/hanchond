package filesmanager

import (
	"errors"
	"io"
	"io/fs"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return io.ReadAll(file)
}

func SaveFile(data []byte, path string) error {
	return SaveFileWithMode(data, path, 0o600)
}

func SaveFileWithMode(data []byte, path string, mode fs.FileMode) error {
	return os.WriteFile(path, data, mode)
}

func DoesFileExist(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
