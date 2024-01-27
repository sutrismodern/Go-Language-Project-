package main

import (
	"fmt"
	"os"
)

func main() {
	type Noktp string
	type Status bool

	var NoktpCt Noktp = "1213431412"
	var StatusKawin Status = true
	fmt.Println(NoktpCt)
	fmt.Fprintln(os.Stdout, []any{StatusKawin}...)

}
