// takes input from user and compares it to strings "Alice" or "Bob"
package main 
import ("bufio"
		"fmt"
		"os"
		)

func main() {
		fmt.Println("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	if(name == "Alice\n" || name == "Bob\n") {
		fmt.Println("Hello, nice to meet you", name)
	} else {
		fmt.Println("Sorry you are not Alice or Bob")
	}
}