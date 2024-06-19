package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"text/tabwriter"
)

func main() {
	correctOnFirstGuess := 0
	correctOnSwap := 0

	for i := 0; i < 1000; i++ {
		choice, otherOption := simulate()
		if choice {
			correctOnFirstGuess++
		} else if otherOption {
			correctOnSwap++
		}
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "Correct on first guess:", "\t", strconv.Itoa(correctOnFirstGuess/10)+"%")
	fmt.Fprintln(w, "Correct on swap:", "\t", strconv.Itoa(correctOnSwap/10)+"%")
	w.Flush()
}

/*
*
Simulate the Monty Hall problem

@return choice bool, otherOption bool
*/
func simulate() (choice bool, otherOption bool) {
	options := shuffle([]bool{true, false, false})

	choice = options[0]

	// Ensure one false is filtered out
	if !choice && !options[1] {
		otherOption = options[2]
	} else {
		otherOption = options[1]
	}

	return choice, otherOption
}

/*
Quick n' dirty Fisher-Yates shuffle

@param items []T
@return []T
*/
func shuffle[T any](items []T) []T {
	for i := range items {
		j := rand.Intn(i + 1)
		items[i], items[j] = items[j], items[i]
	}

	return items
}
