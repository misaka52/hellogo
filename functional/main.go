package main

import (
	"bufio"
	"fmt"
	"hellogo/functional/fibo"
	"io"
	"strings"
)

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()

	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	return strings.NewReader(s).Read(p)
}

func printFileContents(io io.Reader) {
	scanner := bufio.NewScanner(io)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	//var f intGen = fibo.Fibonacci()
	//printFileContents(f)

	fibo.WriteFile("fib.txt")
}
