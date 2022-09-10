package Kalkulator

import(
	"fmt"
)
func Tambah(a,b int) int{
	return a + b
}
func Kurang(a,b int) int{
	return a - b
}
func Kali(a,b int) int{
	return a * b
}
func Bagi(a,b float64) float64{
	return a / b
}

func main(){
	result1 := Tambah(5,3) 
	result2 := Kurang(5,3) 
	result3 := Kali(5,3) 
	result4 := Bagi(5,3) 

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
}
