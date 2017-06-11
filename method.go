package main

import (
	"fmt"
)

type structPeople struct {
	nAge    int
	strName string
}

func (structPeople *structPeople) InitStructPeople(nage int, strname string) {
	structPeople.nAge = nage
	structPeople.strName = strname
}

func (structPeople *structPeople) PrintfStructPeople() {
	fmt.Printf("structPeople.nAge = %d, structPeople.strName = %s\n", structPeople.nAge, structPeople.strName)
}

type structStudent struct {
	structPeople
	nNumID int
	fScore float32
}

func (structStudent *structStudent) InitStructStudent(nage int, strname string, nnumid int, fscore float32) {
	structStudent.InitStructPeople(nage, strname)
	structStudent.nNumID = nnumid
	structStudent.fScore = fscore
}

func (structStudent *structStudent) PrintfStructStudent() {
	structStudent.PrintfStructPeople()
	fmt.Printf("structStudent.nNumID = %d, structStudent.fScore = %f\n", structStudent.nNumID, structStudent.fScore)
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

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func main() {
	structPeople1 := new(structPeople)
	structPeople1.InitStructPeople(20, "test")
	structPeople1.PrintfStructPeople()

	fmt.Println((&intVector{2, 4, 7}).Sum())
	fmt.Println(intVector{3, 5, 8}.Reduction())

	structStudent1 := &structStudent{structPeople{22, "structStudent"}, 1200350116, 99.5}
	structStudent1.PrintfStructStudent()
	structStudent1.PrintfStructPeople()

	structStudent2 := new(structStudent)
	structStudent2.InitStructStudent(23, "酱酱", 1200350116, 99.5)
	structStudent2.PrintfStructStudent()

	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}
