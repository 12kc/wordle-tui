package main

import "fmt"

func main() {
	/* var boardCheck = [3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	} */
	var boardPick = [3][3]string{}
	var userPick string
	fmt.Scan(&userPick)

	boardPick[1][1] = userPick
	fmt.Println(boardPick[1])
}
