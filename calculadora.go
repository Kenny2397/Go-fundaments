package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

// receiver function
func (calc) operate(entrada string, operator string) int {

	entradaLimpia := strings.Split(entrada, operator)
	operator1 := parsear(entradaLimpia[0])
	operator2 := parsear(entradaLimpia[1])

	switch operator {
	case "+":
		// fmt.Println(operator1 + operator2)
		return operator1 + operator2

	case "-":
		// fmt.Println(operator1 - operator2)
		return operator1 - operator2

	case "*":
		// fmt.Println(operator1 * operator2)
		return operator1 * operator2

	case "/":
		// fmt.Println(operator1 / operator2)
		return operator1 / operator2

	default:
		fmt.Println("Invalid Input operator")
		return 0
	}

}

func parsear(entrada string) int {
	operator, _ := strconv.Atoi(entrada)
	return operator
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()
	return operacion
}

func main() {
	entrada := leerEntrada()
	operator := leerEntrada()

	fmt.Println(entrada)
	fmt.Println(operator)

	c := calc{}
	result := c.operate(entrada, operator)
	fmt.Println(result)

	// if err1 != nil || err2 != nil {
	// 	fmt.Println("err1:", err1)
	// 	fmt.Println("err2:", err2)
	// } else {

	// }

	// fmt.Println(operator1 + operator2)
	// fmt.Println(err1, "\n", err2)

}
