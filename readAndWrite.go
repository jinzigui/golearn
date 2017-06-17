package main

import (
	"bufio"
	"fmt"
	"os"
)

type peopleInfo struct {
	nAge    int
	fHegith float32
	fWegith float32
	strName string
}

func (people *peopleInfo) printPeopleInfo() {
	fmt.Printf("your name is %s, ages = %d, hegith = %f, wegith = %f\n",
		people.strName, people.nAge, people.fHegith, people.fWegith)
}

func bufferReadTest() string {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("input people some base info")
	msg, err := inputReader.ReadString('\n')
	if err != nil {

	}
	return msg
}

func main() {
	people := new(peopleInfo)
	fmt.Println("input peopel some base info")
	fmt.Println("now input your name:")
	fmt.Scanf("%s", &people.strName)
	fmt.Println("input your ages:")
	fmt.Scanln(&people.nAge)
	fmt.Println("input your hegith:")
	fmt.Scanln(&people.fHegith)
	fmt.Println("input your wegith:")
	fmt.Scanln(&people.fWegith)
	people.printPeopleInfo()

	fmt.Println(bufferReadTest())
}
