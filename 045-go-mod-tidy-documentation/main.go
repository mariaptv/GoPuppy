
	package main

	import (
		"fmt"
		"github.com/mariaptv/GoPuppy/puppy"
	)

	func main() {
		s1 := puppy.Bark()
		s2 := puppy.Bark()

		fmt.Println(s1)
		fmt.Println(s2)

		fmt.Println(puppy.Bark())
	}
