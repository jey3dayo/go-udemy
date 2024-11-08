package examples

import "fmt"

func FloatExample() {
	// 基本的な float64 の宣言と使用
	fl64 := 3.14
	var fl32 float32 = 1.23 // float32型はあまり使わない

	fmt.Printf("fl64: %f (%T)\n", fl64, fl64)
	fmt.Printf("fl32: %f (%T)\n", fl32, fl32)

	// 型変換の例
	sum := fl64 + float64(fl32)
	fmt.Printf("sum: %f\n", sum)

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Printf("pinf: %v (%T)\n", pinf, pinf)

	ninf := -1.0 / zero
	fmt.Printf("ninf: %v (%T)\n", ninf, ninf)

	nan := zero / zero
	fmt.Printf("nan: %v (%T)\n", nan, nan)

	// 科学的記法
	scientific := 1.23e-4
	fmt.Printf("scientific: %g\n", scientific)

	// 精度の違いを示す
	var precise float64 = 0.1 + 0.2
	fmt.Printf("0.1 + 0.2 = %.20f\n", precise) // 浮動小数点数の精度限界を示す
	fmt.Printf("0.1 + 0.2 = %.4f\n", precise)  // 精度を4桁にして表示
}
