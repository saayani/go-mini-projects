package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A csv file in the format of question, answer.")
	fmt.Println("vim-go")
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

	counter := 0
	for i, problem := range problems {
		fmt.Println("Problem #%d: %d = ", i+1, problem.q)
		var answer string
		fmt.Scanf("%s\n", &answer) // Reference to answer, to have a pointer value, to access it with the variable whenever the value is set
		if answer == problem.a {
			counter++
			//fmt.Println("Correct!")
		}
	}
	fmt.Println("Your score is: %d out of %d", counter, len(problems))
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
