package main

import "fmt"

func main() {
	less := func(a int) int {
		return a * a
	}
	testfunctions()
	sq(less)
}

func sq(less func(a int) int) bool {
	a := []int{2, 1, 4, 6, 8, 4, 9}

	for i := range a {
		a[i] = less(a[i])
	}
	fmt.Println(a)
	return true
}
