package main

import "fmt"

// easy
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder()
	for i := 1; i <= 10; i++ {
		fmt.Printf("1+...+%d %d\n", i, a(i))
	}

	b := adder2(0)
	for i := 1; i <= 10; i++ {
		var s int
		s, b = b(i)
		fmt.Printf("1+...+%d %d\n", i, s)
	}

}
