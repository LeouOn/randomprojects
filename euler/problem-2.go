/* I am new to programming in general and Golang is my first major language. Feel free to add any comments or corrections. Thank you! */
package main 

import "fmt"
func main (){
	a := 0
	b := 1
	ph := 0
	sum := 0
	for b < 4000000{
		ph = a
		a = b
		b = ph + b
		if b%2==0 {
		sum = sum + b
		}else {

		}
	}
	fmt.Println(sum)
}
