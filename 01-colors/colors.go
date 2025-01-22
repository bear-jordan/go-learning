package colors

import (
    "fmt"
    "os"
)

// Define a color -> ansii map
type color struct {
    name string
    color string
}


func MakeMessage(color color) string {
    return fmt.Sprintf("%s%s\033[0m", color.color, color.name)

}

func PrintColor() {
    fmt.Fprint(os.Stderr, "Picking a color\r\n")
    // Define colors
    var colors = []color {
        {name: "black", color: "\033[0;30m"},
        {name: "red", color: "\033[0;31m"},
        {name: "green", color: "\033[0;32m"},
        {name: "yellow", color: "\033[0;33m"},
        {name: "blue", color: "\033[0;34m"},
        {name: "purple", color: "\033[0;35m"},
        {name: "cyan", color: "\033[0;36m"},
        {name: "white", color: "\033[0;37m"},
    }

    // Give a message (default to hello)
    for _, color := range colors {
        message := MakeMessage(color)
        fmt.Println(message)
    }
    // Pick a color
}
