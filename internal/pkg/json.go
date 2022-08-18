package pkg

import (
	"encoding/json"
	"fmt"
)

func ToJson(object any, value *string) error {
	data, err := json.Marshal(object)
	*value = string(data)
	if err != nil {
		return fmt.Errorf("Marshal with error: %w\n", err)
	}
	return nil
}

func ToObject(value string, object *any) error {
	data := []byte(value)
	err := json.Unmarshal(data, object)
	if err != nil {
		return fmt.Errorf("Unmarshal with error: %w\n", err)
	}
	return nil
}
