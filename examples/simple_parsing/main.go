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

func main() {
	os.Clearenv()
	os.Setenv("APP_MAX_WORKERS", "10")

	loadEnv := env.NewLoader[Config]()
	config, _ := loadEnv()

	fmt.Printf("config.AppMaxWorkers: %v\n", config.AppMaxWorkers)
	fmt.Printf("config.AppLoggingType: %v\n", config.AppLoggingType)
}
