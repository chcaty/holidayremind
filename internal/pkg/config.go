package pkg

import (
	"fmt"
	"github.com/spf13/viper"
)

func GetConfigByJson(fileName string, fileType FileType, path string, value any) error {
	viper.SetConfigName(fileName)
	viper.SetConfigType(string(fileType))
	viper.AddConfigPath(path)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return fmt.Errorf("viper get config error %w\n", err)
	}
	err = viper.Unmarshal(value)
	if err != nil {
		return fmt.Errorf("viper unmarshal config error %w\n", err)
	}
	return nil
}

type FileType string

const (
	Json FileType = "json"
	Yml  FileType = "yml"
)

const Path = "../../configs/json"
