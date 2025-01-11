# Unit Converter

Unit Converter é uma aplicação simples em Go para converter unidades de temperatura, peso e distância.

## Funcionalidades

- Conversão de temperatura entre Celsius, Fahrenheit, Kelvin e Rankine.
- Conversão de peso entre Quilogramas, Libras e Onças.
- Conversão de distância entre Quilômetros, Milhas, Jardas e Metros.

## Como Executar

1. Clone o repositório:
    ```sh
    git clone https://github.com/fabiosoliveira/unit-converter.git
    cd unit-converter
    ```

2. Execute a aplicação:
    ```sh
    go run cmd/cli/main.go
    ```

3. Siga as instruções no terminal para escolher o tipo de conversão e o valor a ser convertido.

## Como Testar

Para rodar os testes unitários, execute:
```sh
go test ./pkg/conversion
```

## Licença
Este projeto está licenciado sob a Licença MIT. Veja o arquivo LICENSE para mais detalhes.


## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests.

## Autor

Fabio S. Oliveira