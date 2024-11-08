package examples

import (
	"fmt"
	"time"
)

func printBorder() {
	fmt.Println("-----------")
}

func getCurrentTime() string {
	return time.Now().Format(time.DateTime)
}

// 暗黙的な定義は関数内でしかできない
// var i5 := 500
var i5 int = 500

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func OtherExample() {
	printBorder()
	fmt.Println(getCurrentTime())
	printBorder()

	var i int = 100
	fmt.Println(i)

	var s string = "Hello, World!"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Python"
	fmt.Println(i3, s3)

	// 暗黙的な定義
	// 変数名 := 値
	i4 := 150
	fmt.Println(i4)

	i4 = 400
	fmt.Println(i4)

	// 再定義できない
	// i4 := 500
	// fmt.Println(i4)

	// 型が違うのでエラー
	// i4 = "hello"
	// fmt.Println(i4)

	fmt.Println(i5)

	outer()

	// 関数外では使用できない
	// fmt.Println(s4)

	// 未使用の変数はエラー
	// s5 := "Not Use"
}
