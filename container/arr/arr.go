package main

import "fmt"

func arr() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [3][2]int
	fmt.Println(arr1, arr2, arr3, grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	printArray(arr1)
	fmt.Println(arr1)

	printArrayRef(&arr1)
	fmt.Println(arr1)
	// 报错
	//printArray(arr2)
}

func printArray(arr [5]int) {
	fmt.Println("printArray")
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArrayRef(arr *[5]int) {
	fmt.Println("printArrayRef")
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArraySlice(arr []int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func updateSlice(arr []int) {
	arr[0] = 100
}

func slice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr[2:6])
	// 引用更新
	updateSlice(arr[:])
	fmt.Println(arr)
	arr[0] = 0
	s1 := arr[2:6]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	s2 := arr[1:3]
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
	s3 := append(s2, 11)
	s4 := append(s3, 12)
	fmt.Printf("s3= %v, %d, %d\n", s3, len(s3), cap(s3))
	fmt.Println(s4)
	fmt.Println(arr)
}

func printSlice(arr []int) {
	fmt.Printf("len=%d, cap=%d\n", len(arr), cap(arr))
}

func sliceOp() {
	var arr []int
	for i := 0; i < 100; i++ {
		printSlice(arr)
		arr = append(arr, i*2+1)
	}
	fmt.Println(arr)
}

func remove(arr []int, index int) {
	target := arr[index]
	arr = append(arr[:index], arr[index+1:]...)
	fmt.Printf("remove element %d, after array is %v\n", target, arr)
	printSlice(arr)
}

func sliceOp2() {
	arr := make([]int, 10, 16)
	fmt.Println(arr)
	arr2 := []int{0, 1, 2, 3}
	copy(arr, arr2)
	fmt.Println(arr)
	remove(arr, 2)
}

func main() {
	//arr()
	//slice()
	//sliceOp()
	sliceOp2()
}
