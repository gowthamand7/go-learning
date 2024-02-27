package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//for loop
	fmt.Println("for loop")
	numbers := make([]int, 0)
	i := 0
	for i < 10 {
		i++
		numbers = append(numbers, i)
		fmt.Println(i)
	}

	//for each loop
	fmt.Println("for each loop")
	for index, value := range numbers {
		fmt.Printf("index %d, value %d  %#v \n", index, value, value)
	}

	//while loop
	fmt.Println("while loop")
	j := 0
	for {
		j++
		fmt.Println(j)
		if j == 10 {
			break
		}
	}

	students := []string{"Gowthaman", "Priya", "Thukira"}
	adult := []string{"Gowthaman", "Priya"}

	//break and continue
	fmt.Println("break and continue")
outer:
	for _, student := range students {
		for _, adults := range adult {
			if student == adults {
				fmt.Printf("%s is an adult\n", student)
				continue outer
			}
		}
		fmt.Printf("%s is not an adult\n", student)
	}

	//switch case
	fmt.Println("switch case")
	var x int = getRandomNumber()
	var ok error
	switch r := getRandomNumber(); r {
	case 0, 1, 2, 3, 4:
		_, ok = fmt.Println("Number is less than 5", x)
	case 5, 6, 7, 8, 9:
		fmt.Println("Number is greater than 4")
	default:
		fmt.Println("Number is not less than 5 and not greater than 4")
	}
	fmt.Println(x)
	fmt.Println(ok)

}

func getRandomNumber() int {
	return rand.Intn(11)
}
