package main

import (
	"fmt"
)

func main() {
	fmt.Println(mask(3, 255));
}

// assume A := {0, 1, 2, 3}
func mask(a, b uint8) uint8 {
	b1 := b >> 1;
	switch(a) {
	case 1:
		return (^b1) & b & 0x55;
	case 2:
		return ((^b1) & b | b1 & (^b)) & 0x55;
	case 3:
		return (b1 | b) & 0x55;
	default:
		return 0;
	}
}
