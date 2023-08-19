package main

import "fmt"

func main() {
	//短變量聲明 :=
	//不需要var 也不需要聲明型別

	v1 := 44
	v2 := "hello"
	v3 := 3.14999
	v4 := 3.14111
	fmt.Println(v1, v2, v3)

	fmt.Printf("%.2f , %.2f", v3, v4) //保留小數點後兩位 四捨五入

}
