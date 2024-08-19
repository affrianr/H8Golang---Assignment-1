package main

import (
	"fmt"
	"os"
)

type Student struct {
	Name       string
	Address    string
	Occupation string
	Reason     string
}

func main() {
	students := []Student{
		{Name: "Thomas", Address: "Jakarta", Occupation: "Programmer", Reason: "Ingin belajar Golang"},
		{Name: "Sarah", Address: "Bandung", Occupation: "Designer", Reason: "Tertarik dengan Golang"},
		{Name: "Michael", Address: "Surabaya", Occupation: "Data Analyst", Reason: "Golang digunakan di project perusahaan"},
		{Name: "Jessica", Address: "Medan", Occupation: "Engineer", Reason: "Golang punya performa yang tinggi"},
		{Name: "David", Address: "Yogyakarta", Occupation: "Teacher", Reason: "Mau mengajarkan Golang di kelas"},
		{Name: "Emily", Address: "Semarang", Occupation: "Student", Reason: "Ingin menambah skill programming"},
		{Name: "Daniel", Address: "Malang", Occupation: "DevOps", Reason: "Golang cocok untuk system programming"},
		{Name: "Sophia", Address: "Denpasar", Occupation: "Researcher", Reason: "Golang efisien untuk data processing"},
		{Name: "James", Address: "Makassar", Occupation: "Entrepreneur", Reason: "Golang populer di startup tech"},
		{Name: "Olivia", Address: "Palembang", Occupation: "Consultant", Reason: "Klien banyak menggunakan Golang"},
	}

	if len(os.Args) < 2 {
		fmt.Println("Mohon masukkan nama")
		return
	}

	name := os.Args[1]

	found := false
	for i, student := range students {
		if student.Name == name {
			fmt.Println("ID: ", i)
			fmt.Println("Nama: ", student.Name)
			fmt.Println("Alamat: ", student.Address)
			fmt.Println("Pekerjaan: ", student.Occupation)
			fmt.Println("Alasan:", student.Reason)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Data tidak ditemukan")
	}
}
