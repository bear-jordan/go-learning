package main


import (
    "flag"
    "fmt"
    "os"
)

func MakeMessage(color string, message string) string {
    return fmt.Sprintf("%s%s\033[0m", color, message)

}

func ParseArgs() (string, string, bool) {
    // Init flags
    selectedColor := flag.String("color", "white", "Color to display message")
    verbose := flag.Bool("verbose", false, "Set verbosity")
    flag.Parse()
    
    // Check args are correct
    if flag.NArg() != 1 {
        fmt.Fprintf(os.Stderr, "Usage: `color -color=white [-verbose=true] <message>`\r\n")
        os.Exit(1)
    }

    // Get message
    message := flag.Arg(0)

    return *selectedColor, message, *verbose
}

func main() {
    // Parse args
    selectedColor, message, verbose := ParseArgs()
    if verbose {
        fmt.Fprint(os.Stderr, fmt.Sprintf("Picking a color: `%s`\r\n", selectedColor))
    }

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
    color, ok := validColors[selectedColor]
    if !ok {
        fmt.Fprint(os.Stderr, "Color not found.\r\n")
        color = validColors["white"]
    }
    fmt.Println(MakeMessage(color, message))

    os.Exit(0)
}
