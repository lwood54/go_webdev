package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// NOTE: you can also is the standard library os to write the data directly to file without command line commands
func main() {
	// initialize and define name variable
	name := "Logan Wood"

	// setup template string with variable usage
	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Hello Go</title>
		</head>
		<body>
			<h1>` + name + `</h1>
	</body>
	</html>
	`)

	// create a file with the relative location and specified file name
	newFile, err1 := os.Create("02_file/index.html")
	if err1 != nil {
		log.Fatal("error creating file", err1)
	}
	// defer executes function call when 'surrounding' functions returns, which would be main() in this case
	// this allows the io.Copy() to occur
	defer newFile.Close()

	// writes using pointer to newFile, using a reading of the src, which returns a pointer to Reader
	io.Copy(newFile, strings.NewReader(str))

	// can I reverse this and read from index.html to create index2.html
	file2, err2 := os.Create("02_file/index2.html")
	if err2 != nil {
		log.Fatal("error creating file", err2)
	}
	data, err3 := ioutil.ReadFile("02_file/index.html")
	if err2 != nil {
		log.Fatal("error reading file", err3)
	}
	// try to replace name Logan with Family
	fileString := string(data)
	changedString := replaceString(fileString, "Logan", "Family")
	fmt.Println(changedString)
	io.Copy(file2, strings.NewReader(changedString))

}

func replaceString(fullStr string, strOut string, strIn string) string {
	strIndx := strings.Index(fullStr, strOut)
	strLen := len(strOut)
	return fullStr[0:strIndx] + strIn + fullStr[strIndx+strLen:]
}
