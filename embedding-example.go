package main

import (
	"fmt"
)

type Talker interface {
	Talk(string) string
}

type Listener interface {
	Listen(string)
}

type ConvoParticipant struct {
	name string
	Talker
	Listener
}

func(p ConvoParticipant) Talk(statement string) string {
	fmt.Println(statement)
	return statement
}

func(p ConvoParticipant) Listen(statement string){
	fmt.Printf("You said: \"%v\"", statement)
}


func main() {
	luis := ConvoParticipant{name: "Pepe"}
	manolo := ConvoParticipant{name: "Manolo"}

	statement := luis.Talk("Hola Manolo!")
	manolo.Listen(statement)
}

