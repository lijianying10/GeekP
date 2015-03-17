package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	key := []byte("5e8487e6")
	origtext := []byte("t")

	erytext, err := DesEncrypt(origtext, key)
	if err != nil {
	}
	fmt.Println(erytext)
	fmt.Printf("%v\n", erytext)
	dist := make([]byte, 2048) //开辟存储空间
	base64.StdEncoding.Encode(dist, erytext)
	fmt.Println(string(dist))
	destext, err2 := DesDecrypt(erytext, key)
	if err2 != nil {
	}
	fmt.Println(string(destext))
	fmt.Println(len(origtext), len(string(destext)))
	fmt.Println(string(origtext) == string(destext))
}
