package main

import (
	"fmt"
	"bytes"
	"encoding/gob"
	"log"
)

type testType struct {
	Name	string
	W, X, Y, Z	uint8
}

func main() {
	// testing Golang gobs
	var network bytes.Buffer; // stand-in for network connection
	enc := gob.NewEncoder(&network) // writes to network
	dec := gob.NewDecoder(&network) // reads from network

	err := enc.Encode(testType{"test", 1, 2, 3, 4});
	if err != nil {
		log.Fatal("encode error:", err);
	}

	var test testType;
	err = dec.Decode(&test);
	if err != nil {
		log.Fatal("decode error:", err);
	}

	fmt.Printf("%q: {%d, %d, %d, %d}\n", test.Name, test.W, test.X, test.Y, test.Z);

	// testing lrng8()
	lrng8();
	fmt.Println(lrng8());
}
