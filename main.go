package main


import (
    "flag"
    "fmt"
    "os"
)

func MakeMessage(color string, message string) string {
    return fmt.Sprintf("%s%s\033[0m", color, message)

}

func ParseArgs() string {
    selectedColor := flag.String("color", "white", "Color to display message")
    flag.Parse()

    return *selectedColor
}

func main() {
    // Parse args
    message := os.Args[2]
    selectedColor := ParseArgs()

    fmt.Fprint(os.Stderr, fmt.Sprintf("Picking a color: `%s`\r\n", selectedColor))

    // Define colors
    var validColors = map[string]string {
        "black": "\033[0;30m",
        "red": "\033[0;31m",
        "green": "\033[0;32m",
        "yellow": "\033[0;33m",
        "blue": "\033[0;34m",
        "purple": "\033[0;35m",
        "cyan": "\033[0;36m",
        "white": "\033[0;37m",
    }

    // Pick a color
    if color, ok := validColors[selectedColor]; ok {
        fmt.Println(MakeMessage(color, message))
    } else {
        fmt.Fprint(os.Stderr, "Color not found.\r\n")
    }
}
