package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	readFile()
}

// type elfInventory struct {
//     ListOfCalories []int
//     totalCalories  int
// }

func readFile() {
	filePath := os.Args[1] //"inputfile.txt"
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	//set initial counter
	var calories = 0
	var highestSofar = 0
	//this is bad, figure out how to make it nicer
	var calorieArray []int

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {

			if calories > highestSofar {
				highestSofar = calories
			}
			fmt.Println("calory count for this elf is:", calories, "appending to array of calories")
			calorieArray = append(calorieArray, calories)
			calories = 0
			fmt.Println("calory count is now:", calories)

			fileLines = append(fileLines, "this line was empty")
		} else {
			v := fileScanner.Text()
			if s, err := strconv.Atoi(v); err == nil {
				calories += s
			}

			fileLines = append(fileLines, fileScanner.Text())
		}

	}

	readFile.Close()

	sort.Sort(sort.Reverse(sort.IntSlice(calorieArray)))

	// for _, line := range calorieArray {
	//  fmt.Println(line)
	// }

	fmt.Println(calorieArray[0] + calorieArray[1] + calorieArray[2])
}
