package main

import "fmt"

type ByteSize float64

const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

func main() {
	fmt.Println("KB: ", KB)
	fmt.Println("MB: ", MB)
	fmt.Println("GB: ", GB)
	fmt.Println("TB: ", TB)
}
