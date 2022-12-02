package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile()
}

//points
// 1 = rock
// 2 = paper
// 3 = scissors

// 0 - loss
// 3 - draw
// 6 - win

//a,x = rock
//b, y = paper
//c, z = scissors

//round2
//x = lose
//y = draw
//z = win

func readFile() {
	filePath := "d1input.txt"
	readFile, err := os.Open(filePath)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	if err != nil {
		fmt.Println(err)
	}

	var totalPoints = 0
	var result = ""
	var desiredResult = ""
	for fileScanner.Scan() {

		var round = strings.Split(fileScanner.Text(), " ")

		switch round[1] {
		case "X":
			fmt.Println("loss")
			desiredResult = "loss"
		case "Y":
			fmt.Println("draw")
			desiredResult = "draw"
		case "Z":
			fmt.Println("win")
			desiredResult = "win"
		}

		if round[0] == "A" {

			switch desiredResult {
			case "win":
				fmt.Println("Pick Y")
				totalPoints += 2
				totalPoints += 6
			case "draw":
				fmt.Println("Pick X")
				totalPoints += 1
				totalPoints += 3
			case "loss":
				fmt.Println("Pick Z")
				totalPoints += 3
			}
		}

		if round[0] == "B" {

			switch desiredResult {
			case "win":
				fmt.Println("Pick Z")
				totalPoints += 3
				totalPoints += 6
			case "draw":
				fmt.Println("Pick Y")
				totalPoints += 2
				totalPoints += 3
			case "loss":
				fmt.Println("Pick X")
				totalPoints += 1
			}
		}

		if round[0] == "C" {

			switch desiredResult {
			case "win":
				fmt.Println("Pick X")
				totalPoints += 1
				totalPoints += 6
			case "draw":
				fmt.Println("Pick Z")
				totalPoints += 3
				totalPoints += 3
			case "loss":
				fmt.Println("Pick Y")
				totalPoints += 2
			}
		}

		fmt.Println(round[0], "and", round[1], "result is", result, "total score is", totalPoints)
	}

	readFile.Close()
}
