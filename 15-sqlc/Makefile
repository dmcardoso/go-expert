SHELL:=/bin/sh

# .PHONY setup-environment
setup-environment: 
	brew install golang-migrate
	brew install sqlc

generate-db:
	sqlc generate

# .PHONY create-migration
create-migration:
	migrate create -ext=sql -dir=sql/migrations -seq init

# .PHONY run-migration
run-migration:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up

# .PHONY down-migration
down-migration:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down
