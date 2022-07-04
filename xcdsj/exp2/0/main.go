package main

import "fmt"

type i interface {
	method()
}

type a struct{}
func (_ *a) method1() {}

type b struct{}
func (_ b) method() {}

func main() {

	var o1 i = a{} // a does not implement i (method method has pointer receiver)
	var o2 i = b{}
	fmt.Println(o1, o2)

	(b{}).method()

}