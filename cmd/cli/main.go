package main

import (
	"fmt"

	"github.com/fabiosoliveira/unit-converter/pkg/conversion"
)

func main() {
	//  conversor de unidades (temperatura, peso, distância) usando entrada e saída no terminal.
	// 1. Perguntar ao usuário qual conversão ele deseja fazer
	// 2. Perguntar ao usuário qual valor ele deseja converter
	// 3. Exibir o resultado da conversão

	// Exemplo de execução:
	// Qual conversão você deseja fazer?
	// 1 - Temperatura
	// 2 - Peso
	// 3 - Distância
	// 1
	// Qual valor você deseja converter?
	// 32
	// 32ºF é igual a 0ºC
	// 32ºF é igual a 273.15K
	// 32ºF é igual a 491.67ºR
	// 32ºF é igual a 89.6ºC

	fmt.Println("Qual conversão você deseja fazer?")
	fmt.Println("1 - Temperatura")
	fmt.Println("2 - Peso")
	fmt.Println("3 - Distância")

	var opcao int
	fmt.Scan(&opcao)

	fmt.Println("Qual valor você deseja converter?")
	var valor float64
	fmt.Scan(&valor)

	switch opcao {
	case 1:
		// Temperatura
		fmt.Printf("%.2fºC é igual a %.2fºF\n", valor, conversion.CelsiusParaFahrenheit(valor))
		fmt.Printf("%.2fºC é igual a %.2fK\n", valor, conversion.CelsiusParaKelvin(valor))
		fmt.Printf("%.2fºC é igual a %.2fºR\n", valor, conversion.CelsiusParaRankine(valor))
	case 2:
		// Peso
		fmt.Printf("%.2fKg é igual a %.2fLb\n", valor, conversion.KgParaLb(valor))
		fmt.Printf("%.2fKg é igual a %.2fOz\n", valor, conversion.KgParaOz(valor))
	case 3:
		// Distância
		fmt.Printf("%.2fKm é igual a %.2fMi\n", valor, conversion.KmParaMi(valor))
		fmt.Printf("%.2fKm é igual a %.2fYd\n", valor, conversion.KmParaYd(valor))
		fmt.Printf("%.2fKm é igual a %.2fM\n", valor, conversion.KmParaM(valor))

	default:
		fmt.Println("Opção inválida")
	}
}
