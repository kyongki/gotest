package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const chineseHelloPrefix = "Hao, "

func Hello(name, lang string) string {
	if name == "" {
		name = "Everyone"
	}
	prefix := englishHelloPrefix
	switch lang {
	case "chinese":
		prefix = chineseHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	}
	return prefix + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case "chinese":
		prefix = chineseHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("James", "English"))
}
