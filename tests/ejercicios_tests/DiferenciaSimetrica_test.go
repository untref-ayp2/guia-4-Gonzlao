package tests

import (
	"guia4/ejercicios"
	"guia4/set"
	"testing"
)

func TestDiferenciaSimetrica(t *testing.T) {
	s1 := set.NewSet(1, 2, 3, 4)
	s2 := set.NewSet(3, 4, 5, 6)
	s3 := ejercicios.DiferenciaSimetrica(s1, s2)
	if s3.Size() != 4 {
		t.Errorf("Tamaño %d, debería ser: %d", s3.Size(), 4)
	}
}

func TestNilpotnecia(t *testing.T) {
	a := set.NewSet(1, 2, 3, 4)
	s1 := ejercicios.DiferenciaSimetrica(a, a)

	if s1.Size() != 0 {
		t.Errorf("Tamaño %d, debería ser: %d", s1.Size(), 0)
	}
}

func TestAsociativa(t *testing.T) {
	a := set.NewSet(1, 2, 3, 4)
	b := set.NewSet(3, 4, 5, 6)
	c := set.NewSet(9, 4, 3, 8)

	s1 := ejercicios.DiferenciaSimetrica(ejercicios.DiferenciaSimetrica(a, b), c)
	s2 := ejercicios.DiferenciaSimetrica(a, ejercicios.DiferenciaSimetrica(b, c))

	if !(set.Equal(s1, s2)) {
		t.Errorf("s1: %v no es igual a s2: %v", s1, s2)
	}

}

func TestConmutativa(t *testing.T) {
	a := set.NewSet(1, 2, 3, 4)
	b := set.NewSet(3, 4, 5, 6)

	s1 := ejercicios.DiferenciaSimetrica(a, b)
	s2 := ejercicios.DiferenciaSimetrica(b, a)

	if !(set.Equal(s1, s2)) {
		t.Errorf("s1: %v no es igual a s2: %v", s1, s2)
	}
}

func TestNulo(t *testing.T) {
	a := set.NewSet(1, 2, 3, 4)
	b := set.NewSet[int]()

	s1 := ejercicios.DiferenciaSimetrica(a, b)

	if !(set.Equal(s1, a)) {
		t.Errorf("s1: %v no es igual a A: %v", s1, a)
	}

}
