.PHONY: deps clean build

deps:
	go get -u ./...

clean:
	rm -rf ./declaration/declaration

build:
	GOOS=linux GOARCH=amd64 go build -o declaration/declaration ./declaration

export:
	export AWS_PROFILE=default

create_table:
	aws dynamodb create-table --table-name DeclarationDynamoDBTable \
														--attribute-definitions AttributeName=UserID,AttributeType=N \
														--key-schema AttributeName=UserID,KeyType=HASH \
														--provisioned-throughput ReadCapacityUnits=2,WriteCapacityUnits=2 \
														--endpoint-url http://localhost:8000
