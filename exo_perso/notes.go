package main

import "fmt"

const (
	Faible = iota
	Moyen
	Bon
	Excellent
)

func main() {
	notes := make([]int, 3, 8)
	notes = append(notes, 45, 78, 92, 55, 88, 65, 91)

	for _, note := range notes {
		var categorie int
		var message string

		if note < 50 {
			categorie = Faible
		} else if note < 70 {
			categorie = Moyen
		} else if note < 85 {
			categorie = Bon
		} else {
			categorie = Excellent
		}

		switch categorie {
		case Excellent:
			message = "Bravo !"
			fallthrough
		case Bon:
			if message == "" {
				message = "C'est pas mal !"
			}
			fallthrough
		case Moyen:
			if message == "" {
				message = "À améliorer..."
			}
			fallthrough
		case Faible:
			if message == "" {
				message = "Pas ouf !"
			}
			fmt.Printf("Note: %d - %s\n", note, message)
		}
	}
}
