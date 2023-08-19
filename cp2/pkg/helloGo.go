package main

import (
	"awesomeProject/cp2/pkg/pkg1"
	"awesomeProject/cp2/pkg/pkg2"
	"fmt"
)

func main() {
	var pkg1Money = pkg1.Money
	fmt.Println(pkg1Money)
	var pkg2Money = pkg2.Money
	fmt.Println(pkg2Money)
	//fmt.Println(pkg1.xyz) //can't do this
	fmt.Println(pkg1.Add(pkg1Money, pkg2Money))
	fmt.Println(pkg1.AddMoney(190))
}
