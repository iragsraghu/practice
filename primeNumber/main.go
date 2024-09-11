package main

import "fmt"

func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func GeneratePrimes(limit int) []int {
	primes := []int{}
	for i := 2; i < limit; i++ {
		if IsPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	var limit int
	fmt.Printf("Enter the limit to display prime numbers: ")
	_, err := fmt.Scan(&limit)
	if err != nil {
		fmt.Println("Invalid input, please positive numbers")
		return
	}

	primes := GeneratePrimes(limit)
	if len(primes) == 0 {
		fmt.Println("No prime numbers foound for the given limit ", limit)
	} else {
		fmt.Println("The prime numbers for the given limit", limit, "are ", primes)
	}
}
