package main

// Sum adds two integers and returns the result.
func Sum(a int32, b int32) int32 {
	return a + b
}

func main() {
	// The main function is kept for completeness, but in the Wasm context,
	// Sum function will be invoked directly.
}
