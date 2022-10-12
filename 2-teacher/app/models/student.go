package models

type Group string

const HighPerformer Group = "A"
const LowPerformer Group = "B"

type Student struct {
	ID    int
	Name  string
	Point int
	Group Group
}

// AddBonus to student point based on it's ranking among other students
func (s *Student) AddBonus(ranking int) {
	if ranking < 2 {
		s.addTop2Bonus()
	} else if ranking < 7 {
		s.addTop7Bonus()
	} else if ranking < 11 {
		s.addTop11Bonus()
	}
}

func (s *Student) addTop2Bonus() {
	s.Point += 3
}

func (s *Student) addTop7Bonus() {
	s.Point += 2
}

func (s *Student) addTop11Bonus() {
	s.Point += 1
}
