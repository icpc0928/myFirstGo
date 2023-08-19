package main

import (
	//自定義包的別名
	pk1 "awesomeProject/cp2/pkg/pkg1"
	"fmt"

	pk2 "awesomeProject/cp2/pkg/pkg2"
)

func main() {

	var p1 = pk1.Money
	//var pkg2Money = pkg2.Money		//這樣就找不到這個pkg2了 因為被pk2替代
	var p2 = pk2.Money
	fmt.Println(p1, p2)

}
