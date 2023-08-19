package main

import "fmt"

func main() {
	//強制轉型 型別()

	var sum int = 19
	var count int = 5
	var mean float32
	var i1 int32
	var i2 float32
	mean = float32(sum) / float32(count)
	fmt.Println(mean)
	//i1 = sum / count //這樣是不行的 go 會認為這是有浮點的
	i1 = int32(sum / count)
	fmt.Println(i1)

	i2 = float32(sum / count) //這樣會先將sum/count 取整數 再轉float
	fmt.Println(i2)

	var long1 int64 = 2147483648 //超過int32 最大值: 2147483647
	var int1 int32 = int32(long1)
	fmt.Println(int1) //會影響精度
}
