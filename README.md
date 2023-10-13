# CRUD simples em GO

api simples com CRUD feito em memoria.

## Instalação

Para rodar a aplicação basta apenas clonar o repositório e executar a partir do arquivo main.go

```bash
go run main.api
```

## Uso
A aplicação irá retornar alguns ao fazer um GET para a rota /albums

```bash

# '/albuns' retorna Json com de álbuns
curl localhost:8080/albuns

# '/albuns/2' passando um retorna Json com álbum especifico 
curl localhost:8080/albuns/2

# '/albuns' com método POST adiciona um álbuns a lista
curl -d "id=5&title=Born To Die&artist=Lana Del Rey&price=80.99" -X POST http://localhost:8080/albums
