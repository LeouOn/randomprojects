//elementary-4.go
//Write a program that asks the user for a number n and prints the sum of the numbers 1 to n

package main 

import ("fmt"

		)

func main() {
	var i int
	fmt.Println("Write any number n, this program will give you the sum of all numbers from 0 to n")
	_, err := fmt.Scanf("%d", &i)
 	sum := 0
 	original := i
	if err != nil {
		fmt.Println(err)
	} 
	for ; i > 0; i-- {
		sum +=i
	}
	fmt.Println("The sum of all numbers between", original, "and 0 is:", sum)
}