package json_validator

import (
	"github.com/xeipuuv/gojsonschema"
)

func Validate(schema string, file string) (bool, error) {
	schema_loader := gojsonschema.NewStringLoader(schema)
	file_loader := gojsonschema.NewStringLoader(file)
	result, err := gojsonschema.Validate(schema_loader, file_loader)
	if err != nil {
		return false, err
	}
	return result.Valid(), nil
}
