package main

import ("fmt")

func main() {
    for i := 0; i < 10; i++ {
        // %b automatically converts the number to base 2
        fmt.Printf("%b\n", i);
    }
}
