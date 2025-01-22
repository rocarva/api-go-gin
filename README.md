
# API Go Gin CRUD de Alunos

Este projeto é uma API desenvolvida em Go usando o framework [Gin](https://gin-gonic.com/) para gerenciar um CRUD (Create, Read, Update, Delete) de alunos. A API conecta-se a um banco de dados para armazenar e manipular os dados.

## Configuração e Execução do Projeto

### Pré-requisitos

- [Go](https://golang.org/) instalado na sua máquina.
- Um banco de dados configurado e compatível com o [GORM](https://gorm.io/).
- Dependências do projeto gerenciadas via `go mod`.

### Passos para Executar

1. Clone o repositório:
   ```bash
   git clone https://github.com/rocarva/api-go-gin.git
   cd api-go-gin
   ```

2. Instale as dependências:
   ```bash
   go mod tidy
   ```

3. Configure a conexão com o banco de dados no arquivo `database.go`:
   ```go
   package database

   import (
       "gorm.io/driver/sqlite" // ou o driver do seu banco
       "gorm.io/gorm"
   )

   var DB *gorm.DB

   func ConectaComBancoDeDados() {
       var err error
       DB, err = gorm.Open(sqlite.Open("alunos.db"), &gorm.Config{}) // Ajuste conforme seu banco
       if err != nil {
           panic("Erro ao conectar com o banco de dados")
       }
   }
   ```

4. Execute a aplicação:
   ```bash
   go run main.go
   ```

### Estrutura de Arquivos

- `main.go`: Ponto de entrada do projeto.
- `controllers/controller.go`: Contém os controladores para manipular as requisições da API.
- `models/alunos.go`: Define o modelo `Aluno` usado pelo GORM para interagir com o banco de dados.
- `routes/route.go`: Define as rotas da API e mapeia os endpoints para os controladores.

### Rotas da API

| Método  | Endpoint                  | Descrição                                    |
|---------|---------------------------|--------------------------------------------|
| `GET`   | `/alunos`                 | Retorna todos os alunos cadastrados.       |
| `GET`   | `/:nome/`                 | Retorna uma saudação personalizada.        |
| `POST`  | `/alunos`                 | Cria um novo aluno.                        |
| `GET`   | `/alunos/:id`             | Retorna os dados de um aluno por ID.       |
| `DELETE`| `/alunos/:id`             | Deleta um aluno pelo ID.                   |
| `PATCH` | `/alunos/:id`             | Atualiza os dados de um aluno pelo ID.     |
| `GET`   | `/alunos/cpf/:cpf`        | Retorna os dados de um aluno por CPF.      |

### Exemplo de Requisição

**Criar um novo aluno**
```bash
curl -X POST http://localhost:8080/alunos -H "Content-Type: application/json" -d '{
  "nome": "João Silva",
  "cpf": "123.456.789-00",
  "rg": "12.345.678-9"
}'
```

**Resposta:**
```json
{
  "message": "Aluno criado com sucesso",
  "aluno": {
    "ID": 1,
    "nome": "João Silva",
    "cpf": "123.456.789-00",
    "rg": "12.345.678-9"
  }
}
```

### Tecnologias Utilizadas

- Linguagem: Go
- Framework: Gin
- ORM: GORM
- Banco de Dados: SQLite (ou substitua pelo banco de sua escolha)

## Contribuindo

Se você quiser contribuir com este projeto, sinta-se à vontade para abrir uma issue ou um pull request.

## Licença

Este projeto é licenciado sob a [MIT License](LICENSE).
