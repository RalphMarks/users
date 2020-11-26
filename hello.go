package users

import "fmt"

// Hello to the user
func Hello(name string) string {
	return fmt.Sprintf("Hi, %v o/", name)
}
