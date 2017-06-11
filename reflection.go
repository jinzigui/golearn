package main

import (
	"fmt"
	"reflect"
)

func reflectTest() {
	var nNum int = 10
	var strName string = "酱酱"
	v := reflect.ValueOf(nNum)
	fmt.Printf("nNum value = %v\n", v)
	fmt.Printf("nNum type = %v\n", v.Type()) //type() == kind() ?
	fmt.Printf("nNum kind = %v\n", v.Kind())
	fmt.Printf("nNum interface = %v\n", v.Interface())
	fmt.Printf("nNum canset = %v\n", v.CanSet())
	str := reflect.ValueOf(&strName).Elem()
	fmt.Printf("strName = %v\n", str.Interface())
	fmt.Printf("strName canset = %v\n", str.CanSet())
	str.SetString("渣渣酱")
	fmt.Printf("after setstring = %v\n", str)
}

func main() {
	reflectTest()
}
