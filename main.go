package main

import (
	"fmt"
	"encoding/gob"
	"log"
	"bytes"
	"./lrng8"
	"./cereal"
)

type Test struct {
	Name	string
	W, X, Y, Z	uint8
	Tag [16777216]bool
}

// define decodeToTest function below
func decodeToTest(s []byte) Test {
	t := Test{};
	dec := gob.NewDecoder(bytes.NewReader(s));
	err := dec.Decode(&t);
	if err != nil {
		log.Fatal(err);
	}
	return t;
}

func main() {
	fmt.Println(lrng8.Rand());

	var test Test;
	test.Name = "test";
	test.W = 1;
	test.X = 2;
	test.Y = 3;
	test.Z = 4;
	fmt.Printf("%q: {%d, %d, %d, %d}\n", test.Name, test.W, test.X, test.Y, test.Z);

	testWrite := cereal.EncodeToBytes(test);
	testWrite = cereal.Compress(testWrite);
	cereal.WriteFile(testWrite, "test.dat");

	testRead := cereal.ReadFile("test.dat");
	testRead = cereal.Decompress(testRead);
	tested := decodeToTest(testRead);

	fmt.Printf("%q: {%d, %d, %d, %d}\n", tested.Name, tested.W, tested.X, tested.Y, tested.Z);
}
