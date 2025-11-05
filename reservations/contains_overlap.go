package reservations

// ContainsOverlap erwartet eine Liste von Strings und zwei Paare von jeweils zwei Strings.
// Liefert true, falls die String-Paare sich in der Liste überlappen,
// also falls z.B. t1 zwischen s1 und s2 liegt.
// Anmerkung: Die Reihenfolge der Strings im Paar ist hier nicht relevant.
// Anmerkung: An den Grenzen (also s1 == t1 oder s2 == t2) liegt kein Überlappen vor.
func ContainsOverlap(list []string, s1, s2, t1, t2 string) bool {
	// Hinweis:
	// Verwenden Sie vier int-Variablen und bestimmen Sie zunächst die Positionen
	// der vier Strings in der Liste mittels einer for-Schleife.
	// Danach prüfen Sie die Positionen auf Überlappung.
	// begin:solution
	posS1 := -1
	posS2 := -1
	posT1 := -1
	posT2 := -1
	for pos, item := range list {
		if item == s1 {
			posS1 = pos
		}
		if item == s2 {
			posS2 = pos
		}
		if item == t1 {
			posT1 = pos
		}
		if item == t2 {
			posT2 = pos
		}
	}
	if posS1 == -1 || posS2 == -1 || posT1 == -1 || posT2 == -1 {
		return false
	}
	if (posS1 < posT1 && posT1 < posS2) || (posS1 < posT2 && posT2 < posS2) ||
		(posT1 < posS1 && posS1 < posT2) || (posT1 < posS2 && posS2 < posT2) {
		return true
	}
	return false
	// end:solution
}
