package twofer

import "fmt"

// ShareWith function that prints a message with a passed in name that says who is sharing with whom.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
