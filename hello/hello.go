package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const portuguese = "Portuguese"
const prefixHelloEnglish = "Hello, "
const prefixHelloSpanish = "Hola, "
const prefixHelloFrench = "Bonjour, "
const prefixHelloPortuguese = "Olá, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + "!!!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = prefixHelloFrench
	case spanish:
		prefix = prefixHelloSpanish
	case portuguese:
		prefix = prefixHelloPortuguese
	default:
		prefix = prefixHelloEnglish
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
