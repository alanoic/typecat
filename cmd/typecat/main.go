package main

import (
	"fmt"
	"os"
	typecat "typecat/internal"
)

func main() {
	path, err := typecat.ParseArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	content, err := typecat.ParseFileContent(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := typecat.TransformContent(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", result)
}
