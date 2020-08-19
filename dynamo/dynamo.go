package dynamo

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"../types"
)

const tableName = "fantasy-monsters"

func SaveMonster(monster types.Monster) error {
	monsterMap, marshalErr := dynamodbattribute.MarshalMap(monster)

	if marshalErr != nil {
		fmt.Println("Failed to marshal to dynamo map")
		return marshalErr
	}

	dynamoSession := createDynamoSession()

	input := &dynamodb.PutItemInput{
    Item:      monsterMap,
    TableName: aws.String(tableName),
	}

	_, writeErr := dynamoSession.PutItem(input)

	if writeErr != nil {
			fmt.Println("Failed to write to dynamo")
			return writeErr
	}

	return nil
}

func createDynamoSession() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		},
	))

	return dynamodb.New(sess)
}