# Usar a imagem oficial do Go para construir a aplicação
FROM golang:1.22.1 AS builder

# Definir o diretório de trabalho dentro do container
WORKDIR /app

# Copiar go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixar as dependências
RUN go mod download
RUN go mod verify
RUN go mod tidy

# Copiar o código da aplicação para o diretório de trabalho
COPY . .

# Compilar a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fiap-rocks ./cmd/kitchencontrol/main.go

# Expor a porta em que a aplicação vai rodar
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./fiap-rocks"]
