.PHONY: tidy up down restart logs

# Executa go mod tidy para organizar as dependências
tidy:
	go mod tidy

# Roda o docker-compose up após rodar go mod tidy
up: tidy
	docker-compose up --build

# Para os contêineres do Docker Compose
down:
	docker-compose down

# Reinicia os contêineres
restart: down up

# Exibe os logs dos serviços em execução
logs:
	docker-compose logs -f
