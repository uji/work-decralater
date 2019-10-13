package database

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func NewDBTable() dynamo.Table {
	endpoint := os.Getenv("DYNAMODB_ENDPOINT")
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	fmt.Println(endpoint)

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
	return db.Table(tableName)
}
