# API Booking

API REST escrita em **Go** para gerenciamento de usuários, eventos e autenticação, utilizando **PostgreSQL**, **JWT** e **Docker Compose**.

O projeto foi estruturado para rodar facilmente em ambiente local ou containerizado, com foco em simplicidade, clareza e boas práticas.

---

## 🚀 Tecnologias

* **Go** (>= 1.22 recomendado)
* **PostgreSQL**
* **Docker & Docker Compose**
* **JWT** para autenticação
* **Migrations SQL**

---

## 📁 Estrutura do Projeto

```text
.
├── main.go
├── go.mod
├── command
│   ├── middlewares
│   ├── private
│   │   ├── database
│   │   └── migrations
│   ├── routes
│   └── utils
├── models
├── containers
│   ├── Dockerfile
│   ├── compose.yaml
│   └── docker
│       └── postgres
│           └── init
└── requests.http
```

### Principais diretórios

* **main.go** → ponto de entrada da aplicação
* **command/routes** → definição das rotas HTTP
* **command/middlewares** → middlewares (auth, etc.)
* **command/utils** → JWT, hash, parser, helpers
* **models** → modelos de domínio
* **private/migrations** → migrations SQL do banco
* **containers/** → Dockerfile e docker-compose

---

## ⚙️ Configuração

### Variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto:

```env
APP_PORT=8000
JWT_SECRET=supersecret

DB_HOST=postgres_db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=booking
```

> ⚠️ **Nunca versionar o `.env`**. Use `.env.example` se necessário.

---

## 🐳 Rodando com Docker (recomendado)

A partir da **raiz do projeto**:

```bash
docker compose --env-file .env --file containers/compose.yaml up --build
```

Isso irá:

1. Buildar a imagem da API
2. Subir o PostgreSQL
3. Executar as migrations
4. Expor a API na porta configurada (ex: `8000`)

Para rodar em background:

```bash
docker compose --env-file .env --file containers/compose.yaml up -d
```

Para parar e remover tudo (inclusive volumes):

```bash
docker compose --env-file .env --file containers/compose.yaml down -v
```

---

## ▶️ Rodando sem Docker (local)

Requisitos:

* Go instalado (>= 1.22)
* PostgreSQL rodando

```bash
go mod tidy
go run main.go
```

---

## 🔐 Autenticação

A API utiliza **JWT**.

Fluxo típico:

1. Registro de usuário
2. Login
3. Receber token JWT
4. Enviar token no header:

```http
Authorization: Bearer <token>
```

---

## 📬 Requests de exemplo

O arquivo `requests.http` contém exemplos de chamadas para testes rápidos (compatível com VS Code / JetBrains).

---

## 🧪 Migrations

As migrations ficam em:

```text
command/private/migrations/
```

Elas são aplicadas automaticamente na inicialização do container do Postgres.

---

## 📌 Observações importantes

* O `Dockerfile` fica em `containers/`
* O `build.context` do Docker Compose aponta para a **raiz do projeto**
* `COPY . .` no Dockerfile copia o contexto, não a pasta do Dockerfile

---

## 🧠 Próximos passos (sugestões)

* Versionar API (v1)
* Adicionar testes
* Healthcheck
* CI/CD

---

## 🧑‍💻 Autor

Wilson Nascimento

---

## 📄 Licença

Este projeto é livre para uso e modificação.
