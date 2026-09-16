package main

import "fmt"

type Soldat struct {
	nom     string
	vie     int
	attaque int
}

func afficherEquipe(equipe [6]Soldat) {
	fmt.Println("=== ÉQUIPE ===")
	fmt.Println()
	for _, s := range equipe {
		fmt.Println(s.nom)
		fmt.Printf("Vie : %d\n", s.vie)
		fmt.Printf("Attaque : %d\n\n", s.attaque)
	}
}

func trouverPlusDeVie(equipe [6]Soldat) Soldat {
	plusDeVie := equipe[0]
	for _, s := range equipe {
		if s.vie > plusDeVie.vie {
			plusDeVie = s
		}
	}
	return plusDeVie
}

func trouverPlusDAttaque(equipe [6]Soldat) Soldat {
	plusDAttaque := equipe[0]
	for _, s := range equipe {
		if s.attaque > plusDAttaque.attaque {
			plusDAttaque = s
		}
	}
	return plusDAttaque
}

func calculerVieMoyenne(equipe [6]Soldat) float64 {
	totalPV := 0
	for _, s := range equipe {
		totalPV += s.vie
	}
	return float64(totalPV) / float64(len(equipe))
}

func compterFaibles(equipe [6]Soldat) int {
	compteur := 0
	for _, s := range equipe {
		if s.vie < 800 {
			compteur++
		}
	}
	return compteur
}

func attaquerEquipe(equipe *[6]Soldat, degats int) {
	for i := 0; i < len(equipe); i++ {
		if equipe[i].vie > 0 {
			equipe[i].vie -= degats
			if equipe[i].vie < 0 {
				equipe[i].vie = 0
			}
		}
	}
}

func afficherEtat(equipe [6]Soldat) {
	for _, s := range equipe {
		if s.vie > 0 {
			fmt.Printf("%s : %d PV\n", s.nom, s.vie)
		} else {
			fmt.Printf("%s : KO\n", s.nom)
		}
	}
}

func compterVivants(equipe [6]Soldat, index int) int {
	if index >= len(equipe) {
		return 0
	}

	vivant := 0
	if equipe[index].vie > 0 {
		vivant = 1
	}

	return vivant + compterVivants(equipe, index+1)
}

func peutContinuer(equipe [6]Soldat) bool {
	return compterVivants(equipe, 0) > 0
}

func main() {
	equipe := [6]Soldat{
		{"Arthas", 1200, 250},
		{"Kael", 850, 320},
		{"Thrall", 1500, 180},
		{"Sylvanas", 700, 400},
		{"Garrosh", 1000, 280},
		{"Jaina", 500, 450},
	}

	afficherEquipe(equipe)

	soldatMaxVie := trouverPlusDeVie(equipe)
	soldatMaxAttaque := trouverPlusDAttaque(equipe)
	vieMoyenne := calculerVieMoyenne(equipe)
	nbFaibles := compterFaibles(equipe)

	fmt.Println("=== ANALYSE ===")
	fmt.Println()
	fmt.Printf("Soldat avec le plus de vie : %s\nVie : %d\n\n", soldatMaxVie.nom, soldatMaxVie.vie)
	fmt.Printf("Soldat avec la plus grande attaque : %s\nAttaque : %d\n\n", soldatMaxAttaque.nom, soldatMaxAttaque.attaque)
	fmt.Printf("Vie moyenne : %.2f\n\n", vieMoyenne)
	fmt.Printf("Soldats avec moins de 800 PV : %d\n\n", nbFaibles)

	fmt.Println("=== BATAILLE ===")
	fmt.Println()

	var nbAttaques int
	fmt.Print("Nombre d'attaques ennemies : ")
	fmt.Scan(&nbAttaques)
	fmt.Println()

	for i := 1; i <= nbAttaques; i++ {
		var degats int
		fmt.Printf("Attaque %d : ", i)
		fmt.Scan(&degats)
		fmt.Println()

		attaquerEquipe(&equipe, degats)

		fmt.Printf("=== APRÈS L'ATTAQUE %d ===\n\n", i)
		afficherEtat(equipe)
		fmt.Println()
	}

	vivants := compterVivants(equipe, 0)
	ko := len(equipe) - vivants

	fmt.Println("=== FIN DE LA BATAILLE ===")
	fmt.Println()
	fmt.Printf("Nombre de soldats vivants : %d\n", vivants)
	fmt.Printf("Nombre de soldats KO : %d\n\n", ko)

	if peutContinuer(equipe) {
		fmt.Println("L'équipe peut continuer le combat !")
	} else {
		fmt.Println("Tous les soldats sont KO...")
		fmt.Println("La bataille est terminée !")
	}
}
