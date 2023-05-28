// Virtual Machine 8-bit-AON
package main

import (
	"fmt"
)

func main() {
	var a uint8 = 1
	var b uint8 = 0
	var c uint8 = 0
	var d uint8 = 0
	
	fmt.Printf("a    %08b\n", a)
	fmt.Printf("b    %08b\n", b)
	fmt.Printf("c    %08b\n", c)
	fmt.Printf("d    %08b\n", d)
	fmt.Printf("\n")
	
	fmt.Printf("AND  %08b\n", AND(a, b))
	fmt.Printf("OR   %08b\n", OR(a, b))
	fmt.Printf("NOT  %08b\n", NOT(a))
	fmt.Printf("XOR  %08b\n", XOR(a, b))
	fmt.Printf("AND3 %08b\n", AND3(a, b, c))
	fmt.Printf("OR4  %08b\n", OR4in(a, b, c, d))
	
}


