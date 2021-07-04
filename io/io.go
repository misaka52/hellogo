package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func printFile(filename string) {
	fmt.Println("start print file, filename=", filename)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
	fmt.Println("print end!!!")
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	printFile("abc.txt")
	// 可包含特殊字符，换行
	s := `fabba
		"fdaf"
		81923fdsa

	fds	`
	printFileContents(strings.NewReader(s))
}
