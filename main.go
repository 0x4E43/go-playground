package main

import (
	"os"
)

func main() {
	println("Hello World")
	//Read File
	filePath := "main.yaml"
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	println(string(data))

	//need to read indentation PHASE-1
	//Lets print each bytr

	println("TYPE : %T", data)

	//print key
	var keyList []string
	var key string
	for i := 0; i < len(data); i++ {

		if string(data[i]) == ":" {
			println("key: %v", key)
			keyList = append(keyList, key)
			println("ADDED")
			key = ""
		}
		if string(data[i]) != "\n" && string(data[i]) != " " {
			key += string(data[i])
			println(key)

		}

	}

	for j := 0; j < len(keyList); j++ {
		println("KEY: %v", keyList[j])
	}

}
