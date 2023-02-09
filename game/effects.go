package game


func Nothing (s *State) {}

func BuildTower (s *State) {
	s.Money -= 4
	if (len(s.Village.TabBuild) != 6){
		s.Village.TabBuild = append(s.Village.TabBuild,[2]int{Tower,0})
	}
}

func BuildHouse (s *State) {
	s.Money -= 4
	s.Happiness += 1
	if (len(s.Village.TabBuild) != 6){
		s.Village.TabBuild = append(s.Village.TabBuild,[2]int{House,1})
	}
}

// M = minus, P = Plus
// Hap = Happiness, Mon = Money, Pop = Population
// always in the above order

func addEvents (s *State) {
	s.EventPool = []int{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 ,16 }
}

func P1Hap (s *State) {
	s.Happiness += 1
}

func M2HapP1Mon (s *State) {
	s.Happiness -= 2
	s.Money += 1
}

func M1HapP3Pop (s *State) {
	s.Happiness -= 1
	s.Population += 3
}

func P1HapP1Pop (s *State) {
	s.Population += 1
	s.Happiness += 1
}

func M1HapP1Mon (s *State) {
	s.Money += 1
	s.Happiness -= 1
}

func M1Hap (s *State) {
	s.Happiness -= 1
}

func M1Mon (s *State) {
	s.Money -= 1
}

func P1HapM2Mon (s *State) {
	s.Money -= 2
	s.Happiness += 1
}

func P1HapM1Mon (s *State) {
	s.Money -= 1
	s.Happiness += 1
}

func M5Mon (s *State) {
	s.Money -= 5
}

func M3Hap (s *State) {
	s.Happiness -= 3
}

func P4HapP10MonP3Pop (s *State) {
	s.Happiness += 4
	s.Money += 10
	s.Population += 3
}

func M3HapM6Pop (s *State) {
	s.Happiness -= 3
	s.Population -= 6
}

func M6HapM3Pop (s *State) {
	s.Happiness -= 6
	s.Population -= 3
}

func P4HapM7Mon (s *State) {
	s.Happiness += 4
	s.Money -= 7
}

func P2HapM4Pop (s *State) {
	s.Population -= 4
	s.Happiness += 2
}

func P25Mon (s *State) {
	s.Money += 25
}

func P50Mon (s *State) {
	s.Money += 50
}

func M50Mon (s *State) {
	s.Money -= 50
}

func M20Hap (s *State) {
	s.Happiness -= 20
}

func P1HapM7MonP3Pop (s *State) {
	s.Money -= 7
	s.Happiness += 1
	s.Population += 3
}

func M2HapP1MonM4Pop (s *State) {
	s.Happiness -= 2
	s.Money += 1
	s.Population -= 4
}

func P2HapM4MonP1Pop (s *State) {
	s.Happiness += 2
	s.Money -= 4
	s.Population += 1
}

func P1Mon (s *State) {
	s.Money += 1
}

func M1HapP1MonM1Pop (s *State) {
	s.Happiness -= 1
	s.Money += 1
	s.Population -= 1
}

func M4Mon (s *State) {
	s.Money -= 4
}

func M7HapM2Pop (s *State) {
	s.Happiness -= 7
	s.Population -= 2
}

func M14Mon (s *State) {
	s.Money -= 14
}

func M2Hap (s *State) {
	s.Happiness -= 2
}

func M1HapM2Pop (s *State) {
	s.Happiness -= 1
	s.Population -= 2
}

func M1Pop (s *State) {
	s.Population -= 1
}

func P1HapM5Mon (s *State) {
	s.Happiness += 1
	s.Money -= 5
}
