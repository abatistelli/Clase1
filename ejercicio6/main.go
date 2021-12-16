/*

Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos,
para ello necesitan una aplicación que les permita calcular el descuento en base a 2 variables,
su precio y el descuento en porcentaje. Espera obtener como resultado el valor con el descuento aplicado y luego imprimirlo en consola.
Crear la aplicación de acuerdo a los requerimientos.

*/

package main

import "fmt"

const descuento = 0.25

func main() {

	precio := 1849.50

	precio -= precio * descuento
	fmt.Printf("Precio con descuento = %.2f ", precio)
}
