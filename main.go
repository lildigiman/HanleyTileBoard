// TileBoardGame project main.go
package main

import (
	"fmt"
	"os"
)

var WINNER = [3][3]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}
var tileboard = [3][3]int{{4, 2, 0}, {5, 8, 6}, {7, 1, 3}}
var score = 0

func main() {
	var in string
	for !wonGame() {
		printTiles()
		score++
		fmt.Scan(&in)
		switch in {
		case "w":
			moveUp()
		case "s":
			moveDown()
		case "a":
			moveLeft()
		case "d":
			moveRight()
		case "exit":
			fmt.Printf("You lost the game after %d tries.\n", score)
			os.Exit(0)
		default:
			fmt.Println("Unrecognizable input: Please use WASD")
		}
	}
	fmt.Printf("YOU WONT THE GAME \n Score: %d", score)
}

func moveUp() {
	fmt.Println("Up")
	var i, j = location()
	if i != 0 {
		tileboard[i][j] = tileboard[i-1][j]
		tileboard[i-1][j] = 0
	}
}

func moveDown() {
	fmt.Println("Down")
	var i, j = location()
	if i != 2 {
		tileboard[i][j] = tileboard[i+1][j]
		tileboard[i+1][j] = 0
	}
}

func moveLeft() {
	fmt.Println("Left")
	var i, j = location()
	if j != 0 {
		tileboard[i][j] = tileboard[i][j-1]
		tileboard[i][j-1] = 0
	}
}

func moveRight() {
	fmt.Println("Right")
	var i, j = location()
	if j != 2 {
		tileboard[i][j] = tileboard[i][j+1]
		tileboard[i][j+1] = 0
	}
}

func wonGame() bool {
	if tileboard == WINNER {
		return true
	} else {
		return false
	}
}

func location() (int, int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if tileboard[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func printTiles() {
	for i := 0; i < 3; i++ {
		fmt.Println(tileboard[0:3][i])
	}
}
