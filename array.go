package main

import (
	"fmt"
	"sort"
)

func main() {
	var array [5]int
	for i := 0; i < len(array); i++ {
		array[i] = 2 * i
	}
	for i := 0; i < len(array); i++ {
		fmt.Printf("array at index %d is %d\n", i, array[i])
	}
	for i, _ := range array {
		fmt.Printf("array at index %d is %d\n", i, array[i])
	}
	var strarray = [6]string{2: "test", 4: "jj"}
	for i, _ := range strarray {
		fmt.Printf("strarray at index %d is %s\n", i, strarray[i])
	}

	fmt.Printf("sum = %d\n", sum(&array))
	slices()
	makeslice()
	for_rangetest()
	sliceCopyAndAppend()
	stringChangeIndexChar()
	SliceSort()
}

func sum(array *[5]int) (sum int) {
	for _, v := range array {
		sum += v
	}
	return
}

func slices() {
	var array = [6]int{1, 4, 2, 6, 7, 3}
	var slice []int = array[1:4]
	//var slice2 = []int{2,4,7}
	for i, _ := range slice {
		fmt.Printf("slice index = %d, value = %d\n", i, slice[i])
	}
}

func makeslice() {
	var array = [6]int{1, 4, 2, 6, 7, 3}
	var slice []int = array[1:4]
	printslice(slice, true)
	var slice1 []int = make([]int, 10)
	printslice(slice1, true)
	var slice2 []int = make([]int, 5, 10)
	printslice(slice2, true)
}

func printslice(slice []int, bNeedInit bool) {
	if bNeedInit {
		for i := 0; i < len(slice); i++ {
			slice[i] = 2 * i
		}
	}

	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d]= %d\n", i, slice[i])
	}
}

func for_rangetest() {
	var slice []int = make([]int, 5)
	for _, val := range slice {
		val += 2
	}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d]=%d\n", i, slice[i])
	}
}

func sliceCopyAndAppend() {
	slice_form := []int{1, 3, 2}
	slice_to := make([]int, 10)
	n := copy(slice_to, slice_form)
	fmt.Println(slice_to)
	fmt.Println(n)
	slice_form = append(slice_form, 4, 6, 3)
	fmt.Println(slice_form)

}

func stringChangeIndexChar() {
	str := "golearn"
	byByte := []byte(str)
	byByte[len(byByte)-1] = 'm'
	str = string(byByte)
	fmt.Println(str)
	for _, char := range byByte {
		fmt.Printf("%c----", char)
	}
	fmt.Println()
}

func SliceSort() {
	slice := []int{8, 4, 9, 7, 6}
	sort.Ints(slice)
	printslice(slice, false)
	fmt.Println(sort.IntsAreSorted(slice))
}
