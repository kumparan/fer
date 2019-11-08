package util

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

//CopyFolder :nodoc:
func CopyFolder(source string, dest string) (err error) {

	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := filepath.Join(source, obj.Name())

		destinationfilepointer := filepath.Join(dest, obj.Name())

		if obj.IsDir() {
			err = CopyFolder(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

//CopyFile :nodoc:
func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer func() {
		_ = sourcefile.Close()
	}()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer func() {
		_ = destfile.Close()
	}()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			_ = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}
