package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/DAtek/env"
)

type Point struct {
	X int
	Y int
}

type Config struct {
	AppCenter Point
}

var customParsers = env.ParserMap{
	"Point": func(src string) (any, error) {
		parts := strings.Split(src, ";")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		return Point{x, y}, nil
	},
}

func main() {
	os.Clearenv()
	os.Setenv("APP_CENTER", "12;5")

	loadEnv := env.NewLoader[Config](customParsers)
	config, _ := loadEnv()

	fmt.Printf("config.AppCenter: %v\n", config.AppCenter)
}
