// Logic CPU
package main


// 2AND
func AND(a uint8, b uint8) uint8 {
	return a*b
}

// 2OR
func OR(a uint8, b uint8) uint8 {
	if (a + b) != 0 {
		return 1
	} else {
		return 0
	}
}

// NOT
func NOT(a uint8) uint8 {
	return 1 - a
}

// 2XOR
func XOR(a uint8, b uint8) uint8 {
	f1 := NOT(a)
	f2 := NOT(b)
	f3 := AND(f1, b)
	f4 := AND(a, f2)
	f5 := NOT(f3)
	f6 := NOT(f4)
	f7 := AND(f5, f6)
	return NOT(f7)
}

// 3AND
func AND3(a uint8, b uint8, c uint8) uint8 {
	f1 := AND(a, b)
	return AND(c, f1)
}

// 4OR
func OR4in(a uint8, b uint8, c uint8, d uint8) uint8 {
	f1 := OR(a, b)
	f2 := OR(c, d)
	return OR(f1, f2)
}


