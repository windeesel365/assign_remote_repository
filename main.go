package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/windeesel365/golang-quiz-git-game/data"
)

func shuffle(slice []string) {
	rand.Seed(time.Now().UnixNano())
	for i := range slice {
		j := rand.Intn(i + 1) //ใช้คู่กับ rand.Seed(time.Now().UnixNano()) เพื่อให้ได้ j มีค่าระหว่าง 0,1,2,3,..,i
		slice[i], slice[j] = slice[j], slice[i]
		//สลับกันด้วยindex
	}
}

func main() {
	//bufio.NewReader(os.Stdin)
	reader := bufio.NewReader(os.Stdin)
	correctAnswers := 0
	questionsToAsk := 5

	fmt.Println("**Welcome to git :) command quiz game!**")
	fmt.Println("Select correct git command for each description:")
	fmt.Println()

	for i := 0; i < questionsToAsk; i++ {
		questionIndex := rand.Intn(len(data.Commands))
		correctCommand := data.Commands[questionIndex]

		//ทำ options 4 ทางเลือก มีที่ถูก 1 อัน correctCommand.Command ใส่ก่อนในslice
		options := []string{correctCommand.Command}
		for len(options) < 4 {
			randomOption := data.Commands[rand.Intn(len(data.Commands))].Command
			if !contains(options, randomOption) {
				options = append(options, randomOption)
			}
		}
		shuffle(options)

		//ask question  until valid input A/B/C/D
		var isCorrect bool // isCorrect เป็น false อยู่ ยังไม่ valid
		for !isCorrect {   //falseก็  trueเพื่อให้loopไปเรื่อยๆ จนกว่า isCorrectเป็น true

			fmt.Printf("Question %d: %s\n", i+1, correctCommand.Description)
			for j, option := range options {
				fmt.Printf("%c: %s\n", 'A'+j, option)
			}

			//reader.ReadString เพื่อ read user input
			fmt.Print("Your answer (A/B/C/D): ")
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(strings.ToUpper(answer))

			if answer != "" && answer[0] >= 'A' && answer[0] <= 'D' {
				selectedOption := options[answer[0]-'A']
				if selectedOption == correctCommand.Command {
					fmt.Printf("Correct !!\n\n")
					correctAnswers++
				} else {
					fmt.Printf("Incorrect. The correct answer was: %s\n\n", correctCommand.Command)
				}
				isCorrect = true
			} else {
				fmt.Printf("invalid input. Please answer with A, B, C, or D.\n\n")
				//ไม่ valid ก็ไม่ valid ต่อไป ถามใหม่
			}
		}
	}

	fmt.Printf("You finished quiz! You answered %d out of %d questions correctly\n", correctAnswers, questionsToAsk)
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
