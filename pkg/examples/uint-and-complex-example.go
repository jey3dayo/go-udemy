package examples

import "fmt"

func UintAndComplexExample() {
	// uint の例
	var ui uint = 42
	var ui8 uint8 = 255 // 0 から 255 まで
	var ui16 uint16 = 65535
	var ui32 uint32 = 4294967295
	var ui64 uint64 = 18446744073709551615

	fmt.Printf("uint: %d (%T)\n", ui, ui)
	fmt.Printf("uint8: %d (%T)\n", ui8, ui8)
	fmt.Printf("uint16: %d (%T)\n", ui16, ui16)
	fmt.Printf("uint32: %d (%T)\n", ui32, ui32)
	fmt.Printf("uint64: %d (%T)\n", ui64, ui64)

	// 複素数の例
	var c64 complex64 = 3.14 + 2i
	c128 := complex(1.23, 4.56) // complex128 型

	fmt.Printf("complex64: %v (%T)\n", c64, c64)
	fmt.Printf("complex128: %v (%T)\n", c128, c128)

	// 複素数の実部と虚部
	fmt.Printf("実部: %f\n", real(c128))
	fmt.Printf("虚部: %f\n", imag(c128))

	// 複素数の演算
	sum := c64 + complex64(c128)
	fmt.Printf("複素数の和: %v\n", sum)

	// 実数と複素数の和
	sum2 := c64 + 2.0
	fmt.Printf("複素数の和2: %v\n", sum2)

	sum3 := c64 + 2.0i
	fmt.Printf("複素数の和3: %v\n", sum3)

	sum4 := c64 + (1.0 + 2.1i)
	fmt.Printf("複素数の和4: %v\n", sum4)
}
