package game

import (
	"log"
	"errors"
	"strings"
	"os"
	"fmt"
	"strconv"
	"time"
)

func CheckError (e error) {
	if e != nil {
		log.Fatal(e)
	}
}

type State struct {
	King_age	int
	Date		time.Time

	Happiness	int
	Money		int
	Population	int
	EventPool	[]int

	Village		*Village
	EventList	[]*Event
	ChoiceList	[]*Choice
	Effects		map[int]func(s *State)
}

func EndGame (s *State) bool {
	if s.Happiness <= 0 { return true }
	if s.Money <= 0 { return true }
	if s.Population <= 0 { return true }
	return false
}

func NewState () *State {
	return &State{
		15,
		time.Date(481, time.January, 1, 12, 0, 0, 0, time.UTC),
		100,
		100,
		100,
		[]int{ 0 },
		NewVillage(),
		LoadEvents(),
		LoadChoices(),
		LoadEffects(),
	}
}

func LoadEffects () map[int]func(s *State) {
	m := make(map[int]func(s *State))
	m[0] = addEvents
	m[1] = addEvents
	m[2] = P1Hap
	m[3] = M2HapP1Mon
	m[4] = M1HapP3Pop
	m[5] = P1HapP1Pop
	m[6] = M1HapP1Mon
	m[7] = M1Hap
	m[8] = M1Mon
	m[9] = P1HapM2Mon
	m[10] = Nothing
	m[11] = M1Hap
	m[12] = P1HapM1Mon
	m[13] = M1Hap
	m[14] = M5Mon
	m[15] = P4HapP10MonP3Pop
	m[16] = M3HapM6Pop
	m[17] = M6HapM3Pop
	m[18] = P4HapM7Mon
	m[19] = P2HapM4Pop
	m[20] = Nothing
	m[21] = P25Mon
	m[22] = P50Mon
	m[23] = M50Mon
	m[24] = M20Hap
	m[25] = P1HapM7MonP3Pop
	m[26] = M2HapP1MonM4Pop
	m[27] = P2HapM4MonP1Pop
	m[28] = P1Hap
	m[29] = P1Mon
	m[30] = M1HapP1MonM1Pop
	m[31] = M4Mon
	m[32] = M7HapM2Pop
	m[33] = M14Mon
	m[34] = M2Hap
	m[35] = M1HapM2Pop
	m[36] = M1Hap
	m[37] = M1Pop
	m[38] = BuildTower
	m[39] = M3Hap
	m[40] = BuildHouse
	m[41] = P1HapM5Mon 
	return m
}

func SaveGame (name string, s *State) {
	save := ""
	path := "save/" + name + ".sav"
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	CheckError(err)
	err = f.Truncate(0)
	CheckError(err)

	if s == nil {
		save += "15;01-01-481;5;5;5;\n0;\n"
	} else {
		save += strconv.Itoa(s.King_age) + ";"
		save += s.Date.Format("02-01-2006") + ";"
		save += strconv.Itoa(s.Happiness) + ";"
		save += strconv.Itoa(s.Money) + ";"
		save += strconv.Itoa(s.Population) + ";"
		save += "\n"
		for i := 0; i < len(s.EventPool); i++ {
			save += strconv.Itoa(s.EventPool[i]) + ";"
		}

		save += "\n"
		save += s.Village.SaveVillage()
	}
	_, err = f.WriteString(save)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadSave (name string) (*State, error)  {
	path := "save/" + name
	f, err := os.Open(path)
	if err != nil { return nil, err }
	info, err := f.Stat()
	if err != nil { return nil, err }
	size := info.Size()
	list := make([]byte, size)
	_, err = f.Read(list)
	if err != nil {
		return nil, err
	}

	data := string(list)
	things := strings.Split(data, "\n")
	tab := strings.Split(things[0], ";")
	events := strings.Split(things[1], ";")
	builds := things[2]

	if len(tab) <= 5 {
		return nil, errors.New("Incorrect file")
	}
	var state State
	state.King_age, err = strconv.Atoi(tab[0])
	if err != nil {
		return nil, err
	}

	state.Date, err = time.Parse("02-01-2006", tab[1])
	if err != nil {
		return nil, err
	}

	state.Happiness, err = strconv.Atoi(tab[2])
	if err != nil {
		return nil, err
	}

	state.Money, err = strconv.Atoi(tab[3])
	if err != nil {
		return nil, err
	}

	state.Population, err = strconv.Atoi(tab[4])
	if err != nil {
		return nil, err
	}

	state.Village = LoadVillage(builds)
	for i := 0; i < len(events) - 1; i++ {
		j, err := strconv.Atoi(events[i])
		if err != nil {
			return nil, errors.New("Invalid save")
		}
		state.EventPool = append(state.EventPool, j)
	}

	state.EventList = LoadEvents()
	state.ChoiceList = LoadChoices()
	state.Effects = LoadEffects()
	return &state, nil
}

func GetSaves () []string {
	path := "save/"
	saves := []string{}
	f, err := os.Open(path)
	CheckError(err)
	c, err := f.ReadDir(0)
	CheckError(err)
	for i := 0; i < len(c); i++ {
		ext := strings.Split(c[i].Name(), ".")
		if len(ext) < 2 {
			continue
		}
		if ext[1] != "sav" {
			continue
		}
		saves = append(saves, c[i].Name())
	}
	return saves
}

func RemoveSaves (fileName string){
	err := os.Remove("save/" + fileName)
	if err != nil {
		log.Fatalf("Failed to removes file :", err)
		return 
	}
	fmt.Println("File removed successfully :",fileName)
}

