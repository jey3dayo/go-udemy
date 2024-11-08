package examples

import "fmt"

func ByteExample() {
	// バイトスライスの宣言
	var byteA []byte = []byte{72, 73}
	byteB := []byte{74, 75}

	// byteA
	fmt.Println(byteA)
	fmt.Println(string(byteA))

	// byteB
	fmt.Println(string(byteB))

	// byteC
	c := []byte("Hiyo")
	fmt.Println(c)
	fmt.Println(string(c))

	// バイトスライスを文字列に変換
	data := []byte{72, 73, 74, 75}
	str := string(data)
	fmt.Printf("文字列: %s\n", str)

	// 文字列をバイトスライスに変換
	original := "こんにちは"
	bytes := []byte(original)
	fmt.Printf("バイトスライス: %v\n", bytes)

	// バイトスライスの長さ
	length := len(bytes)
	fmt.Printf("バイトスライスの長さ: %d\n", length)

	// バイトの比較
	isEqual := data[0] == bytes[0]
	fmt.Printf("最初のバイトが等しいか: %v\n", isEqual)

	// バイトスライス結合
	combined := append(data, bytes...)
	fmt.Printf("結合されたバイトスライス: %v\n", combined)
	fmt.Printf("結合された文字列: %v\n", string(combined))
}
