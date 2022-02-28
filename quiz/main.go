package main

import (
	"GolangEx"
	"encoding/csv"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"math/rand"
	"os"
	"strings"
	"time"
)

var logger = GolangEx.GetLogger()

func main() {

	fileName := flag.String("fileName", "problems.csv", "Used for passing the problem name")
	timeLimit := flag.Int("limit", 5, "Timer for each Question")
	suffleFlag := flag.Bool("suffle", false, "Suffle the question default is falsr")
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
	if *suffleFlag == true {
		ShuffleProblems(&problems)
	}
	count := 0
	fmt.Println(*timeLimit, *suffleFlag)
	fmt.Println("Press Enter to Start the Timer")
	fmt.Scanf("%v")
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	//go spinner(100 * time.Millisecond)
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

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func ShuffleProblems(problems *[]problem) {
	size := len(*problems)
	rand.Seed(time.Now().UnixNano())
	for i := size - 1; i >= 0; i-- {
		idx := rand.Intn(i + 1)
		(*problems)[i], (*problems)[idx] = (*problems)[idx], (*problems)[i]
	}
}
