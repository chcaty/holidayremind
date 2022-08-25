package uxconfig

import (
	"fmt"
	"github.com/spf13/viper"
)

func GetValue(fileName string, fileType fileType, path string, value any) error {
	viper.SetConfigName(fileName)
	viper.SetConfigType(string(fileType))
	viper.AddConfigPath(path)
	err := viper.ReadInConfig() // Find and read the uxconfig file
	if err != nil {             // Handle errors reading the uxconfig file
		return fmt.Errorf("viper get uxconfig file error %w\n", err)
	}
	err = viper.Unmarshal(value)
	if err != nil {
		return fmt.Errorf("viper unmarshal uxconfig error %w\n", err)
	}
	return nil
}
