//Write a function with one variadic parameter that finds the greatest number in a list of numbers.
//Finally Works!

package main
import "fmt"
func findgreatest(xs []int)(int){
	greatest := 0 
	for _, v := range xs{	

	if v>greatest{
		greatest=v
	}else{

	}}
	return greatest
}
func main(){
x := []int{
  48,96,86,68,
  57,82,63,70,
  37,34,83,27,
  19,97, 9,17,
}
y := []int{
	96,	87,	55,	48,	82,
11,	94,	15,	51,	78,
48,	68,	56,	35,	49,
82,	97,	61,	13,	66,
31,	80,	11,	68,	42,
91,	62,	53,	62,	41,
10,	73,	91,	59,	39,
92,	60,	44,	79,	71,
67,	51,	83,	1,	43,
95,	87,	61,	86,	9,
96,	87,	87,	51,	97,
68,	39,	75,	48,	37,
16,	69,	91,	23,	44,
20,	36,	10,	21,	5,
44,	41,	76,	46,	96,
96,	4,	77,	63,	93,
54,	8,	35,	78,	59,
10,	89,	47,	99,	83,
27,	81,	71,	23,	60,
30,	64,	96,	4,	25,

}
n := findgreatest(x)
fmt.Println("The greatest number in array x is: ", n)
fmt.Println("The greatest number in array y is: ", findgreatest(y))
}
