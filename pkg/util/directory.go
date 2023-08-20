package util

import (
	"os"
	"path/filepath"
)

func GetListOfFilesFromDirectory(fileExtension string, directory string) (map[string]os.FileInfo, error) {
	toReturn := map[string]os.FileInfo{}
	err := filepath.Walk(directory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == fileExtension {
				toReturn[info.Name()] = info
			}
			return nil
		})
	return toReturn, err
}
