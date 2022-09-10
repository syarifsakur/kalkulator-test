package main

import(
	"fmt"
)
func tambah(a,b int) int{
	return a + b
}
func kurang(a,b int) int{
	return a - b
}
func kali(a,b int) int{
	return a * b
}
func bagi(a,b int) int{
	return a / b
}

func main(){
	result1 := tambah(5,3) 
	result2 := kurang(5,3) 
	result3 := kali(5,3) 
	result4 := bagi(5,3) 

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
}
