package reservations

import "fmt"

func ExampleContainsBoth() {
	list := []string{"Hamburg", "Hannover", "Göttingen", "Kassel", "Frankfurt", "Mannheim", "Karlsruhe"}

	fmt.Println(ContainsBoth(list, "Hamburg", "Kassel"))
	fmt.Println(ContainsBoth(list, "Kassel", "Hamburg"))
	fmt.Println(ContainsBoth(list, "Berlin", "Kassel"))
	fmt.Println(ContainsBoth(list, "Hamburg", "Stuttgart"))
	fmt.Println(ContainsBoth(list, "Berlin", "Stuttgart"))

	// Output:
	// true
	// true
	// false
	// false
	// false
}
