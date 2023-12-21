# dynamo_client
Cliente para interactuar con AWS DynamoDB, para utilizar en las Lambdas de Labs & Support.  


## Uso
Para importar un paquete especifico
`go get github.com/Bancar/uala-labssupport-go-core/dynamo`  

Ejecutar el siguiente comando para traerte una version especifica del paquete  
`go get github.com/Bancar/uala-labssupport-go-core/dynamo/dynamo@vX.Y.Z`
  
En caso de querer usar una version que aun no esta en main se debe ejecutar el siguiente comando, reemplazando `branch` con el nombre de la rama requerida  
`go get github.com/Bancar/uala-labssupport-go-core/dynamo_client@branch`


## Ejemplos
- Ej
```go
package main

import (
	"context"
	"fmt"

	dynamo_client "github.com/Bancar/uala-labssupport-go-core/dynamo-client"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Item struct {
	Id   string
	Name string
}

func getItem() {
	dc := dynamo_client.New(context.TODO())

	key := map[string]types.AttributeValue{
		"Id": &types.AttributeValueMemberS{Value: "abv-123-a4d"},
	}
	res, err := dc.GetItem("devShapes", key)
	if err != nil {
		return
	}

	fmt.Print(res)
}

func deleteItem() {
	dc := dynamo_client.New(context.TODO())

	key := map[string]types.AttributeValue{
		"Id": &types.AttributeValueMemberS{Value: "abv-123-a4d"},
	}
	res, err := dc.DeleteItem("devShapes", key)
	if err != nil {
		return
	}

	fmt.Print(res)
}

func putItem() {
	dc := dynamo_client.New(context.TODO())

	item := Item{
		Id:   "abc-123",
		Name: "some item",
	}
	res, err := dc.PutItem("devShapes", item)
	if err != nil {
		return
	}

	fmt.Print(res)
}

func updateItem() {
	dc := dynamo_client.New(context.TODO())

	req := models.UpdateRequest{
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: "abc-123"},
		},
		UpdateExpression: "set Name = :name",
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":name": &types.AttributeValueMemberS{Value: "other value"},
		},
	}

	res, err := dc.UpdateItem("devShapes", req)
	if err != nil {
		return
	}

	fmt.Print(res)
}

func query() {
	dc := dynamo_client.New(context.TODO())

	req := models.QueryFilter{
		IndexName:              "name-index",
		KeyConditionExpression: "Name = :name",
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":name": &types.AttributeValueMemberS{Value: "common name"},
		},
	}
	res, err := dc.QueryWithPage("devShapes", req, 20, nil)
	if err != nil {
		return
	}
	fmt.Print(res)

	if res.LastEvaluatedKey != nil {
		res, err = dc.QueryWithPage("devShapes", req, 20, res.LastEvaluatedKey)
		if err != nil {
			return
		}
		fmt.Print(res)
	}
}

func queryAll() {
	dc := dynamo_client.New(context.TODO())

	req := models.QueryFilter{
		IndexName:              "name-index",
		KeyConditionExpression: "Name = :name",
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":name": &types.AttributeValueMemberS{Value: "common name"},
		},
	}
	res, err := dc.QueryAll("devShapes", req)
	if err != nil {
		return
	}
	fmt.Print(res)
}
```