package main

import "fmt"

func main() {
	//指針
	var x int = 1000
	fmt.Printf("x = %d \n", x)
	fmt.Printf("x 內存地址: %x\n", &x)

	var ptr = &x
	fmt.Printf("指針變量的地址: %x \n", ptr)
	fmt.Printf("指針變量的值是: %d \n", *ptr)
	var ptrptr = &ptr
	fmt.Printf("變量2的地址: %x \n", ptrptr)
	fmt.Printf("變量2的地址: %x \n", &ptrptr)
	var ptr3 = *ptrptr
	fmt.Printf("變量3的值: %d \n", *ptr3)

	//空指針
	//如果指針變量沒有初始化, 那麼它的內存地址為0
	//表示變量沒有分配內存空間, 這就是空指針
	var emt *int
	fmt.Printf("指針emt 的值是: %x \n", emt)
	//判斷指針是否回空指針
	if emt == nil {
		fmt.Println(" emt == nil")
	}

	//二級指針
	var p1 int
	var p2 *int
	var p3 **int
	var p4 ***int
	p1 = 300
	p2 = &p1 //獲取p1地址 一級指針
	p3 = &p2 //獲取p2地址 二級指針
	p4 = &p3 //三級指針

	fmt.Printf("p1: %d\n", p1)
	fmt.Printf("p2: %x\n", p2)
	fmt.Printf("*p2: %d\n", *p2)
	fmt.Printf("p3: %x\n", p3)
	fmt.Printf("*p3: %x\n", *p3)
	fmt.Printf("**p3: %d\n", **p3)
	fmt.Printf("p4: %x\n", p4)
}
