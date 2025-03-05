# Projeto para finalização do curso de Clean Arch e K8S

## Descrição

A Kitchen Control API é uma aplicação para gerenciar clientes, produtos, pedidos e itens de pedidos. Esta API fornece endpoints para criar, buscar, atualizar e deletar registros.

## Tecnologias

- Go
- GORM
- MySQL
- Gin Web Framework

## Instalação

### Pré-requisitos

- wsl 2, macos ou linux
- Go 1.22.1 ou superior
- Docker
- Git
- make

### Passos para Instalação

1. Clone o repositório:
    ```bash
    git clone https://github.com/caiojorge/fiap-challenge-ddd.git
    cd fiap-challenge-ddd
    ```

2. Instale as dependências:
    ```bash
    go mod tidy
    ```

3. Execute as migrações do banco de dados:
    ```bash
    make fiap-run (ou docker-compose up -d)
    
    ```
- O arquivo init-db esta conectado ao docker, e deve ser executado para criar o banco de dados caso não exista

4. Execute os testes:
    ```bash
    make test (ou go test -v -cover ./...)
    ```

5. Inicie o servidor e acesse o swagger:
    ```bash
    make fiap-run (ou docker-compose up -d)
    http://localhost:8080/kitchencontrol/api/v1/docs/index.html

    ```
6. Para desligar a aplicação:
    ```bash
    make fiap-stop (ou docker-compose up -d)
    ``` 
7. Para acessar o DB
    ```bash
    http://localhost:8282/
    ```
8. Para gerar CPFs
    ```bash
    https://www.geradordecpf.org/
    ```

## Uso

### Endpoints (acesso via swagger)

- http://localhost:8080/kitchencontrol/api/v1/docs/index.html
- Acessar o swagger, e toda documentação de uso está lá.
- https://www.geradordecpf.org/


## Documentação
- O link para acesso ao miro será enviado aos professores via plataforma da fiap

## Licença
Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## SWAG no desenvolvimento apenas.
- Se o comando swag --version não funcionar, executar os passos abaixo:
    - Para gerar a documentação no padrão open api, será necessário instalar o swag

    ```bash
    go install github.com/swaggo/swag/cmd/swag@latest
    echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
    source ~/.bashrc
    swag --version
    ```
    - no macos, ao invés de .bashrc, use .zhrc

## Algumas dicas
    ### WSL
        - Para deletar os Identifier 
        ```bash
            find . -name "*.Identifier" -type f -delete 
            wsl --install -d Debian 
            wsl -l -v
        ```
    