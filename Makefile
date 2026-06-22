# ===== CONFIG =====
include .env
export


# ===== MIGRATION =====
migrate-up:
	migrate -path $(MIGRATE_PATH) -database "$(MIGRATE_DATABASE_URL)" up

migrate-down:
	migrate -path $(MIGRATE_PATH) -database "$(MIGRATE_DATABASE_URL)" down 1

migrate-force:
	migrate -path $(MIGRATE_PATH) -database "$(MIGRATE_DATABASE_URL)" force $(version)

migrate-version:
	migrate -path $(MIGRATE_PATH) -database "$(MIGRATE_DATABASE_URL)" version

migrate-create:
		migrate create -ext sql -dir migrations $(name)

# ===== SQLC =====
sqlc:
	sqlc generate

# ===== RUN APP =====
run:
	go run cmd/api/main.go

# ===== BUILD =====
build:
	go build cmd/api/main.go

# ===== CLEAN =====
clean:
	rm -rf bin/