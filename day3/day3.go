package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type letterPriority struct {
	letter string
	priority  int
  }

type sackAndGroup struct {
	rucksack string
	groupNumber  int
  }

  var lowercasePrios []letterPriority
  var upperCasePrios []letterPriority

  var sackGroups []sackAndGroup

func main() {
	preparePriorities()
	readFile()
}

var totalSum = 0;
var totalSumPart2 = 0;

func readFile() {
	filePath := "d3input.txt"
	readFile, err := os.Open(filePath)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	if err != nil {
		fmt.Println(err)
	}
	
	//var currentGroup = 1;
	var index = 0;
	var sack1 = ""
	var commonCharacter []string
	//var prevIdex = 0;
	for fileScanner.Scan() {

		var rucksackString = fileScanner.Text()
		index++
		fmt.Println("starting loop at index", index)


		if(index == 1) {
			fmt.Println("index 1", rucksackString)
			sack1 = rucksackString
			
		} else if index == 2 {
			fmt.Println("index 2", rucksackString)

			for _, char := range sack1 {
				//fmt.Printf("character %c starts at byte position %d\n", char)

				if(strings.Contains(rucksackString, string(char))) {
					commonCharacter = append(commonCharacter,string(char))
					fmt.Println(string(char), "is in both sacks")
					//break
				}
			}
			
		} else if index == 3 {
			////part 2
			fmt.Println("index 3", rucksackString)

				for _, char := range commonCharacter {
				if(strings.Contains(rucksackString, string(char))){
					fmt.Println(char, "is also in sack 3")
					totalSumPart2 += getPriority(char)
					fmt.Println(totalSumPart2, "is now")
					break
				}
			}
		
			index = 0
			commonCharacter = nil
		}

		// if(index >= prevIdex + 3) {
		// 	prevIdex = index
		// 	currentGroup++
		// }

		// sackGroups = append(sackGroups, sackAndGroup{rucksackString,currentGroup})


		//get length of the string
		//halfLength := len(rucksackString) / 2

		//rucksack1 := rucksackString[0:halfLength]
		//rucksack2 := rucksackString[halfLength:]
    	//fmt.Println(rucksack1, rucksack2, "original", rucksackString)

		// for _, char := range rucksack1 {
		// 	//fmt.Printf("character %c starts at byte position %d\n", char, pos)

		// 	if(strings.Contains(rucksack2, string(char))) {
		// 		fmt.Println(string(char), "is in both sacks")
		// 		totalSum += getPriority(char)
		// 		fmt.Println(totalSum, "is now")
		// 		break
		// 	}
		// }
	}
	readFile.Close()

	fmt.Println(totalSumPart2, "is at the end")

	for _, sack := range sackGroups {
			fmt.Println("sack ",sack.rucksack, "is in group", sack.groupNumber)
		}
		
}

func getPriority(characterValue string) int {

	var arrayTocheck = lowercasePrios
	var returnValue = 0
	if(strings.ToUpper(characterValue) == characterValue) {
		arrayTocheck = upperCasePrios
	}

	for _, prioObject := range arrayTocheck {
		if(prioObject.letter == characterValue) {
			fmt.Println("priority of",characterValue, "is", prioObject.priority )
			returnValue += prioObject.priority
		}
	}	
	return returnValue
}

func preparePriorities() {
	var priorityLower = 1
	var priorityUpper = 27
	for i := 'a'; i <= 'z'; i++ {
		R := unicode.ToUpper(i)
		//fmt.Println(R, priority)
		stringR := fmt.Sprintf("%c", i)
		stringUpper := fmt.Sprintf("%c", R)
		lowercasePrios = append(lowercasePrios, letterPriority{stringR,priorityLower})
		upperCasePrios = append(upperCasePrios, letterPriority{stringUpper,priorityUpper})

		priorityUpper++
		priorityLower++
	}	
}

