//elementary-5.go
//Write a program that asks the user for a number n and prints the sum of the numbers 1 to n
//Modify the previous program such that only multiples of three or five are considered in the sum, e.g. 3, 5, 6, 9, 10, 12, 15 for n=17
package main 

import ("fmt"

		)

func main() {
	var i int
	fmt.Println("Write any number n, this program will give you the sum of all multiples of 3 from 0 to n")
	_, err := fmt.Scanf("%d", &i)
 	sum := 0
 	original := i
	if err != nil {
		fmt.Println(err)
	} 
	for ; i > 0; i-- {
		if i%3==0{
		sum +=i
	}
		else {}
	}
	fmt.Println("The sum of all numbers between", original, "and 0 is:", sum)
}