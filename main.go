package main

import (
	"fmt"
	"os"
	"bufio"
	"time"
	"math/rand"
	"strings"
)


var complimentsByCategory = map[string][]string{
	"motivational": {
		"You are doing great",
		"Your consistency will pay off",
		"Your determination is admirable",
		"You are resilient",
		"Your perseverance is motivating",
		"You are unstoppable",
	},
	"confidence": {
		"You have a brilliant mind",
		"Your confidence is empowering",
		"You are courageous",
		"You are talented",
		"You are genuine",
		"You are wise",
	},
	"kindness": {
		"Your kindness inspires others",
		"You make people feel valued",
		"You are compassionate",
		"Your honesty is appreciated",
		"Your empathy is powerful",
		"Your kindness makes a difference",
	},
}


func Generator(compliments map[string][]string, category string) (string, error) {
	slice, exist := compliments[category]
	if !exist || len(slice) == 0 {
		return "", fmt.Errorf("category doesn't exist or invalid compliments")
	}
	// Seed number generator with current time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	indexR := r.Intn(len(slice))

	return slice[indexR], nil
}

func UserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))
	if err != nil {
		fmt.Println("Error reading user input")
		os.Exit(1)
	}
	return input
}


func main() {

	//input := UserInput()
	//sliceOfIndex := Generator(complimentsByCategory, input)
	fmt.Println("\033[41m<><> Welcome to the Compliment Generator <><>\033[0m")
	fmt.Println("Available categories: ")
	fmt.Println("1. Motivational")
	fmt.Println("2. Confidence")
	fmt.Println("3. Kindness")
	fmt.Println("Type a category and press ENTER (or type 'quit' to exit)")

	for {
    fmt.Print(">>> ")
    input := UserInput() // read new input each time
    if input == "quit" {
        fmt.Println("Goodbye!")
        break
    }

    compliment, err := Generator(complimentsByCategory, input)
	if err == nil{
		fmt.Println("\033[32m" + compliment + "\033[0m")
	}else{
		fmt.Println("error: unknown category")
	}
    
}

}


