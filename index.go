package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
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
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		ExitProject("Unble to Read CSV")
	}

	problems := parseProblem(lines)
	fmt.Println(problems)

	var score int

	for i, que := range problems {
		fmt.Printf("\nProblem #%d : %s = ", i+1, que.q)

		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == que.a {
			score++
			fmt.Println("Correct!")
		}
	}

	fmt.Printf("You Scored %d out of %d.", score, len(lines))
}

func parseProblem(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

// Struct for Question Answer pair
type problem struct {
	q string
	a string
}

// Funtion for Exiting the project on Wrong CSV file input
func ExitProject(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
