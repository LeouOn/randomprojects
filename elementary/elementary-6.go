//elementary-6.go
//WWrite a program that asks the user for a number n and gives him the possibility to choose between computing the sum and computing the product of 1,…,n.
// program is broken, something wrong with scanf
package main 

import ("fmt"

		)

func main() {
	var i, sum, product int
	var sumorproduct string
	fmt.Println("Write any number n")
	_, err := fmt.Scan("%d", &i)
 	original := i
	if err != nil {
		fmt.Println(err)
	} 
	fmt.Println("Do you want the product or sum of all numbers from 1 to ", original, "?")
	_, err = fmt.Scan("%q", sumorproduct)
	for ; i >= 1; i-- {
		if sumorproduct == "sum"{
		sum += i
	} else if sumorproduct =="roduct"{
		product *= i
	}
	}
	fmt.Println("The sum of all numbers between", original, "and 0 is:", sum)
	fmt.Println("The product of all numbers between", original, "and 0 is:", product)
	fmt.Println(i, sumorproduct)
}