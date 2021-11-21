package main

import (
	"encoding/base64"
	"flag"
	"fmt"
)

func main() {
	var action string
	var str string
	var err error
	var pass string

	flag.StringVar(&action, "action", "", "decode | encode")
	flag.StringVar(&str, "str", "", "string for decode | encode")
	flag.Parse()
	if str == "" {
		fmt.Println("Params action and str are required!2")
		return
	}

	switch action {
	case "decode":
		pass, err = decode(str)
	case "encode":
		pass, err = encode(str)
	default:
		fmt.Println("Params action and str are required!1")
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Password: ", pass)
}
func decode(password string) (string, error) {
	var byteArray []byte
	var err error
	byteArray, err = base64.StdEncoding.DecodeString(password)
	if err != nil {
		return "", err
	}
	for k := range byteArray {
		byteArray[k] = byteArray[k]<<1 | byteArray[k]>>7
	}
	return string(byteArray), err
}

func encode(password string) (string, error) {
	var byteArray []byte
	var err error
	byteArray = []byte(password)
	for k := range byteArray {
		byteArray[k] = byteArray[k]>>1 | byteArray[k]<<7
	}
	password = base64.StdEncoding.EncodeToString(byteArray)
	return password, err
}
