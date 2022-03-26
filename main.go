package main

import (
	"fmt"
	"net/http"
	"os"
)

// type bot interface {
// 	getGreeting() string
// }

// type englishBot struct{}
// type spanishBot struct{}

func main() {
	// eb := englishBot{}
	// sb := spanishBot{}

	// PrintGreeting(eb)
	// PrintGreeting(sb)

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bs := make([]byte, 9999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}

// func PrintGreeting(b bot) {
// 	fmt.Println(b.getGreeting())
// }

// func (englishBot) getGreeting() string {
// 	// VERY custom logic for generating english greeting
// 	return "Hello!"
// }

// func (spanishBot) getGreeting() string {
// 	return "Hola!"
// }