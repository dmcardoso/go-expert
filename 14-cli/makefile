SHELL:=/bin/sh

setup-environment: 
	go install github.com/spf13/cobra-cli@latest
	cobra-cli init

connect-sqlite:
	sqlite3 data.db

configure-sqlite:
	touch data.db
	sqlite3 data.db -init database-init.sql

run:
	go run main.go
