package solution

import (
	"github.com/spf13/viper"
)

//подготовка чтения из конфига
func ReadConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

const MaxRoutines = 10

type Index struct {
}
