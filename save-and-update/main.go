package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"encoding/json"
	"../types"
)


func main() {
	lambda.Start(Handler)
}


func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var monster types.Monster

	err := json.Unmarshal([]byte(req.Body), &monster)

	if err != nil {
		return response("Couldn't unmarshal json into monster struct", http.StatusBadRequest), nil
	}

	return response(monster.Name, http.StatusOK), nil
}

func response(body string, statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse {
		StatusCode: statusCode,
		Body: string(body),
		Headers: map[string]string {
			"Access-Control-Allow-Origin": "*",
		},
	}
}