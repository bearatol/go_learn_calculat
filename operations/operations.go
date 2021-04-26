package operations

import "math"

func Sum(a, b float64) (res float64) {
	res = a + b
	return
}

func Raznost(a, b float64) (res float64) {
	res = a - b
	return
}

func Proizvedenie(a, b float64) (res float64) {
	res = a * b
	return
}

func Delenie(a, b float64) (res float64) {
	res = a / b
	return
}

func Ostatok(a, b float64) (res float64) {
	//res = a % b
	res = math.Mod(a, b)
	return
}
