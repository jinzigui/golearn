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

func main() {
	structPeople := new(structTest)
	structPeople.InitStruct(20, "test")
	structPeople.PrintfStruct()
}
