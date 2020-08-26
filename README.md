<h1 align="center">
  <img style="background-color: #312e38; border-radius: 10px;" alt="starwars-logo" src="https://www.google.com/url?sa=i&url=https%3A%2F%2Fpt.wikipedia.org%2Fwiki%2FStar_Wars&psig=AOvVaw3QK9RX6EImtUMeDcO7FrP2&ust=1598566035622000&source=images&cd=vfe&ved=0CAIQjRxqFwoTCLDTg9jwuesCFQAAAAAdAAAAABAD" />
</h1>

## 🔖 Sobre o projeto

O projeto **Starwars**. Tendo como objetivo, uma API para CRUD de planetas.

## 💻 Tecnologias

- <img width="20px" src="https://img.icons8.com/color/2x/golang.png" /> [Golang](https://golang.org/ "Golang")
- <img width="20px" src="https://img.icons8.com/dusk/2x/docker.png" /> [Docker](https://www.docker.com/ "Docker")

## ▶️ Getting Started

- **Passo 1️⃣** : git clone do projeto [B2W](https://github.com/rafaelsanzio/b2w "B2W")
- **Passo 2️⃣** : executar a instalação do [Go](https://golang.org/ "Go") e [Docker](https://www.docker.com/ "Docker")

```bash
   # Navegando até a pasta do projeto
   $ cd b2w

   # Instalando as dependências do projeto
   $ go get .

   # Criando imagem do banco de dados (mongoDB) da aplicação
   $ docker run --name starwars -p 27017:27017 -d -t mongo

   #Iniciando banco de dados da aplicação
   $ docker start starwars

   # Gerando arquivo dump.json
   $ go run main.go
```

## ⚙️ Exemplificando rotas

```json
/* Requisição de criação de um planeta */
🟢 POST - /planets
params: {
 "name": "Naboo",
"weather": "hot",
"ground": "terrible"
}

/* Requisições para buscar um planeta a partir do id */
🟣 GET - /planets/{id}

/* Requisições para listagem de todos os planetas*/
🟣 GET - /planets - Lista todos os planetas
🟣 GET - /planets?name={name} - Lista todos os planetas filtrando por nome

/* Requisição deleção de um planeta */
🔴 DELETE - /planets/{id}

/* Requisição atualização de parâmetros de um planeta*/
🟠 PUT - /planets/{id}
params: {
 "name": "Naboo",
  "weather": "hot",
  "ground": "terrible"
}
```

## ㊗ ️ Considerações

- Projeto desenvolvido by:

  - <a href="https://github.com/rafaelsanzio">
      <img src="https://img.shields.io/badge/-Rafael%20Sanzio-000000?style=flat&logo=GitHub&logoColor=#000000" />
    </a>

  - <a href="https://www.linkedin.com/in/rafael-sanzio-012778143/">
      <img src="https://img.shields.io/badge/-Rafael%20Sanzio-0077B5?style=flat&logo=LinkedIN&logoColor=#000000" />
    </a>
