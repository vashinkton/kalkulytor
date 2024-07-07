package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func RomanToArabic(roman string) (int, error) {
	romanNumeralMap := map[string]int{
		"M": 1000, "CM": 900, "D": 500, "CD": 400,
		"C": 100, "XC": 90, "L": 50, "XL": 40,
		"X": 10, "IX": 9, "V": 5, "IV": 4,
		"I": 1,
	}

	arabic := 0
	for i := 0; i < len(roman); {
		if i+1 < len(roman) && romanNumeralMap[roman[i:i+2]] != 0 {
			arabic += romanNumeralMap[roman[i:i+2]]
			i += 2
		} else {
			arabic += romanNumeralMap[string(roman[i])]
			i++
		}
	}
	return arabic, nil
}

func Calculate(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("неверный оператор")
	}
}

// Проверка на входное число не больше 10
func CheckRange(number int) bool {
	return number >= 1 && number <= 10
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите выражение:")
	for scanner.Scan() {
		input := scanner.Text()
		re := regexp.MustCompile(`(?i)(\d+|\w+)\s*([\+\-\*\/])\s*(\d+|\w+)`)
		matches := re.FindStringSubmatch(input)

		if matches == nil || len(matches) != 4 {
			fmt.Println("Неверный формат ввода")
			continue
		}

		operand1, operand2 := matches[1], matches[3]
		operator := matches[2]

		var num1, num2 int
		var err error

		if regexp.MustCompile(`(?i)^m{0,4}(cm|cd|d?c{0,3})(xc|xl|l?x{0,3})(ix|iv|v?i{0,3})$`).MatchString(operand1) {
			num1, err = RomanToArabic(strings.ToUpper(operand1))
			if err != nil {
				fmt.Println("Ошибка при конвертации римского числа:", err)
				continue
			}
		} else {
			num1, err = strconv.Atoi(operand1)
			if err != nil {
				fmt.Println("Ошибка при конвертации арабского числа:", err)
				continue
			}
		}

		// Диапазон первого числа
		if !CheckRange(num1) {
			fmt.Println("Число вне допустимого диапазона (1-10)")
			continue
		}

		if regexp.MustCompile(`(?i)^m{0,4}(cm|cd|d?c{0,3})(xc|xl|l?x{0,3})(ix|iv|v?i{0,3})$`).MatchString(operand2) {
			num2, err = RomanToArabic(strings.ToUpper(operand2))
			if err != nil {
				fmt.Println("Ошибка при конвертации римского числа:", err)
				continue
			}
		} else {
			num2, err = strconv.Atoi(operand2)
			if err != nil {
				fmt.Println("Ошибка при конвертации арабского числа:", err)
				continue
			}
		}

		// Диапазон второго числа
		if !CheckRange(num2) {
			fmt.Println("Число вне допустимого диапазона (1-10)")
			continue
		}

		result, err := Calculate(num1, num2, operator)
		if err != nil {
			fmt.Println("Ошибка при вычислении:", err)
			continue
		}
		fmt.Printf("Результат: %d\n", result)
	}
}
