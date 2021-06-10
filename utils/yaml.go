package utils

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
)

func YAMLtoJSON(b []byte) ([]byte, error) {
	if len(b) == 0 {
		return nil, nil
	}

	var in map[string]interface{}
	if err := yaml.Unmarshal(b, &in); err != nil {
		return nil, err
	}

	return json.Marshal(in)
}
