// Exo 1 : Hello Go !
package main

import "fmt"
func main(){
	fmt.Println("Hello, Go world !")

	// Exo 2: Variables & types
	entier := 10
	chaine := "Bonjour"
	booleen := true
	flotant := 3.15

	fmt.Printf("%T, %v\n",entier,entier) //%T pour le type de la variable , %v pour la valeur de la variable
	fmt.Printf("%T, %v\n",chaine,chaine)
	fmt.Printf("%T, %v\n",booleen,booleen)
	fmt.Printf("%T, %v\n",flotant,flotant)

	// Exo 3 : Les constantes 
	const PI = 3.14
	const MaxUsers = 100000

	fmt.Printf("%v et %v sont les valeurs des constantes : PI et MaxUsers\n",PI,MaxUsers)

	// Exo 4 : Conditions simples
	age := 201
	if age >= 18 && age <=45{
		fmt.Printf("Vous êtes majeur !\n")
	}else if age > 45 && age <= 200{
		fmt.Printf("Vous êtes senior !\n")
	}else if age < 18 && age >= 0 {
		fmt.Printf("Vous êtes mineur !\n")
	}else{
		fmt.Printf("Vous n'existez pas !\n")
	}

	// Exo 5 : if avec initialisation
	if x := 5; x%2 == 0{
		fmt.Printf("%v est un nombre paire !\n",x)
	}else{
		fmt.Printf("%v est un nombre impaire !\n",x)
	}
	fmt.Printf("\n")

	// Exo 6 : Boucle Simple
	// afficher les nombres de 1 a 20
	for i := 0 ; i<20; i++{
		fmt.Printf("%v\n",i)
	} 
	fmt.Printf("\n")
	//afficher uniquement les nbr impaires
	for i := 0 ; i<20 ; i++{
		if i%2 != 0{
			fmt.Printf("%v\n",i)
		}
	} 
	fmt.Printf("\n")

	// afficher la somme de 1 à 5
	som := 0
	for i := 0 ; i<6 ; i++{
		som = som + i
	}
	fmt.Printf("La somme de 1 à 5 : %v\n",som)

	// Exo 7 : Break/Continue
	fmt.Println("Multiple de 0 à 1M :")
	for i:=1; i<1000000; i++{
		if i % 123457 == 0{
			fmt.Println(i,"Est divisible par 123457")
			break
		}
	}
	fmt.Printf("\n")

	// Exo 8 : Boucles imbriqués... étoiles.
	for i:=1; i<=4; i++{
		for j:=0; j<=i ; j++{
			if j == i{
				fmt.Printf("\n")
				continue
			}
			fmt.Printf("*")
		}
	} //La complexité est de 0(n^2)
	
}

