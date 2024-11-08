package examples

import "fmt"

func ArrayExample() {
	var arr1 [3]int
	fmt.Printf("初期化された配列: %v => %T\n", arr1, arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Printf("文字列の配列: %v\n", arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Printf("暗黙的に宣言された配列: %v\n", arr3)

	arr4 := [...]string{"C", "D", "E"}
	fmt.Printf("要素数を省略した配列: %v => %T\n", arr4, arr4)

	// 整数型の配列を宣言
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("配列: %v\n", numbers)

	fmt.Println("arr2", arr2[0], arr2[1], arr2[2])

	arr2[2] = "C"
	fmt.Println(arr2)

	var arr5 [4]int
	// 型が違うのでエラー
	// arr5 = arr1

	fmt.Println("len(arr5) =>", len(arr5))

	// 配列の要素にアクセス
	fmt.Printf("最初の要素: %d\n", numbers[0])
	fmt.Printf("最後の要素: %d\n", numbers[len(numbers)-1])

	// 配列の長さを取得
	length := len(numbers)
	fmt.Printf("配列の長さ: %d\n", length)

	// 配列の要素を変更
	numbers[2] = 35
	fmt.Printf("変更後の配列: %v\n", numbers)

	// 配列の要素をループで表示
	fmt.Println("配列の要素:")
	for i, value := range numbers {
		fmt.Printf("インデックス %d: %d\n", i, value)
	}
}
