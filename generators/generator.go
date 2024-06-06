package generators

import "errors"

type Generator struct {
}

func (gen *Generator) GenerateTemplate(templateType string) error {
	switch templateType {
	case "mvc":
		return MvcGenerator{}.Generate(templateType)
	case "api":
		return ApiGenerator{}.Generate(templateType)
	default:
		return errors.New("Invalid template type: "+templateType)
	}
}