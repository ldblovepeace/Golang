package main

import (
	"fmt"
	"reflect"
)


func main(){
	var i float64 = 3.4

	t := reflect.ValueOf(i)
	fmt.Println("type:", t.Type())
	fmt.Printf("The value of i is %f\n", t.Float())
//下面这段reflect的机制，不能直接用t.SetFloat()修改i的值
	q := reflect.ValueOf(&i)
	v := q.Elem()
	v.SetFloat(1.0)
	fmt.Printf("The value of i is %f now\n", i)
}