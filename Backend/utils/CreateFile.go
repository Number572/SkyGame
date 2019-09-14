package utils

import (
	"os"
)

func CreateFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		panic("Database not created")
	}
	file.Close()
}
