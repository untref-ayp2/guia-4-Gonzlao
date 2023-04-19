package ejercicios

import (
	"guia4/set"
)

func Interseccion[T comparable](conjuntos ...*set.Set[T]) *set.Set[T] {
	s1 := conjuntos[0]
	for _, c := range conjuntos {
		s1 = set.Intersection(s1, c)
	}

	return s1
}
