# Variables
CONTAINER_NAME=some-postgres
DB_USER=postgres
DB_PASSWORD=XYZ
DB_NAME=tpespecial
POSTGRES_IMAGE=postgres:16

.PHONY: all setup run clean

# Default command: setup database and run the app
all: setup run

# Setup the database (container + schema)
setup:
	@echo "Arranco el contenedor"
	@docker start $(CONTAINER_NAME) > /dev/null 2>&1 || \
	    docker run --name $(CONTAINER_NAME) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -p 5432:5432 -d $(POSTGRES_IMAGE)
	@echo "Esperando a que PostgreSQL esté listo"
	@until docker exec $(CONTAINER_NAME) pg_isready -U $(DB_USER) -q; do sleep 1; done
	@echo "Iniciando la configuración de la base de datos"
	@docker exec $(CONTAINER_NAME) psql -U $(DB_USER) -tc "SELECT 1 FROM pg_database WHERE datname = '$(DB_NAME)'" | grep -q 1 || \
		docker exec $(CONTAINER_NAME) psql -U $(DB_USER) -c "CREATE DATABASE $(DB_NAME)"
	@echo "Applying schema to '$(DB_NAME)'..."
	@cat ./db/schema/schema.sql | docker exec -i -e PGPASSWORD=$(DB_PASSWORD) $(CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME)
	@echo "Setup completo"

# Run the Go application
run:
	@echo "Ejecutando la aplicación Go"
	@go run .

# Stop and remove the container
clean:
	@echo "Limpiando el contenedor"
	@docker stop $(CONTAINER_NAME) > /dev/null 2>&1 || true
	@docker rm $(CONTAINER_NAME) > /dev/null 2>&1 || true