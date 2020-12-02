package game

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

type Game struct {
	field Field
	score int
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

	keyboardChan := make(chan termbox.Key)

	// Watch for player input, and hand out events.
	go WatchPlayerInput(keyboardChan)

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
	for {
		Display()
		switch <- keyboardChan {
		case termbox.KeyArrowUp:
			fmt.Println("up")

		case termbox.KeyArrowDown:
			fmt.Println("down")

		case termbox.KeyArrowLeft:
			fmt.Println("left")

		case termbox.KeyArrowRight:
			fmt.Println("right")

		case termbox.KeyEsc:
			return
		}
	}

}

// WatchPlayerInput watches for player input event
func WatchPlayerInput(keyboardChan chan termbox.Key) {
	termbox.SetInputMode(termbox.InputEsc)
	for {
		e := termbox.PollEvent()
		keyboardChan <- e.Key
	}
}
