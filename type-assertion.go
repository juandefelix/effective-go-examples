package main

import "fmt"

type Stringer interface {
    String() string
}

type Person struct {
    Name string
    Age int
}

func (p Person) String() string {
    return fmt.Sprintf("Name: %v, Age: %v", p.Name, p.Age)
}

func findType(value interface{}) {
    switch str := value.(type) {
    case string:
        fmt.Println("This is a string: ", str)
    case Stringer:
        fmt.Println("This is a Stringer: ", str.String())
    }
}

func main() {
    luis := Person{ Name: "Luis", Age: 40}

    findType(luis)
    findType("luis")
}
