package main

//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
import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// game tic tac toe

type Game struct {
	Board      [9]string
	Player     string
	TurnNumber int
}

func main() {
	fmt.Println("welcome to tic tac toe")

	var game Game
	game.Player = "O"

	gameOver := false
	var winner string

	for gameOver != true {
		PrintBoard(game.Board)
		move := askToPlay()
		err := game.Play(move)
		if err != nil {
			fmt.Println(err)
			continue
		}
		gameOver, winner = CheckForWinner(game.Board, game.TurnNumber)
	}

	PrintBoard(game.Board)
	if winner == "" {
		fmt.Println("its a draw")
	} else {
		fmt.Printf("%s won", winner)
	}
}

func CheckForWinner(b [9]string, n int) (bool, string) {
	test := false
	i := 0

	// horizontal test
	for i < 9 {
		test = b[i] == b[i+1] && b[i+1] == b[i+2] && b[i] != ""
		if !test {
			i += 3
		} else {
			return true, b[i]
		}
	}
	i = 0
	// vertical test

	for i < 3 {
		test = b[i] == b[i+3] && b[i+3] == b[i+6] && b[i] != ""
		if !test {
			i += 3
		} else {
			return true, b[i]
		}
	}

	// diagonal 1 test
	if b[0] == b[4] && b[4] == b[8] && b[0] != "" {
		return true, b[i]
	}

	// diagonal 2 test
	if b[2] == b[4] && b[4] == b[6] && b[2] != "" {
		return true, b[i]
	}

	if n == 9 {
		return true, ""
	}

	return false, ""

}

func ClearScreen() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

func (game *Game) SwitchPlayer() {
	if game.Player == "O" {
		game.Player = "X"
		return
	}

	game.Player = "O"

}

func (game *Game) Play(pos int) error {
	if game.Board[pos-1] == "" {
		game.Board[pos-1] = game.Player
		game.SwitchPlayer()
		game.TurnNumber += 1
		return nil
	}
	return errors.New("try another position")
}

func askToPlay() int {
	var moveInt int
	fmt.Println("enter pos to play: ")
	fmt.Scan(&moveInt)
	return moveInt
}

func PrintBoard(b [9]string) {
	ClearScreen()
	for i, v := range b {
		if v == "" {
			fmt.Printf(" ")
		} else {
			fmt.Printf(v)
		}

		if i > 0 && (i+1)%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf("|")
		}
	}
}
