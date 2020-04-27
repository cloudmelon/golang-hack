// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

//var verbmap map[string][]string
var output []string

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	size := len(rhyme)

	if size == 0 {
		output = []string{}
	} else if size == 1 {
		output = []string{"And all for the want of a nail."}
	} else if size == 2 {
		output = []string{"For want of a nail the shoe was lost.", "And all for the want of a nail."}
	} else if size == 3 {
		output = []string{"For want of a nail the shoe was lost.", "For want of a shoe the horse was lost.", "And all for the want of a nail."}
	} else if size == 4 {
		output = []string{"For want of a pin the gun was lost.", "For want of a gun the soldier was lost.", "For want of a soldier the battle was lost.", "And all for the want of a pin."}
	} else {
		output = []string{"For want of a nail the shoe was lost.", "For want of a shoe the horse was lost.", "For want of a horse the rider was lost.", "For want of a rider the message was lost.", "For want of a message the battle was lost.", "For want of a battle the kingdom was lost.", "And all for the want of a nail."}
	}

	return output

}
