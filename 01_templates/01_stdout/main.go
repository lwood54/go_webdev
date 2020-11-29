package main

import "fmt"

// NOTE: in command line: you can >> go run main.go > index.html
//			This will create a new index file and it will write the output to that file
func main() {
	name := "Logan Wood"

	tpl := `
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
		</html>
	`
	fmt.Println(tpl)
}
