package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

func main() {
	var action, str, filepath, pass string
	var err error

	flag.StringVar(&action, "action", "", "decode | encode")
	flag.StringVar(&str, "str", "", "string for decode | encode")
	flag.StringVar(&filepath, "filepath", "", "path to .sync file")
	flag.Parse()

	if filepath != "" {
		parsefile(filepath)
		return
	}

	if str == "" {
		fmt.Println("Error: parameters action and str are required")
		printUsage()
		return
	}

	switch action {
	case "decode":
		pass, err = decode(str)
	case "encode":
		pass, err = encode(str)
	default:
		fmt.Println("Error: parameters action and str are required")
		printUsage()
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Pass: ", pass)
}

func printUsage() {
	fmt.Printf("Usage: %s [OPTIONS] argument ...\n", os.Args[0])
	flag.PrintDefaults()
}

func parsefile(filepath string) {
	cfg, err := ini.Load(filepath)
	if err != nil {
		fmt.Printf("Error: Fail to read file: %v\n", err)
		printUsage()
		os.Exit(0)
	}
	secs := cfg.Sections()
	for _, v := range secs {
		fmt.Println("Name: ", v.Key("Name").String())
		fmt.Println("Host: ", v.Key("Host").String())
		fmt.Println("Port: ", v.Key("Port").String())
		fmt.Println("User: ", v.Key("User").String())
		pass, err := decode(v.Key("Password").String())
		if err == nil {
			fmt.Println("Pass: ", pass)
		} else {
			fmt.Println("Password [NOT PARSED]: ", v.Key("Password").String())
		}

		fmt.Println()
	}
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
