package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slack-api/domain"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var layout = "2001-01-01 00:00:00 +0000 UTC"

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// get db session
	endpoint := os.Getenv("DYNAMODB_ENDPOINT")
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	fmt.Println(endpoint)
	fmt.Println(tableName)

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Endpoint:    aws.String(endpoint),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", ""),
		DisableSSL:  aws.Bool(true),
	})

	if err != nil {
		panic(err)
	}

	db := dynamo.New(sess)
	declTable := db.Table(tableName)

	// bind request body
	reqBody := request.Body
	jsonBytes := ([]byte)(reqBody)
	decl := new(domain.Declaration)
	if err := json.Unmarshal(jsonBytes, decl); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	// save
	decl.CreatedAt = time.Now().UTC()
	if err := declTable.Put(decl).Run(); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
