package main

import "fmt"

func main(){

	var dia int

	fmt.Printf("Ingrese el día: ")
	fmt.Scanf("%d",&dia)

	switch dia{

		case 1:
			fmt.Printf("Lunes")
		case 2:
			fmt.Printf("Martes")
		case 3:
			fmt.Printf("Miercoles")
		case 4:
			fmt.Printf("Jueves")
		case 5:
			fmt.Printf("Viernes")
		case 6:
			fmt.Printf("Sabado")
		case 7:
			fmt.Printf("Domingo")
	
	}

}
