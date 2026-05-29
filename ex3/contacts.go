package main

import "fmt"

// Personne
type Personne struct {
	Prenom, Nom string
	Age         int
	Email       string
}

func (p Personne) NomComplet() string {
	return fmt.Sprintf("%s %s", p.Prenom, p.Nom)
}

func (p Personne) Presentation() string {
	return fmt.Sprintf("Je m'appelle %s, j'ai %d ans. Email : %s", p.NomComplet(), p.Age, p.Email)
}

// Adresse
type Adresse struct {
	Rue, Ville, CodePostal string
}

func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

//Employé

type Employe struct {
	Personne
	Adresse
	Poste   string
	Salaire float64
}

func (e Employe) FicheEmploye() {
	fmt.Println("=== Fiche Employé ===")
	fmt.Println(e.Presentation())
	fmt.Printf("Poste : %s\n", e.Poste)
	fmt.Printf("Adresse : %s\n", e.Adresse.Format())
	fmt.Printf("Salaire : %.2f €\n", e.Salaire)
}

func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire += e.Salaire * pct / 100
}

// Etudiant
type Etudiant struct {
	Personne
	Promo   string
	Moyenne float64
}

func (et Etudiant) MentionObtenue() string {
	switch {
	case et.Moyenne >= 16:
		return "TB"
	case et.Moyenne >= 14:
		return "B"
	case et.Moyenne >= 12:
		return "AB"
	default:
		return "P"
	}
}

func main() {
	e1 := Employe{
		Personne: Personne{"Test", "Personne", 32, "test@test.com"},
		Adresse:  Adresse{"12 rue de la Paix", "Paris", "75000"},
		Poste:    "Développeur",
		Salaire:  3500,
	}
	e2 := Employe{
		Personne: Personne{"Bob", "Bob", 45, "bob@bob.com"},
		Adresse:  Adresse{"8 avenue des Lilas", "Lyon", "69000"},
		Poste:    "Chef de projet",
		Salaire:  4200,
	}

	et1 := Etudiant{
		Personne: Personne{"Clara", "Petit", 21, "clara@uni.fr"},
		Promo:    "L3 Info",
		Moyenne:  15.5,
	}
	et2 := Etudiant{
		Personne: Personne{"Allan", "Test", 20, "allan@uni.fr"},
		Promo:    "M2 Info",
		Moyenne:  11.0,
	}

	e1.AugmenterSalaire(10)
	e1.FicheEmploye()
	fmt.Println()
	e2.FicheEmploye()
	fmt.Println()
	fmt.Println("=== Étudiant 1 ===")
	fmt.Println(et1.Presentation())
	fmt.Printf("Promo : %s | Moyenne : %.2f — Mention : %s\n", et1.Promo, et1.Moyenne, et1.MentionObtenue())
	fmt.Println()
	fmt.Println("=== Étudiant 2 ===")
	fmt.Println(et2.Presentation())
	fmt.Printf("Promo : %s | Moyenne : %.2f — Mention : %s\n", et2.Promo, et2.Moyenne, et2.MentionObtenue())
}
