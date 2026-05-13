package archive

import (
	"os"
)

func ReadArchive(path string) ([]byte, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return data, err
}