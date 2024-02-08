package main

import (
	"bufio"
	"fmt"
	"goCalc/internal"
	"os"
	"strconv"
	"strings"
)

type Number interface {
	int | float64
}

func readNumber[T Number](prompt string, reader *bufio.Reader) (T, error) {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		var result T
		if num, err := strconv.ParseFloat(input, 64); err == nil {
			result = T(num)
			return result, nil
		} else {
			fmt.Println("Введите числовое значение")
		}
	}
}

func calcLogic[T Number](reader *bufio.Reader) {
	num1, err := readNumber[T]("Число A: ", reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	num2, err := readNumber[T]("Число B: ", reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Выберите операцию (+, -, *, /): ")
	operator, _ := reader.ReadString('\n')
	operator = strings.TrimSpace(operator)

	operation := internal.GetOperation[T](rune(operator[0]))
	if operation == nil {
		fmt.Println("Недопустимая операция")
		return
	}

	result, resultErr := operation.Perform(num1, num2)
	if resultErr != nil {
		fmt.Println(resultErr)
		return
	}

	fmt.Printf("%v %c %v = %v\n", num1, operator[0], num2, result)
}

func main() {
	fmt.Println("Консольный калькулятор:")
	reader := bufio.NewReader(os.Stdin)
	for {
		calcLogic[float64](reader)
	}
}
