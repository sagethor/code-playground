package lrng8

import (
	"testing"
)

func TestLrng8_32(t *testing.T) {
	lrng8_32();
	avg := 0.0;
	for i := 0; i < len(buffer); i++ {
		avg += float64(buffer[i]);
	}
	avg = avg / 256;
	// less than .007% probability of average less than 110, assuming normal distribution. 
	if avg < 110 {
		t.Errorf("lrng8_32() average is improbably low: %f", avg);
	}
	// less than .007% probability of average more than 145, assuming normal distribution.
	if avg > 145 {
		t.Errorf("lrng8_32() average is improbably high: %f", avg);
	}
}

func TestLrng8_64(t *testing.T) {
	lrng8_64();
	avg := 0.0;
	for i := 0; i < len(buffer); i++ {
		avg += float64(buffer[i]);
	}
	avg = avg / 256;
	// less than .007% probability of average less than 110, assuming normal distribution. 
	if avg < 110 {
		t.Errorf("lrng8_64() average is improbably low: %f", avg);
	}
	// less than .007 probability of average more than 145, assuming normal distribution.
	if avg > 145 {
		t.Errorf("lrng8_64() average is improbably high: %f", avg);
	}
}
