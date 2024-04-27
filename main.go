package main

import "fmt"

func main() {
	var max_number int
	fmt.Print("Enter max number: ")
	_, err := fmt.Scan(&max_number)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	for i := 1; i < max_number+1; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
