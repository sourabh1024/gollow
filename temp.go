package main

import (
	bytes2 "bytes"
	"encoding/gob"
	"fmt"
	"log"
	"strconv"
)

type X interface {
	GetKey() string
}

type A struct {
	A int64
}

func (a A) GetKey() string {
	return strconv.FormatInt(a.A, 10)
}

type B struct {
	S string
}

func (b B) GetKey() string {
	return b.S
}

//func Save(x []X) []byte {
//	bytes, _ := json.Marshal(x)
//	return bytes
//}
//
//func Retrieve(bytes []byte) []X {
//	var p []X
//	err := json.Unmarshal(bytes, &p)
//	if err != nil {
//		fmt.Println(err)
//	}
//	return p
//}

// interfaceEncode encodes the interface value into the encoder.
func interfaceEncode(enc *gob.Encoder, p X) {
	// The encode will fail unless the concrete type has been
	// registered. We registered it in the calling function.

	// Pass pointer to interface so Encode sees (and hence sends) a value of
	// interface type. If we passed p directly it would see the concrete type instead.
	// See the blog post, "The Laws of Reflection" for background.
	err := enc.Encode(&p)
	if err != nil {
		log.Fatal("encode:", err)
	}
}

// interfaceDecode decodes the next interface value from the stream and returns it.
func interfaceDecode(dec *gob.Decoder) X {
	// The decode will fail unless the concrete type on the wire has been
	// registered. We registered it in the calling function.
	var p X
	err := dec.Decode(&p)
	if err != nil {
		log.Fatal("decode:", err)
	}
	return p
}

func main() {

	a := A{
		A: 1,
	}

	b := B{
		S: "2",
	}

	x := []X{
		&a, &b,
	}

	gob.Register(&a)
	gob.Register(&b)

	var bytes bytes2.Buffer
	// Create an encoder and send some values.
	enc := gob.NewEncoder(&bytes)
	interfaceEncode(enc, x[0])
	interfaceEncode(enc, x[1])

	fmt.Println(bytes)

	// Create a decoder and receive some values.
	dec := gob.NewDecoder(&bytes)

	result := interfaceDecode(dec)
	fmt.Println(result.GetKey())

	result = interfaceDecode(dec)
	fmt.Println(result.GetKey())

}
