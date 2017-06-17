package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
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

func scanfTest() {
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
}

func bufferReadTest() string {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("input people some base info")
	msg, err := inputReader.ReadString('\n')
	if err != nil {

	}
	return msg
}

func readFileTest() {
	file, openFileerr := os.Open("test.txt")
	if openFileerr != nil {
		fmt.Println("open file failure")
		return
	}

	fileReader := bufio.NewReader(file)
	for {
		strText, readerr := fileReader.ReadString('\n')
		fmt.Println(strText)
		if readerr == io.EOF {
			break
		}
	}
	file.Close()
	//os open file, bufio.NewReader control file reader, reader read text

	//another read way
	buf, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("open file failure")
	}
	fmt.Printf("%s", string(buf))
}

func writeFileTest() {
	file, fileerr := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if fileerr != nil {
		fmt.Println("open file failure")
		return
	}
	writeReader := bufio.NewWriter(file)
	strMsg := "just write test\n"
	writeReader.WriteString(strMsg)
	writeReader.Flush()

	file.WriteString("file control writestring fun\n") //call writeString func
	file.Close()
}

func main() {
	scanfTest()
	fmt.Println(bufferReadTest())
	writeFileTest()
	readFileTest()
}
