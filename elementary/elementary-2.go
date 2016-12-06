// takes input from user and spits it back
package main 
import ("bufio"
		"fmt"
		"os"
		)

func main() {
	fmt.Println("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello, nice to meet you", name)
}