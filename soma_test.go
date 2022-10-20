package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 10)

	if total != 25 {
		t.Errorf("Resultado da soma inválido: esperado %d porém retornou %d", 25, total)
	}
}
