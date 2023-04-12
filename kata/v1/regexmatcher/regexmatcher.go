package main

import (
	"fmt"
	"regexp"
)

func main() {
	//matchRegex("test string peach peeach pch")
	matchRegex2()
}

func matchRegex(input string) error {
	r := regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)
	fmt.Println(r.ReplaceAllString(input, "<fruit>"))
	fmt.Println("regexp:", r.FindString("regexp"))
	return nil
}
func matchRegex2() {
	rgx, _ := regexp.MatchString("p([a-z]+)ch", "test string peach peeach pch")
	fmt.Println("matched(boolean)?:", rgx)
	r := regexp.MustCompile("p([a-z]+)ch")

	fmt.Println("MatchString(boolean)?:", r.MatchString("Test string peach peeach pinch punch pair."))
	fmt.Println("FindString(string)?:", r.FindString("Test string peach peeach pinch punch pair."))
	fmt.Println("FindStringIndex([]int)?:", r.FindStringIndex("peach peeach pinch punch pair."))
	fmt.Println("FindStringSubmatch([]string)?:", r.FindStringSubmatch("pinch peach peeach pinch punch pair."))
	fmt.Println("FindStringSubmatchIndex([]int)?:", r.FindStringSubmatchIndex("pinch peach peeach pinch punch pair."))
	fmt.Println("FindAllString([]string)?:", r.FindAllString("pinch peach peeach pinch punch pair.", -1))
	fmt.Println("limit FindAllString([]string)?:", r.FindAllString("pinch peach peeach pinch punch pair.", 2))
}
