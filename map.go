package main

import (
	"fmt"
)

type Data struct {
	mData map[string]string
}

func main1111() {
	fmt.Println("hello Map")
	md := Data{make(map[string]string)}
	md.SET("Hello", "World")
	//This is to add dummy commit
	//This is another commit
	fmt.Println(md.mData["Hello"])
}

func (d *Data) SET(key string, value string) {
	d.mData[key] = value
}
