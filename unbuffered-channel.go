package main

import "fmt"

func expensiveCalculation(channel chan int) {
	// run your expensice calculation here
	fmt.Println("finished the expensive calculation")

	channel <- -1
}

func main() {
	channel := make(chan int)
	go expensiveCalculation(channel)

	<- channel // will be blocked untl the expensive calculation has finished
	fmt.Println("after the expensive calculation") // this will run after the expensive calculation has finished

}
