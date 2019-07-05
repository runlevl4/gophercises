package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

var numRight int
var numWrong int
var qa []QA

type QA struct {
	question string
	answer   int
}

func main() {
	testfile := flag.String("f", "problems.csv", "Test file")
	flag.Parse()
	loadTestFile(*testfile)
	doQuiz()
}

func loadTestFile(file string) {
	csvFile, err := os.Open(file)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", file))
	}
	reader := csv.NewReader(csvFile)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		tmp, _ := strconv.Atoi(line[1])
		qa = append(qa, QA{
			question: line[0],
			answer:   tmp,
		})
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func doQuiz() {
	var answer int
	fmt.Println("Your quiz starts now...\n")

	for i, _ := range qa {
		fmt.Println(qa[i].question)
		fmt.Scan(&answer)
		if answer == int(qa[i].answer) {
			numRight++
			fmt.Println("Correct!")
		} else {
			numWrong++
			fmt.Println("Whoops! Try harder next time.")
		}
	}

	per := math.Round((float64(numRight) / float64(len(qa))) * 100)
	if per > 70 {
		fmt.Printf("\nYou scored %.0f%% on this test with %d correct answers! Great job! :-D", per, numRight)
	} else {
		fmt.Printf("\nWith %d incorrect answers, you scored %.0f%% on this test. :-(", numWrong, per)
	}
}
