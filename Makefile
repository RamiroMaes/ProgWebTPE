# Variables
CONTAINER_NAME=some-postgres
DB_USER=postgres
DB_PASSWORD=XYZ
DB_NAME=tpespecial
POSTGRES_IMAGE=postgres:16
HTTP_PORT ?= 8080
# Apuntamos a nuestro script de BASH que tiene las lineas con metodos curl
TEST_FILE ?= requests.bash

.PHONY: all setup run e2e test clean

# Setup DB (contenedor + create DB si falta + schema)
setup:
	@echo "Arranco el contenedor"
	@docker start $(CONTAINER_NAME) >/dev/null 2>&1 || \
	docker run --name $(CONTAINER_NAME) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -p 5432:5432 -d $(POSTGRES_IMAGE)
	@echo "Esperando a PostgreSQL"
	@until docker exec $(CONTAINER_NAME) pg_isready -U $(DB_USER) -q; \
	do sleep 1; done
	@echo "Creando base si esta no existe"
	@docker exec $(CONTAINER_NAME) psql -U $(DB_USER) -tc "SELECT 1 FROM pg_database WHERE datname='$(DB_NAME)'" | \
	grep -q 1 || docker exec $(CONTAINER_NAME) psql -U $(DB_USER) -c "CREATE DATABASE $(DB_NAME)"
	@echo "Aplicando el schema a la base"
	@cat ./db/schema/schema.sql | \
	docker exec -i -e PGPASSWORD=$(DB_PASSWORD) $(CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME)
	@echo "Setup OK"


# E2E: DB + server bg + BASH + stop
test todo: setup
	@echo "→ Arrancando server en background..."
	@go run ./cmd/server > .server.log 2>&1 & echo $$! > .server.pid
	@echo "→ Esperando HTTP en :$(HTTP_PORT)..."
	@for i in $$(seq 1 40); do curl -s http://localhost:$(HTTP_PORT)/ >/dev/null 2>&1 && break; \
	sleep 1; done
	@echo "→ Ejecutando Tests (curl)..."
	@chmod +x $(TEST_FILE)
	@./$(TEST_FILE)
	@status=$$?; echo "→ Deteniendose el server"; kill $$(cat .server.pid) 2>/dev/null || true; \
	rm -f .server.pid; exit $$status

# Limpieza de contenedor
clean:
	@echo "Limpiando contenedor"
	@docker stop $(CONTAINER_NAME) >/dev/null 2>&1 || true
	@docker rm  $(CONTAINER_NAME) >/dev/null 2>&1 || true