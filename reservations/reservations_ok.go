package reservations

// ReservationsOk erwartet eine Route (Liste von Stationen (Strings))
// eine Liste von Reservierungen (jeweils ein Paar von Strings).
// Liefert true, falls alle Reservierungen in der Route erfüllt werden können,
// also falls für jede Reservierung beide Stationen in der richtigen Reihenfolge
// in der Route vorkommen und falls es keine sich überlappenden Reservierungen gibt.
func ReservationsOk(route []string, reservations [][2]string) bool {
	// Hinweis:
	// Verwenden Sie die ContainsInOrder-Funktion, um zu prüfen,
	// ob beide Stationen einer Reservierung in der richtigen Reihenfolge in der Route vorkommen.
	// Verwenden Sie die ContainsOverlap-Funktion, um zu prüfen,
	// ob sich zwei Reservierungen überlappen.
	// Verwenden Sie zwei verschachtelte for-Schleifen:
	// Die äußere Schleife durchläuft alle Reservierungen,
	// die innere Schleife prüft für jede Reservierung alle anderen Reservierungen auf Überlappung.
	// begin:solution
	for i, res1 := range reservations {
		s1 := res1[0]
		s2 := res1[1]
		if !ContainsInOrder(route, s1, s2) {
			return false
		}
		for j, res2 := range reservations {
			if i != j {
				t1 := res2[0]
				t2 := res2[1]
				if ContainsOverlap(route, s1, s2, t1, t2) {
					return false
				}
			}
		}
	}
	return true
	// end:solution
}
