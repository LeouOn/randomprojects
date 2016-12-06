package main 
import "fmt" 
func main() {
x := []int{
  48,96,86,68,
  57,82,63,70,
  37,34,83,27,
  19,97, 9,17,
}
smallest := x[0]
largest := x[0]
for _, value := range x {
	if value<smallest{
		smallest = value
	}else {

	}
	fmt.Println("current smallest:", smallest)

}
for _, value := range x{
	if value>largest{
		largest = value
	} else {

	}
	fmt.Println("current largest:", largest)
}
fmt.Println("The smallest number is ", smallest, " and the largest number is", largest)
}