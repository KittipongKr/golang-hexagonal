package json

import (
	jsoniter "github.com/json-iterator/go"
)

func JsoniterMarshalIndent(input interface{}, output interface{}) error {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	b, _ := json.MarshalIndent(input, "", "  ")
	err := json.Unmarshal(b, output)
	if err != nil {
		return err
	}

	return nil
}
