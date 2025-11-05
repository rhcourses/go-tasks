package listprops

// PrimeValues gibt erwartet eine Liste von Zahlen und liefert die Anzahl der Primzahlen in der Liste.
func PrimeValues(list []int) int {
	// Hinweis:
	// Verwenden Sie eine for-Schleife, um die Liste zu durchlaufen.
	// Prüfen Sie in jeder Iteration, ob das aktuelle Element eine Primzahl ist.
	// Schreiben Sie sich dazu ggf. eine Hilfsfunktion, z.B. "IsPrime".
	// Erhöhen Sie den Zähler entsprechend.
	// begin:solution
	count := 0
	for _, el := range list {
		if IsPrime(el) {
			count++
		}
	}
	return count
	// end:solution
}

// begin:solution:helpers

// IsPrime ist eine Hilfsfunktion, die prüft, ob eine Zahl eine Primzahl ist.
// Liefert true, falls die Zahl eine Primzahl ist, sonst false.
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// end:solution:helpers
