# Aufgaben zur Validierung von Reservierungen

Dieses Verzeichnis enthält Aufgaben, die sich mit der Validierung von
Reservierungen befassen, wie sie z.B. bei der Buchung von Zugtickets auftreten können.

## Herangehensweise

I.W. muss hier für gegebene Reservierungen überprüft werden, ob sie zu einer
vorgegebenen Route passen. Dabei sind verschiedene Bedingungen zu überprüfen.

1. Beide Stationen einer Reservierung müssen auf der Route liegen.
2. Die Reihenfolge der Stationen in der Reservierung muss mit der Reihenfolge auf der Route übereinstimmen.
3. Es dürfen sich keine Reservierungen überschneiden, d.h. es darf keine zwei
   Reservierungen geben, bei denen die Endstation der einen Reservierung
   zwischen den beiden Stationen der anderen Reservierung liegt.

## Aufgaben-Übersicht

Es gibt eine Reihe von vorbereitenden Aufgaben, die allgemein für Listen und Strings
formuliert sind, die aber für die Verwendung im Rahmen der Reservierungs-Validierung gedacht sind.

| Aufgabe             | Beschreibung                                                       |
| ------------------- | ------------------------------------------------------------------ |
| contains_both       | Prüfen, ob beide Stationen einer Reservierung auf der Route liegen |
| contains_in_order   | Prüfen, ob die Stationen in der richtigen Reihenfolge liegen       |
| contains_overlap    | Prüfen, ob sich zwei Reservierungen überschneiden                  |
| reservations_ok     | Prüfen, ob eine Liste von Reservierungen gültig ist                |

## Anmerkung für Fortgeschrittene

Die Liste von Reservierungen in der letzten Aufgabe ist als zweidimensionales Array
modelliert, wobei jede Reservierung selbst wieder als Liste von zwei Strings
besteht (Start- und Endstation). Fortgeschrittene können sich überlegen, wie man
dafür einen besser passenden Datentyp definieren kann (z.B. eine eigene Struct oder
Paare von Strings verwenden).
