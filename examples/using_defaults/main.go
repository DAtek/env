package main

import (
	"fmt"
	"os"

	"github.com/DAtek/env"
)

type Config struct {
	AppMaxWorkers  int
	AppLoggingType *string
}

type DefaultConfig struct {
	AppMaxWorkers int
}

func main() {
	os.Clearenv()

	loadEnv := env.NewLoader[Config]()
	defaultConfig := DefaultConfig{AppMaxWorkers: 8}
	config, _ := loadEnv(defaultConfig)

	fmt.Printf("config.AppMaxWorkers: %v\n", config.AppMaxWorkers)
	fmt.Printf("config.AppLoggingType: %v\n", config.AppLoggingType)
}
