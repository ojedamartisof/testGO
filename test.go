package main
//x - slice
//a la funcion le paso los parametros.
//desde el primero hasta el penultimo
//func main(){
//	z := []int{3,2,4,5,9}
//	fmt.Println(prueba(z))
//}
func prueba( x []int) int {
	sol := -1 //los dos puntos son cuando la variable no esta declarada
	for i := 0; i <= len(x) -2; i++{
		for j := i + 1; j <= len(x) -1; j++{
			if x[i] == x[j]{
				sol = x[i]
				break //si encontro el duplicado se va
			}
		}
			if  sol != -1 { //si encuentra el duplicado se va, no sigue buscando
				break
			}
	}
	return sol
}

