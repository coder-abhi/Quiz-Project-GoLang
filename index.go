package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	//                      Flag Name, Value       , Usage
	csvFileName := flag.String("csv", "problems.csv", "Its is in formate of Quesion and Answers")
	timeLimit := flag.Int("limit", 30, "Time limit for Quiz")
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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, que := range problems {
		fmt.Printf("\nProblem #%d : %s = ", i+1, que.q)
		answerCh := make(chan string)
		go func() {

			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("You Scored %d out of %d.", score, len(lines))
			return

		case answer := <-answerCh:
			if answer == que.a {
				score++
				fmt.Println("Correct!")
			}

		}
	}

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
