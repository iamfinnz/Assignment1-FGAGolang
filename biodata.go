package main

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	Nama string
	Alamat string
	Pekerjaan string
	Alasan string
}

var dataTeman = map[int]Teman{
	1: {"Ilham Arifin", "Pekanbaru", "Freelancer", "Ingin berkarir sebagai Backend Developer"},
	2: {"Ibnu Muchda", "Rokan Hilir", "Mahasiswa", "Ingin memperdalam ilmu backend"},
	3: {"Afri Yoga", "Rokan Hulu", "Web Developer", "Ingin mengembangkan website backend"},
	4: {"Ichsanul Ikhwan", "Bukit Tinggi", "Animator", "Ingin switch career menjadi Backend Developer"},
	5: {"Ryan Chandra", "Pekanbaru", "Web Developer", "Ingin memperdalam bahasa Golang"},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Silahkan pakai perintah go run biodata.go (nomor 1-5)")
		return
	}

	nomor, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Nomor harus bilangan bulat")
		return
	}

	teman, found := dataTeman[nomor]
	if !found{
		fmt.Println("Data orang yang dicari tidak ditemukan")
		return
	}

	fmt.Println("Nama : ", teman.Nama)
	fmt.Println("Alamat : ", teman.Alamat)
	fmt.Println("Pekerjaan : ", teman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang : ", teman.Alasan)
}