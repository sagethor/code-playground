package main

import (
	"fmt"
	"encoding/gob"
	"log"
	"os"
	"./lrng8"
)

type Test struct {
	Name	string
	W, X, Y, Z	uint8
	Tag [16777216]bool
}

// define decodeToTest function below

func main() {
	// initialization
	lrng8.Rand();

	// testing
	var test Test;
	test.Name = "test";
	test.W = 1;
	test.X = 2;
	test.Y = 3;
	test.Z = 4;
	fmt.Printf("%q: {%d, %d, %d, %d}\n", test.Name, test.W, test.X, test.Y, test.Z);

	// finish creating cereal module and use function to test

}
