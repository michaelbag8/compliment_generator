package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var compliments = []string{
	"You are doing great",
	"Your consistency will pay off",
	"You have a brilliant mind",
	"Your smile brightens the room",
	"You are stronger than you think",
	"Your kindness inspires others",
	"You are full of creativity",
	"You make people feel valued",
	"Your determination is admirable",
	"You are a great listener",
	"Your positivity is contagious",
	"You bring out the best in others",
	"You are courageous",
	"Your ideas are refreshing",
	"You are dependable",
	"Your laughter is uplifting",
	"You are thoughtful",
	"Your energy motivates people",
	"You are resilient",
	"Your presence is comforting",
	"You are trustworthy",
	"Your hard work shows",
	"You are compassionate",
	"Your perspective is unique",
	"You are inspiring",
	"Your patience is remarkable",
	"You are resourceful",
	"Your honesty is appreciated",
	"You are supportive",
	"Your confidence is empowering",
	"You are generous",
	"Your dedication is impressive",
	"You are adaptable",
	"Your enthusiasm is refreshing",
	"You are wise",
	"Your humility is admirable",
	"You are uplifting",
	"Your optimism is encouraging",
	"You are reliable",
	"Your empathy is powerful",
	"You are talented",
	"Your spirit is radiant",
	"You are thoughtful",
	"Your perseverance is motivating",
	"You are genuine",
	"Your creativity sparks joy",
	"You are inspiring change",
	"Your kindness makes a difference",
	"You are unstoppable",
	"Your presence is a gift",
}

func Generator(compliments []string) string {
	if len(compliments) == 0 {
		fmt.Println("Compliment can no be empty")
	}
	// Seed number generator with current time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	indexR := r.Intn(len(compliments))

	return compliments[indexR]
}

func UserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading user input")
		os.Exit(1)
	}
	return input
}

func main() {
	fmt.Println("\033[41m<><> Welcome to the Compliment Generator <><>\033[0m")

	for {
		input := UserInput()
		randomC := Generator(compliments)
		fmt.Println("=== Press ENTER to receive a compliment === || \"quit\" to exit ")
	

		// to trim the \n from terminal
		input = strings.TrimSpace(input)
		if input == "quit" {
			fmt.Println("Goodbye!")
			break
		}

		fmt.Println("\033[32m" + randomC + "\033[0m")

	}
}
