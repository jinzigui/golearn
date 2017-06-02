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
	for i,_:=range array{	
		fmt.Printf("array at index %d is %d\n", i, array[i])
	}
	var strarray = [6]string{2:"test", 4:"jj"}
	for i,_:=range strarray{	
		fmt.Printf("strarray at index %d is %s\n", i,strarray[i])
	}
	
	fmt.Printf("sum = %d\n", sum(&array))
}

func sum(array *[5] int)(sum int){
	for _, v:= range array{
		sum += v
	}
	return
}
