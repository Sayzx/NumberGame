package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculGame() {
	//  générer un nombre comprit entre 1 et 100
	n := rand.Intn(100)
	nmax := 100
	nmin := 1
	fmt.Println("Trouvez le nombre entre 1 et 100")
	var choix int
	fmt.Scan(&choix)
	for choix != n {
		if choix > n {
			fmt.Println("Le nombre est plus petit")
			nmax = choix
		} else if choix < n {
			fmt.Println("Le nombre est plus grand")
			nmin = choix
		}
		fmt.Println("Trouvez le nombre entre", nmin, "et", nmax)
		fmt.Scan(&choix)
		for choix == n {
			fmt.Println("Bravo ! Vous avez trouvé le nombre !")
			fmt.Println("Voulez vous rejouer ?")
			var replay int
			fmt.Scan(&replay)
			if replay == 1 {
				CalculGame()
			} else {
				fmt.Println("Bonne Continuation le jeux va s'arreter")
				break
			}

		}
	}
}

func MemoryGame() {
	worldliste := []string{"Hello", "World", "Chaise", "Table", "Ordinateur", "Souris", "Clavier", "Tapis", "Lampe", "Livre", "Stylo", "Crayon", "Feuille", "Cahier", "Ciseaux", "Gomme", "Crayon de papier", "Crayon de couleur", "Crayon de bois", "Crayon"}
	fmt.Println(worldliste)
}

func DivGame() {

}

func PointeurGAME() {
	var ingame int
	var choixage int
	var age int
	var theorie int
	ingame = 1
	for ingame == 1 {
		theorie = age/2 + 7
		fmt.Println("Bienvenue sur la théorie du pointeur")
		fmt.Println("Elle va te dire avec qu'elle age maximum tu peu sortir avec d'etre dit comme pointeur")
		fmt.Println("Indique ton age: ")
		fmt.Scan(&age)
		fmt.Println(" ")
		fmt.Println("Vous pouvez donc sortir avec des personnes de: ")
		fmt.Print(theorie)
		fmt.Print("ans max")
		fmt.Println(" ")
		fmt.Println(" ")
		time.Sleep(2 * time.Second)
		fmt.Println("Voulez vous tester avec un autres age ?")
		fmt.Println(" 1 - OUI | 2 - NON")
		fmt.Scan(&choixage)
		if choixage == 1 {
			ingame = 1
		} else if choixage == 2 {
			ingame = 0
			fmt.Println("Bonne Continuation le jeux va s'arreter")
			break
		}

	}

}

func main() {
	fmt.Println("Bienvenue sur !| Number Games |!")
	fmt.Println("Choisissez un jeu :")
	fmt.Println("1 - Jeux de bingo | 2 - Jeux de mémoire | 3 - Jeux de division ou 4 BONUS")
	var choix int
	fmt.Scan(&choix)
	if choix == 1 {
		CalculGame()
	} else if choix == 2 {
		MemoryGame()
	} else if choix == 3 {
		DivGame()
	} else if choix == 4 {
		PointeurGAME()
	} else if choix != 1 || choix != 2 || choix != 3 {
		fmt.Println("Choix incorrect, veuillez choisir un jeu :")
		fmt.Println("1 - Jeux de bingo | 2 - Jeux de mémoire | 3 - Jeux de division")
		fmt.Scan(&choix)
	}
}
