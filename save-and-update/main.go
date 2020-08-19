package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"encoding/json"
	"../types"
	"../dynamo"
)


func main() {
	lambda.Start(Handler)
}


func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var monster types.Monster

	marshalErr := json.Unmarshal([]byte(req.Body), &monster)

	if marshalErr != nil {
		return response("Couldn't unmarshal json into monster struct", http.StatusBadRequest), nil
	}

	dynamoErr := dynamo.SaveMonster(monster)

	if dynamoErr != nil {
		return response(dynamoErr.Error(), http.StatusInternalServerError), nil
	}

	return response("Successfully wrote monster to log.", http.StatusOK), nil
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