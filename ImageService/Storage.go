package main

import (
	"io"
	"os"
)

func saveFile(name string, file io.Reader, folder string) error {
	newFile, err := os.Create("/tmp/storage/" + folder + "/" + name)
	defer newFile.Close()
	if err != nil {
		logError(err)
		return err
	}
	_, err = io.Copy(newFile, file)
	return err
}
