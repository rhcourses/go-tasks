package reservations

import "fmt"

func ExampleReservationsOk() {
	route := []string{"Hamburg", "Hannover", "Göttingen", "Kassel", "Frankfurt", "Mannheim", "Karlsruhe"}
	res_ok_1 := [][2]string{
		{"Hamburg", "Hannover"},
		{"Kassel", "Frankfurt"},
		{"Mannheim", "Karlsruhe"},
	}
	res_ok_2 := [][2]string{
		{"Hamburg", "Hannover"},
		{"Kassel", "Frankfurt"},
		{"Frankfurt", "Karlsruhe"},
	}
	res_missing := [][2]string{
		{"Hamburg", "Hannover"},
		{"Kassel", "Stuttgart"},
		{"Mannheim", "Karlsruhe"},
	}
	res_wrong_order := [][2]string{
		{"Hamburg", "Hannover"},
		{"Frankfurt", "Kassel"},
		{"Mannheim", "Karlsruhe"},
	}
	res_overlap := [][2]string{
		{"Hamburg", "Kassel"},
		{"Hannover", "Frankfurt"},
		{"Mannheim", "Karlsruhe"},
	}

	fmt.Println(ReservationsOk(route, res_ok_1))
	fmt.Println(ReservationsOk(route, res_ok_2))
	fmt.Println(ReservationsOk(route, res_missing))
	fmt.Println(ReservationsOk(route, res_wrong_order))
	fmt.Println(ReservationsOk(route, res_overlap))

	// Output:
	// true
	// true
	// false
	// false
	// false
}
