package main
import "fmt"

func main(){
	var array[5] int
	for i:=0; i < len(array); i++{
		array[i] = 2*i
	}
	for i:= 0; i< len(array);i++{
		fmt.Printf("array at index %d is %d\n", i, array[i])
	}
}
