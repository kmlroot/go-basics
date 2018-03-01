package numbers

// Returns 3 integers
func GetVariables() (int, int, int) {
	return 1, 2, 3
}

// Returns 3 floats
func GetFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

// Sum 3 integer numbers and returns the result
func Sum(a int, b int) int {
	return a + b
}
