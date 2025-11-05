package listprops

// OddValues gibt erwartet eine Liste von Zahlen und liefert die Anzahl der ungeraden Zahlen in der Liste.
func OddValues(list []int) int {
	// Hinweis:
	// Verwenden Sie eine for-Schleife, um die Liste zu durchlaufen.
	// Prüfen Sie in jeder Iteration, ob das aktuelle Element ungerade ist.
	// Erhöhen Sie den Zähler entsprechend.
	// Alternativ: Sie können auch die EvenValues-Funktion verwenden.
	// begin:solution
	count := 0
	for _, el := range list {
		if el%2 != 0 {
			count++
		}
	}
	return count
	// end:solution
}
