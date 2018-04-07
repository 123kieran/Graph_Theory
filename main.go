//Author: Kieran O'Halloran

package main

import (
	"bufio"
	"fmt"
	"os"
	//Import folder for using shunt.go & nfa.go
	files "./src/files"
)

func main() {
	option := 0
	exit := true

	//"While" this loop keeps the programme running until it is exited
	for exit {

		//Initial selection
		fmt.Print("\n***Please select an option:***\n1)Infix Expression to Postfix Expresssion\n2)Postix Regular Expresssion to NFA\n	3)Exit\n")
		fmt.Scanln(&option)

		//Enter inifix to postfix
		if option == 1 {
			fmt.Print("Please enter the Inifix expression to be converted:")
			//Read in string and use the function created in shunt.go to trim the string
			reader := bufio.NewReader(os.Stdin)
			expression, _ := reader.ReadString('\n')
			expression = files.TrimEndString(expression)

			//Output expression
			fmt.Println("Infix:  ", expression)
			//Pass expression into Intopost function and output result
			fmt.Println("Postfix: ", files.Intopost(expression))

			//Enter Postfix to NFA
		} else if option == 2 {

			fmt.Print("Please enter the Postfix expression to be converted:")
			////Read in string and use function created in shunt.go to trim the string
			reader := bufio.NewReader(os.Stdin)
			expression, _ := reader.ReadString('\n')
			expression = files.TrimEndString(expression)

			//Output expression
			fmt.Println("Postfix:  ", expression)
			//Pass expression into Poregtonfa function and output result
			fmt.Println("NFA: ", files.Poregtonfa(expression))
			fmt.Print("\n")

			//Ask user for basic string
			fmt.Print("Please enter a regular string to see if it matches the NFA:")
			////Read in string and use function created in shunt.go to trim the string
			regString, _ := reader.ReadString('\n')
			regString = files.TrimEndString(regString)
			regString = files.Intopost(regString)

			//Pass expression and basic string into PoPostFixMatchmatch

			//If returns true, strings match
			if files.PostFixMatch(expression, regString) == true {
				//Output result
				fmt.Println("The Regular string, ", regString, " matches the expression: ", expression)

				//If returns false, strings do not match
			} else if files.PostFixMatch(expression, regString) == false {
				fmt.Print("Sorry, String does not match")
				//Catch any errors
			} else {
				fmt.Print("An error occured")
			}

			//Exit program
		} else if option == 3 {
			fmt.Print("\nProgram Terminated\n")
			//Stop loop
			exit = false
		} else {
			//Catch any input errors
			fmt.Print("\nPlease enter a valid option.\n")
		}
	}
} //main
