SHELL:=/bin/sh

setup-environment: 
	pip3 install localstack
	pip3 install awscli-local
	go mod tidy

start-localstack:
	localstack start -d

stop-localstack:
	localstack stop

empty-bucket:
	awslocal s3 rm s3://go-expert --recursive

delete-bucket:
	awslocal s3api delete-bucket --bucket go-expert

localstack-status:
	localstack status services

create-bucket:
	awslocal s3api create-bucket --bucket go-expert

list-buckets:
	awslocal s3api list-buckets

generate-files:
	go run cmd/generator/main.go

upload-files:
	go run cmd/uploader/main.go
