package main

import "fmt"

func main() {

	var l map[string]int
	m := map[string]int{}
	o := map[int]string{}
	//m = make(map[string]int)
	m["Thukira"]++
	l = m
	fmt.Println(m)
	fmt.Println(l)
	fmt.Println(o)

	/**
	names := os.Args[1:]

	fmt.Println(hello.SayHello(names))
	//fmt.Printf("the average is %f\n", types.Average())
	types.Strings("Hello World!", "World", "Gowthaman")
	types.Array()
	**/

}
