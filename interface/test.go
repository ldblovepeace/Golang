package main

import(
	"fmt"
)

type Human struct{
	name string
	age int
	phone string
}

type Employee struct{
	Human
	company string
	money float32
}

type Student struct{
	Human
	school string
	loan float32
}

func (h Human) Sayhi(){
	fmt.Printf("Hi, i'm %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string){
	fmt.Println("lalalala...", lyrics)
}

func (e Employee) Sayhi(){
	fmt.Printf("Hi, i'm %s, work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

type Men interface{
	Sayhi()
	Sing(lyrics string)
}

func main(){
	mike := Student{Human{"Mike", 25, "123456777"}, "MIT", 0.00}
	// paul := Student{Human{"Paul", 21, "123456788"}, "Harvard", 100}
	// sam := Employee{Human{"Sam", 30, "123456789"}, "Golang inc.", 1000}
	// tom := Employee{Human{"Tom", 36, "123456711"}, "Golang inc.", 1000}

	var i Men

	i = mike
	i.Sayhi()
	i.Sing("qwerty")
}