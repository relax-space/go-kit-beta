package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type ProjectDto struct {
	Databases map[string][]string `json:"databases"` //mysql,redis,mongo,sqlserver
}

func ReadViper(env string, config interface{}) error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}

	if env != "" {
		f, err := os.Open("config." + env + ".yml")
		if err != nil {
			return fmt.Errorf("Fatal error config file: %s \n", err)
		}
		defer f.Close()
		viper.MergeConfig(f)
	}

	if err := viper.Unmarshal(config); err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}
	return nil
}

func ReadToDto() (p *ProjectDto, err error) {
	b := []byte(`databases:
  mysql:
    - fruit
  sqlserver:
    - fruit2`)
	vip := viper.New()
	vip.SetConfigType("yaml")
	if err = vip.ReadConfig(bytes.NewBuffer(b)); err != nil {
		return
	}
	p = &ProjectDto{}
	if err = vip.Unmarshal(p); err != nil {
		return
	}
	return
}
