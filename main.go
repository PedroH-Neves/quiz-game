package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	fileName := flag.String("filename", "problems.csv", "csv de duas colunas perguntas e respostas")
	quizTime := flag.Int("timer", 30, "specify the quiz time")
	flag.Parse()

	fmt.Println("Welcome to the quiz game!!!")
	fmt.Println("The default time to answer the questions is 30 seconds")
	fmt.Println("Press ENTER to start the quiz")
	var enter string
	fmt.Scanln(&enter)
	fmt.Println("Time is running")

	timer := time.NewTimer(time.Duration(*quizTime) * time.Second)

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

questionLoop:
	for _, record := range records {
		question := record[0]
		answer := record[1]

		fmt.Println("Question:", question)

		answerCh := make(chan string)

		go func() {
			var userAnswer string
			fmt.Scan(&userAnswer)
			answerCh <- userAnswer
		}()

		select {
		case <-timer.C:
			fmt.Println("Sorry, the time is up!!")
			break questionLoop
		case userAnswer := <-answerCh:
			if userAnswer == answer {
				correctAnswers++
			}

		}
	}
	wrongAnswers := totalQuestions - correctAnswers
	fmt.Printf("This is your score:\n Right answers:%d\n Wrong answers:%d\n", correctAnswers, wrongAnswers)
}
