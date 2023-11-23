package main

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
// 	println("Hello World")
// 	//Read File
// 	filePath := "main.yaml"
// 	data, err := os.ReadFile(filePath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	println(string(data))

// 	//need to read indentation PHASE-1
// 	//Lets print each bytr

// 	println("TYPE : %T", data)

// 	//print key
// 	var keyList []string
// 	var key string
// 	for i := 0; i < len(data); i++ {

// 		if string(data[i]) == ":" {
// 			println("key: %v", key)
// 			keyList = append(keyList, key)
// 			println("ADDED")
// 			key = ""
// 		}
// 		if string(data[i]) != "\n" && string(data[i]) != " " {
// 			key += string(data[i])
// 			println(key)

// 		}

// 	}

// 	for j := 0; j < len(keyList); j++ {
// 		println("KEY: %v", keyList[j])
// 	}

// 	//Example file writing
// 	writeFilePath := "hello.txt"

// 	writeData := []byte("Hello World")
// 	fs, err := os.OpenFile(writeFilePath, os.O_APPEND, os.ModeAppend)
// 	flag, err := fs.Write(writeData)
// 	println(flag)
// 	// err = os.WriteFile(writeFilePath, writeData, os.ModeAppend) //TODO: Append is not working
// 	checkError(err)
// 	fs.Close()

// }

func main() {
	testConsole()
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}

}

func testConsole() {
	var value string
	// fmt.Println("Hello World")
	// fmt.Println("Enter Any thing")
	// fmt.Scanf("%s", &value)
	// fmt.Println("You have entered : ", value)
	fmt.Print("joker # ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	value = scanner.Text()

	fmt.Println("Value: ", value)

}
