package examples

import "fmt"

func IntExample() {
	i := 3
	fmt.Println(i)

	var i2 int64 = 100
	fmt.Println(i2)

	// 型が違うのでエラー
	// fmt.Printf("%d\n", i+i2)

	// 型を合わせる
	fmt.Printf("%d\n", i+int(i2))

	fmt.Printf("i: %T, i2: %T\n", i, i2)
	fmt.Printf("i+i2->%v: %T\n", int64(i)+i2, int64(i)+i2)
}
