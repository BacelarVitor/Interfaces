package main

type englishBot struct{}
type spanishBot struct{}

func main() {

}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating english greeting
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}