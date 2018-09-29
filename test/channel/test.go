package main

import (
	"fmt"
)


func sum(a []int, c chan int){
	sum := 0
	for _, v := range a{
		sum += v
	}
	c <- sum
}

func main(){
	a := []int{1,2,3,4,5,6,7}

	c := make(chan int)
	// c := make(chan int, 1)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	c1 := make(chan int, 3)

	c1 <- 1
	c1 <- 2
	fmt.Println(<-c1)
	fmt.Println(<-c1)
}