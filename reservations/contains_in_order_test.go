package reservations

import "fmt"

func ExampleContainsInOrder() {
	list := []string{"Hamburg", "Hannover", "Göttingen", "Kassel", "Frankfurt", "Mannheim", "Karlsruhe"}

	fmt.Println(ContainsInOrder(list, "Hamburg", "Kassel"))
	fmt.Println(ContainsInOrder(list, "Kassel", "Hamburg"))
	fmt.Println(ContainsInOrder(list, "Berlin", "Kassel"))
	fmt.Println(ContainsInOrder(list, "Hamburg", "Stuttgart"))
	fmt.Println(ContainsInOrder(list, "Berlin", "Stuttgart"))

	// Output:
	// true
	// false
	// false
	// false
	// false
}
