package service

import (
	"context"

	"github.com/Bancar/uala-labssupport-go-core/dynamo/pkg/models"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	ddb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type client interface {
	Query(context.Context, *ddb.QueryInput, ...func(*ddb.Options)) (*ddb.QueryOutput, error)
	GetItem(context.Context, *ddb.GetItemInput, ...func(*ddb.Options)) (*ddb.GetItemOutput, error)
	PutItem(context.Context, *ddb.PutItemInput, ...func(*ddb.Options)) (*ddb.PutItemOutput, error)
	UpdateItem(context.Context, *ddb.UpdateItemInput, ...func(*ddb.Options)) (*ddb.UpdateItemOutput, error)
	DeleteItem(context.Context, *ddb.DeleteItemInput, ...func(*ddb.Options)) (*ddb.DeleteItemOutput, error)
}

type DynamoDB struct {
	c client
}

func New(c client) DynamoDB {
	return DynamoDB{c: c}
}

func (d DynamoDB) GetItem(tableName string, key map[string]types.AttributeValue) (*ddb.GetItemOutput, error) {
	return d.c.GetItem(context.TODO(), &ddb.GetItemInput{
		TableName: &tableName,
		Key:       key,
	})
}

func (d DynamoDB) DeleteItem(tableName string, key map[string]types.AttributeValue) (*ddb.DeleteItemOutput, error) {
	return d.c.DeleteItem(context.TODO(), &ddb.DeleteItemInput{
		TableName: &tableName,
		Key:       key,
	})
}

func (d DynamoDB) PutItem(tableName string, inp interface{}) (*ddb.PutItemOutput, error) {
	item, err := attributevalue.MarshalMap(inp)
	if err != nil {
		return nil, err
	}

	return d.c.PutItem(context.TODO(), &ddb.PutItemInput{
		TableName: &tableName,
		Item:      item,
	})
}

func (d DynamoDB) UpdateItem(tableName string, req models.UpdateRequest) (*ddb.UpdateItemOutput, error) {
	return d.c.UpdateItem(context.TODO(), &ddb.UpdateItemInput{
		TableName:                 &tableName,
		Key:                       req.Key,
		UpdateExpression:          &req.UpdateExpression,
		ExpressionAttributeNames:  req.ExpressionAttributeNames,
		ExpressionAttributeValues: req.ExpressionAttributeValues,
	})
}

func (d DynamoDB) RawQuery(req ddb.QueryInput) (*ddb.QueryOutput, error) {
	return d.c.Query(context.TODO(), &req)
}

func (d DynamoDB) QueryWithPage(tableName string, filter models.QueryFilter, limit int32, lastKey map[string]types.AttributeValue) (*ddb.QueryOutput, error) {
	return d.RawQuery(ddb.QueryInput{
		TableName:                 &tableName,
		ExclusiveStartKey:         lastKey,
		Limit:                     &limit,
		IndexName:                 &filter.IndexName,
		KeyConditionExpression:    &filter.KeyConditionExpression,
		FilterExpression:          &filter.FilterExpression,
		ProjectionExpression:      &filter.ProjectionExpression,
		ExpressionAttributeNames:  filter.ExpressionAttributeNames,
		ExpressionAttributeValues: filter.ExpressionAttributeValues,
		ScanIndexForward:          &filter.ScanIndexForward,
	})
}

func (d DynamoDB) QueryAll(tableName string, filter models.QueryFilter) ([]map[string]types.AttributeValue, error) {
	var lastKey map[string]types.AttributeValue = nil
	var items []map[string]types.AttributeValue

	for {
		out, err := d.RawQuery(ddb.QueryInput{
			TableName:                 &tableName,
			ExclusiveStartKey:         lastKey,
			IndexName:                 &filter.IndexName,
			KeyConditionExpression:    &filter.KeyConditionExpression,
			FilterExpression:          &filter.FilterExpression,
			ProjectionExpression:      &filter.ProjectionExpression,
			ExpressionAttributeNames:  filter.ExpressionAttributeNames,
			ExpressionAttributeValues: filter.ExpressionAttributeValues,
			ScanIndexForward:          &filter.ScanIndexForward,
		})
		if err != nil {
			return items, err
		}
		items = append(items, out.Items...)

		if out.LastEvaluatedKey == nil {
			break
		} else {
			lastKey = out.LastEvaluatedKey
		}
	}

	return items, nil
}
