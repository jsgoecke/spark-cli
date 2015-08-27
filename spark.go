package main

import (
    "fmt"
    "time"
)

func main() {
	fmt.Println("Retrieving audience contact information...\n")
	time.Sleep(3 * time.Second)
	fmt.Println("Initiating command confirmation...\n")
	time.Sleep(20 * time.Second)
	fmt.Println("Confirmation received.\n")
	time.Sleep(3 * time.Second)
	fmt.Println("Introducing myself...\n")
	time.Sleep(7 * time.Second)
	fmt.Println("Go ahead. Answer it. I'd really like to say hello.\n")
	time.Sleep(21 * time.Second)
	fmt.Println("2,234 voice calls answered.\n")
	time.Sleep(3 * time.Second)
	fmt.Println("521 minutes total duration.\n")
	time.Sleep(3 * time.Second)
	fmt.Println("Command 'introduce yourself' executed successfully.\n")
	time.Sleep(21 * time.Second)
	fmt.Println("Shall I call everyone in the Data Center session? y/n\n")
	time.Sleep(2 * time.Second)
}