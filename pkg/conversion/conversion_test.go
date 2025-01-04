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

func TestFahrenheitParaRankine(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{32, 491.67},
		{212, 671.67},
	}

	for _, test := range tests {
		result := FahrenheitParaRankine(test.input)
		result = math.Round(result*100) / 100 // Arredondar para 2 casas decimais
		if result != test.expected {
			t.Errorf("FahrenheitParaRankine(%f) = %f; want %f", test.input, result, test.expected)
		}
	}
}

func TestFahrenheitParaKelvin(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{32, 273.15},
		{0, 255.372},
		{100, 310.928},
	}

	for _, test := range tests {
		result := FahrenheitParaKelvin(test.input)
		result = math.Round(result*1000) / 1000 // Arredondar para 3 casas decimais
		if result != test.expected {
			t.Errorf("FahrenheitParaKelvin(%f) = %f; expected %f", test.input, result, test.expected)
		}
	}
}

func TestFahrenheitParaCelsius(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{32, 0},
		{0, -17.7778},
		{100, 37.7778},
	}

	for _, test := range tests {
		result := FahrenheitParaCelsius(test.input)
		result = math.Round(result*10000) / 10000 // Arredondar para 4 casas decimais
		if result != test.expected {
			t.Errorf("FahrenheitParaCelsius(%f) = %f; expected %f", test.input, result, test.expected)
		}
	}
}
