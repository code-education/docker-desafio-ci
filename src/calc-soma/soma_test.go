package main

import "testing"

func TestSoma(t *testing.T) {
    resultado := soma(5, 5)
    if resultado != 10 {
       t.Errorf("Soma está incorreta, resultado: %d, esperado: %d.", resultado, 10)
    }
}