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
			if i == 0 && j == 0 || i == 0 && j == 1 || i == 0 && j == 2 || i == 0 && j == 3 {
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
	fmt.Println(LimitesMatrixUp(matriz))

}
