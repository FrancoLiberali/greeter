package main

import "fmt"

// SayHello returns a greeting to the person called "name"
func SayHello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
