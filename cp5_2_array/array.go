package main

import "fmt"

func main() {
	//指定陣列長度
	var arr1 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//陣列長度根據元素個數推斷出來
	var arr2 = [...]int{1, 2, 3, 4, 5}

	//短變量形式聲明
	//指定長度
	arr3 := [10]int{1, 2, 3, 4, 5, 6, 7, 8}
	//不指定長度
	arr4 := [...]int{1}

	fmt.Println(len(arr1), len(arr2), len(arr3), len(arr4))

	//訪問arr
	var n [len(arr1)]int
	var i int
	for i = 0; i < len(n); i++ {
		n[i] = i + 100
	}
	//尋訪陣列n
	for i = 0; i < len(n); i++ {
		fmt.Printf("Element[%d] = %d\n", i, n[i])
	}
}
