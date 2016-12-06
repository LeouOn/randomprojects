// Write a function which takes an integer and halves it and returns true if it was even or false if it was odd. 
// For example half(1) should return (0, false) and half(2) should return (1, true).

package main 
import "fmt" 
func half(xs int) (int, bool){
if xs%2==0 && xs!=0{
	return xs, true
} else {
	return xs, false
}
}
func main() {
	fmt.Println(half(7))
	fmt.Println(half(4))
	fmt.Println(half(1234234))
	fmt.Println(half(91))
}