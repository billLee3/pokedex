package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type cliCommand struct {
    name string
    description string
    callback func() error
}

func main(){
    scanner := bufio.NewScanner(os.Stdin)
	command := map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},	
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	}

    for true{
	fmt.Print("Pokedex > ")
	isInput := scanner.Scan()
	if isInput == false{
	    break
	}
	userInput := scanner.Text()
	cleanInputs := CleanInput(userInput)
	cleanInput := cleanInputs[0]
	_, ok := command[cleanInput]
	if !ok{
	    fmt.Println("Unkown Command")
		continue
	}
	if err := command[cleanInput].callback(); err != nil{
		fmt.Println("Callback failed")
		}
    }
}

func CleanInput(text string) []string{
    output := strings.ToLower(text)
    fields := strings.Fields(output)
    //fmt.Printf("Split fields: %q\n", fields)
    return fields
}

func commandExit() error{
   fmt.Println("Closing the Pokedex... Goodbye!\n")
   os.Exit(0)
   return nil
}

func commandHelp() error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: \n")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}
