package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	strTest := "mac ip = 127.0.0.1, window ip = 192.168.0.1, iphone ip = 0.0.0.1"
	regexpRule := "[0-9]+.[0-9]+"
	f := func(str string) string {
		vaule, _ := strconv.ParseFloat(str, 32)
		return strconv.FormatFloat(vaule*2, 'f', 2, 32)
	}
	if match, _ := regexp.Match(regexpRule, []byte(strTest)); match {
		fmt.Println("macth")
	}
	recmp, _ := regexp.Compile(regexpRule)
	str1 := recmp.ReplaceAllString(strTest, "110")
	fmt.Println(str1)
	fmt.Println(f(strTest))
	str2 := recmp.ReplaceAllStringFunc(strTest, f)
	fmt.Println(str2)
}
