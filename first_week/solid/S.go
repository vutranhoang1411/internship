package main

import "fmt"

//wrong way
type Employee struct {
}

func (e Employee) TaskA() {
	fmt.Printf("Employee doing task A")
}
func (e Employee) TaskB() {
	fmt.Printf("Employee doing task B")
}

// right way
type Employee2 interface {
	work()
}
//seperate taskA and taskB into 2 struct instead of bind them together
type EmployeeTaskA struct {
}
type EmployeeTaskB struct {
}

func (e EmployeeTaskA) work() {
	fmt.Printf("Employee doing task A")
}
func (e EmployeeTaskB) work() {
	fmt.Printf("Employee doing task B")
}
