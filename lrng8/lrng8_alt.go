package lrng8

import (
	"encoding/binary"
)
// this file adds alternative multipliers for use

func A_lrng8_64() {
	_ = buffer[255];
	for j := 0; j < 32; j++ {
		seed64 *= 2862933555777941757;
		seed64++;
		binary.LittleEndian.PutUint64(buffer[j*8:], seed64);
	}
}

func B_lrng8_64() {
	_ = buffer[255];
	for j := 0; j < 32; j++ {
		seed64 *= 3202034522624059733;
		seed64++;
		binary.LittleEndian.PutUint64(buffer[j*8:], seed64);
	}
}

func C_lrng8_64() {
	_ = buffer[255];
	for j := 0; j < 32; j++ {
		seed64 *= 3935559000370003845;
		seed64++;
		binary.LittleEndian.PutUint64(buffer[j*8:], seed64);
	}
}
