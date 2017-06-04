package main

import (
	"fmt"
)

type structTest struct {
	nAge    int
	strName string
}

func (structPeople *structTest) InitStruct(nage int, strname string) {
	structPeople.nAge = nage
	structPeople.strName = strname
}

func (structPeople *structTest) PrintfStruct() {
	fmt.Printf("structTest.nAge = %d, structTest.strName = %s\n", structPeople.nAge, structPeople.strName)
}

type intVector []int

func (vector *intVector) Sum() (sum int) {
	for _, num := range *vector {
		sum += num
	}
	return
}

func (vector intVector) Reduction() (result int) {
	for _, num := range vector {
		result -= num
	}
	return
}

func main() {
	structPeople := new(structTest)
	structPeople.InitStruct(20, "test")
	structPeople.PrintfStruct()
	fmt.Println((&intVector{2, 4, 7}).Sum())
	fmt.Println(intVector{3, 5, 8}.Reduction())
}
