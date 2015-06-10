package main

import "fmt"

func main(){

	var num,i int
	i = 100

	for (i >= 1){

		if (i%2 == 0){

			num = num + i
		}
		
		i = i - 1

	}

	fmt.Printf("La suma de los numeros es: %d",num)

}
