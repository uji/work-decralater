package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type Declaration struct {
	UserID int `json:"UserID" dynamo:"UserID"`
	// Date      time.Time `dynamo:"Date" json:"Date"`
	// StartAt   time.Time `dynamo:"StartAt" json:"StartAt"`
	// EndAt     time.Time `dynamo:"EndAt" json:"EndAt"`
	// Breaktime time.Time `dynamo:"Breaktime" json:"Breaktime"`
	Place     string    `dynamo:"Place" json:"Place"`
	Comment   string    `dynamo:"Comment" json:"Comment"`
	CreatedAt time.Time `dynamo:"CreatedAt"`
}

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
	decl := new(Declaration)
	if err := json.Unmarshal(jsonBytes, decl); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	decl.CreatedAt = time.Now().UTC()
	fmt.Println(decl)

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
