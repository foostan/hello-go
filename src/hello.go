package main

import "fmt"

func main() {
	ms := []string{"Hello", "World", "and", "Go"}

	for _, m := range ms {
		fmt.Println(m)
	}
}
