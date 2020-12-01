package main

import(
    "fmt"
    "github.com/nsf/termbox-go"
)


func main() {
    // Initualizes termbox library
    err := termbox.Init()

    // If an error occur, then panic
    if err != nil {
        // If fail, stop all function and go to defer functions
        panic(err)
    }

    // Ensure that termbox is always close
    defer termbox.Close()

    fmt.Println()
}
