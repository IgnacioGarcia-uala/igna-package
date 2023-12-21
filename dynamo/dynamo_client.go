package dynamo_client

import (
	"context"
	"log"

	"github.com/IgnacioGarcia-uala/igna-package/dynamo/pkg/models"
	"github.com/IgnacioGarcia-uala/igna-package/dynamo/pkg/service"
	"github.com/aws/aws-sdk-go-v2/config"
	ddb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoClient interface {
	/* Obtener un elemento por su clave primaria
	Params:
	- TableName: string con el nombre de la tabla a la cual contiene el elemento
	- Key: mapa con los nombres de los atributos que forman la clave primaria, y los valores que debera tener
	Output:
	1. Resultado de la busqueda, contiene el Item
	2. Error, puede producirse al intentar obtener el Item
	*/
	GetItem(tableName string, key map[string]types.AttributeValue) (*ddb.GetItemOutput, error)

	/* Eliminar un elemento por su clave primaria
	Params:
	- TableName: string con el nombre de la tabla a la cual contiene el elemento
	- Key: mapa con los nombres de los atributos que forman la clave primaria, y los valores que debera tener
	Output:
	1. Resultado de la eliminacion, contiene el Item
	2. Error, puede producirse al intentar eliminar el Item
	*/
	DeleteItem(tableName string, key map[string]types.AttributeValue) (*ddb.DeleteItemOutput, error)

	/* Insertar un elemento o Actualizarlo por su clave primaria
	Params:
	- TableName: string con el nombre de la tabla a la cual contiene el elemento
	- Item: estructura con los valores a insertar, se recomienda definir el alias de cada atributo. Ver mas en
	https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/dynamodbattribute/#:~:text=%60dynamodbav%60%20struct
	Output:
	1. Resultado de la insercion, contiene el Item
	2. Error, puede producirse al intentar insertar el Item
	*/
	PutItem(tableName string, item interface{}) (*ddb.PutItemOutput, error)

	/* Insertar un elemento o Actualizarlo por su clave primaria,
	permite agregar, eliminar o modificar atributos del Item
	Params:
	- TableName: string con el nombre de la tabla a la cual contiene el elemento
	- UpdateRequest: estructura la clave del Item, la expresion de actualizacion y los valores que contentra
	Output:
	1. Resultado de la actualizacion, contiene el Item
	2. Error, puede producirse al intentar actualizar el Item
	*/
	UpdateItem(tableName string, request models.UpdateRequest) (*ddb.UpdateItemOutput, error)

	/* Realizar una Consulta
	Params:
	- QueryInput: estructura que contiene la tabla y los filtros a aplicar. Ver mas en
	https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/dynamodb#QueryInput
	Output:
	1. Resultado de la consulta, contiene los Items
	2. Error, puede producirse al intentar realizar la consulta
	*/
	RawQuery(request ddb.QueryInput) (*ddb.QueryOutput, error)

	/* Realizar una Consulta, con paginado
	Params:
	- TableName: string con el nombre de la tabla a la cual realizar la consulta
	- QueryFilter: filtros a aplicar en la tabla
	- Limit: entero que contiene el maximo numero de items a evaluar(no a retornar)
	- LastKey: ultimo elemento evaluado en la anterior consulta, utilizado para el paginado
	Output:
	1. Resultado de la consulta, contiene los Items
	2. Error, puede producirse al intentar realizar la consulta
	*/
	QueryWithPage(tableName string, request models.QueryFilter, limit int32, lastKey map[string]types.AttributeValue) (*ddb.QueryOutput, error)

	/* Realizar una Consulta, en una unica pagina
	Params:
	- TableName: string con el nombre de la tabla a la cual realizar la consulta
	- QueryFilter: filtros a aplicar en la tabla
	Output:
	1. Listado de todos los Items que matcheen con el filtro
	2. Error, puede producirse al intentar realizar la consulta
	*/
	QueryAll(tableName string, request models.QueryFilter) ([]map[string]types.AttributeValue, error)
}

func New(ctx context.Context) DynamoClient {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("configuration error, %s", err.Error())
	}
	return service.New(ddb.NewFromConfig(cfg))
}
