package main 
import "fmt"
/* func average(xs []float64) float64{
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}
func main() {
	xs := []float64{98,93,77,82,83}
	fmt.Println(average(xs))*/
/*
func f() (int, int){
	return 5, 6
}
func main(){
	x, y := f()
	fmt.Println(x, y)
}
*/
func add(args ...int) int{
	total := 0
	for _, v := range args {
		total +=v
	}
	return total
}
func main(){
	fmt.Println(add(1,2,3))
}