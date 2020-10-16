package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	// Get arguments
	min := flag.Int("min", 0, "The minimum number that computer will be played.")
	max := flag.Int("max", 100, "The maximum number that computer will be played.")
	attempts := flag.Int("attempts", 5, "The maximum attempt that you have to guess.")

	// Pares the flags
	flag.Parse()

	// Welcome message
	fmt.Println("*** Hello to the guessing game ***")

	// Intro
	fmt.Printf("You have %d attempts to guess the number between %d and %d\n(You can change these settings by -h switch)\n", *attempts, *min, *max)

REPLY:
	rand.Seed(time.Now().UnixNano())

	var rndNumber = rand.Intn(*max-*min+1) + *min

	for i := 0; i < *attempts; i++ {
		var guess int
		fmt.Println("What is your guess:")
		fmt.Scanf("%d\n", &guess)

		switch {
		case guess == rndNumber:
			fmt.Println("You WIN!")

			fmt.Println("Do you want to play again? Y/N")
			var command string
			fmt.Scanf("%s\n", &command)

			switch {
			case command == "y" || command == "Y":
				goto REPLY
			case command == "n" || command == "N":
				fmt.Println("Bye Bye!")
				os.Exit(0)
			}

		case guess > rndNumber && i < *attempts-1:
			fmt.Println("You have to guess a lower number")
		case guess < rndNumber && i < *attempts-1:
			fmt.Println("You have to guess a higher number")
		default:
			fmt.Println("You LOSE. You are a LOOOOOOOOOOOOSER!")
		}
	}
}
