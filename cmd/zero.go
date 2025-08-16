package cmd

import (
	"fmt"
	"os"
	"strconv"
)

func Add(first string, second string) (result string) {
	num1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: First value is invalid")
		return
	}

	num2, err := strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Second value is invalid")
		return
	}

	return fmt.Sprintf("%.2f", num1+num2)
}

func Subtract(first string, second string) (result string) {
	num1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: First value is invalid")
		return
	}

	num2, err := strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Second value is invalid")
		return
	}

	return fmt.Sprintf("%.2f", num1-num2)
}

func Multiply(first string, second string, shouldRoundUp bool) (result string) {
	num1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: First value is invalid")
		return
	}

	num2, err := strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Second value is invalid")
		return
	}

	if shouldRoundUp {
		return fmt.Sprintf("%.2f", num1*num2)
	}

	return fmt.Sprintf("%f", num1*num2)
}
