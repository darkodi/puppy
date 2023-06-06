package puppy

import (
	"fmt"

	"github.com/darkodi/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBarks() string {
	return dog.WhenGrownUp(Bark())
}
func FromV1() {
	fmt.Println("I am from version 1.1.0")
}
