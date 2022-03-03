package main

import "fmt"


func SumNumber[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

func main() {
	ints := map[string]int64{
		"key1": 10,
		"key2": 100,
	}
	fmt.Printf("Sum int64 value: %v\n", SumNumber(ints))

	floats := map[string]float64{
		"key1": 10.0,
		"key2": 100.0,
	}
	fmt.Printf("Sum float64 value: %v\n", SumNumber(floats))
}
