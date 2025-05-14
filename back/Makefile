POSTGRES_URL ?= "postgres://postgres:docker@:5432/insecticide?sslmode=disable" 
SCHEMA_PATH ?= "file://schema.hcl"
MIGRATIONS_PATH ?= "file://migrations"

run: 
	go run src/cmd/main.go

migrate:
	atlas migrate diff $(name) \
		--dir $(MIGRATIONS_PATH) \
		--to $(SCHEMA_PATH) \
		--dev-url $(POSTGRES_URL) \

migrate-up:
	atlas migrate apply \
		--url $(POSTGRES_URL) \
		--dir $(MIGRATIONS_PATH) \
		--allow-dirty

migrate-status:
	atlas migrate status \
		--url $(POSTGRES_URL) \
		--dir $(MIGRATIONS_PATH)
	

