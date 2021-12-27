package cereal

import (
	"testing"
	"encoding/gob"
	"log"
	"bytes"
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

// this is an integration test
func TestCereal(t *testing.T) {
	var test Test;
	test.Name = "test";
	test.W = 1;
	test.X = 2;
	test.Y = 3;
	test.Z = 4;

	testWrite := EncodeToBytes(test);
	testWrite = Compress(testWrite);
	WriteFile(testWrite, "test.dat");

	testRead := ReadFile("test.dat");
	testRead = Decompress(testRead);
	tested := decodeToTest(testRead);

	if tested.Name != "test" {
		t.Errorf("encode-compress-write-read-decompress-decode failed, expected string test, got: %q", tested.Name);
	}
	if tested.W != 1 {
		t.Errorf("encode-compress-write-read-decompress-decode failed, expected W = 1, got: %d", tested.W);
	}
	if tested.X != 2 {
		t.Errorf("encode-compress-write-read-decompress-decode failed, expected X = 2, got: %d", tested.X);
	}
	if tested.Y != 3 {
		t.Errorf("encode-compress-write-read-decompress-decode failed, expected Y = 3, got: %d", tested.Y);
	}
	if tested.Z != 4 {
		t.Errorf("encode-compress-write-read-decompress-decode failed, expected Z = 4, got: %d", tested.Z);
	}
}
