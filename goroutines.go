package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(from string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("msg = %s, i= %d", from, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
func main() {
	c := make(chan string)
	go f("gorountine and channel", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("printf msg:%q\n", <-c)
	}
	fmt.Println("done")
}
