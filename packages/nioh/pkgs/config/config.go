package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// CfgProvider ...
type CfgProvider struct {
	UnmarshalKey      func(key string, cfg interface{}) error
	UnmarshalExactKey func(key string, cfg interface{}) interface{}
}

// ReadConfig to invoke reading of local config
func ReadConfig(cf string) *CfgProvider {
	var instance = viper.New()
	instance.AutomaticEnv()

	if cf != "" {
		instance.SetConfigFile(cf)
		err := instance.ReadInConfig()
		if err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("fatal error reading config file: %s", err))
		}
	}

	return &CfgProvider{
		UnmarshalKey: func(key string, cfg interface{}) error {
			if partialCfg := instance.Sub(key); partialCfg != nil {
				if err := partialCfg.UnmarshalExact(&cfg); err != nil {
					return fmt.Errorf("[warn]: Unable to bind config with given key \"%s\": %s", key, err)

				}
				return nil
			}

			return fmt.Errorf("[warn]: Unable to retrieve any config with given key:  %s", key)
		},
	}
}
