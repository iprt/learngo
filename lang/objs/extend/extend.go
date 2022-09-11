package main

import "fmt"

type human struct {
	name  string
	age   int
	phone string
}

func (s *human) sayHi() {
	fmt.Printf("name is %s ;age is %d;phone is %s\n", s.name, s.age, s.phone)
}

type student struct {
	// 继承！！！
	human
	school string
}

type employee struct {
	human
	company string
}

// 重写！！！
func (e *employee) sayHi() {
	fmt.Println("before rewrite")
	e.human.sayHi()
	fmt.Println("after rewrite")
}

func main() {
	mark := student{human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := employee{human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.sayHi()
	fmt.Println()
	sam.sayHi()
}
