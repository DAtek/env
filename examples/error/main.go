package main

import (
	"fmt"
	"os"

	"github.com/DAtek/env"
)

type Point struct {
	X int
	Y int
}

type Config struct {
	AppCenter     Point
	AppMaxWorkers int
	AppDbUrl      string
}

func main() {
	os.Clearenv()
	os.Setenv("APP_CENTER", "12;6")
	os.Setenv("APP_MAX_WORKERS", "more than ever")

	loadEnv := env.NewLoader[Config]()
	_, err := loadEnv()

	fmt.Println(err)
	fmt.Println()

	errorCollection := err.(*env.ErrorCollection)
	fmt.Printf("errorCollection: %v\n", errorCollection.Errors)
}
