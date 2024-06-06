package generators

import "github.com/ortizdavid/go-nopain/filemanager"

type MvcGenerator struct {
}

func (mvc MvcGenerator) Generate(templateType string) error {
	var fileManager filemanager.FileManager
	fileManager.CreateSingleFolder("")
	return nil
}