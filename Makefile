# Variables (del fuente 1)
CONTAINER_NAME=some-postgres
DB_USER=postgres
DB_PASSWORD=XYZ
DB_NAME=tpespecial
POSTGRES_IMAGE=postgres:16
HTTP_PORT ?= 8080
TEST_FILE ?=requests.bash

# Lista de todas las reglas "phony"
.PHONY: all setup container-start container-wait db-create db-schema test clean

# Regla 1: Inicia un contenedor existente o crea uno nuevo
container-start:
	@echo "→ Iniciando contenedor $(CONTAINER_NAME)..."
	@docker start $(CONTAINER_NAME) >/dev/null 2>&1 || \
	docker run --name $(CONTAINER_NAME) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -p 5432:5432 -d $(POSTGRES_IMAGE) # (Basado en fuente 2)

# Regla 2: Espera a que el contenedor de PostgreSQL esté listo
container-wait: container-start
	@echo "→ Esperando a PostgreSQL..."
	@until docker exec $(CONTAINER_NAME) pg_isready -U $(DB_USER) -q; \
	do sleep 1; done # (Basado en fuente 2, 3)

# Regla 3: Crea la base de datos si no existe
db-create: container-wait
	@echo "→ Creando base de datos $(DB_NAME) (si no existe)..."
	@docker exec $(CONTAINER_NAME) psql -U $(DB_USER) -tc "SELECT 1 FROM pg_database WHERE datname='$(DB_NAME)'" | \
	grep -q 1 || docker exec $(CONTAINER_NAME) psql -U $(DB_USER) -c "CREATE DATABASE $(DB_NAME)" # (Basado en fuente 3, 4)

# Regla 4: Aplica el schema.sql a la base de datos
db-schema: db-create
	@echo "→ Aplicando el schema a la base..."
	@cat ./db/schema/schema.sql | \
	docker exec -i -e PGPASSWORD=$(DB_PASSWORD) $(CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME) # (Basado en fuente 4, 5)

# Regla 5: Setup completo
setup: db-schema
	@echo "Setup OK"

# Regla 6: Test
# Test: DB + server bg + BASH + stop
test: setup
	@echo "→ Arrancando server en background..."
	@go run ./cmd/server > .server.log 2>&1 & echo $$! > .server.pid # (Basado en fuente 6)
	@echo "→ Esperando HTTP en :$(HTTP_PORT)..."
	@for i in $$(seq 1 40); do curl -s http://localhost:$(HTTP_PORT)/ >/dev/null 2>&1 && break; \
	sleep 1; done # (Basado en fuente 6, 7)
	@echo "→ Ejecutando Tests (curl)..."
	@chmod +x $(TEST_FILE)
	@./$(TEST_FILE) # (Basado en fuente 7)
	@status=$$?; echo "→ Deteniendose el server"; kill $$(cat .server.pid) 2>/dev/null || true; \
	rm -f .server.pid; exit $$status # (Basado en fuente 8)

# Limpieza de contenedor
clean:
	@echo "Deteniendo servidor en :$(HTTP_PORT) (si existe)..."
	@lsof -t -i:$(HTTP_PORT) | xargs kill 2>/dev/null || true
	@echo "Limpiando contenedor"
	@docker stop $(CONTAINER_NAME) >/dev/null 2>&1 || true # (Basado en fuente 9)
	@docker rm  $(CONTAINER_NAME) >/dev/null 2>&1 || true # (Basado en fuente 9)