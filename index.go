package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	//                      Flag Name, Value       , Usage
	csvFileName := flag.String("csv", "problems.csv", "Its is in formate of Quesion and Answers")
	flag.Parse()
	/* Use of Flag :-
	It will now show anything when we run it simply
	like - go run index.go

	But It will show 3rd parameter usage when we run it as help
	like - go run index.go -h (or) -help
	*/
	_ = csvFileName

	file, err := os.Open(*csvFileName)
	if err != nil {
		ExitProject("It cann't read the file!")
	}
	_ = file

}

// Funtion for Exiting the project on Wrong CSV file input
func ExitProject(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
