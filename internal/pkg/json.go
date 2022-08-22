package pkg

import (
	"encoding/json"
	"fmt"
)

func ToJson(object any, value *string) error {
	data, err := json.Marshal(object)
	*value = string(data)
	if err != nil {
		return fmt.Errorf("json marshal with error: %w\n", err)
	}
	return nil
}

func ToBytes(object any, value *[]byte) error {
	data, err := json.Marshal(object)
	*value = data
	if err != nil {
		return fmt.Errorf("json marshal with error: %w\n", err)
	}
	return nil
}

func ToObjectByString(value string, object any) error {
	data := []byte(value)
	err := json.Unmarshal(data, object)
	if err != nil {
		return fmt.Errorf("json unmarshal with error: %w\n", err)
	}
	return nil
}

func ToObjectByBytes(value []byte, object any) error {
	err := json.Unmarshal(value, object)
	if err != nil {
		return fmt.Errorf("josn unmarshal with error: %w\n", err)
	}
	return nil
}
