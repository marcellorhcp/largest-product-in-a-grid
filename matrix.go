package main

import "fmt"

var matriz = [][]int{
	{1, 2, 3, 4},
	{5, 6, 7, 8},
	{9, 10, 11, 12},
	{13, 14, 15, 16},
}

func multiplicaMatrix(s [][]int) int {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			s[i][j] * s[i][j]
		}
	}
}
func main() {
	fmt.Println(matriz)

}
