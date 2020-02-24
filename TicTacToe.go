package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type move struct {
	row int
	col int
}

func main() {
	// Local constants

	// Local variables
	board := buildBoard()
	lastMove := move{-1, -1}
	numOfMoves := 0
	scanner := bufio.NewReader(os.Stdin)

	/****** start main() ******/

	for !checkSolved(board, lastMove) && numOfMoves != 9 {
		printBoard(board)

		// User's turn
		if numOfMoves%2 == 0 {
			fmt.Print("-> ")
			input, _ := scanner.ReadString('\n')
			mv := formatMove(input)
			if checkMove(board, mv) {
				lastMove = mv
			}

			for !checkMove(board, mv) {
				fmt.Print("-> ")
				input, _ := scanner.ReadString('\n')
				mv := formatMove(input)
				if checkMove(board, mv) {
					lastMove = mv
				}
			}

			board[lastMove.row][lastMove.col] = "X"

		} else {
			// AI turn
			row := rand.Intn(3)
			col := rand.Intn(3)
			mv := move{row, col}

			for !checkMove(board, mv) {
				row := rand.Intn(3)
				col := rand.Intn(3)
				mv = move{row, col}
			}

			lastMove = mv
			board[lastMove.row][lastMove.col] = "O"
		}

		numOfMoves++
		fmt.Println()
	}

	if numOfMoves == 9 && !checkSolved(board, lastMove) {
		fmt.Println("It's a draw!")
	} else if numOfMoves%2 == 1 {
		fmt.Println("You won!")
	} else {
		fmt.Println("AI won!")
	}

}

func buildBoard() [3][3]string {
	// Local constants

	// Local variables
	var board [3][3]string

	/****** start buildBoard() ******/

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = "-"
		}
	}

	return board
}

func printBoard(board [3][3]string) {
	// Local constants

	// Local variables

	/****** start printBoard() ******/

	for i := 0; i < 4; i++ {
		if i == 0 {
			fmt.Print(" ABC")
		} else {
			for j := 0; j < 4; j++ {
				if j == 0 {
					fmt.Print(i)
				} else {
					fmt.Print(board[i-1][j-1])
				}
			}
		}
		fmt.Println()
	}
}

func checkSolved(board [3][3]string, mv move) bool {
	// Local constants

	// Local variables

	/****** start checkSolved() ******/

	if mv.row == -1 || mv.col == -1 {
		return false
	}

	// Check col
	for i := 0; i < 3; i++ {
		if board[i][mv.col] != board[mv.row][mv.col] {
			break
		} else if i == 2 {
			return true
		}
	}

	// Check row
	for i := 0; i < 3; i++ {
		if board[mv.row][i] != board[mv.row][mv.col] {
			break
		} else if i == 2 {
			return true
		}
	}

	// Check diag
	if mv.row == mv.col {
		for i := 0; i < 3; i++ {
			if board[i][i] != board[mv.row][mv.col] {
				break
			} else if i == 2 {
				return true
			}
		}
	}

	// Check reverse diag
	if mv.row+mv.col == 2 {
		for i := 0; i < 3; i++ {
			if board[i][2-i] != board[mv.row][mv.col] {
				break
			} else if i == 2 {
				return true
			}
		}
	}

	return false
}

func formatMove(input string) move {
	// Local constants

	// Local variables
	x := []rune(input)
	r := string(x[1])
	c := string(x[0])
	var row int
	var col int

	/****** start formatMove() ******/

	// Convert row
	if r == "1" {
		row = 0
	} else if r == "2" {
		row = 1
	} else if r == "3" {
		row = 2
	}

	// Convert col
	if c == "A" {
		col = 0
	} else if c == "B" {
		col = 1
	} else if c == "C" {
		col = 2
	}

	return move{row, col}

}

func checkMove(board [3][3]string, mv move) bool {
	// Local constants

	// Local variables

	/****** start checkMove() ******/

	if board[mv.row][mv.col] == "-" {
		return true
	}

	return false
}
