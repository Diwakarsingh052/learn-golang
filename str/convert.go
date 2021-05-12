package str

import (
	"strconv"
)

func ConvertStrToInt(s string) (int, error) { // error should be last returned value

	i, err := strconv.Atoi(s) // most of time your error var name is err

	if err != nil {
		return 0, err // always return default value of a type in case of err
	}

	i = i * 2

	return i, nil

}
