package main

import "fmt"

func ReverseString(str string) string {
	sliceRunes := []rune(str)
	for i, j := 0, len(sliceRunes)-1; i < len(sliceRunes)/2; i, j = i+1, j-1 {
		sliceRunes[i], sliceRunes[j] = sliceRunes[j], sliceRunes[i]
	}
	return string(sliceRunes)
}
func main() {
	var str string
	fmt.Printf("Enter the string to reverse: ")
	fmt.Scan(&str)

	reverseString := ReverseString(str)
	fmt.Printf("Reverse string for is: %v\n", reverseString)
}
