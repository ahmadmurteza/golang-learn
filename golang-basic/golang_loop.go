package golangbasic

import "fmt"

func GolangLoop() {
	// use break
	fmt.Println("=============== use break to stop foreach at 7")
	for i := 1; i < 11; i++ {
		fmt.Println(i)
		if i == 7 {
			break
		}
	}

	// use continue
	fmt.Println("=============== use continue to skip foreach if is ood (ganjil)")
	for i := 1; i < 11; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
	}

	// use continue and break at sametime
	fmt.Println("=============== use continue and break at same time")
	for i := 1; i < 11; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}
		fmt.Println(i)
	}

	// label for loop
outerLoop:
	for i := 0; i < 11; i++ {
	innerLoop:
		for j := 0; j < 11; j++ {
			if j == 4 {
				break innerLoop
			}
			if i == 3 {
				break outerLoop
			}

			fmt.Println("matriks[", i, "][", j, "]", "\n")
		}
	}

}
