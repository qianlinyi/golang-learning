package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v1 := reflect.ValueOf(&x)
	v2 := reflect.ValueOf(x)
	t := reflect.TypeOf(&x)
	fmt.Println(v1, v2.Interface(), v2, t)
	defer func() { // go使用defer和recover处理panic
		if r := recover(); r != nil {
			fmt.Println(r)
			v1.Elem().SetFloat(7.1)
			fmt.Println("x:", v1.Elem().Interface(), v1.Elem())
		}
	}()
	v2.SetFloat(7.1) // panic 对应反射第三定律
}
