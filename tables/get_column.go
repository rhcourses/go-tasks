package tables

// GetColumn erwartet ein zweidimensionales Array (Tabelle) und eine Spaltennummer.
// Liefert eine Liste mit den Werten der angegebenen Spalte.
// Falls die Zeilen unterschiedliche Längen haben, wird für fehlende Werte ein leerer String geliefert.
func GetColumn(table [][]string, col int) []string {
	// Hinweis:
	// Verwenden Sie eine for-Schleife, um die Zeilen der Tabelle zu durchlaufen.
	// Greifen Sie in jeder Zeile auf das Element der angegebenen Spalte zu.
	// Prüfen Sie dabei jeweils die Länge der Zeile, um Indexfehler zu vermeiden.
	// begin:solution
	var result []string
	for _, row := range table {
		if col < len(row) {
			result = append(result, row[col])
		} else {
			result = append(result, "")
		}
	}
	return result
	// end:solution
}
