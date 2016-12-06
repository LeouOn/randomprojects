//pointers
package main
import "fmt"
/* This will not change the value of X in the main function
func zero(x int){
	x = 0
} */
/*
func zero(xPtr *int) {
	*xPtr = 0
}
func main(){
	x:= 5
	zero(&x)
	fmt.Println(x) //x is still five
} */
func one(xPtr *int) {
	*xPtr =1
}
func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}