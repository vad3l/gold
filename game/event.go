package game

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
	"strings"
)


type Event struct {
	Id		int
	Title		string
	Description	string
	Choices		[]int
}

func LoadEvents () []*Event {
	file, err := os.Open("./data/dialog/events.csv")
	if err != nil {
		fmt.Println("Error opening file (events.csv) :", err)
		return nil
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Create a slice to store the contents of the file
	lines := make([]string, 0)

	// Read each line of the file
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// Check for any errors that may have occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file (events.csv) :", err)
		return nil
	}

	
	tab := []*Event{}

	for _, value := range lines {
		res := strings.Split(value,";")
		tabChoice := []int{}
		
		for i := 3 ; i < len(res); i++ {
			casted, err := strconv.Atoi(res[i])
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return nil
			}
			tabChoice = append(tabChoice,casted)
		}

		id, err := strconv.Atoi(res[0])
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return nil
			}
		
		tab = append(tab,&Event{
			id,
			res[1],
			strings.ReplaceAll(res[2], `\n`, "\n"),
			tabChoice,
		}) 
	}

	return tab	
}
