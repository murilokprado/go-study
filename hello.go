package main

import "fmt"

const prefixHelloEnglish = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return prefixHelloEnglish + name + "!!!"
}

func main() {
	fmt.Println(Hello("World"))
}
