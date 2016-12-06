// The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function which can find fib(n).

package main
import "fmt"
/* This is what a one function Fibonacci generator looks like
func main(){
	for i, j := 0, 1; j < 100; i, j = i+j, i{
	fmt.Println(i)
	} */

//Now let's do it in two functions
func fibo(max int){
	for i, j := 0, 1; j < max; i, j = i+j, i{
		fmt.Println(i)
	}
}
func main(){
	fibo(50)
	fibo(10)
}
	