/* I am new to programming in general and Golang is my first major language. Feel free to add any comments or corrections. Thank you! */

package main

import "fmt"

func main() {
	sum := 0
	for i :=0; i<1000; i++{
		if i%3==0 {
		sum = sum + i
	} else if i%5==0 {
		sum = sum + i
	} else {

	}
	}
	fmt.Println(sum)
}