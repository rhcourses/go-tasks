package reservations

import "fmt"

func ExampleContainsOverlap() {
	list := []string{"Hamburg", "Hannover", "Göttingen", "Kassel", "Frankfurt", "Mannheim", "Karlsruhe"}

	// Überlappende Paare
	fmt.Println(ContainsOverlap(list, "Hamburg", "Kassel", "Göttingen", "Frankfurt"))
	fmt.Println(ContainsOverlap(list, "Göttingen", "Frankfurt", "Hamburg", "Kassel"))

	// Nicht überlappende Paare (normale Fälle)
	fmt.Println(ContainsOverlap(list, "Hamburg", "Hannover", "Göttingen", "Frankfurt"))
	fmt.Println(ContainsOverlap(list, "Hamburg", "Kassel", "Frankfurt", "Karlsruhe"))
	fmt.Println(ContainsOverlap(list, "Frankfurt", "Karlsruhe", "Hamburg", "Kassel"))
	fmt.Println(ContainsOverlap(list, "Karlsruhe", "Mannheim", "Frankfurt", "Göttingen"))

	// Nicht überlappende Paare (Grenzfälle, bei denen sich die Paare an den Enden berühren)
	fmt.Println(ContainsOverlap(list, "Karlsruhe", "Mannheim", "Mannheim", "Göttingen"))
	fmt.Println(ContainsOverlap(list, "Mannheim", "Göttingen", "Karlsruhe", "Mannheim"))

	// Output:
	// true
	// true
	// false
	// false
	// false
	// false
	// false
	// false
}
