package common

import (
	"os"
)

func LoadData(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	return string(data), err
}
