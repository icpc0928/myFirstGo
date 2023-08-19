package main

import "fmt"

func main() {
	fmt.Println("HELLO")

	var v1 = "hello" //聲明變量可省略型別
	var v2 = 20
	var v3 = 3.14
	var v4 string //聲明變量暫時不附值 則需先給型別
	var v5 int
	var v6 float64
	v4 = v1
	v5 = v2
	v6 = v3
	fmt.Println(v4, v5, v6)
	var v7, v8, v9 int = 7, 8, 9      //一次聲明多個變量(相同類型 其實可省略型別)
	var v10, v11, v12 = 12, "h", 3.14 //一次聲明多個變量且不同類型
	fmt.Println(v7, v8, v9)
	fmt.Println(v11, v12, v10)

}
