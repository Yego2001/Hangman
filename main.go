package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//go:embed words.txt
var file string

func main() {
	name()

	startTime := time.Now()

	words := strings.Fields(file)

	index := rand.Intn(len(words))
	secretWord := words[index]
	errors := 0
	maxErrors := 8
	guess := make(map[int]bool)

	n := len(secretWord)

	for i := 0; i < n; i++ {
		fmt.Printf("_ ")
	}

	fmt.Println("               ")

	fmt.Printf("          \n💫Введи букву:💫 ")

	for {
		var vvod string
		fmt.Scanln(&vvod)
		fmt.Print("\033[H\033[2J")
		name()

		v := vvod[0]

		guess[int(v)] = true

		ok := false

		for i := 0; i < len(secretWord); i++ {
			if int(secretWord[i]) == int(v) {
				ok = true
			}
		}

		if ok == false {
			errors++
		}

		drawHangman(errors)

		if errors == maxErrors {
			fmt.Printf("\x1b[38;2;254;253;198m      \n😔Ты проиграл😔 \n💫Ничего страшного, Ты Умничка💫 \n\x1b[0m")
			fmt.Println("\x1b[38;2;254;253;198mДанное слово:\x1b[0m", secretWord)
			duration := time.Since(startTime)
			minutes := int(duration.Minutes())
			seconds := int(duration.Seconds()) % 60
			fmt.Printf("\n\x1b[38;2;255;192;203m\n⌛Ты отгадывал это слово:⌛ %d мин %d сек\n\x1b[0m", minutes, seconds)

			break
		}

		printWord(secretWord, guess)

		count := 0

		for i := 0; i < len(secretWord); i++ {
			if guess[int(secretWord[i])] {
				count++
			}
		}

		if count == len(secretWord) {
			fmt.Printf("\x1b[38;2;254;253;198m          \n😊Ты победил😊 \n⭐Молодец!⭐ \n          \x1b[0m")

			duration := time.Since(startTime)
			minutes := int(duration.Minutes())
			seconds := int(duration.Seconds()) % 60
			fmt.Printf("\x1b[38;2;255;192;203m\n⌛Ты отгадывал это слово:⌛ %d мин %d сек\n\x1b[0m", minutes, seconds)

			break
		}
	}
}

func drawHangman(errors int) {
	if errors == 1 {

		fmt.Println("         ")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("         ")
	} else if errors == 2 {

		fmt.Println("|---")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("         ")
	} else if errors == 3 {

		fmt.Println("|---")
		fmt.Println("|   0")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("         ")
	} else if errors == 4 {

		fmt.Println("|---")
		fmt.Println("|   0")
		fmt.Println("|   |")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("         ")
	} else if errors == 5 {

		fmt.Println("|---")
		fmt.Println("|   0")
		fmt.Println("|  /|")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("         ")
	} else if errors == 6 {

		fmt.Println("|---")
		fmt.Println("|   0")
		fmt.Println("|  /|\\ ")
		fmt.Println("|")
		fmt.Println("|")
		fmt.Println("         ")
	} else if errors == 7 {

		fmt.Println("|---")
		fmt.Println("|   0")
		fmt.Println("|  /|\\ ")
		fmt.Println("|  / ")
		fmt.Println("|")

	} else if errors == 8 {

		fmt.Println("|---")
		fmt.Println("|   0")
		fmt.Println("|  /|\\ ")
		fmt.Println("|  / \\")
		fmt.Println("|")

	}
}

func printWord(secretWord string, guess map[int]bool) {

	for i := 0; i < len(secretWord); i++ {
		if guess[int(secretWord[i])] {
			fmt.Print(string(secretWord[i]), " ")
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}

func name() {
	fmt.Println("")
	fmt.Println("\x1b[38;2;255;192;203m-----ВИСЕЛИЦА-----\x1b[0m")
	fmt.Println("")
}
