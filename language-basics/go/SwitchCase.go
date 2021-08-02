package main

import "fmt"

func main() {
        five := 5

        fmt.Println("Fetching number")
        fmt.Println("Number fetched: ", five)
        fmt.Println("Verifying identity of fetched number")

        switch five {
        case 3:
                fmt.Println("Fetched number is: 3")
        case 4:
                fmt.Println("Fetched number is: 4")
        case 5:
                fmt.Println("Fetched number is: 5")
        case 6:
                fmt.Println("Fetched number is: 6")
        default:
                fmt.Println("Could not identify fetched number")
        }
}
