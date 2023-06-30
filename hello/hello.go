package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/lqwangxg/go/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}
	// Request a greeting message.
	messages, err := greetings.Hellos(names)

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}
	lenArgs :=len(os.Args)
	if(lenArgs<3){
		fmt.Printf("os.Args Length = %d, 3 integer numbers are required.\n", lenArgs) 
		return;
	}
	t, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])
	y, _ := strconv.Atoi(os.Args[3])
	
	
	for i, arg := range os.Args {
		fmt.Printf("[%d] %s\n", i, arg)
	}
	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(messages)
	msg := greetings.RandomYSeed(t,b,y)
	fmt.Println(msg)
}
