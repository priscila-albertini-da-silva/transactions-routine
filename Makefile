# Nome do binário final
BINARY_NAME=main

# Nome da imagem Docker
IMAGE_NAME=app-go

# Comando padrão para Go
GO=go

# Comando padrão para Docker
DOCKER=docker
DOCKER_COMPOSE=docker-compose

# Build local da aplicação Go
build:
	$(GO) build -o $(BINARY_NAME) .

# Roda a aplicação localmente
run:
	$(GO) run $(BINARY_NAME).$(GO) run

# Compila e sobe os contêineres com o Docker Compose
docker-up:
	$(DOCKER_COMPOSE) up --build

# Para os contêineres
docker-down:
	$(DOCKER_COMPOSE) down

# Build da imagem Docker
docker-build:
	$(DOCKER) build -t $(IMAGE_NAME) .

# Limpa binários gerados
clean:
	rm -f $(BINARY_NAME)

# Testa a aplicação Go
test:
	$(GO) test ./...

# Exibe ajuda com os comandos disponíveis
help:
	@echo "Comandos disponíveis:"
	@echo "  make build          Compila a aplicação Go localmente"
	@echo "  make run            Executa o programa compilado"
	@echo "  make docker-up      Sobe os contêineres com Docker Compose"
	@echo "  make docker-down    Para os contêineres"
	@echo "  make docker-build   Build da imagem Docker"
	@echo "  make test           Roda os testes da aplicação Go"
	@echo "  make clean          Remove arquivos compilados"

