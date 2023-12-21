package models

import "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

// Referencia https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/dynamodb#QueryInput
type QueryFilter struct {
	IndexName                 string
	KeyConditionExpression    string
	FilterExpression          string
	ProjectionExpression      string
	ExpressionAttributeNames  map[string]string
	ExpressionAttributeValues map[string]types.AttributeValue
	ScanIndexForward          bool
}

// Referencia https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/dynamodb#UpdateItemInput
type UpdateRequest struct {
	Key                       map[string]types.AttributeValue
	UpdateExpression          string
	ExpressionAttributeNames  map[string]string
	ExpressionAttributeValues map[string]types.AttributeValue
}
