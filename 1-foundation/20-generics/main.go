package main

func SumInteger(m map[string]int) int {
	var sum int

	for _, v := range m {
		sum += v
	}

	return sum
}

func SumFloat(m map[string]float64) float64 {
	var sum float64

	for _, v := range m {
		sum += v
	}

	return sum
}

func Sum[T Number](m map[string]T) T {
	var sum T

	for _, v := range m {
		sum += v
	}

	return sum
}

/** Constraint */
type Number interface {
	~int | ~float64
}

// type Number interface {
// 	int | float64
// }
type MyNumber int

func Compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	/** Generics
	 */

	m := map[string]int{"Daniel": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Daniel": 1000.0, "João": 2000.0, "Maria": 3000.0}
	m3 := map[string]MyNumber{"Daniel": 1000, "João": 2000, "Maria": 3000}

	println(SumInteger(m))
	// error
	// println(SumInteger(m3))
	println(SumFloat(m2))

	println(Sum(m))
	println(Sum(m2))
	println(Sum(m3))
	println(Compare(10, 10))
}
