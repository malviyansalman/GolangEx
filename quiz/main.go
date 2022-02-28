package main

import (
	"GolangEx"
	"encoding/csv"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"os"
)

var logger = GolangEx.GetLogger()

func main() {

	fileName := flag.String("fileName", "problems.csv", "Used for passing the problem name")
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
	for i, prob := range problems {
		fmt.Printf("Problem # %d : %s = \n", i+1, prob.ques)
		var ans string
		fmt.Scanf("%s", &ans)
		if ans == prob.ans {
			count += 1
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
			ans:  line[1],
		}
	}
	return problems
}

type problem struct {
	ques string
	ans  string
}
