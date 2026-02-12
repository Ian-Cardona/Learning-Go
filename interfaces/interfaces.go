package main

import "fmt"

type MyInt int

func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	i2, ok := i.(int)
	if !ok {
		fmt.Errorf("unexpected type for %v", 1)
	}
	fmt.Println(i2 + 1)
}
