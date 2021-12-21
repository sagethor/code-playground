var buffC uint8;
var buff [256]byte;

func lrng8() uint8 {
	if (buffC == 255) {
		if (bitsUintSize == 64) {
			defer lrng8_64();
		} else {
			defer lrng8_32();
		}
	}
		return buff[buffC];
}

var seed32 uint32 = uint32(nanow());
func lrng8_32() {
	_ = buff[255];
	for i := 0; i < 64; i++ {
		seed *= 69069;
		seed32++;
		binary.LittleEndian.PutUint32(buff[i*4:], seed32);
	}
}

var seed64 uint64 = uint64(nanow());
func lrng8_64() {
	_ = buff[255];
	for i := 0; i < 32; i++ {
		seed64 *= 6364136223846793005;
		seed64++;
		binary.LittleEndian.PutUint64(buff[j*8:], seed64);
	}
}
