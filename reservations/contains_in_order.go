package reservations

// ContainsInOrder erwartet eine Liste von Strings und zwei Strings s1 und s2.
// Liefert true, falls sowohl s1 als auch s2 in der Liste enthalten sind, und s1 vor s2 kommt.
func ContainsInOrder(list []string, s1 string, s2 string) bool {
	// Hinweis:
	// Verwenden Sie zwei boolesche Variablen, um zu verfolgen,
	// ob s1 und s2 in der Liste gefunden wurden.
	// Verwenden Sie eine for-Schleife, um die Liste zu durchlaufen.
	// In jeder Iteration prüfen Sie, ob das aktuelle Element gleich s1 oder s2 ist.
	// Falls s2 vor s1 gefunden wird, können Sie sofort false zurückgeben.
	// Sobald beide gefunden wurden, können Sie true zurückgeben.
	// begin:solution
	foundS1 := false
	foundS2 := false
	for _, item := range list {
		if item == s1 {
			foundS1 = true
		}
		if item == s2 {
			if !foundS1 {
				return false
			}
			foundS2 = true
		}
		if foundS1 && foundS2 {
			return true
		}
	}
	return false
}
