package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var (
	imOnce sync.Once
	im     immutable
)

type (
	// ImmutableConfig ...
	ImmutableConfig interface {
		GetPort() int
		GetKey() string
	}

	immutable struct {
		Port int    `mapstructure:"PORT"`
		Key  string `mapstructure:"KEY"`
	}
)

func (i *immutable) GetPort() int {
	return i.Port
}

func (i *immutable) GetKey() string {
	return i.Key
}

func GetDefaultImmutableConfig() ImmutableConfig {
	var outer error
	var success = true

	env := os.Getenv("APP_ENV")
	pwd := os.Getenv("APP_PWD")

	if env == "test" && pwd == "" {
		panic(fmt.Errorf("APP_PWD env is required in test env"))
	}

	imOnce.Do(func() {
		v := viper.New()
		v.SetConfigName("app.config")

		if env == "test" {
			v.AddConfigPath(pwd)
		} else {
			v.AddConfigPath(".")
		}

		v.SetEnvPrefix("vp")
		v.AutomaticEnv()

		if err := v.ReadInConfig(); err != nil {
			success = false
			outer = fmt.Errorf("failed to read app.config file due to %s", err)
		}

		v.Unmarshal(&im)
	})

	if !success {
		panic(outer)
	}

	return &im
}
