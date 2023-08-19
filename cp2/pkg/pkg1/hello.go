package pkg1

var Money = 10 //大寫也是公用的
var xyz = 10

// Add 方法聲明 大寫表示public 公開der
// 後面的int 表示回傳型別
func Add(a, b int) int {
	return addToV(a, b)
}

func addToV(a, b int) int {
	return a + b
}

func AddMoney(a int) int {
	return a + Money
}
func addXyz(a int) int {
	return a + xyz
}
