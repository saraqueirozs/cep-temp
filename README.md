# API de Clima por CEP

Sistema em Go que recebe um CEP, identifica a cidade e retorna o clima atual em Celsius, Fahrenheit e Kelvin. O sistema é publicado no Google Cloud Run.

## Funcionalidades

- Busca a cidade pelo CEP (ViaCEP).
- Obtém dados do clima (WeatherAPI).
- Converte temperatura para diferentes unidades.

## Requisitos

- **Go**: Versão 1.22 ou superior.
- **Docker**: contêiner.

### Variáveis de Ambiente

- `PORT`: Porta do servidor (padrão: 8080).
- `API_KEY`: Chave da WeatherAPI.

### Localmente

1. Configure as variáveis de ambiente.
2. Inicie a aplicação:
   ```bash
   go run main.go

docker build -t api-clima-cep e docker run -p 8080:8080 -e PORT=8080 -e API_KEY="sua_chave_api".

use com parametro: GET /?cep=12345678

A API está disponível publicamente no Cloud Run:
https://cep-temp-450904054836.us-central1.run.app/


