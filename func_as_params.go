package main

import "fmt"

type Filter func(string) string

func sayHello(word string, filter Filter) string {
	filteredWord := filter(word)
	return filteredWord
}

func wordCheck(word string) string {
	if word == "Anjing" {
		return "..."
	} else {
		return word
	}
}

func main() {
	fmt.Println(sayHello("Anjing", wordCheck))
	fmt.Println(sayHello("Fulan", wordCheck))
}
