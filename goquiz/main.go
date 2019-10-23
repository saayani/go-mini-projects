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
	csvFilename := flag.String("csv", "problems.csv", "A csv file in the format of question, answer.")
	flag.Parse()
	file, err := os.Open(*csvFilename) // csvFilename is a pointer to a string
	if err != nil {
		fmt.Println("Couldn't open the file: %s", *csvFilename)
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to the parse the csv file")
	}
	problems := parseLines(lines)
	fmt.Println(problems)
	timeLimit := flag.Int("limit", 5, "the time limit for the quiz in seconds")
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	counter := 0
	for i, problem := range problems {
		fmt.Println("Problem #%d: %d = ", i+1, problem.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer) // Reference to answer, to have a pointer value, to access it with the variable whenever the value is set
			// Use a closure to use data defined outside
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println("Your score is: %d out of %d", counter, len(problems))
			return
		case answer := <-answerCh:
			if answer == problem.a {
				counter++
				fmt.Println("Correct!")
			}
		}
	}
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
