package goarea

import "math"

// Pi representa o valor da razão entre a circunferência
// de qualquer círculo e seu diâmetro
const Pi = 3.1416

// Circ é responsável por calcular a área de uma circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsável por calcular a área de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visível!
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
