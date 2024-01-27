package main

import "fmt"

func main() {
	var names [3]string //angka 3 menunjukkan jumlah array

	names[0] = "nama1"
	names[1] = "nama2"
	names[2] = "nama3"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//cara langsung
	var nama = [4]string{
		"namaa",
		"namab",
		"namac",
		"namad",
	}
	fmt.Println(nama)

	fmt.Println(len(names)) //untuk mengetahui jumlah data []
	fmt.Println(len(nama))

}
