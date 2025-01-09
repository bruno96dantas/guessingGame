package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the number guessing game!")
	fmt.Println("Your goal is to guess the secret number between 0 and 100. You will have 10 attempts for this.")

	randNumber := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	attempts := [10]int64{}

	for i := range attempts {
		fmt.Println("What's your guess?")
		scanner.Scan()
		attempt := scanner.Text()
		attempt = strings.TrimSpace(attempt)

		attemptInt, err := strconv.ParseInt(attempt, 10, 64)
		if err != nil {
			fmt.Println("Your guess must be an integer between 0 and 100.")
			return
		}

		switch {
		case attemptInt < randNumber:
			fmt.Println("You made a mistake, the secret number is greater than", attemptInt)
		case attemptInt > randNumber:
			fmt.Println("You made a mistake, the secret number is less than", attemptInt)
		case attemptInt == randNumber:
			fmt.Printf(
				"Congratulations! You won! The number is: %d\n"+
					"You got it right in %d attempts.\n"+
					"These were your attempts: %v\n",
				randNumber, i+1, attempts[:i],
			)
			return
		}

		attempts[i] = attemptInt
	}

	fmt.Printf(
		"You lost! Unfortunately you didn't get the number right, which was: %d\n"+
			"You had 10 tries.\n"+
			"These were your tries: %v\n",
		randNumber, attempts,
	)
}
