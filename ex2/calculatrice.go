package ex2

import "fmt"

func operer(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("Division par 0 impossible")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("Erreur lors de l'opération : %s", op)
	}
}

func creerOperation(op string) func(float64, float64) (float64, error) {
	return func(a, b float64) (float64, error) {
		return operer(a, b, op)
	}
}

func main() {
	fmt.Println("=== Calculatrice Go ===")
	fmt.Println("Quitter : quit")

	for {
		fmt.Print("> ")

		var token string
		fmt.Scan(&token)

		if token == "quit" {
			break
		}

		// Premier input
		var a float64
		_, err := fmt.Sscanf(token, "%f", &a)
		if err != nil {
			fmt.Println("Erreur : premier opérande invalide")
			continue
		}

		// Opérateur et deuxième nombre
		var op string
		var b float64
		fmt.Scan(&op)

		if op == "quit" {
			break
		}

		fmt.Scan(&b)

		// Opération
		operation := creerOperation(op)
		resultat, errOp := operation(a, b)

		if errOp != nil {
			fmt.Printf("Erreur : %v\n", errOp)
		} else {
			fmt.Printf("= %.2f\n", resultat)
		}
	}

	fmt.Println("Au revoir !")
}
