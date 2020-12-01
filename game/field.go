package game

import (
    "github.com/nsf/termbox-go"

)


type Field struct {
    food    Food
    snake   Snake
    height  int
    width   int
}
