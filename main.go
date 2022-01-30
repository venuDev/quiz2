package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type quizQuestion struct {
	question string
	answer   string
}

func main() {

	csvFName := flag.String("csv", "problems.csv", "file name that contains questions,answer in CSV format")

	flag.Parse()

	f, err := os.Open(*csvFName)

	if err != nil {
		safeexit(fmt.Sprintf("An error occured while opening the %s file. please find below error details.\n", *csvFName), err)

	}

	defer f.Close()

	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		safeexit("Failed to load lines", err)
	}
	ansCounter := 0
	for i, q := range parseQuizQuestions(lines) {

		fmt.Printf("Problem #%d: %s ?\n", i+1, q.question)
		var ans string
		fmt.Scanf("%s\n", &ans)
		if q.answer == strings.TrimSpace(ans) {
			ansCounter++
		}
		fmt.Printf("You Scored %d out of %d \n", ansCounter, len(lines))
	}

}

func parseQuizQuestions(lines [][]string) []quizQuestion {
	tmpqq := make([]quizQuestion, len(lines))

	for i, line := range lines {
		tmpqq[i] = quizQuestion{
			question: line[0],
			answer:   line[1],
		}
	}
	return tmpqq
}

func safeexit(msg string, err error) {

	fmt.Println(msg)

	if err != nil {
		log.Fatalln(err)
	}

}
