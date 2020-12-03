package game

import (
	"os"
	"time"
	"fmt"

	"github.com/nsf/termbox-go"
)

type Game struct {
	field Field
	score int
}
type Coordinate struct {
	x int
	y int
}

// StartGame starts the game of snake.
func StartGame() {
	fmt.Println("Starting game")

	// Initualizes termbox library
	err := termbox.Init()

	// If an error occur, then panic
	if err != nil {
		// If fail, stop all function and go to defer functions
		panic(err)
	}
	// Ensure that termbox always closes
	defer termbox.Close()

	/*
	   Order of events:
	    1. Show intro screen.
	    2. Get player input so we can decide to start game or whatnot.
        3. Start game when player inputs start.
        4. Show snake screen.
        5. Render snake.
        6. Start moving snake and randomly giving food.
        7. Take player input for direction.
		8. Check if player loses
	*/

	game := Game {
		field: InitField(),
		score: 0,
	}
	// Watch for player input.
	go WatchPlayerInput(&game)
	game.field.Display()
	for {
		game.field.snake.move()
		game.field.Display()

		time.Sleep(50*time.Millisecond)
	}

}

// WatchPlayerInput watches for player input event
func WatchPlayerInput(game *Game) {
	termbox.SetInputMode(termbox.InputEsc)
	for {
		e := termbox.PollEvent()
		switch e.Key {
		case termbox.KeyArrowUp:
			game.field.snake.direction = UP

		case termbox.KeyArrowDown:
			game.field.snake.direction = DOWN

		case termbox.KeyArrowLeft:
			game.field.snake.direction = LEFT

		case termbox.KeyArrowRight:
			game.field.snake.direction = RIGHT

		case termbox.KeyEsc:
			os.Exit(0)
			return
		}
	}
}
