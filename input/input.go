package input

import (
	"bufio"
	"fmt"
	"os"
)

func CreateUrl() string {
	fmt.Print("Please enter the name of the Pokemon.\n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return "http://localhost:8081/?name=" + scanner.Text()
}
