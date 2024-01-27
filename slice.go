package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"Septemer",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7] //data ke 4-7 mei-juli
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	months[5] = "Diganti apa...," //mengubah slice
	fmt.Println(slice1)

}
