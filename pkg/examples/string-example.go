package examples

import "fmt"

func StringExample() {
	// 文字列の基本的な宣言
	var hi string = "Hi"
	var hello string = "こんにちは"
	var world string = "世界"
	var si string = "300"

	fmt.Printf("hello: %s\n", hello)
	fmt.Printf("world: %s\n", world)
	fmt.Printf("hello is %T\n", hello)
	fmt.Printf("si is %T\n", si)

	fmt.Println(`test
		これは
					test`)

	fmt.Println("こちら\"です\"")

	fmt.Println(hi[0])
	fmt.Println(string(hi[0]))

	// 文字列の結合
	greeting := hello + ", " + world + "!"
	fmt.Printf("greeting: %s\n", greeting)

	// 文字列の長さ
	length := len(greeting)
	fmt.Printf("greetingの長さ: %d\n", length)

	// 部分文字列の取得
	substring := greeting[0:5]
	fmt.Printf("部分文字列: %s\n", substring)

	// 文字列の比較
	isEqual := hello == world
	fmt.Printf("hello == world: %v\n", isEqual)

	// 文字列の繰り返し
	repeated := hello + " " + hello
	fmt.Printf("繰り返し: %s\n", repeated)
}
