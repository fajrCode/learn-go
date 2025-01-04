package main

import "fmt"

type Filter func(string) bool

func sayHello(name string, filter Filter) string {
	if filter(name) {
		return "hello ..."
	} else {
		return "hello " + name
	}
}

func main() {
	filtered := func(name string) bool {
		return name == "Bad Word"
	}

	fmt.Println(sayHello("Bad Word", filtered))

}
