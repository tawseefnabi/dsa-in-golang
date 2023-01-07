package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	fmt.Println("Enter first number: ")
	var first int16
	fmt.Scanln(&first)

	fmt.Println("Enter second number: ")
	var second int16
	fmt.Scanln(&second)
	fmt.Println(reflect.TypeOf(first).Kind())
	if reflect.TypeOf(first).Kind() != reflect.Int16 || reflect.ValueOf(second).Kind() != reflect.Int16 {
		log.Fatal("Only numbers are allowed")
	}
	fmt.Println("Sum is: ", first+second)
}
