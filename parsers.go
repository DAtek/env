package env

import (
	"strconv"
	"strings"
)

type Parser func(src string) (any, error)
type ParserMap map[string]Parser

var baseParsers = ParserMap{
	"string": func(src string) (any, error) {
		return src, nil
	},
	"int": func(s string) (any, error) {
		return strToInt(s, 64, func(i int64) any { return int(i) })
	},
	"int8": func(s string) (any, error) {
		return strToInt(s, 8, func(i int64) any { return int8(i) })
	},
	"int16": func(s string) (any, error) {
		return strToInt(s, 16, func(i int64) any { return int16(i) })
	},
	"int32": func(s string) (any, error) {
		return strToInt(s, 32, func(i int64) any { return int32(i) })
	},
	"int64": func(s string) (any, error) {
		return strToInt(s, 64, func(i int64) any { return i })
	},
	"uint": func(s string) (any, error) {
		return strToUint(s, 64, func(u uint64) any { return uint(u) })
	},
	"uint8": func(s string) (any, error) {
		return strToUint(s, 8, func(a uint64) any { return uint8(a) })
	},
	"uint16": func(s string) (any, error) {
		return strToUint(s, 16, func(a uint64) any { return uint16(a) })
	},
	"uint32": func(s string) (any, error) {
		return strToUint(s, 32, func(a uint64) any { return uint32(a) })
	},
	"uint64": func(s string) (any, error) {
		return strToUint(s, 64, func(a uint64) any { return a })
	},
	"float32": func(s string) (any, error) {
		f, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return f, err
		}

		return float32(f), nil
	},
	"float64": func(s string) (any, error) {
		return strconv.ParseFloat(s, 64)
	},
	"bool": func(src string) (any, error) {
		lowercase := strings.ToLower(src)
		for _, trueValue := range trueValues {
			if trueValue == lowercase {
				return true, nil
			}
		}

		return false, nil
	},
}

var trueValues = []string{"t", "true", "y", "yes", "1"}

func strToInt(s string, bitSize int, cast func(int64) any) (any, error) {
	v, err := strconv.ParseInt(s, 10, bitSize)
	if err != nil {
		return v, err
	}

	return cast(v), nil
}

func strToUint(s string, bitSize int, cast func(uint64) any) (any, error) {
	v, err := strconv.ParseUint(s, 10, bitSize)
	if err != nil {
		return v, err
	}

	return cast(v), nil
}
