# CLI de Teste de Carga HTTP em Go / Go HTTP Stress Test CLI

## Português

### CLI de Teste de Carga HTTP em Go

Uma ferramenta de linha de comando (CLI) em Go para realizar testes de carga em serviços web.

### Objetivo

Este projeto tem como objetivo permitir a realização de testes de carga (stress test) em uma URL específica, simulando múltiplas requisições HTTP com diferentes níveis de concorrência.

### Como funciona

A aplicação executa um número total de requisições HTTP distribuídas de forma concorrente para uma determinada URL. Ao final do teste, é gerado um relatório com estatísticas da execução.

### Parâmetros (via CLI)

- `--url`: URL do serviço a ser testado.
- `--requests`: Número total de requisições a serem feitas.
- `--concurrency`: Número de chamadas simultâneas (nível de concorrência).

### Exemplo de execução

```
docker run <sua-imagem-docker> --url=http://google.com --requests=1000 --concurrency=10
```

**Observação**: substitua <sua-imagem-docker> pelo nome da imagem construída.

### O que é medido no relatório
Ao final da execução, será exibido um relatório contendo:

- Tempo total gasto na execução.
- Quantidade total de requisições realizadas.
- Quantidade de requisições com status HTTP 200 OK.
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).


## English

A command-line interface (CLI) tool written in Go to perform load testing on web services.

### Objective

This project aims to enable stress testing on a specific URL by simulating multiple HTTP requests with configurable concurrency levels.

### How it works

The application performs a total number of HTTP requests, distributed concurrently to a specified URL. At the end of the test, a report with execution statistics is generated.

### CLI Parameters

- `--url`: URL of the service to be tested.
- `--requests`: Total number of HTTP requests to perform.
- `--concurrency`: Number of concurrent requests (concurrency level).

### Example of execution

```
docker run <your-docker-image> --url=http://google.com --requests=1000 --concurrency=10
```

Note: Replace <your-docker-image> with the name of the image you built.

### What is included in the report
At the end of execution, the tool will print a report containing:

 - Total time taken to complete the test.
 - Total number of requests sent.
 - Number of requests that returned HTTP 200 OK.
 - Distribution of other HTTP status codes (e.g., 404, 500, etc.).