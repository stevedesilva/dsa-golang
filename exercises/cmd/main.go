package main

import (
	"fmt"
	"net/http"
)

func main() {
	list := make([]string, 0, 1)
	fmt.Printf("> main : %v %p\n", list, &list)
	A(list)
	fmt.Printf("< main : %v %p\n", list, &list)
	http.NewServeMux()

}

func A(list []string) {
	fmt.Printf(">> : %v %p\n", list, &list)
	list = append(list, "foo")
	fmt.Printf("<< : %v %p\n", list, &list)
}
