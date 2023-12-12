package main

import (
	"fmt"
)

type Data struct {
	mData map[string]string
}

func main() {
	fmt.Println("hello Map")
	md := Data{make(map[string]string)}
	md.SET("Hello", "World")

	fmt.Println(md.mData["Hello"])
}

func (d *Data) SET(key string, value string) {
	d.mData[key] = value
}
