package main

import (
	"encoding/json"
	"slack-api/domain"
	"slack-api/infra/database"
	"slack-api/interface/repository"
	"slack-api/usecase/service"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	declTable := database.NewDBTable()
	r := repository.NewDeclarationRepository(&declTable)
	s := service.NewDeclarationService(r)

	// bind request body
	reqBody := request.Body
	jsonBytes := ([]byte)(reqBody)
	decl := new(domain.Declaration)
	if err := json.Unmarshal(jsonBytes, decl); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if err := s.Create(decl); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
