package main

import "fmt"

func main() {
	text := "bojour vibe"
	age := 60

	if age < 50 {
		fmt.Printf(text)
		fmt.Printf("vous etes jeune")
	} else {
		fmt.Printf(text)
		fmt.Printf("vous etes vieux")
	}

}
