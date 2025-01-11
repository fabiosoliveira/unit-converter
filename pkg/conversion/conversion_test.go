package conversion

import (
	"math"
	"testing"
)

func TestKmParaM(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{1, 1000},
		{0.5, 500},
		{0, 0},
	}

	for _, test := range tests {
		result := KmParaM(test.input)
		if result != test.expected {
			t.Errorf("KmParaM(%f) = %f; expected %f", test.input, result, test.expected)
		}
	}
}

func TestKmParaYd(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{1, 1093.61},
		{0.5, 546.805},
		{0, 0},
	}

	for _, test := range tests {
		result := KmParaYd(test.input)
		if result != test.expected {
			t.Errorf("KmParaYd(%f) = %f; expected %f", test.input, result, test.expected)
		}
	}
}

func TestKmParaMi(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{1, 0.621371},
		{0.5, 0.3106855},
		{0, 0},
	}

	for _, test := range tests {
		result := KmParaMi(test.input)
		if result != test.expected {
			t.Errorf("KmParaMi(%f) = %f; expected %f", test.input, result, test.expected)
		}
	}
}

func TestKgParaOz(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{1, 35.274},
		{0.5, 17.637},
		{0, 0},
	}

	for _, test := range tests {
		result := KgParaOz(test.input)
		if result != test.expected {
			t.Errorf("KgParaOz(%f) = %f; expected %f", test.input, result, test.expected)
		}
	}
}

func TestKgParaLb(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{1, 2.20462},
		{0.5, 1.10231},
		{0, 0},
	}

	for _, test := range tests {
		result := KgParaLb(test.input)
		if result != test.expected {
			t.Errorf("KgParaLb(%f) = %f; expected %f", test.input, result, test.expected)
		}
	}
}

func TestCelsiusParaKelvin(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{0, 273.15},
		{-273.15, 0},
		{100, 373.15},
	}

	for _, test := range tests {
		result := CelsiusParaKelvin(test.input)
		if result != test.expected {
			t.Errorf("CelsiusParaKelvin(%f) = %f; expected %f", test.input, result, test.expected)
		}
	}
}

func TestCelsiusParaFahrenheit(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{0, 32},
		{-40, -40},
		{100, 212},
	}

	for _, test := range tests {
		result := CelsiusParaFahrenheit(test.input)
		if result != test.expected {
			t.Errorf("CelsiusParaFahrenheit(%f) = %f; expected %f", test.input, result, test.expected)
		}
	}
}

func TestCelsiusParaRankine(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{0, 491.67},
		{-273.15, 0},
		{100, 671.67},
	}

	for _, test := range tests {
		result := CelsiusParaRankine(test.input)
		result = math.Round(result*100) / 100 // Arredondar para 2 casas decimais
		if result != test.expected {
			t.Errorf("CelsiusParaRankine(%f) = %f; expected %f", test.input, result, test.expected)
		}
	}
}
