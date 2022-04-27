package main

import "fmt"

var matriz = [][]int{
	{50, 70, 3, 45},
	{90, 88, 17, 68},
	{28, 55, 31, 72},
	{13, 89, 34, 77},
}

func LimitesMatrixUp(mt [][]int) bool {
	var possible bool

	for i := 0; i < len(mt); i++ {
		for j := 0; j < len(mt); j++ {
			if i == 0 {
				possible = false
			} else {
				possible = true
			}
			fmt.Println("debug:", mt[i][j], ":", possible)
		}
	}
	return possible
}
func LimitesMatrixDown(mt [][]int) bool {
	var possible bool

	for i := 0; i < len(mt); i++ {
		for j := 0; j < len(mt); j++ {
			if i == len(mt)-1 {
				possible = false
			} else {
				possible = true
			}
			fmt.Println("debug:", mt[i][j], ":", possible)
		}
	}
	return possible
}

func LimitesMatrixRight(mt [][]int) bool {
	var possible bool

	for i := 0; i < len(mt); i++ {
		for j := 0; j < len(mt); j++ {
			if j == len(mt)-1 {
				possible = false
			} else {
				possible = true
			}
			fmt.Println("debug:", mt[i][j], ":", possible)
		}
	}
	return possible
}

func LimitesMatrixLeft(mt [][]int) bool {
	var possible bool

	for i := 0; i < len(mt); i++ {
		for j := 0; j < len(mt); j++ {
			if j == 0 {
				possible = false
			} else {
				possible = true
			}
			fmt.Println("debug:", mt[i][j], ":", possible)
		}
	}
	return possible
}

func main() {
	fmt.Println("-\t-\tCIMA\t-\t-")
	fmt.Println(LimitesMatrixUp(matriz))
	fmt.Println("-\t-\tBAIXO\t-\t-")
	fmt.Println(LimitesMatrixDown(matriz))
	fmt.Println("-\t-\tDIREITA\t-\t-")
	fmt.Println(LimitesMatrixRight(matriz))
	fmt.Println("-\t-\tESQUERDA\t-\t-")
	fmt.Println(LimitesMatrixLeft(matriz))
}
