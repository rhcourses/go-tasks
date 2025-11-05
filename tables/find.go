package tables

// Find erwartet eine Liste und einen Wert.
// Sucht den Wert in der Liste und liefert die Position.
// Liefert -1, falls der Wert nicht in der Liste vorkommt.
func Find(list []string, v string) int {
	// Hinweis:
	// Verwenden Sie eine for-Schleife, um die Liste zu durchsuchen.
	// Idealerweise verwenden Sie eine "range"-Schleife, um sowohl
	// auf die Position als auch auf das Element zuzugreifen.
	// begin:solution
	for pos, el := range list {
		if el == v {
			return pos
		}
	}
	return -1
	// end:solution
}
