package main

import (
	"fmt"

	"github.com/bearatol/go_learn_calculat/operations"
)

func main() {
	var (
		a, b, res float64
		oper, err string
	)

	fmt.Print("Введите пример с двумя значениями: ")
	fmt.Scanf("%f %s %f %v", &a, &oper, &b)

	switch oper {
	case "+":
		res = operations.Sum(a, b)
	case "-":
		res = operations.Raznost(a, b)
	case "*":
		res = operations.Proizvedenie(a, b, a)
	case "/":
		res = operations.Delenie(a, b)
	case "%":
		res = operations.Ostatok(a, b)
	default:
		err = "Не верный пример"
	}
	fmt.Println("help")
	if err == "" {
		fmt.Printf("%v %v %v = %v", a, oper, b, res)
	} else {
		fmt.Println(err)
	}

	fmt.Println("111")

}
