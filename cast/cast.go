package cast

// Suite of casting functions to speed up solutions
// This is NOT idiomatic Go... but AOC isn't about that...

import (
	"fmt"
	"strconv"
)

// ToInt will cast a given arg into an int type.
// Supported types are:
//   - string
func ToInt(arg interface{}) int {
	var val int
	var err error
	switch arg.(type) {
	case string:
		val, err = strconv.Atoi(arg.(string))
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
	case rune:
		val, err = strconv.Atoi(string(arg.(rune)))
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
	case float32,
		float64:
		val = int(arg.(float64))
	default:
		panic(fmt.Sprintf("unhandled type for int casting %T", arg))
	}
	return val
}

// ToString will cast a given arg into a string type.
// Supported types are:
//   - int (all)
//   - byte
//   - rune
func ToString(arg interface{}) string {
	var str string
	switch arg.(type) {
	case int,
		uint,
		int8,
		int16,
		uint16,
		uint32,
		int64,
		uint64:
		str = strconv.FormatInt(int64(arg.(int)), 10)
	case byte: // uint8
		b := arg.(byte)
		str = string(rune(b))
	case rune: // int32
		str = string(arg.(rune))
	default:
		panic(fmt.Sprintf("unhandled type for string casting %T", arg))
	}
	return str
}

func ToSlice(m map[interface{}]interface{}) []interface{} {
	v := make([]interface{}, 0, len(m))

	for _, value := range m {
		v = append(v, value)
	}

	return v
}

const (
	ASCIICodeCapA   = int('A') // 65
	ASCIICodeCapZ   = int('Z') // 65
	ASCIICodeLowerA = int('a') // 97
	ASCIICodeLowerZ = int('z') // 97
)

// ToASCIICode returns the ascii code of a given input
func ToASCIICode(arg interface{}) int {
	var asciiVal int
	switch arg.(type) {
	case string:
		str := arg.(string)
		if len(str) != 1 {
			panic("can only convert ascii Code for string of length 1")
		}
		asciiVal = int(str[0])
	case byte:
		asciiVal = int(arg.(byte))
	case rune:
		asciiVal = int(arg.(rune))
	}

	return asciiVal
}

// ASCIIIntToChar returns a one character string of the given int
func ASCIIIntToChar(code int) string {
	return string(rune(code))
}
