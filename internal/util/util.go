package util

import (
	"log"
	"os"
	"path/filepath"
)

func FindRootDirectory() string {
	currentProcessPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	for dir := currentProcessPath; dir != "/"; dir = filepath.Dir(dir) {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
	}
	log.Fatal("root directory not found (missed go.mod?)")
	return ""
}
