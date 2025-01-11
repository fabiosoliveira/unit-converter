package conversion

func KmParaM(valor float64) float64 {
	return valor * 1000
}

func KmParaYd(valor float64) float64 {
	return valor * 1093.61
}

func KmParaMi(valor float64) float64 {
	return valor * 0.621371
}

func KgParaOz(valor float64) float64 {
	return valor * 35.274
}

func KgParaLb(valor float64) float64 {
	return valor * 2.20462
}

func CelsiusParaKelvin(valor float64) float64 {
	return valor + 273.15
}

func CelsiusParaFahrenheit(valor float64) float64 {
	return valor*9/5 + 32
}

func CelsiusParaRankine(valor float64) float64 {
	return (valor + 273.15) * 9 / 5
}
