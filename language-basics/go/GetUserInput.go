package main

import (
    "fmt"
)

func main() {

        name := ""

        fmt.Println("What is your name?")
        fmt.Print(">") // Print doesn't insert new line

        fmt.Scanln(&name)

        fmt.Print("Hello ", name, "!\n")

}
