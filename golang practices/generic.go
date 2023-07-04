package main

import "fmt"

type number interface {
	int64 | float64
}

func main() {

	ints := map[string]int64{"first": 34, "second": 87}
	floats := map[string]float64{"first": 34.2, "second": 87.3}
	fmt.Println(" int non generic sum ", intsum(ints))
	fmt.Println("float no generic sum ", floatsum(floats))
	fmt.Println("float no generic sum ", intfloatsum(floats))
	fmt.Println("float no generic sum ", intfloatsum(ints))

}
func intsum(m map[string]int64) int64 {
	var s int64 = 0
	for _, v := range m {
		s = s + v
	}
	return s
}
func floatsum(m map[string]float64) float64 {
	var f float64 = 0
	for _, v := range m {
		f = f + v
	}
	return f
}
func intfloatsum[k comparable, v number](m map[k]v) v {
	var s v
	for _, v := range m {
		s += v
	}
	return s
}
