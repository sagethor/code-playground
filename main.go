package main

import (
	"fmt"
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
	var test testType;
	test.Name = "test";
	test.W = 1;
	test.X = 2;
	test.Y = 3;
	test.Z = 4;

	// we can gob to any general thing...
	file, err := os.Create(test.Name + ".gob");
	if err != nil {
		log.Fatal("file create error:", err);
	}
	enc := gob.NewEncoder(file);
	enc.Encode(test);
	file.Close();
	fmt.Printf("%q: {%d, %d, %d, %d}\n", test.Name, test.W, test.X, test.Y, test.Z);

	// recycling, use = instead of :=
	file, err = os.Open(test.Name + ".gob");
	if err != nil {
		log.Fatal("file open error:", err);
	}
	dec := gob.NewDecoder(file);
	err = dec.Decode(&test);
	if err != nil {
		log.Fatal("decode error:", err);
	}
	file.Close();
	fmt.Printf("%q: {%d, %d, %d, %d}\n", test.Name, test.W, test.X, test.Y, test.Z);

	// testing lrng8()
	lrng8();
	fmt.Println(lrng8());
}
