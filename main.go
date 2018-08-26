package main

import ("fmt")

/*******************
* Expected output
* ----------------
* 0
* 1
* 10
* 11
* 100
* 101
* 110
* 111
* 1000
* 1001
*******************/
func main() {
    for i := 0; i < 10; i++ {
        // %b automatically converts the number to base 2
        fmt.Printf("%b\n", i);
    }
}
