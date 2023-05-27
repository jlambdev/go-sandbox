package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Useful for type hints & refactoring, e.g. we may want to read from JSON in the future
type problem struct {
    q string
    a string
}

func exit(msg string) {
    fmt.Println(msg)
    os.Exit(1)
}

func parseLines(lines [][]string) []problem {
    ret := make([]problem, len(lines)) // we know how big the value needs to be (number of rows)
    for i, line := range lines {
        ret[i] = problem{
            q: line[0],
            a: strings.TrimSpace(line[1]),
        }
    }
    return ret
}

/*
   Part 2 requirements:
    - add a timer that can be updated via a flag, default to 30 seconds
    - add an option to shuffle the quiz order each run
*/

func main() {
    // The `:=` syntax is shorthand for declaring and initialising a variable
    csvPath := flag.String("csv", "problems.csv", "A CSV file in the format of 'question,answer' (default \"problems.csv\")")
    timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds (default 30 seconds)")
    flag.Parse()

    // `*` dereferences the pointer and gets the value stored at that address in memory
	file, err := os.Open(*csvPath)
    if err != nil {
        exit(fmt.Sprintf("Error opening CSV file: %s\n", *csvPath))
    }

	// `defer` puts the function on the call stack and is guaranteed to
	// run next, regardless of how the `main` function exits
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        exit("Error reading CSV records")
    }

    var score int = 0

    problems := parseLines(records)

    timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

    // You can use this to explicitly break specific sections
    problemloop:
        for i, p := range problems {
            fmt.Printf("Problem #%d: %s = ", i + 1, p.q)

            answer, err := strconv.ParseInt(p.a, 0, 32) // https://gobyexample.com/number-parsing
            if err != nil {
                exit(fmt.Sprintf("Error parsing answer '%s' from record\n", p.a))
            }

            answerCh := make(chan int64)

            go func() {
                var userInput string
                _, err = fmt.Scanln(&userInput) // `&` symbol is used to get the memory address
                if err != nil {
                    exit(fmt.Sprintf("Error reading input '%s'\n", userInput))
                }

                i, err := strconv.ParseInt(userInput, 0, 32)
                if err != nil {
                    exit("Error parsing integer from user input")
                }

                answerCh <- i
            }()

            select {
            case <-timer.C:
                break problemloop
            case i :=<-answerCh:
                if i == answer {
                    fmt.Println("Correct!")
                    score++
                } else {
                    fmt.Println("Incorrect :(")
                }
            }
        }

    fmt.Printf("\nYou scored %d out of %d questions correctly.\n", score, len(problems))
}
