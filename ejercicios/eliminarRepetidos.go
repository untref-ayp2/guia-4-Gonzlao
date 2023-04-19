package ejercicios

import "guia4/set"

func EliminarRepetidos[T comparable](arreglo []T) []T {
	aux := set.NewSet(arreglo...)

	return aux.Values()
}
