package app_test

import (
	"testing"

	"github.com/alimgiray/verbose-octo-palm-tree/teacher/app/models"

	"github.com/alimgiray/verbose-octo-palm-tree/teacher/app"

	"github.com/stretchr/testify/assert"
)

var solver *app.Solver

func init() {
	solver = &app.Solver{}
	solver.
		NumberOfStudents(15).
		HighPerformers(10).
		MinPoints(-5).
		MaxPoints(5).
		Init()
}

func TestMinMax(t *testing.T) {
	_, err := solver.ValidateMinMax("-10")
	assert.NotNil(t, err, "It should give an error for wrong points")

	_, err = solver.ValidateMinMax("10")
	assert.NotNil(t, err, "It should give an error for wrong points")

	num, err := solver.ValidateMinMax("1")
	assert.Nil(t, err, "It shouldn't give an error for point in range")
	assert.Equal(t, 1, num, "It should be able to correctly parse given point")
}

func TestStudentID(t *testing.T) {
	_, err := solver.ValidateUserID("-1")
	assert.NotNil(t, err, "It should give an error for wrong student ID")

	_, err = solver.ValidateUserID("19")
	assert.NotNil(t, err, "It should give an error for wrong student ID")

	num, err := solver.ValidateUserID("1")
	assert.Nil(t, err, "It shouldn't give an error for correct student ID")
	assert.Equal(t, 1, num, "It should be able to correctly parse given student ID")
}

func TestStudentPoints(t *testing.T) {
	students := solver.GetStudents()
	student := students[0]

	oldPoint := student.Point
	bonusPoints := 5
	expectedPoint := oldPoint + bonusPoints

	solver.GivePointsToStudent(student.ID, bonusPoints)

	assert.Equal(t, expectedPoint, student.Point, "New points should be added correctly to the student")
}

func TestBonusPoints(t *testing.T) {
	students := solver.GetStudents()
	student := students[0]
	oldPoint := student.Point

	topPerformerBonus := 3
	expectedPoint := oldPoint + topPerformerBonus

	solver.NextHour()
	assert.Equal(t, expectedPoint, student.Point, "Bonus points should be added correctly to the student")
}

func TestGroupChange(t *testing.T) {
	students := solver.GetStudents()
	highPerformer := students[9]
	lowPerformer := students[10]

	assert.Equal(t, models.LowPerformer, lowPerformer.Group, "Student 10 should be in the LowPerformer (B) group")

	for i := 0; i < ((highPerformer.Point-lowPerformer.Point)/5)+1; i++ {
		solver.GivePointsToStudent(lowPerformer.ID, 5)
	}

	assert.Equal(t, models.LowPerformer, highPerformer.Group, "Student 10 should be promoted to HighPerformer (A) group")
	assert.Equal(t, models.HighPerformer, lowPerformer.Group, "Student 10 should be demoted to LowPerformer (B) group")
}
