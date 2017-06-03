package main

import (
	"fmt"
)

func mapTest() {
	var map1 map[int]int
	map1 = map[int]int{1: 9, 2: 8, 3: 7}
	printmap(map1)
	if _, isNotNil := map1[3]; isNotNil {
		delete(map1, 3)
		printmap(map1)
	}
	map2 := make(map[int]int)
	map2 = map1
	printmap(map2)
}

func printmap(maplist map[int]int) {
	for key, value := range maplist {
		fmt.Printf("maplist[%d] = %d\n", key, value)
	}
}

func main() {
	mapTest()
}
