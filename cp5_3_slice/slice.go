package main

import "fmt"

func main() {
	//切片 slice

	//聲明為數組 只是無須指定其大小
	//使用make 方法

	//make:  make([]T, len, cap)     T:切片數據類型, len: 切片長度, cap: 切片容量

	//切片
	strSlice1 := []string{"hello", "world", "!"}
	intSlice1 := []int{1, 2, 3, 4, 5, 6}
	//使用make創建
	var strSlice2 = make([]string, 10, 20) //指定容量20
	var intSlice2 = make([]int, 2)
	intSlice2 = append(intSlice2, 1)

	fmt.Println(len(strSlice1), len(intSlice1), len(strSlice2), len(intSlice2))
	fmt.Println(cap(strSlice2)) //容量
	fmt.Println(cap(intSlice2))

	var cap2 int = 0
	var tempCap = cap(intSlice2)
	limit := 100000
	for cap2 < limit {
		intSlice2 = append(intSlice2, 1)
		if tempCap != cap(intSlice2) {
			fmt.Println(cap(intSlice2), len(intSlice2))
		}
		tempCap = cap(intSlice2)
		cap2++
	}

}
