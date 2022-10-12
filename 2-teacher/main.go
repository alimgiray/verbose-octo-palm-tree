package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/alimgiray/verbose-octo-palm-tree/teacher/app"
	"github.com/alimgiray/verbose-octo-palm-tree/teacher/app/models"

	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	printHelp()

	solver := app.Solver{}
	solver.
		NumberOfStudents(15).
		HighPerformers(10).
		MinPoints(-5).
		MaxPoints(5).
		Init()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter Command: ")
		scanner.Scan()

		command := scanner.Text()

		switch command {
		case "nh":
			solver.NextHour()

		case "nd":
			solver.NextDay()

		case "ls":
			students := solver.GetStudents()
			printStudents(students)

		case "q":
			os.Exit(0)

		case "h":
			printHelp()

		default:
			fields := strings.Fields(command)
			if len(fields) == 3 && fields[0] == "ap" {
				id, err := solver.ValidateUserID(fields[1])
				if err != nil {
					fmt.Println(err)
					continue
				}
				points, err := solver.ValidateMinMax(fields[2])
				if err != nil {
					fmt.Println(err)
					continue
				}
				solver.GivePointsToStudent(id, points)
			} else {
				fmt.Println("Wrong command")
			}
		}
	}
}

func printHelp() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Command", "Meaning", "Explanation"})
	t.AppendRows([]table.Row{
		{"nh", "Next Hour", "Advances time to next hour"},
		{"nd", "Next Day", "Advances time to next day"},
		{"ls", "List Students", "Prints current student points"},
		{"ap {id} {points}", "Add Points", "Adds specified to the student with given id"},
		{"h", "Help", "Prints this list"},
		{"q", "Quit", "Ends the program"},
	})
	t.Render()
}

func printStudents(students []*models.Student) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Name", "Points", "Group"})
	for _, student := range students {
		t.AppendRows([]table.Row{
			{student.ID, student.Name, student.Point, student.Group},
		})
	}
	t.Render()
}
