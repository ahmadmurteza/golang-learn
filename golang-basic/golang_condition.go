package golangbasic

import "fmt"

func GolangCondition() {
	// if condition
	fmt.Println("=============== golang if else")
	var point = 8

	if point == 10 {
		fmt.Println("=============== Lulus dengan nilai sempurna")
	} else if point > 7 {
		fmt.Println("Lulus nilai kkm")
	} else if point > 4 {
		fmt.Println("Hampir lulus")
	} else {
		fmt.Println("Tidak lolos")
	}

	// temporary variable di if else
	fmt.Println("=============== golang if else with temporary variable")
	var point2 = 8840.0

	if percent := point2 / 100; percent == 100 { // temporary variable
		fmt.Println("Lulus dengan nilai sempurna")
	} else if percent > 70 {
		fmt.Println("Lulus")
	} else if percent > 40 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Println("Tidak lulus")
	}

	// switch case condition
	fmt.Println("=============== Switch condition")
	var point3 = 6

	switch {
	case (point3 == 10) || (point3 < 10):
		fmt.Println("Nilai sangat bagus")
	case point3 == 8:
		fmt.Println("Nilai bagus")
	case (point == 7) || (point == 6) || (point == 5):
		fmt.Println("Nilai cukup")
	default:
		{
			fmt.Println("Your ok")
			fmt.Println("Try again")
		}
	}

}
