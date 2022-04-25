package main

import "fmt"

var matriz = [][]int{
	{50, 70, 3, 45},
	{90, 88, 17, 68},
	{28, 55, 31, 72},
	{13, 89, 34, 77},
}

func multiplicaMatrix(s [][]int) {

	for i := 0; i <= len(s)-1; i++ {
		for j := 0; j <= len(s)-1; j++ {

			//fmt.Println(s[i][j], "*", s[i][j+1], "=", s[i][j]*s[i][j+1]*s[i][j+2]*s[i][j+3])
			//fmt.Println("debug:", s[i][j])
			fmt.Println("debug:", s[i][j], "*", s[i][j+1], "=", s[i][j]*s[i][j+1])
		}
	}

}
func main() {
	multiplicaMatrix(matriz)

}
