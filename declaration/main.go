package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type Declaration struct {
	ID     int `dynamo:"id" json:"-"`
	UserID int `dynamo:"user_id" json:"user_id"`
	// Date      time.Time `dynamo:"date" json:"date"`
	// StartAt   time.Time `dynamo:"start_at" json:"start_at"`
	// EndAt     time.Time `dynamo:"end_at" json:"end_at"`
	// Breaktime time.Time `dynamo:"breaktime" json:"breaktime"`
	// Place     string    `dynamo:"place" json:"place"`
	// Comment   string    `dynamo:"comment" json:"comment"`
	// CreatedAt time.Time `dynamo:"created_at"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// get db session
	endpoint := os.Getenv("DYNAMODB_ENDPOINT")
	// tableName := os.Getenv("DYNAMODB_TABLE_NAME")
	fmt.Println(request.RequestContext)

	sess := session.Must(session.NewSession())
	config := aws.NewConfig().WithRegion("us-east-1")
	if len(endpoint) > 0 {
		config = config.WithEndpoint(endpoint)
	}
	db := dynamo.New(sess, config)
	declTable := db.Table("declaration")

	// bind request body
	reqBody := request.Body
	jsonBytes := ([]byte)(reqBody)
	decl := Declaration{}
	if err := json.Unmarshal(jsonBytes, &decl); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if err := declTable.Put(decl).Run(); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
