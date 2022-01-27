package mask

import (
	"fmt"
)

// assume A := {0, 1, 2, 3}
// extend for uint64
func hillmask(a, b uint8) uint8 {
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

// puzzle pirates distilling game example
// array should have 10, one per each column
// consider using the final value for gamestate / score
func swapmask(arr []uint64) {
	for i = 0; i < 9; i++ {
		h, l uint64;
		if (i & 1) == 1 {
			h = arr[i+1];
			l = arr[i];
		} else {
			h = arr[i];
			l = arr[i+1];
		}
		// split up state vars for readability
		h0 := h & 0x5555;
		h1 := h & 0xb0b0;
		l0 := l & 0x5555;
		l1 := l & 0xb0b0;

		c := (^h1) & (^h0) & l1 & (^l0) |
		(^h1) & h0 & (^l1) & (^l0) |
		h1 & (^h0) & (^l1) & l0;

		// okay this is probably being a little too cute here
		// we store the swapmask in the upper 32, raw data only uses lower 32
		arr[i] = arr[i] & 0xffff | c;
	}
}
