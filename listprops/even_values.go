package listprops

// EvenValues gibt erwartet eine Liste von Zahlen und liefert die Anzahl der geraden Zahlen in der Liste.
func EvenValues(list []int) int {
	// Hinweis:
	// Verwenden Sie eine for-Schleife, um die Liste zu durchlaufen.
	// Prüfen Sie in jeder Iteration, ob das aktuelle Element gerade ist.
	// Erhöhen Sie den Zähler entsprechend.
	// begin:solution
	count := 0
	for _, el := range list {
		if el%2 == 0 {
			count++
		}
	}
	return count
	// end:solution
}
