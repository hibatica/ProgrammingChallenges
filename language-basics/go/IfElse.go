package main

import "fmt"

func main() {
        if 1 > 2 {
                fmt.Println("1 is greater than 2")
        } else if 1 > 0 {
                fmt.Println("1 is greater than 0")
        } else {
                fmt.Println("There seems to be a problem!")
        }
}
