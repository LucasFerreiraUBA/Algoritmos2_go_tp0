package tp0

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	*x, *y = *y, *x
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	if len(vector) == 0 {
		return -1
	}
	var indiceMayor int
	for indice := indiceMayor; indice < len(vector); indice++ {
		if vector[indice] > vector[indiceMayor] {
			indiceMayor = indice
		}
	}
	return indiceMayor
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	long1, long2 := len(vector1), len(vector2)
	menorLongitud := ObtenerMenor(long1, long2)
	for indice := 0; indice < menorLongitud; indice++ {
		if vector1[indice] > vector2[indice] {
			return 1
		} else if vector1[indice] < vector2[indice] {
			return -1
		}
	}
	// Son iguales hasta el momento, pero puede que algún arreglo sea más largo que otro
	if long1 > long2 {
		return 1
	} else if long1 < long2 {
		return -1
	}
	return 0
}

func ObtenerMenor(l1, l2 int) int {
	if l1 >= l2 {
		return l2
	}
	return l1
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	if len(vector) <= 1 {
		return
	}
	inicio, fin := 0, len(vector)
	for i := inicio; i < fin; i++ {
		for j := i + 1; j < fin; j++ {
			if vector[j] < vector[i] {
				Swap(&vector[i], &vector[j])
			}
		}
	}
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	return 0
}

// EsPalindromo devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA.
func EsPalindromo(cadena string) bool {
	return false
}
