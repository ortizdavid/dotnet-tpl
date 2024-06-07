package helpers

import (
	"os"
	"path/filepath"
)

func GetCurrentFolder() string {
	cwd, _:= os.Getwd()
	folderName := filepath.Base(cwd)
	return folderName
}