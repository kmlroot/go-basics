package numbers

import "errors"

// Returns 3 integers
func GetVariables() (int, int, int) {
	return 1, 2, 3
}

// Returns 3 floats
func GetFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

// Sum 3 integer numbers and returns the result
func Sum(a interface{}, b interface{}) (int, error) {
	switch a.(type) {
	case string:
		return 0, errors.New("The A value is a string")
	}

	switch b.(type) {
	case string:
		return 0, errors.New("The B value is a string")
	}

	return a.(int) + b.(int), nil
}
