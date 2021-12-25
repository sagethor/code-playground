package cereal

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func EncodeToBytes(p interface{}) []byte {
	buf := bytes.Buffer{};
	enc := gob.NewEncoder(&buf);
	err := enc.Encode(p);
	if err != nil {
		log.Fatal(err);
	}
	fmt.Println("uncompressed size (bytes): ", len(buf.Bytes()));
	return buf.Bytes();
}
func Compress(s []byte) []byte {
	buf := bytes.Buffer{};
	zip := gzip.NewWriter(&buf);
	zip.Write(s);
	zip.Close();
	fmt.Println("compressed size (bytes): ", len(zip.Bytes()));
	return zip.Bytes();
}
func Decompress(s []byte) []byte {
	rdr, _ := gzip.NewReader(bytes.NewReader(s));
	data, err := ioutil.ReadAll(rdr);
	if err != nil {
		log.Fatal(err);
	}
	rdr.Close();
	fmt.Println("uncompressed size (bytes): ", len(data));
	return data;
}

/*
func DecodeToStructTest(s []byte) Test {
	t := Test{};
	dec := gob.NewDecoder(bytes.NewReader(s));
	err := dec.Decode(&t);
	if err != nil {
		log.Fatal(err);
	}
	return t;
}
*/

func WriteFile(s []byte, file string) {
	f, err := os.Create(file);
	if err != nil {
		log.Fatal(err);
	}
	f.Write(s);
}
func ReadFile(path string) []byte {
	f, err := os.Open(path);
	if err != nil {
		log.Fatal(err);
	}
	data, err := ioutil.ReadAll(f);
	if err != nil {
		log.Fatal(err);
	}
	return data;
}

/* USAGE FLOW
toWrite := EncodeToBytes(t);
toWrite = Compress(toWrite);
WriteFile(toWrite, "test.dat");

toRead := ReadFromFile("test.dat");
toRead = Decompress(toRead);
test := DecodeToStructTest(toRead);

fmt.Println(test);
*/

