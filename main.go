package main

import "fmt"

type Number interface {
	int | int8 | int32 | int64 | float32 | float64
}

func SumNumber[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
	    s += v
	}
	return s
}

func main() {
	ints := map[string]int{
		"key1": 10,
		"key2": 100,
	}
	fmt.Printf("Sum int value: %v\n", SumNumber(ints))

	int8s := map[string]int{
		"key1": 15,
		"key2": 150,
	}
	fmt.Printf("Sum int8 value: %v\n", SumNumber(int8s))

	int32s := map[string]int32{
		"key1": 20,
		"key2": 200,
	}
	fmt.Printf("Sum int32 value: %v\n", SumNumber(int32s))

	int64s := map[string]int64{
		"key1": 30,
		"key2": 300,
	}
	fmt.Printf("Sum int64 value: %v\n", SumNumber(int64s))

	float32s := map[string]float64{
		"key1": 40.1,
		"key2": 400.1,
	}
	fmt.Printf("Sum float32 value: %v\n", SumNumber(float32s))


	float64s := map[string]float64{
		"key1": 50.2,
		"key2": 500.2,
	}
	fmt.Printf("Sum float64 value: %v\n", SumNumber(float64s))
}
