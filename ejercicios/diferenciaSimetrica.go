package ejercicios

import (
	"guia4/set"
)

func DiferenciaSimetrica[T comparable](s1, s2 *set.Set[T]) *set.Set[T] {
	d1 := set.Difference(s1, s2)
	d2 := set.Difference(s2, s1)

	return set.Union(d1, d2)
}
