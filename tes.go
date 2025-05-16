package main

import "fmt"

func main() {
	var tobe, toba float64

	fmt.Scan(&tobe)

	if tobe > 1000 {
		toba = tobe * 0.8
	} else if tobe >= 500 && tobe <= 1000 {
		toba = tobe * 0.85
	} else {
		toba = tobe * 0.95
	}

	fmt.Print(toba)
}
