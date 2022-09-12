package dicetower

import (
	"dice-roller/ui"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func Engine() {
	fmt.Println(ui.NumberOfDice)
	numberOfDice := getNumberOfDice(os.Stdin)

	fmt.Println(ui.DiceMenu)
	dieSize := getDieSize(os.Stdin)

	rollResults := []int{}
	total := 0

	fmt.Printf("Rolling %d d%ds...\n", numberOfDice, dieSize)
	for i := numberOfDice; i > 0; i-- {
		score := rollDie(dieSize)
		rollResults = append(rollResults, score)
		total += score
		time.Sleep(1 * time.Microsecond)
	}

	fmt.Printf("\nRoll results: %v \n", rollResults)
	fmt.Printf("Total: %d\n", total)
}

func rollDie(sides int) int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := sides

	return rand.Intn(max) + min
}

func getNumberOfDice(stdin io.Reader) int {
	var quantity int64

	_, err := fmt.Scanf("%d\n", &quantity)
	if err != nil {
			log.Fatal(err)
	} else if int(quantity) < 1 {
		fmt.Printf("number of dice (%d) must be an integer greater than 0", int(quantity))
		os.Exit(6)
	}

	return int(quantity)
}

func getDieSize(stdin io.Reader) int {
	var size int64
	var validDice = map[int]int{1:4,2:6,3:8,4:10,5:12,6:20,7:100}

	_, err := fmt.Scanf("%d\n", &size)
	if err != nil {
			log.Fatal(err)
	} else if !contains(validDice, int(size)) {
		fmt.Printf("number of sides (%d) must be %v", int(size), validDice)
		os.Exit(7)
	}

	return validDice[int(size)]
}

func contains(validDice map[int]int, input int) bool {
	for k := range validDice {
		if input == k {
			return true
		}
	}

	return false
}
