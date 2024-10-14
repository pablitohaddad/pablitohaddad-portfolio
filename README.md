# Portfólio do [pablitohaddad](https://www.linkedin.com/in/pablohaddad/)

## Descrição
Este projeto é um portfólio interativo desenvolvido em **Go**, utilizando o framework **Gin** e o banco de dados **PostgreSQL**. O portfólio foi criado para exibir meus **projetos**, **certificações**, **tecnologias utilizadas**, e para permitir o envio de mensagens de contato.

## Funcionalidades
- **Usuários**: Gerenciamento de usuários com dados como nome, e-mail, bio, e senha.
- **Projetos**: Exibição e gerenciamento de projetos em que estou trabalhando ou que já concluí.
- **Certificações**: Registro de certificações obtidas, com informações como título e instituição.
- **Contato**: Um formulário para enviar mensagens de contato diretamente através do portfólio.
- **Tecnologias**: Exibição das tecnologias utilizadas nos projetos, com uma relação muitos-para-muitos entre projetos e tecnologias.
- **Notificação**: Um microserviço responsável por registrar as mudanças feitas nas entidades do sistema (ex. criação, atualização, exclusão).

## Tecnologias Utilizadas (Até agora)
<p align="center">
  <img align="center" alt="Golang" height="60" width="90" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original.svg" style="margin: 0 15px;">
  <img align="center" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" alt="Gin" width="70" height="60" style="margin: 0 15px;">
  <img align="center" src="https://upload.wikimedia.org/wikipedia/commons/2/29/Postgresql_elephant.svg" alt="PostgreSQL" width="90" height="75" style="margin: 0 15px;">
</p>

## Arquitetura do Projeto

<p align="center">
  <img src="https://github.com/user-attachments/assets/777aaef6-08b3-4a22-91b2-a6d317bc349c" alt="Arquitetura do Projeto" width="500">
</p>




## Como Executar o Projeto

### Pré-requisitos
- [Go](https://golang.org/doc/install) instalado na sua máquina.
- [PostgreSQL](https://www.postgresql.org/download/) instalado e rodando.
- [MongoDB](https://www.mongodb.com/try/download/community) instalado (para o serviço de notificação).
- [Docker](https://docs.docker.com/get-docker/) (opcional para facilitar o setup do ambiente).
- [RabbitMQ](https://www.rabbitmq.com/download.html) (para o sistema de mensageria).

### Passos para Executar o Projeto

#### 1. Clone o Repositório
```bash
git clone https://github.com/seu-usuario/seu-repositorio.git
cd seu-repositorio
```
### 2. Instale as Dependências
```bash
go mod tidy
```
### 3. Crie o banco de dados

```bash
psql -U postgres
CREATE DATABASE meu_portfolio_db;
```
### 4. Conecte-se no banco de dados alterando as informações no path database/connection.go
```
dsn := "host=localhost user=postgres password=suaSenha dbname=meu_portfolio_db port=5432 sslmode=disable"
```
4. Rodar as Migrações
Execute a aplicação para rodar as migrações e criar as tabelas no banco de dados:
```
go run main.go
```
5. Teste a Aplicação
Após iniciar o servidor, você pode acessar a aplicação no endereço:
```
http://localhost:8080
```

## Acompanhe o meu desenvolvimento! Até agora eu já fiz...

#### 1. **Conexão com o Banco de Dados**
   - [X] Implementar a conexão com o banco de dados **PostgreSQL** utilizando **GORM**.
   - [X] Testar a conexão para garantir que todas as tabelas estão sendo criadas corretamente no banco de dados.

#### 2. **Funcionalidades CRUD**
   - [ ] **Usuários**: Implementar as operações CRUD (Create, Read, Update, Delete) para gerenciar usuários.
   - [ ] **Projetos**: Implementar as operações CRUD para gerenciar projetos.
   - [ ] **Certificações**: Implementar as operações CRUD para gerenciar certificações.
   - [ ] **Contatos**: Implementar as operações CRUD para gerenciar mensagens de contato.
   - [ ] **Tecnologias**: Implementar as operações CRUD para gerenciar tecnologias.

#### 3. **Endpoints**
   - [ ] Desenvolver os endpoints RESTful para cada uma das funcionalidades acima, como:
     - **GET** `/users`: Retorna todos os usuários.
     - **POST** `/users`: Cria um novo usuário.
     - **GET** `/projects`: Retorna todos os projetos, etc.
   
#### 4. **Autenticação e Autorização**
   - [ ] Criar endpoints para autenticação de usuários (login/logout).
   - [ ] Implementar proteção das rotas com **JWT** para garantir que apenas usuários autenticados possam acessar e modificar dados.

#### 5. **Microserviço de Notificações**
   - [ ] O microserviço de notificações está configurado para consumir mensagens do sistema monolítico via **RabbitMQ** e registrar operações CRUD no **MongoDB**.
   - [ ] Configurar a troca de mensagens adequadamente.
   - [ ] Testar a integração com o microserviço de notificações.
"""
