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

func FahrenheitParaRankine(valor float64) float64 {
	return valor + 459.67
}

func FahrenheitParaKelvin(valor float64) float64 {
	return (valor-32)*5/9 + 273.15
}

func FahrenheitParaCelsius(valor float64) float64 {
	return (valor - 32) * 5 / 9
}
