package main

import (
	"fmt"
)

type Talker interface {
	Talk() string
}

type Listener interface {
	Listen(string)
}

type Convo struct {
	moderator string
	Talker
	Listener
}

type ListeningPerson struct {
	name string
}

type TalkingPerson struct {
	name string
}

func(p TalkingPerson) Talk() string {
	statement := "Hola Manola!"
	fmt.Println(statement)
	return statement
}

func(p ListeningPerson) Listen(statement string){
	fmt.Println("I hear you!")
}


func main() {
	talker := TalkingPerson{"Pepe"}
	listener := ListeningPerson{"Manolo"}
	convo := Convo{"Luis", talker, listener}

	fmt.Println("Conversation moderated by: ", convo.moderator)
	statement := convo.Talk()
	convo.Listen(statement)

}

