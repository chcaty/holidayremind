package pkg

import (
	"fmt"
	"github.com/spf13/viper"
)

func GetConfigByJson(value any, fileName string) error {
	viper.SetConfigName(fileName)
	viper.SetConfigType("json")
	viper.AddConfigPath("../../configs/json")
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
