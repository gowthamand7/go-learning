package main

import "fmt"

func testfunctions() {
	a := [...]int{1, 2, 3, 4, 5}
	doArray(a)
	fmt.Printf("outside the function a@ %p\n", a)

	b := []int{1, 2, 3, 4, 5}
	doSlice(b)

	c := map[string]int{
		"Gowthaman": 1,
		"Priya":     2,
		"Thukira":   3,
	}

	doMap(c)
}

func doArray(a [5]int) {
	a[0] = 0
	fmt.Printf("inside the function a@ %p\n", a)
}

func doSlice(b []int) {
	b[0] = 0
	fmt.Printf("inside the function b@ %p\n", b)
	b = append(b, 10)
	fmt.Printf("inside the function after reassiging b@ %p\n", b)
}

func doMap(c map[string]int) {
	c["Main"] = 4
	fmt.Printf("inside the function c@ %p\n", c)
}
