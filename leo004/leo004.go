package main

import "fmt"

func main() {
	//常量 const 不得修改(只讀) like final

	const v1 = int32(134)
	const v2 = "222"
	const v3 = 2456 //不呼叫不會有事

	fmt.Println(v1)
	//v1 = 345 //can't do this
	fmt.Println(v1)
	fmt.Println(v2)

}
