package main

import "fmt"

func isPrime(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	if number > 1 {
		return true
	}
	return false
}

func findPrimeGame() {
	fmt.Println("Prime numbers less than 20:")
	for number := 0; number < 20; number++ {
		if isPrime(number) {
			fmt.Printf("%d ", number)
		}
	}
}

func guessNumberGame() {
	val := 0
	for {
		fmt.Print("Enter number:")
		fmt.Scanf("%d", &val)
		switch {
		case val < 0:
			panic("You entered a negative number!")
		case val == 0:
			fmt.Println("0 is neither negative nor positive")
		default:
			fmt.Println("You entered:", val)
		}
	}
}

func main() {
	//findPrimeGame()
	//guessNumberGame()
}
