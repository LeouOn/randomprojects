//Q: How do you get the memory address of a variable?
//A: Using pointers we can get to the memory address of a variable. say you have an integer x, you may be able to get the address using &x.
//Q: How do you assign a value to a pointer?
//A: you have to be able to "dereference a pointer variable" to write to that pointer, say you have pointer 'Ptr', you must write '*Ptr =' before the variable you want to write to that pointer
//Q: How do you create a new pointer?
//A: you write the pointer name 'name' followed by *(type) like name *int or blue *string... etc.

/* package main
import "fmt"
func square(x *float64){
	*x = *x * *x
}
func main() {
	x := 1.5
	square(&x)
	fmt.Println(x) //2.25

} */

package main
import "fmt"
func swap(x, y *int){
*x, *y = *y, *x
}
func main() {
	x := 1
	y := 2
	fmt.Println(x,y)
	swap(&x, &y)
	fmt.Println("The New Value of X is ", x) //2
	fmt.Println("The New Value of Y is ", y) //1	
}