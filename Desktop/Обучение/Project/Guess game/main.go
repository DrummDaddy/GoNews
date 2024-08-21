package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Я загадал случайное число от 1 до 100.")
	fmt.Println("Сможешь угадать?")

	reader := bufio.NewReader(os.Stdin)
	succes := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("У тебя есть", 10-guesses, "что бы угадать.")
		fmt.Print("Попробуй угадать: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Ой! Твое число меньше загаданного.")

		} else if guess > target {
			fmt.Println("Ой!Твое число больше загаданного.")

		} else {
			succes = true
			fmt.Println("Отлично! Ты угадал!")
			break
		}

	}
	if !succes {
		fmt.Println("Увы, ты не отгадал мое число. Это было:", target)
	}

}
