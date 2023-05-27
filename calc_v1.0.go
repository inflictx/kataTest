package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите математическую операцию: ")
	input, _ := reader.ReadString('\n')
	result, err := calculate(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Результат:", result)
}

func calculate(input string) (int, error) {
	// Удаляем пробелы в начале и конце строки
	input = strings.TrimSpace(input)

	// Разделяем строку на операнды и оператор
	operands := strings.Split(input, " ")
	if len(operands) != 3 {
		return 0, fmt.Errorf("неверный формат математической операции")
	}

	// Парсим операнды
	a, err := parseOperand(operands[0])
	if err != nil {
		return 0, err
	}
	b, err := parseOperand(operands[2])
	if err != nil {
		return 0, err
	}

	// Проверяем, что используются числа одной системы счисления
	isArabic := isArabicNumeral(operands[0]) && isArabicNumeral(operands[2])
	isRoman := isRomanNumeral(operands[0]) && isRomanNumeral(operands[2])
	if !isArabic && !isRoman {
		return 0, fmt.Errorf("используются одновременно разные системы счисления")
	}

	// Выполняем операцию
	var result int
	operator := operands[1]
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		result = a / b
	default:
		return 0, fmt.Errorf("неверный оператор")
	}

	// Проверяем условия для римских чисел
	if isRoman {
		if result <= 0 {
			return 0, fmt.Errorf("результат работы с римскими числами меньше единицы")
		}
		return result, nil
	}

	return result, nil
}

func parseOperand(operand string) (int, error) {
	if isArabicNumeral(operand) {
		// Парсим арабское число
		num, err := strconv.Atoi(operand)
		if err != nil {
			return 0, fmt.Errorf("неверный формат числа")
		}
		if num < 1 || num > 10 {
			return 0, fmt.Errorf("число должно быть от 1 до 10")
		}
		return num, nil
	} else if isRomanNumeral(operand) {
		// Парсим римское число
		num, err := romanToArabic(operand)
		if err != nil {
			return 0, fmt.Errorf("неверный формат числа")
		}
		return num, nil
	}

	return 0, fmt.Errorf("неверный формат числа")
}

func isArabicNumeral(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isRomanNumeral(s string) bool {
	validRomanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, rn := range validRomanNumerals {
		if s == rn {
			return true
		}
	}
	return false
}

func romanToArabic(s string) (int, error) {
	romanNumerals := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	result := 0
	prev := 0
	for i := len(s) - 1; i >= 0; i-- {
		val := romanNumerals[s[i]]
		if val < prev {
			result -= val
		} else {
			result += val
			prev = val
		}
	}

	if arabicToRoman(result) != s {
		return 0, fmt.Errorf("неверный формат числа")
	}

	return result, nil
}

func arabicToRoman(num int) string {
	arabicNumerals := []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	romanNumerals := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}

	result := ""
	for i := len(arabicNumerals) - 1; i >= 0; i-- {
		for num >= arabicNumerals[i] {
			result += romanNumerals[i]
			num -= arabicNumerals[i]
		}
	}

	return result
}
