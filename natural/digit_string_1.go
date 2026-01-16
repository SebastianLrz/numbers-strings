package natural

// DigitString1 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Einer-Stelle einer Zahl >= 21 vorkommen würde.
// Außerdem wird bei Ziffern != 0 das Wort "und" angehängt.
//
// Beispiel: Für 3 soll der String "dreiund" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
// Diese Funktion muss nur für den Normalfall (Zahlen >= 21) funktionieren.
func DigitString1(digit int) string {

	if digit == 1 {
		return "einund"
	} else if digit == 2 {
		return "zweiund"
	} else if digit == 3 {
		return "dreiund"
	} else if digit == 4 {
		return "vierund"
	} else if digit == 5 {
		return "fünfund"
	} else if digit == 6 {
		return "sechsund"
	} else if digit == 7 {
		return "siebenund"
	} else if digit == 8 {
		return "achtund"
	} else if digit == 9 {
		return "neunund"
	} else if digit == 0 {
		return ""
	} else {
		return "Fehler"
	}

}
