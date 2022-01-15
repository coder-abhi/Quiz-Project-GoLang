package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	topic := flag.String("topic", "general", "On Which Topic you want to start Quiz, options -> math, general")

	timeLimit := flag.Int("limit", 15, "Time limit for Quiz")
	flag.Parse()

	var problemSet []problem

	switch *topic {
	case "math":
		fmt.Println("In Swithch cse")
		problemSet = MathQuiz(*timeLimit)

	case "general":
		problemSet = GeneralQuiz(*timeLimit)
	}
	// fmt.Println(*csvFileName)

	var score int

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// for loop for printing question and scanning every answer
	for i, que := range problemSet {
		fmt.Printf("\nProblem #%d : %s = ", i+1, que.q)
		answerCh := make(chan string)
		go func() {

			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		// select funtion works like which channel give value first will execute
		select {
		case <-timer.C: // Execute when limit time will over
			fmt.Printf("\nYou Scored %d out of %d.\n", score, len(problemSet))
			return

		case answer := <-answerCh: // Execute when answer is given
			if answer == que.a {
				score++
				fmt.Println("Correct!")
			} else {
				fmt.Println("Wrong !!!")
			}
		}
	}

}

func MathQuiz(limit int) []problem {

	min := 0
	max := 100

	var mathSet [][]string

	i := 0
	for i < 10 {

		k := 0
		var mathQue string
		ans := 0
		for k < 2 {
			rand.Seed(time.Now().UnixNano())
			digit := rand.Intn(max-min+1) + min
			ans += digit
			mathQue += (strconv.Itoa(digit) + " ")
			k++
		}

		mathQue = strings.Replace(mathQue, " ", " + ", 1)
		mathAns := strconv.Itoa(ans)
		mathSet = append(mathSet, []string{mathQue, mathAns})
		i++
	}
	return parseProblem(mathSet)
}

func GeneralQuiz(limit int) []problem {

	var fileList []string
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Something is wrong, Program Can't Read files")
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".csv") {

			fileList = append(fileList, file.Name())
		}
	}

	fmt.Println("Choose Your Topic :")
	var inputFile int

	for i, file := range fileList {
		fmt.Println(i+1, file)
	}

	fmt.Scanln(&inputFile)
	csvFileName := fileList[inputFile-1]

	fmt.Printf("\n\n Your Selected Topic : %s\t\tSelected Time Limit : %d\n", csvFileName, limit)

	file, err := os.Open(csvFileName)
	if err != nil {
		ExitProject("It cann't read the file!")
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		ExitProject("Unble to Read CSV")
	}

	return parseProblem(lines)
}

// Converting input in our formate ( in GeneralQuiz() )
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
