package golangbasic

import "fmt"

func GolangArray() {
	// array seperti ini memiliki batas 4 elemen
	// tapi tidak papa jika hanya memiliki 2 elemen
	var names [4]string

	names[0] = "Ahmad"
	names[1] = "Murteza"
	names[2] = "Akbari"
	// names[3] = "Den"
	// names[4] = "error" // jika elemennya melebihi dari 4 maka akan error

	fmt.Println(names)

	// array bisa juga dibuat dengan lebih mudah seperti berikut
	// tetap aturan sebelumnya jika array melebihi 4 elemen maka akan error
	var fruits = [4]string{
		"banana",
		"apple",
		"grape",
		"melon",
	}
	// var fruits = [4]string{"banana", "apple", "grape", "melon", "error"} // ini error karena memiliki 4 elemen
	fmt.Println(fruits)

	// array bisa dibuat dengan nilai tanpa jumlah dengan ...
	var numbers = [...]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// array multidemensi bisa dibuat dengan berikut
	var numbersMulti1 = [2][3]int{[3]int{1, 2, 3}, [3]int{3, 2, 1}}
	var numbersMulti2 = [2][3]int{{1, 2, 3}, {3, 2, 1}} // ini cara lebih singkat nya
	fmt.Println(numbersMulti1)
	fmt.Println(numbersMulti2)

	// untuk memanggil satu persatu kita bisa memggunakan perulangan contohnya kita pakai variabel fruits diatas
	fmt.Println("======= memanggil array dengan for")
	for i := 0; i < len(fruits); i++ {
		fmt.Println("Buah elemen", i, ":", fruits[i])
	}

	// cara lebih singkat bisa menggunakan for range
	fmt.Println("======= memanggil array dengan for range")
	for i, fruit := range fruits {
		fmt.Println("Buah elemen", i, ":", fruit)
	}

	// cara alokasi elemen array dengan make
	var animals = make([]string, 2)
	animals[1] = "Lion"
	animals[3] = "Fish"

	fmt.Println(animals)
}
