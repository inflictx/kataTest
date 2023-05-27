package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
	"I": 1,
	"II": 2,
	"III": 3,
	"IV": 4,
	"V": 5,
	"VI": 6,
	"VII": 7,
	"VIII": 8,
	"IX": 9,
	"X": 10,
}

type Operation struct {
	Input  string
	Result interface{}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	history := make([]Operation, 0, 3) // Создаем пустой срез для хранения истории операций
	for {
		fmt.Print("Введите математическую операцию: ")
		input, _ := reader.ReadString('\n')
		result, err := calculate(input)
		if err != nil {
			fmt.Println(err)
			continue // Пропускаем текущую итерацию цикла и продолжаем считывание новой операции
		}
		fmt.Println("Результат:", result)

		op := Operation{
			Input:  input,
			Result: result,
		}

		history = append(history, op) // Добавляем текущую операцию в историю

		// Если история содержит более трех операций, удаляем самую старую
		if len(history) > 3 {
			history = history[1:]
		}

		fmt.Println("История операций:")
		for _, op := range history {
			fmt.Printf("%s = %v\n", op.Input, op.Result)
		}
		fmt.Println()
	}
}

func calculate(input string) (interface{}, error) {
	input = strings.TrimSpace(input)
	operands := strings.Split(input, " ")
	if len(operands) != 3 {
		return nil, fmt.Errorf("неверный формат математической операции")
	}

	a, err := parseOperand(operands[0])
	if err != nil {
		return nil, err
	}
	b, err := parseOperand(operands[2])
	if err != nil {
		return nil, err
	}

	if isArabicNumeral(operands[0]) != isArabicNumeral(operands[2]) {
		return nil, fmt.Errorf("используются одновременно разные системы счисления")
	}

	operator := operands[1]
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return nil, fmt.Errorf("деление на ноль")
		}
		return a / b, nil
	default:
		return nil, fmt.Errorf("неверный оператор")
	}
}

func parseOperand(operand string) (int, error) {
	if isArabicNumeral(operand) {
		num, err := strconv.Atoi(operand)
		if err != nil || num < 1 || num > 10 {
			return 0, fmt.Errorf("неверный формат числа")
		}
		return num, nil
	} else if isRomanNumeral(operand) {
		return romanToArabic(operand)
	}

	return 0, fmt.Errorf("неверный формат числа")
}

func isArabicNumeral(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isRomanNumeral(s string) bool {
	_, exists := romanNumerals[s]
	return exists
}

func romanToArabic(s string) (int, error) {
	num, exists := romanNumerals[s]
	if !exists {
		return 0, fmt.Errorf("неверный формат числа")
	}
	return num, nil
}
