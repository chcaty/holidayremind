package pkg

import "errors"

// GetMapValue 根据Key获取对应的Value
func GetMapValue[T any, K comparable](config map[K]T, key K, value *T) error {
	if result, ok := config[key]; ok {
		*value = result
		return nil
	}
	return errors.New("map not contain this key")
}
