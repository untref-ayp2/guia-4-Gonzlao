package ejercicios

import (
	"guia4/set"

	"github.com/agrison/go-commons-lang/stringUtils"
)

func Letras(s string) *set.Set[string] {
	s1 := set.NewSet[string]()

	for _, letra := range s {
		if stringUtils.IsAlpha(string(letra)) {
			s1.Add(string(letra))
		}
	}
	return s1
}
