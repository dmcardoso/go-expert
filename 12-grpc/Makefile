SHELL:=/bin/sh

generate-proto: 
	protoc --go_out=. --go-grpc_out=. proto/course_category.proto

connect-sqlite:
	sqlite3 db.sqlite

configure-sqlite:
	touch db.sqlite
	sqlite3 db.sqlite -init database-init.sql

run:
	go run cmd/grpc-server/main.go

connect-evans:
	evans -r repl
