package main

import (
	"fmt"
	"os"
)

func readFileContent(filename string) string {
	data, _ := os.ReadFile(filename)
	return string(data)
}

func main() {
	content := readFileContent("nonexistent_file.txt")
	fmt.Println("文件内容：", content)
}
