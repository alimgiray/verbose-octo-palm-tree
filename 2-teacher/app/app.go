package app

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strconv"

	"github.com/alimgiray/verbose-octo-palm-tree/teacher/app/models"

	faker "syreclabs.com/go/faker"
)

const endOfDay = 17

type Solver struct {
	hour                int
	day                 int
	students            []*models.Student
	maxPoints           int
	minPoints           int
	highPerformersLimit int
}

// MaxPoints for setting the maximum point the teacher can gave to students.
func (s *Solver) MaxPoints(max int) *Solver {
	s.maxPoints = max
	return s
}

// MinPoints for setting the minimum point the teacher can gave to students.
func (s *Solver) MinPoints(min int) *Solver {
	s.minPoints = min
	return s
}

// HighPerformers for setting limit for high performers (group A) ranking
func (s *Solver) HighPerformers(limit int) *Solver {
	s.highPerformersLimit = limit
	return s
}

// NumberOfStudents for setting number of students to work with.
func (s *Solver) NumberOfStudents(count int) *Solver {
	s.students = make([]*models.Student, count)
	return s
}

// Init the program
func (s *Solver) Init() {
	fmt.Println("Initializing the program")
	fmt.Println("Hours are limited between 9-16")
	s.day = 0
	s.hour = 9
	s.createStudents()
	s.printCurrentDateTime()
}

// NextHour advances time to next hour
func (s *Solver) NextHour() {
	s.hour += 1
	s.addBonusPoints()
	if s.hour == endOfDay {
		fmt.Println("Advancing to the next day")
		s.NextDay()
	} else {
		s.printCurrentDateTime()
	}
}

// NextDay advances time to next day
func (s *Solver) NextDay() {
	hourDiff := endOfDay - s.hour
	for i := 0; i < hourDiff; i++ {
		s.addBonusPoints()
	}

	s.hour = 9
	s.day += 1
	s.printCurrentDateTime()
	if s.day%7 == 0 {
		s.resetStudentPoints()
	}
}

// GetStudents returns current state of students
func (s *Solver) GetStudents() []*models.Student {
	return s.students
}

// printCurrentDateTime
func (s *Solver) printCurrentDateTime() {
	fmt.Printf("It's day %d and hour %d:00\n", s.day, s.hour)
}

// GivePointsToStudent adds points to given student
func (s *Solver) GivePointsToStudent(id, points int) {
	for _, student := range s.students {
		if student.ID == id {
			student.Point += points
			break
		}
	}
	s.determineGroups()
}

// addBonusPoints to students every hour
func (s *Solver) addBonusPoints() {
	for i, student := range s.students {
		student.AddBonus(i)
	}
	s.determineGroups()
}

// resetStudentPoints for the new week
func (s *Solver) resetStudentPoints() {
	fmt.Println("Resetting student points for the new week")
	for _, student := range s.students {
		student.Point = 0
		student.Group = ""
	}
}

// createStudents for initializing students
func (s *Solver) createStudents() {
	fmt.Printf("Creating %d new students with random initial points\n", len(s.students))
	for i := 0; i < len(s.students); i++ {
		s.students[i] = &models.Student{
			ID:    i + 1,
			Name:  faker.Name().Name(),
			Point: rand.Intn(s.maxPoints-s.minPoints) + s.minPoints,
			Group: "",
		}
	}
	s.determineGroups()
}

// determineGroups group students as high performers and low performers
func (s *Solver) determineGroups() {
	sort.Slice(s.students, func(i, j int) bool {
		return s.students[i].Point > s.students[j].Point
	})
	for i, student := range s.students {
		if i < s.highPerformersLimit {
			student.Group = models.HighPerformer
		} else {
			student.Group = models.LowPerformer
		}
	}
}

// ValidateMinMax to ensure user entered correct input
func (s *Solver) ValidateMinMax(input string) (int, error) {
	num, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return 0, errors.New("Please enter a number")
	}
	if int(num) < s.minPoints || int(num) > s.maxPoints {
		return 0, fmt.Errorf("Invalid number, please enter values between %d and %d", s.minPoints, s.maxPoints)
	}
	return int(num), nil
}

// ValidateUserID to ensure user entered correct input
func (s *Solver) ValidateUserID(input string) (int, error) {
	id, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return 0, errors.New("Please enter a number")
	}
	found := false
	for _, student := range s.students {
		if student.ID == int(id) {
			found = true
			break
		}
	}
	if !found {
		return 0, errors.New("Invalid number, please enter a valid student ID")
	}
	return int(id), nil
}
