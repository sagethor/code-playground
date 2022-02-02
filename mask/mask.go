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

// col_odd = random & 0x0003ffff
// col_even = random & 0x0000ffff

// puzzle pirates distilling game example
// array should have 10, one per each column
// consider using the final value for gamestate / score
func swapmask(board []uint32) []uint32 {
	swaps [9]uint32;
	c0, c1, h0, h1, l0, l1 uint32;
	for i = 0; i < 9; i++ {
		// even leading column
		if (i & 1) == 0 {
			h0 = board[i] & 0x5555;
			h1 = (board[i] & 0xAAAA) >> 1;
			l0 = board[i+1] & 0x5555;
			l1 = (board[i+1] & 0xAAAA) >> 1;
			c0 = (^h1) & (^h0) & l1 & (^l0);
			h0 = l0; h1 = l1;
			l0 = r0 >> 2 & 0x5555;
			l1 = (r0 >> 2 & 0xAAAA) >> 1;
			c1 = (^h1) & (^h0) & l1 & (^l0);
			swaps[i] = c0 | c1 << 1;
			// odd leading column
		}	else {
			// h0 = board[i] & 0x5555;
			// h1 = (board[i] & 0xAAAA) >> 1
			l0 = board[i+1] >> 2 & 0x5555;
			l1 = (board[i+1] >> 2 & 0xAAAA) >> 1;
			c1 = (^h1) & (^h0) & l1 & (^l0);
			l0 = h0; l1 = h1;
			h0 = board[i+1] & 0x5555;
			h1 = (board[i+1] & 0xAAAA) >> 1;
			c0 = (^h1) & (^h0) & l1 & (^l0);
			swaps[i] = c0 | c1 << 1;
		}
	}
	return swaps;
}

// condensed version (test against above)
// for this and the above we can look at returning []uint16 instead
func condmask(board []uint32) []uint32 {
	swaps [9] uint32;
	c0, h0, h1, l0, l1 uint32;
	for i = 0; i < 9; i++ {
		// even leading column
		if (i & 1) == 0 {
			h0 = board[i] & 0x5555 & (board[i+1] & 0xAAAA) << 1;
			h1 = (board[i] & 0xAAAA) >> 1 & board[i+1] & 0xAAAA;
			l0 = board[i+1] & 0x5555 & (board[i] >> 2 & 0x5555) << 1;
			l1 = (board[i+1] & 0xAAAA) >> 1 & board[i] >> 2 & 0xAAAA;
			swaps[i] = (^h1) & (^h0) & l1 & (^l0);
		// odd leading column
		} else {
			h0 = board[i+1] & 0x5555 & (board[i] & 0x5555) << 1;
			h1 = (board[i+1] & 0xAAAA) >> 1 & board[i] & 0xAAAA;
			l0 = board[i] & 0x5555 & (board[i+1] >> 2 & 0x5555) << 1;
			l1 = (board[i] & 0xAAAA) >> 1 & board[i+1] >> 2 & 0xAAAA;
			swaps[i] = (^h1) & (^h0) & l1 & (^l0);
		}
	}
}
