/*

Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin.

Por otro lado también es necesario:

-Saber cuántos de sus empleados son mayores a 21 años.
-Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
-Eliminar a Pedro del mapa.

*/

package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	mayores := 0
	fmt.Printf("La edad de Benjamin es: %d\n", employees["Benjamin"])

	for _, element := range employees {
		if element > 21 {
			mayores++
		}
	}

	fmt.Printf("Hay %d empleados mayores de 21 años\n", mayores)

	employees["Federico"] = 25

	fmt.Println(employees)

	delete(employees, "Pedro")

	fmt.Println(employees)
}
