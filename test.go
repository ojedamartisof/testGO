package main

import "fmt"

func main() {
	z := []int{2, 1, 3, 5, 3, 2}
	fmt.Println(prueba(z))
}

func prueba(x []int) int {
	sol := len(x) // pongo como solucion inicial un indice grande, que no exista en el arreglo
	for i := 0; i <= len(x)-2; i++ {
		for j := i + 1; j <= len(x)-1; j++ {
			if x[i] == x[j] && j < sol { // busco siempre el menor indice
				sol = j // guardo el menor indice
			}
		}
	}
	if sol != len(x) { // pregunto si la solucion encontrada es diferente a la inicial

		return x[sol] // retorno el elemento del arreglo en ese indice
	}

	// en caso que la solucion siga siendo la incial es porque no encontro un duplicado entonces devuelvo -1
	return -1

}
