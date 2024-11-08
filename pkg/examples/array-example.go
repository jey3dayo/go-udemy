package examples

import "fmt"

func ArrayExample() {
	// 整数型の配列を宣言
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

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
