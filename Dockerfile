# Usa a imagem oficial do Go para build
FROM golang:1.23 AS builder

# Define o diretório de trabalho no container
WORKDIR /app

# Copia os arquivos go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixa as dependências do projeto
RUN go mod download

# Copia todos os arquivos do projeto para o diretório de trabalho
COPY . .

# Executa o build para gerar o binário
RUN go build -o /server

# Usa uma imagem leve para rodar o binário (opcional)
FROM gcr.io/distroless/base-debian10

# Define o diretório de trabalho no container
WORKDIR /

# Copia o binário gerado pelo estágio anterior
COPY --from=builder /server /server

# Porta exposta
EXPOSE 8080

# Executa o binário
CMD ["/server"]