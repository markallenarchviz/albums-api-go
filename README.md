# CRUD simples em GO

api simples com CRUD feito em memoria.

## Instalação

Para rodar a aplicação basta clonar o repositório e executar a partir do arquivo main.go

```bash
go run main.go
```

## Uso
A aplicação possui três rotas: GET /albums, GET /albums/:id e POST /albums

```bash

# '/albuns' retorna Json com uma lista de álbuns
curl localhost:8080/albums

# '/albuns/2' passando um ID retorna um Json com álbum especifico 
curl localhost:8080/albums/2

# '/albuns' com método POST adiciona um álbum a lista
curl -X POST localhost:8080/albums -H 'Content-Type: application/json' -d '{"id": "5","title": "Born To Die","artist": "Lana Del Rey","price": 80.99}'
