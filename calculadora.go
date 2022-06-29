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
// func (calc) operate(entrada string, operator string) int {

// }
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()
	fmt.Println(operacion)

	operador := "/"

	valores := strings.Split(operacion, operador)
	fmt.Println(valores)

	fmt.Println(valores[0] + valores[1])

	// operador1, _ := strconv.Atoi(valores[0])

	operador1, err1 := strconv.Atoi(valores[0])
	operador2, err2 := strconv.Atoi(valores[1])

	if err1 != nil || err2 != nil {
		fmt.Println("err1:", err1)
		fmt.Println("err2:", err2)
	} else {
		switch operador {
		case "+":
			fmt.Println(operador1 + operador2)

		case "-":
			fmt.Println(operador1 - operador2)

		case "*":
			fmt.Println(operador1 * operador2)

		case "/":
			fmt.Println(operador1 / operador2)

		default:
			fmt.Println("Invalid Input operator")

		}
	}

	// fmt.Println(operador1 + operador2)
	// fmt.Println(err1, "\n", err2)

}
