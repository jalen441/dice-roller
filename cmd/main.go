package main

import (
	"dice-roller/dicetower"
	"dice-roller/ui"
	"fmt"
	"log"
)

func main() {
	fmt.Printf("\x1bc")
	fmt.Print(ui.Intro)
	for {
		dicetower.Engine()
		tryAgain := tryAgain()
		if !tryAgain {
			break
		}
	}
}

func tryAgain() bool {
	tryAgain := ""

	fmt.Println(ui.Repeat)
	_, err := fmt.Scanf("%v\n", &tryAgain)
	if err != nil {
		log.Fatal(err)
	} 
	if (tryAgain == "Y" || tryAgain == "y") { 
    return true
	}

	return false
}