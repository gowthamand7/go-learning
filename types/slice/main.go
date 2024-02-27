package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	aa := [][]byte{}

	for _, item := range items {
		nitem := item[:]
		aa = append(aa, nitem)
	}

	fmt.Println(aa)

	users := []User{
		{
			Name: "Gowtham",
			Age:  23,
		},
		{
			Name: "Sachin",
			Age:  24,
		},
		{
			Name: "Rohit",
			Age:  25,
		},
	}

	gow := &users[0]

	nu := User{
		Name: "Priya",
		Age:  30,
	}

	fmt.Printf(" Before reallocation : %p \n", &users[0])
	gow.Age = 26
	fmt.Printf(" After reallocation : %p \n", &users[0])

	users = append(users, nu)

	fmt.Println(users)

	print := func(a []int) {
		fmt.Printf("Value : %+v, Lenght : %d, Capacity : %d, Address : %p, Type : %T, Nil : %t\n", a, len(a), cap(a), &a, a, a == nil)
	}

	var a []int
	b := []int{}
	c := []int{1, 2, 3, 4, 5}
	d := make([]int, 0, 6)

	print(a)
	print(b)
	print(c)
	print(d)
	d = append(d, 0, 1, 2, 3, 4, 5, 6, 7)
	print(d)

	e := c[:2:2]
	print(e)

}
