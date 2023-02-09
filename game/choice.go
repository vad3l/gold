package game

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type Choice struct {
	Id		int
	Title		string
	Description	string
}

func LoadChoices () []*Choice {
	file, err := os.Open("data/dialog/choices.csv")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return nil
	}	
	
	tab := []*Choice{}
	for _, value := range records {
		id, err := strconv.Atoi(value[0])
			if err != nil {
				log.Fatal(err)
				return nil
			}
		
		tab = append(tab,&Choice{
			id,
			value[1],
			value[2],
		}) 
	}

	return tab	
}
