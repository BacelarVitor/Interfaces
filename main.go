package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	PrintGreeting(eb)
	PrintGreeting(sb)
}

func PrintGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating english greeting
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}