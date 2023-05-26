package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

/*
   Requirements:
    - read problems.csv by default, but allow a csv path to be passed in
    - iterate through the rows and keep track of how many questions they got right
    - support a -h flag, which will give help options for the user (-csv arg, -limit arg)
*/

func main() {
    // The `:=` syntax is shorthand for declaring and initialising a variable
	file, err := os.Open("problems.csv")
    if err != nil {
        fmt.Println("Error opening CSV file:", err)
        return
    }

	// `defer` puts the function on the call stack and is guaranteed to
	// run next, regardless of how the `main` function exits
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Error reading CSV records:", err)
        return
    }

    var score int = 0

    for row_num, record := range records {
        var question = record[0]

        var answer = record[1]
        a, err := strconv.ParseInt(answer, 0, 64) // https://gobyexample.com/number-parsing
        if err != nil {
            fmt.Println("Error parsing answer from record:", err)
            return
        }

        fmt.Printf("Problem #%d: %s = ", row_num + 1, question)

        var user_input string
        _, err = fmt.Scanln(&user_input) // `&` symbol is used to get the memory address
        if err != nil {
            fmt.Println("Error reading input:", err)
            return
        }

        i, err := strconv.ParseInt(user_input, 0, 64)
        if err != nil {
            fmt.Println("Error parsing integer from user input:", err)
            return
        }

        if i == a {
            fmt.Println("Correct!")
            score++
        } else {
            fmt.Println("Incorrect :(")
        }
    }

    fmt.Printf("Game over! You scored %d points.\n", score)
}
