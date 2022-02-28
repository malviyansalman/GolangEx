package main

import (
	"GolangEx"
	"encoding/csv"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"os"
	"strings"
	"time"
)

var logger = GolangEx.GetLogger()

func main() {

	fileName := flag.String("fileName", "problems.csv", "Used for passing the problem name")
	timeLimit := flag.Int("limit", 5, "Timer for each Question")
	flag.Parse()
	file, err := os.Open(*fileName)
	if err != nil {
		logger.Info("Unable to Open file", zap.Any("Error", err))
	}
	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()
	if err != nil {
		logger.Info("Unable to read from CSV file", zap.Any("Error", err))
	}
	problems := parseCSV(lines)
	count := 0
	fmt.Println(*timeLimit)
	fmt.Println("Press Enter to Start the timer")
	fmt.Scanf("%v")
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	for i, prob := range problems {
		fmt.Printf("Problem # %d : %s = \n", i+1, prob.ques)
		ansChan := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s", &ans)
			ansChan <- ans
		}()
		select {
		case <-timer.C:
			fmt.Printf("You score %d out of %d\n", count, len(problems))
			return
		case ans := <-ansChan:
			if ans == prob.ans {
				count += 1
			}
		}
	}
	fmt.Printf("You score %d out of %d\n", count, len(problems))
}

func parseCSV(lines [][]string) []problem {
	size := len(lines)
	problems := make([]problem, size)
	for i, line := range lines {
		problems[i] = problem{
			ques: line[0],
			ans:  strings.TrimSpace(line[1]),
		}
	}
	return problems
}

type problem struct {
	ques string
	ans  string
}
