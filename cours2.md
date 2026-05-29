Déclaration 2 variables :

x, y := 10, “hello”

Echanges de variables :

```jsx
x, y := y, x
```

Déclaration plusieurs variables :

```jsx
var (
	serveur string = "localhost"
	port int = 8000
	actif bool = true
)
```

:= que dans une fonction

Variable non utilisé ⇒ ne compile pas

IOTA :

```jsx
type NiveauLog int
const (
	DEBUG NiveauLog = iota
	INFO
	WARNING
	ERROR
	CRITICAL
)
```

Retour multiple : retourner plusieurs valeurs ⇒ résultat + erreur

```jsx
resultat, err = diviser(10,3)
return // naked return

resultat2, _ = diviser(10,3) // on ignore l'erreur
```

closures : fonction qui se souvient de l’environnement dans lequel elle a été crée

```jsx
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
```

boucle while, for, for … range

Structure :

```go
// Definition d'un type struct
type Produit struct {
	Nom        string
	Prix       float64
	Stock      int
	Actif      bool
	Categories []string // un slice dans une struct, c'est possible!
}

// --- 3 facons d'initialiser une struct ---

// 1. Positionnelle (fragile si on ajoute des champs)
p1 := Produit("Iphone", 999.99, 50, true, []string{"high-tech"})

// 2. Nommée (RECOMMANDÉE - plus robuste aux changements)
p2 := Produit{
	Nom:   "MacBook Pro",
	Prix:  2499.99,
	Stock: 10,
	Actif: true,
	Categories: []string{"high-tech", "PC"},
}

// 3. Valeur zéro puis affectation
var p3 Produit
p3.Nom = "AirPods"
p3.Prix = 199.99
p3.Stock = 100
p3.Actif = true
```

Embed/ composition :

Classes embriqués ⇒ Voiture contient Marque/Roues et Moteur qui contient des pièces, …

```go
type Client struct {
	Personne // embed
	Adresse // embed
	Email string
	NumeroClient int
}
```

visibilité ⇒ déterminé par le 1er caractères

majuscule ⇒ exporté
miniscule ⇒ pas exporté

defer : toujours exécuté avant le return, libérer ressources/garbage collector

packages standards incontournables :

```go
fmt => format + affichage
strings => manipulation strings
strconv => conversions strings <=> types
math => calculs maths
sort => tri
time => date
```