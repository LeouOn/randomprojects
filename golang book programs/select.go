//Q1: How do you specify the direction of a channel type? 
//You specify the direction of a channel through the arrows <- is write only -> is read only
//Q2: Write your own Sleep function using time.After.
//
//Q3: What is a buffered channel? How would you create one with a capacity of 20?
//A buffered channel is asynchronous, so they will continually send messages into the buffer and out. You would have to make a second parameter after the channel type like int, 20.

package main 
import ("fmt"
		"time")
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

go func() {
	for {
		c1 <- "from 1"
		time.Sleep(time.Second * 2)
	}
	}()

go func() {
	for {
		c2 <- "from 2"
		time.Sleep(time.Second * 3)
	}
}()

go func() {
	for {
		select {
		case msg1 := <- c1:
			fmt.Println("Message 1", msg1)
		case msg2 := <- c2:
			fmt.Println("Message 2", msg2)
		case <- time.After(time.Second):
			fmt.Println("timeout")
		}
	}
}()

var input string
fmt.Scanln(&input)
}