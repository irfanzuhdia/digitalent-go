package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	Name    string
	Address string
	Job     string
	Reason  string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Gunakan: go run main.go [index atau nomor absen]")
		return
	}
	studentsList := []Student{
		{
			Name:    "Irfan",
			Address: "Jl. Raya Pahlawan No. 1",
			Job:     "Programmer",
			Reason:  "Belajar Golang"},
		{
			Name:    "Zuhdi",
			Address: "Jl. Raya Pahlawan No. 2",
			Job:     "Fresh Graduate",
			Reason:  "Persiapan bekerja"},
		{
			Name:    "Abdillah",
			Address: "Jl. Raya Pahlawan No. 3",
			Job:     "Data Scientist",
			Reason:  "Mengisi waktu luang"},
	}
	findStudents(os.Args, studentsList)
}

func findStudents(osArgs []string, studentsList []Student) {
	for i := 1; i < len(osArgs); i++ {
		j, err := strconv.Atoi(os.Args[i])
		if err != nil {
			fmt.Print(`Error: "`, os.Args[i], `" is not a number`)
			fmt.Println()
			fmt.Println()
		} else {
			if j > len(studentsList) {
				fmt.Println("No :", j, " not found")
				fmt.Println()
			} else {
				fmt.Println("No:", j)
				fmt.Println("Name:", studentsList[j-1].Name)
				fmt.Println("Address:", studentsList[j-1].Address)
				fmt.Println("Job:", studentsList[j-1].Job)
				fmt.Println("Reason:", studentsList[j-1].Reason)
				fmt.Println()
			}
		}
	}
}
