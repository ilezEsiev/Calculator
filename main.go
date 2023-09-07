package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Введите арифметическое выражение след. типа (e.g., '5 + 3' or 'V * II'): ")
	str, _ := bufio.NewReader(os.Stdin).ReadString('\r')
	str = strings.Trim(str, "\r")
	str = strings.Trim(str, "\n")

	parts := strings.Split(str, " ")

	if len(parts) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	num1 := parts[0]
	operator := parts[1]
	num2 := parts[2]

	arabic1, roman1, err1 := parseNumber(num1)
	arabic2, roman2, err2 := parseNumber(num2)

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid number format. Please use Arabic (1-10) or Roman (I-X) num.")
		return
	}

	if (roman1 != "" && roman2 == "") || (roman2 != "" && roman1 == "") {
		fmt.Println("Ошибка, одно из чисел арабское другое римское")
		return
	}

	result := 0
	switch operator {
	case "+":
		result = arabic1 + arabic2
	case "-":
		result = arabic1 - arabic2
	case "*":
		result = arabic1 * arabic2
	case "/":
		if arabic2 == 0 {
			fmt.Println("Нельзя делить на ноль.")
			return
		}
		result = arabic1 / arabic2
	default:
		fmt.Println("Неправильный оператор/Invalid operator. Please use +, -, *, or /.")
		return
	}

	if roman1 != "" {
		if result <= 0 {
			fmt.Println("Ошибка")
			return
		}
		romanResult := arabicToRoman(result)
		fmt.Println("Result:", romanResult)
	} else {
		fmt.Println("Result:", result)
	}
}

func parseNumber(s string) (int, string, error) {

	arabicNumerals := map[string]int{
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5,
		"6": 6, "7": 7, "8": 8, "9": 9, "10": 10,
	}

	if val, ok := arabicNumerals[s]; ok {
		return val, "", nil
	}

	romanNumerals := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}

	if val, ok := romanNumerals[s]; ok {
		return val, s, nil
	}

	return 0, "", fmt.Errorf("invalid number format")
}

func arabicToRoman(num int) string {

	romanSymbols := []string{"X", "IX", "V", "IV", "I"}
	romanValues := []int{10, 9, 5, 4, 1}

	result := ""
	for i := 0; i < len(romanSymbols); i++ {
		for num >= romanValues[i] {
			result += romanSymbols[i]
			num -= romanValues[i]
		}
	}
	return result
}
