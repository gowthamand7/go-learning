package types

import "fmt"

func Array() {

	a := [3]int{1, 2, 3}
	fmt.Printf("a is %T and the value us %[1]d\n", a)

	b := [3]string{"one", "two", "three"}
	fmt.Printf("a is %T and the value us %[1]d\n", b)

	c := [...]User{{"Gowthaman", "<EMAIL>"}, {"Priya", "<EMAIL>"}}
	fmt.Printf("a is %T and the value us %[1]d\n", c)

	e := []int{1, 2, 3}
	e = append(e, 4)
	fmt.Printf("a is %T and the value us %[1]d\n", e)

	f := a[2:3]
	f = append(f, 4)
	fmt.Printf("a is %T and the value us %[1]d\n", f)

}

type User struct {
	name  string
	email string
}
