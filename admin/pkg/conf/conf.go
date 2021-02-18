package conf

import (
	"github.com/spf13/viper"
)

var (
	Instance *Conf

	DefaultConfigPath string = "./conf"
)

type Conf struct {
	ConfMap map[string]*viper.Viper
}

func New(filename string) (*viper.Viper, error) {
	if Instance != nil {
		if _, ok := Instance.ConfMap[filename]; ok {
			return Instance.ConfMap[filename], nil
		}
	} else {
		Instance = &Conf{
			ConfMap: make(map[string]*viper.Viper),
		}
	}

	v := viper.New()
	v.SetConfigName(filename)
	v.SetConfigType("json")
	v.AddConfigPath(DefaultConfigPath)

	err := v.ReadInConfig()
	if err != nil { return nil, err }

	Instance.ConfMap[filename] = v
	return Instance.ConfMap[filename], nil
}
