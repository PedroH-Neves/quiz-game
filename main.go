package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	fileName := flag.String("filename", "problems.csv", "csv de duas colunas perguntas e respostas")
	flag.Parse()

	fmt.Println("Welcome to the quiz game!!!")

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	totalQuestions := len(records)
	correctAnswers := 0

	for _, record := range records {
		question := record[0]
		answer := record[1]
		var userAnswer string
		fmt.Println("Question:", question)
		fmt.Scan(&userAnswer)
		correctAnswer := userAnswer == answer
		if correctAnswer {
			correctAnswers += 1
		}
		fmt.Println("Your answer is:", userAnswer)
	}
	result := totalQuestions - correctAnswers
	fmt.Printf("This is your score:\n Right answers:%d\n Wrong answers:%d\n", correctAnswers, result)
}
