package examples

import (
	"fmt"
	"strconv"
)

func CastExample() {
	var i int = 1
	fl64 := float64(i)
	fmt.Printf("i => %T \n", i)
	fmt.Printf("f64 => %T \n", fl64)

	i2 := int(fl64)
	fmt.Printf("f64 => %T \n", i2)

	fl := 10.5
	i3 := int(fl)
	fmt.Printf("i3 => %T \n", i3)
	fmt.Println(i3)

	var s string = "100"
	fmt.Printf("s => %T \n", s)

	i4, _ := strconv.Atoi(s)
	fmt.Printf("i4 => %T \n", i4)
}
