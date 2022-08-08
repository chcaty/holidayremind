package common

import (
	"encoding/json"
	"fmt"
)

func MapToJson(m any) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Marshal with error: %+v\n", err)
		return "", nil
	}
	return string(jsonByte), nil
}
