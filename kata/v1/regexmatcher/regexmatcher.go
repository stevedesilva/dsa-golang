package main

import (
	"fmt"
	"regexp"
)

func main() {
	matchRegex("test string peach peeach pch")
}

func matchRegex(input string) error {
	r := regexp.MustCompile("p([a-z]+)ch")

	fmt.Println("regexp:", r)
	fmt.Println(r.ReplaceAllString(input, "<fruit>"))
	return nil
}
