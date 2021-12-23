package main

import (
	"fmt"
	"bytes"
	"encoding/gob"
	"log"
	"os"
)

type testType struct {
	Name	string
	W, X, Y, Z	uint8
	Tag [16777216]bool
}

func main() {
	// testing Golang gobs
	var network bytes.Buffer; // stand-in for network connection
	enc := gob.NewEncoder(&network) // writes to network
	dec := gob.NewDecoder(&network) // reads from network

	var a [16777216]bool;
	err := enc.Encode(testType{"test", 1, 2, 3, 4, a});
	if err != nil {
		log.Fatal("encode error:", err);
	}

	var test testType;
	err = dec.Decode(&test);
	if err != nil {
		log.Fatal("decode error:", err);
	}

	fmt.Printf("%q: {%d, %d, %d, %d}\n", test.Name, test.W, test.X, test.Y, test.Z);

	// test gobs to and from file (okay we're assuming the above worked & recycling)
	file, err := os.Create("test.gob");
	if err != nil {
		log.Fatal("file create error:", err);
	}

	enc = gob.NewEncoder(file);
	enc.Encode(test);

	file.Close();

	var get testType;

	file, err = os.Open("test.gob");
	if err != nil {
		log.Fatal("file open error:", err);
	}
	
	dec = gob.NewDecoder(file);
	err = dec.Decode(&get);
	if err != nil {
		log.Fatal("decode error:", err);
	}

	file.Close();

	//fmt.Println(get);

	// testing lrng8()
	lrng8();
	fmt.Println(lrng8());
}
