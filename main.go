package main

import (
	"fmt"
	"os"
)
import "strings"

//func readString() string {
//	rdr := bufio.NewReader(os.Stdin)
//	str, _ := rdr.ReadString('\n')
//	return str
//}

func main() {
	s := os.Args[1]

	fmt.Print(len(strings.Fields(s)))
}
