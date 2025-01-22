/*
Author: Bear Jordan
Description: Utility to color standard input
Usage: Usage: `color -color=<color> [-verbose=true] <message>
*/

package main

import (
    "flag"
    "fmt"
    "os"
)

func MakeMessage(color, message string) string {
    return fmt.Sprintf("%s%s\033[0m", color, message)
}

func ParseArgs() (string, string, bool, bool) {
    // Init args
    selectedColor := flag.String("color", "white", "Color to display message")
    verbose := flag.Bool("verbose", false, "Set verbosity")
    help := flag.Bool("help", false, "Display help menu")
    flag.Parse()

    // Validate args
    if *help || flag.NArg() != 1 {
        return *selectedColor, "", *verbose, true
    }

    return *selectedColor, flag.Arg(0), *verbose, *help
}

func ShowHelp(validColors map[string]string) {
    // Print out help menu
    fmt.Fprintln(os.Stderr, "Usage: `color -color=<color> [-verbose=true] \"<message>\"`")
    fmt.Fprintln(os.Stderr, "\nValid colors:")
    for color := range validColors {
        fmt.Fprintf(os.Stderr, "  - %s\n", color)
    }
}

func main() {
    // Init colors
    validColors := map[string]string{
        "black":  "\033[0;30m",
        "red":    "\033[0;31m",
        "green":  "\033[0;32m",
        "yellow": "\033[0;33m",
        "blue":   "\033[0;34m",
        "purple": "\033[0;35m",
        "cyan":   "\033[0;36m",
        "white":  "\033[0;37m",
    }

    // Parse args
    selectedColor, message, verbose, help := ParseArgs()

    // Show help if needed
    if help {
        ShowHelp(validColors)
        os.Exit(0)
    }

    // Validate selected color
    if verbose {
        fmt.Fprintf(os.Stderr, "Picking a color: `%s`\n", selectedColor)
    }

    color, ok := validColors[selectedColor]
    if !ok {
        fmt.Fprintln(os.Stderr, "Color not found.")
        color = validColors["white"]
    }

    // Print message
    fmt.Println(MakeMessage(color, message))
}
