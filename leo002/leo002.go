package main

import "fmt"

func main() {
	var v1 = 20
	var v2 = "hello, world"
	var v3 = 10.12345
	fmt.Printf("%d\n", v1)
	fmt.Printf("變量類型: %T\n", v1) //顯示型別
	fmt.Printf("%s\n", v2)
	fmt.Printf("變量類型: %T\n", v2) //顯示型別
	fmt.Printf("v3: %s \n", v3)  //浮點藥用%f 型別對不上print會顯示他是啥型別
	fmt.Printf("v3: %f \n", v3)
	fmt.Printf("變量類型: %T\n", v3) //顯示型別
}
