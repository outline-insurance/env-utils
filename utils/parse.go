package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// ParseEnvFile parses a jsonc env file and returns a map with the contents
func ParseEnvFile(filePath string) (*map[string]string, error) {
	envFile, err := os.Open(filePath)
	if err != nil {
		return nil, errors.Wrapf(err, "while opening json file %s", filePath)
	}
	defer envFile.Close()

	fileContents, err := ioutil.ReadAll(envFile)
	if err != nil {
		return nil, errors.Wrapf(err, "while reading json file %s", filePath)
	}

	var envMap map[string]string
	err = json.Unmarshal(fileContents, &envMap)
	if err != nil {
		return nil, errors.Wrapf(err, "while unmarshalling json file %s", filePath)
	}
	return &envMap, nil
}
