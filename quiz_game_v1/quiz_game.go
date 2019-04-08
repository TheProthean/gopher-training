package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type results struct {
	correct int
	total   int
}

func main() {

	csvFileName := flag.String("csvfile", "problems.csv", "Name of csv file, where the quiz questions are located.")
	f, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Println("Error occured: ", err)
		os.Exit(1)
	}
	csvReader := csv.NewReader(bufio.NewReader(f))
	res := results{
		correct: 0,
		total:   0,
	}

	for {
		record, err1 := csvReader.Read()
		if err1 == io.EOF {
			break
		} else if err1 != nil {
			fmt.Println("Error occured: ", err)
			os.Exit(1)
		}
		askQuestion(record, &res)
	}
	fmt.Printf("Correct answers: %d. Total questions: %d.\n", res.correct, res.total)
}

func askQuestion(quizRecord []string, stats *results) {
	fmt.Printf("%s: ", quizRecord[0])
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	answer := scanner.Text()
	if strings.Compare(strings.ToLower(answer), quizRecord[1]) == 0 {
		stats.correct++
		fmt.Println("Correct!")
	} else {
		fmt.Println("Wrong!")
	}
	stats.total++
}
