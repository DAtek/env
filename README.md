[![codecov](https://codecov.io/gh/DAtek/env/graph/badge.svg?token=MAjETYy681)](https://codecov.io/gh/DAtek/env) [![Go Report Card](https://goreportcard.com/badge/github.com/DAtek/env)](https://goreportcard.com/report/github.com/DAtek/env)


# env
Lightweight library, which parses environmental variables into structs.

## Features
- Supports custom parsers
- Optional fields
- Default values
- Generics
- Returns friendly error message
- Returned error can be parsed into struct for further processing if needed

## Examples

### Simple

```go
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

```

Output:
```
config.AppMaxWorkers: 10  
config.AppLoggingType: <nil>
```


### Using defaults

```go
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

```

Output:
```
config.AppMaxWorkers: 8
config.AppLoggingType: <nil>
```

### Using custom parsers

```go
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

```

Output:
```
config.AppCenter: {12 5}
```

### Error
```go
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

```

Output:
```
Parser missing for environmental variable 'APP_CENTER'. Required type: 'Point'
Environmental variable 'APP_MAX_WORKERS' has wrong type. Required type: 'int'
Environmental variable 'APP_DB_URL' is unset

errorCollection: [{APP_CENTER unsupported_type Point PARSER_NOT_FOUND} {APP_MAX_WORKERS wrong_type int strconv.ParseInt: parsing "more than ever": invalid syntax} {APP_DB_URL required string <nil>}]
```