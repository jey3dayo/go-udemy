package examples

import "fmt"

func BooleanExample() {
	// 基本的なブール値の宣言
	// 変数の宣言を修正
	var isTrue bool = true
	var isFalse bool = false

	fmt.Printf("isTrue: %v\n", isTrue)
	fmt.Printf("isFalse: %v\n", isFalse)

	// 論理演算
	and := isTrue && isFalse
	or := isTrue || isFalse                            // 修正: t と f を isTrue と isFalse に変更
	not := !isTrue                                     // 修正: t を isTrue に変更
	xor := (isTrue || isFalse) && !(isTrue && isFalse) // XOR演算の追加

	fmt.Printf("AND演算 (isTrue && isFalse): %v\n", and)
	fmt.Printf("OR演算 (isTrue || isFalse): %v\n", or)
	fmt.Printf("NOT演算 (!isTrue): %v\n", not)
	fmt.Printf("XOR演算 ((isTrue || isFalse) && !(isTrue && isFalse)): %v\n", xor)

	// 比較演算
	x := 5
	y := 10
	z := 20
	isGreater := x > y
	isEqual := x == y
	isNotEqual := x != y
	isLessThan := x < z

	fmt.Printf("x > y: %v\n", isGreater)
	fmt.Printf("x == y: %v\n", isEqual)
	fmt.Printf("x != y: %v\n", isNotEqual)
	fmt.Printf("x < z: %v\n", isLessThan)
	fmt.Printf("x >= y: %v\n", x >= y) // x が y 以上であるかを確認
}
