package lrng8

import "time"
import "encoding/binary"
import "math/bits"
// consider getting rid of these private globals?
var counter uint8 = 255;
var buffer [256]uint8;

// initial filling of buffer
func init() {
	Rand();
}

func nanow() int64 {
	return time.Now().UnixNano();
}

func Rand() uint8 {
	if (counter == 255) {
		if (bits.UintSize == 64) {
			defer lrng8_64();
		} else {
			defer lrng8_32();
		}
	}
	return buffer[counter];
}

var seed32 uint32 = uint32(nanow());
func lrng8_32() {
	_ = buffer[255];
	for i := 0; i < 64; i++ {
		seed32 *= 69069;
		seed32++;
		binary.LittleEndian.PutUint32(buffer[i*4:], seed32);
	}
}

var seed64 uint64 = uint64(nanow());
func lrng8_64() {
	_ = buffer[255];
	for j := 0; j < 32; j++ {
		seed64 *= 6364136223846793005;
		seed64++;
		binary.LittleEndian.PutUint64(buffer[j*8:], seed64);
	}
}
