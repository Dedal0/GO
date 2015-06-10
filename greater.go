package main

import "fmt"

func main(){

	var n1,n2,n3 int

	fmt.Println("Ingrese un numero: ")
	fmt.Scanf("%d",&n1)
	fmt.Println("Ingrese un numero: ")
	fmt.Scanf("%d",&n2)
	fmt.Println("Ingrese un numero: ")
	fmt.Scanf("%d",&n3)

	
	if (n1 > n2){

		if (n1 > n3){

			if (n2 > n3){

				fmt.Printf("%d es el mayor %d es el medio y %d es el menor",n1,n2,n3);

			}else{

				fmt.Printf("%d es el mayor %d es el medio y %d es el menor",n1,n3,n2);
			}
		}else{

			fmt.Printf("%d es el mayor %d es el medio y %d es el menor",n3,n1,n2);
		}
	}else{

		if (n2 > n3){

			if (n3 > n1){

				fmt.Printf("%d es el mayor %d es el medio y %d es el menor",n2,n3,n1);
			}else{

				fmt.Printf("%d es el mayor %d es el medio y %d es el menor",n2,n1,n3);
			}
		}else{
			
			fmt.Printf("%d es el mayor %d es el medio y %d es el menor",n3,n2,n1);
		}
	}

}
