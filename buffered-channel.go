package main

import "fmt"

const maxOutstanding = 10

func expensiveCalculation(myAttribute int, channel chan int) {
	// run your expensice calculation here
	fmt.Println("executing some code with", myAttribute)

	<- channel
}

func main() {
	channel := make(chan int, maxOutstanding) // Buffered channel up to 10 spots

	for i := 0; i < 200; i++ {
		channel <- i // this will wait until we have a free spot in the buffered channel
		func() {
			go expensiveCalculation(i, channel)
		}()
	}
}
